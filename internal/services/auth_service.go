package services

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/internal/middlewares"
	monify "monify/protobuf"
	"time"
)

type AuthService struct {
	Secret string
	monify.UnimplementedAuthServiceServer
}

func (s AuthService) EmailLogin(ctx context.Context, req *monify.EmailLoginRequest) (*monify.EmailLoginResponse, error) {
	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	query, err := db.Query(`
	SELECT user_id 
	FROM email_login 
	WHERE email = $1`, req.Email)

	if err != nil {
		return nil, err
	}

	if !query.Next() {
		return nil, status.Error(codes.NotFound, "email not found.")
	}

	var userId uuid.UUID
	err = query.Scan(&userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal err.")
	}

	//generate refresh token and insert into database
	var refreshToken string = uuid.New().String()
	_, err = db.ExecContext(ctx, `
	UPDATE user_identity SET refresh_token=$1 WHERE user_id = $2
	`, refreshToken, userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal err.")
	}

	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userId.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	ss, err := token.SignedString(s.Secret)
	if err != nil {
		return nil, err
	}

	return &monify.EmailLoginResponse{
		UserId:       userId.String(),
		AccessToken:  ss,
		RefreshToken: refreshToken,
	}, nil
}
