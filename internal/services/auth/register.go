package auth

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/internal/middlewares"
	monify "monify/protobuf"
)

func emailExists(ctx context.Context, email string, db *sql.DB) (bool, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT user_id
		FROM email_login
		WHERE email = $1
	`, email)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}

func createUser(ctx context.Context, db *sql.DB, email string) (uuid.UUID, error) {
	userId := uuid.New()
	_, err := db.ExecContext(ctx, `
		INSERT INTO user_identity (user_id) VALUES ($1)
	`, userId)
	if err != nil {
		return uuid.Nil, err
	}
	_, err = db.ExecContext(ctx, `INSERT INTO email_login(email, user_id) VALUES ($1, $2)`, email, userId)
	if err != nil {
		return uuid.Nil, err
	}
	return userId, nil
}

func (s Service) EmailRegister(ctx context.Context, req *monify.EmailRegisterRequest) (*monify.EmailRegisterResponse, error) {
	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	exists, err := emailExists(ctx, req.Email, db)
	if err != nil {
		logger.Error("failed to query email", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal err.")
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "Email already exists.")
	}

	userId, err := createUser(ctx, db, req.Email)
	if err != nil {
		logger.Error("failed to create user", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal err.")
	}
	return &monify.EmailRegisterResponse{UserId: userId.String()}, nil
}