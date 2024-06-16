package auth

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/lib"
	monify "monify/protobuf/gen/go"
	"time"
)

func matchEmailUser(ctx context.Context, email string, password string) (uuid.UUID, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	query, err := db.QueryContext(ctx, `
	SELECT user_id, password 
	FROM email_login 
	WHERE email = $1`, email)
	defer query.Close()
	if err != nil {
		return uuid.Nil, err
	}

	if !query.Next() {
		return uuid.Nil, nil
	}

	var userId uuid.UUID
	var hashedPassword string
	err = query.Scan(&userId, &hashedPassword)
	if err != nil {
		logger.Error("", zap.Error(err))
		return uuid.Nil, status.Error(codes.Internal, "internal err.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return uuid.Nil, status.Error(codes.PermissionDenied, "Password incorrect.")
	}

	return userId, nil
}

func GenerateAccessToken(ctx context.Context, userId uuid.UUID, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userId.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	return token.SignedString([]byte(secret))
}

func (s Service) EmailLogin(ctx context.Context, req *monify.EmailLoginRequest) (*monify.EmailLoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "Email and password is required.")
	}
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)

	userId, err := matchEmailUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	if userId == uuid.Nil {
		return nil, status.Error(codes.NotFound, "Email not found.")
	}

	//generate refresh token and insert into database
	refreshToken := GenerateRefreshToken(userId)
	if err = InsertRefreshToken(ctx, userId, refreshToken); err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal err.")
	}

	ss, err := GenerateAccessToken(ctx, userId, s.Secret)

	if err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal err.")
	}

	return &monify.EmailLoginResponse{
		UserId:       userId.String(),
		AccessToken:  ss,
		RefreshToken: refreshToken,
	}, nil
}

/// infra

func InsertRefreshToken(ctx context.Context, userId uuid.UUID, refreshToken string) error {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	_, err := db.ExecContext(ctx, `
	UPDATE user_identity SET refresh_token=$1 WHERE user_id = $2
	`, refreshToken, userId)
	return err
}
