package auth

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/internal/middlewares"
	monify "monify/protobuf"
	"time"
)

func matchEmailUser(ctx context.Context, email string, password string, db *sql.DB) (uuid.UUID, error) {
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
		return uuid.Nil, err
	}

	_ = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return uuid.Nil, errors.New("incorrect password")
	}

	return userId, nil
}

func generateAndInsertRefreshToken(ctx context.Context, userId uuid.UUID, db *sql.DB) (string, error) {
	var refreshToken string = uuid.New().String()
	_, err := db.ExecContext(ctx, `
	UPDATE user_identity SET refresh_token=$1 WHERE user_id = $2
	`, refreshToken, userId)
	return refreshToken, err
}

func GenerateAccessToken(ctx context.Context, userId uuid.UUID, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userId.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	return token.SignedString([]byte(secret))
}

func (s Service) EmailLogin(ctx context.Context, req *monify.EmailLoginRequest) (*monify.EmailLoginResponse, error) {
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)

	userId, err := matchEmailUser(ctx, req.Email, req.Password, db)
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal err.")
	}

	if userId == uuid.Nil {
		return nil, status.Error(codes.NotFound, "Email not found.")
	}

	//generate refresh token and insert into database
	refreshToken, err := generateAndInsertRefreshToken(ctx, userId, db)
	if err != nil {
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
