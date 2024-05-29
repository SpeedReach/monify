package auth

import (
	"context"
	"database/sql"
	"monify/internal/middlewares"
	monify "monify/protobuf/gen/go"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func get_userID(ctx context.Context, db *sql.DB, refreshToken string) (uuid.UUID, error) {

	query, err := db.QueryContext(ctx, `
	SELECT user_id
	FROM user_identity
	WHERE refresh_token= $1`, refreshToken)
	defer query.Close()
	if err != nil {
		return uuid.Nil, err
	}
	if !query.Next() {
		return uuid.Nil, nil
	}
	var userId uuid.UUID
	err = query.Scan(&userId)
	if err != nil {
		return uuid.Nil, err
	}
	return userId, nil

}

func (s Service) RefreshToken(ctx context.Context, req *monify.RefreshTokenRequest) (*monify.RefreshTokenResponse, error) {
	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	userId, err := get_userID(ctx, db, req.RefreshToken)

	//它給我refresh token，我要給refresh token + access token
	refreshToken, err := generateAndInsertRefreshToken(ctx, userId, db)
	if err != nil {
		logger.Error("failed to get refresh token", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal err.")
	}
	accessToken, err := GenerateAccessToken(ctx, userId, s.Secret)

	if err != nil {
		logger.Error("failed to get access token", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal err.")

	}
	return &monify.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
