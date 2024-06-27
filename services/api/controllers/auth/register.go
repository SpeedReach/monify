package auth

import (
	"context"
	"database/sql"
	"monify/lib"
	monify "monify/protobuf/gen/go"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func emailExists(ctx context.Context, email string) (bool, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
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

func CreateUser(ctx context.Context, email string, password string) (uuid.UUID, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	if email == "" || password == "" {
		return uuid.Nil, status.Error(codes.InvalidArgument, "Email and password is required.")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return [16]byte{}, err
	}

	userId := uuid.New()
	_, err = db.ExecContext(ctx, `
		INSERT INTO user_identity (user_id, nick_id) VALUES ($1, $2)
	`, userId, userId.String())
	if err != nil {
		return uuid.Nil, err
	}

	_, err = db.ExecContext(ctx, `INSERT INTO email_login(email, user_id, password) VALUES ($1, $2, $3)`, email, userId, string(hashedPassword))
	if err != nil {
		return uuid.Nil, err
	}
	return userId, nil
}

func (s Service) EmailRegister(ctx context.Context, req *monify.EmailRegisterRequest) (*monify.EmailRegisterResponse, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	exists, err := emailExists(ctx, req.Email)
	if err != nil {
		logger.Error("failed to query email", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal err.")
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "Email already exists.")
	}

	userId, err := CreateUser(ctx, req.Email, req.Password)
	if err != nil {
		logger.Error("failed to create user", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal err.")
	}
	return &monify.EmailRegisterResponse{UserId: userId.String()}, nil
}
