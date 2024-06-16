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
	"monify/lib"
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

func (s Service) GenerateInviteCode(ctx context.Context, req *monify.GenerateInviteCodeRequest) (*monify.GenerateInviteCodeResponse, error) {
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}

	var groupId uuid.UUID
	var err error
	if groupId, err = uuid.Parse(req.GroupId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid group ID")
	}

	var hasPerm bool
	if hasPerm, err = CheckPermission(ctx, groupId, userId); err != nil {
		logger.Error("Failed to check permission", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	if !hasPerm {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")
	}

	inviteCode := generateInviteCode()
	if _, err = db.Exec(`
		INSERT INTO group_invite_code (group_id, invite_code) VALUES ($1, $2)
	`, groupId, inviteCode); err != nil {
		logger.Error("Failed to insert invite code", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
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
