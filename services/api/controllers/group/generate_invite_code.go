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
	"monify/lib/group"
	monify "monify/protobuf/gen/go"
	"time"
)

const (
	inviteCodeChars  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	timeDeterLength  = 4
	randomLength     = 2
	inviteCodeLength = timeDeterLength + randomLength
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

	code, err := getExistsInviteCode(ctx, groupId)
	if err != nil {
		logger.Error("Failed to get invite code", zap.Error(err))
		return nil, err
	}
	if code != (group.InviteCode{}) && !code.IsExpired() {
		_, err := db.ExecContext(ctx, `DELETE FROM group_invite_code WHERE invite_code = $1`, code.GroupId)
		if err != nil {
			logger.Error("Failed to delete invite code", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}
		return &monify.GenerateInviteCodeResponse{InviteCode: code.Code}, nil
	}

	// generate invite code, we retry when the invite code already exists and is active
	retries := 0
	var inviteCode string

	for retries < 3 {
		retries++
		inviteCode = generateInviteCode()
		tx, err := db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelRepeatableRead,
		})
		if err != nil {
			logger.Error("Failed to start transaction", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}
		defer tx.Rollback()

		//check if invite code already exists and is active
		row := tx.QueryRowContext(ctx, "SELECT created_at FROM group_invite_code WHERE invite_code = $1", inviteCode)
		var createdAt time.Time
		err = row.Scan(&createdAt)
		if err != nil && err != sql.ErrNoRows {
			logger.Error("", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}

		switch err {
		// if invite code does not exist, we insert the invite code
		case sql.ErrNoRows:
			_, err = tx.ExecContext(ctx, `
				INSERT INTO group_invite_code (group_id, invite_code) VALUES ($1, $2)
				`, groupId, inviteCode)
			if err != nil {
				logger.Error("Failed to insert invite code", zap.Error(err))
				return nil, status.Error(codes.Internal, "Internal")
			}
			break
		case nil:
			// if invite code exists and is active, we wait and retry
			if createdAt.Add(group.ExpiresInterval).After(time.Now()) {
				// valid invite code already exists retry
				time.Sleep(time.Millisecond * 83)
				continue
			}
			// if invite code exists but is expired, we update the invite code
			_, err = tx.ExecContext(ctx, `
			UPDATE group_invite_code SET created_at = $1, group_id = $2 WHERE invite_code = $3`, time.Now(), groupId, inviteCode)
			if err != nil {
				logger.Error("Failed to update invite code", zap.Error(err))
				return nil, status.Error(codes.Internal, "Internal")
			}
			break
		default:
			logger.Error("Failed to scan invite code", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}

		if err = tx.Commit(); err != nil {
			logger.Error("Failed to commit transaction", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}
		break
	}

	return &monify.GenerateInviteCodeResponse{InviteCode: inviteCode}, err
}

func indexToChar(index int) byte {
	return inviteCodeChars[index]
}
func generateInviteCode() string {
	charsCount := len(inviteCodeChars)
	seed := time.Now().UnixMilli() % int64(group.ExpiresInterval)
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

func getExistsInviteCode(ctx context.Context, groupId uuid.UUID) (group.InviteCode, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	inviteCode := group.InviteCode{GroupId: groupId}
	err := db.QueryRowContext(ctx, "SELECT invite_code, created_at FROM group_invite_code WHERE group_id = $1", groupId).Scan(&inviteCode.Code, &inviteCode.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return group.InviteCode{}, nil
		}
		logger.Error("Failed to get invite code", zap.Error(err))
		return group.InviteCode{}, status.Error(codes.Internal, "Internal")
	}
	return inviteCode, nil
}
