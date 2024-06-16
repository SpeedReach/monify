package middlewares

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type AuthMiddleware struct {
	JwtSecret string
}

func (m AuthMiddleware) ExtractUserId(ctx context.Context, req any, info *grpc.UnaryServerInfo) (uuid.UUID, error) {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return uuid.Nil, nil
	}
	auths := md.Get("authorization")
	if len(auths) == 0 {
		return uuid.Nil, nil
	}
	auth := auths[0]

	if strings.HasPrefix(auth, "Bearer ") {
		tokenStr := auth[7:]
		token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.JwtSecret), nil
		})
		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
			userId, err := uuid.Parse(claims.Subject)
			if err != nil {
				return uuid.Nil, status.Error(codes.Unauthenticated, err.Error())
			}

			return userId, nil
		} else {
			switch err {
			case nil:
				return uuid.Nil, status.Error(codes.Unauthenticated, "Unauthenticated")
			default:
				return uuid.Nil, status.Error(codes.Unauthenticated, err.Error())
			}
		}
	}
	return uuid.Nil, nil
}
