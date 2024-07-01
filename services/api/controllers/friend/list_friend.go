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

func (s Service) ListFriend(ctx context.Context, req *monify.FriendEmpty) (*monify.ListFriendResponse, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	query, err := db.QueryContext(ctx, `
		SELECT user1_id, user2_id, name, relation_id
		FROM friend, user_identity
		WHERE (user1_id = $1 AND user2_id = user_id) OR (user2_id = $2 AND user1_id = user_id)`, userId, userId)
	if err != nil {
		logger.Error("Select friend of userId error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	defer query.Close()

	var friends []*monify.Friend
	for {
		if !query.Next() {
			break
		}
		var friend monify.Friend
		var user1Id uuid.UUID
		var user2Id uuid.UUID
		if err = query.Scan(&user1Id, &user2Id, &friend.Name, &friend.RelationId); err != nil {
			logger.Error("Scan error.", zap.Error(err))
			return nil, status.Error(codes.Internal, "")
		}
		if user1Id == userId {
			friend.FriendId = user2Id.String()
			friends = append(friends, &friend)
		} else {
			friend.FriendId = user1Id.String()
			friends = append(friends, &friend)
		}
	}

	return &monify.ListFriendResponse{
		Friends: friends,
	}, nil

}
