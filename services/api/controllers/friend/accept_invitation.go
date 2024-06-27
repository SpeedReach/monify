package friend

import (
	"context"
	"database/sql"
	"monify/lib"
	monify "monify/protobuf/gen/go"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) AcceptInvitation(ctx context.Context, req *monify.AcceptInvitationRequest) (*monify.FriendEmpty, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	userId := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if userId == uuid.Nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}

	//START transaction
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()

	// insert  the friend relation
	relationId := uuid.New()
	var user1Id uuid.UUID
	var user2Id uuid.UUID
	// make sure user1_id < user2_id
	if req.User1Id > req.User2Id {
		user1Id, err = uuid.Parse(req.User2Id)
		if err != nil {
			return nil, status.Error(codes.Internal, "Convert string to uuid error.")
		}
		user2Id, err = uuid.Parse(req.User1Id)
		if err != nil {
			return nil, status.Error(codes.Internal, "Convert string to uuid error.")
		}
	} else {
		user1Id, err = uuid.Parse(req.User1Id)
		if err != nil {
			return nil, status.Error(codes.Internal, "Convert string to uuid error.")
		}
		user2Id, err = uuid.Parse(req.User2Id)
		if err != nil {
			return nil, status.Error(codes.Internal, "Convert string to uuid error.")
		}
	}
	// insert
	_, err = tx.ExecContext(ctx, `INSERT INTO friend (relation_id, user1_id, user2_id) VALUES ($1, $2, $3)`, relationId, user1Id, user2Id)
	if err != nil {
		logger.Error("Insert friend relation error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	// mark that the invitation is accepted
	_, err = tx.ExecContext(ctx, `UPDATE friend_invite SET is_accepted = TRUE WHERE invite_id = $1`, req.InviteId)
	if err != nil {
		logger.Error("Update friend_invite is_accepted error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	//COMMIT transaction
	if err = tx.Commit(); err != nil {
		logger.Error("failed to commit", zap.Error(err))
		return nil, err
	}

	return &monify.FriendEmpty{}, nil

}
