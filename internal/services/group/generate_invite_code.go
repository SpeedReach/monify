package group

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"math/rand"
	"monify/internal/middlewares"
	monify "monify/protobuf/gen/go"
	"time"
)

const (
	inviteCodeChars  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	timeDeterLength  = 4
	randomLength     = 2
	inviteCodeLength = timeDeterLength + randomLength
	expiresInterval  = int64(time.Minute * 10)
)

func checkPermission(ctx context.Context, db *sql.DB, groupId uuid.UUID, userId uuid.UUID) (bool, error) {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) FROM group_member WHERE group_id = $1 AND user_id = $2
	`, groupId, userId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s Service) GenerateInviteCode(ctx context.Context, req *monify.GenerateInviteCodeRequest) (*monify.GenerateInviteCodeResponse, error) {
	userId, ok := ctx.Value(middlewares.UserIdContextKey{}).(uuid.UUID)
	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid group ID")
	}
	hasPerm, err := checkPermission(ctx, db, groupId, userId)
	if err != nil {
		logger.Error("Failed to check permission", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	if !hasPerm {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")
	}
	inviteCode := generateInviteCode()
	_, err = db.Exec(`
		INSERT INTO group_invite_code (group_id, invite_code) VALUES ($1, $2)
	`, groupId, inviteCode)
	return &monify.GenerateInviteCodeResponse{InviteCode: inviteCode}, err
}

func indexToChar(index int) byte {
	return inviteCodeChars[index]
}
func generateInviteCode() string {
	charsCount := len(inviteCodeChars)
	seed := time.Now().UnixMilli() % expiresInterval
	inviteCodeRange := int(math.Pow(float64(charsCount), timeDeterLength))
	code := int(seed) % inviteCodeRange
	inviteCode := ""

	for i := 0; i < timeDeterLength; i++ {
		index := code % charsCount
		code /= charsCount
		inviteCode += string(indexToChar(index))
	}

	for i := 0; i < randomLength; i++ {
		inviteCode += string(indexToChar(rand.Int() % charsCount))
	}
	return inviteCode
}
