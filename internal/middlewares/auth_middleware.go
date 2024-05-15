package middlewares

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type AuthMiddleware struct {
	JwtSecret string
}

func (m AuthMiddleware) PreHandler(ctx context.Context, req any, info *grpc.UnaryServerInfo) error {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return nil
	}
	auths := md.Get("authorization")
	if len(auths) == 0 {
		return nil
	}
	auth := auths[0]

	if strings.HasPrefix(auth, "Bearer ") {
		tokenStr := auth[7:]
		token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.JwtSecret), nil
		})
		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
			ctx = context.WithValue(ctx, UserIdContextKey{}, claims.Subject)
		} else {
			switch err {
			case nil:
				return status.Error(codes.Unauthenticated, "Unauthenticated")
			default:
				return status.Error(codes.Unauthenticated, err.Error())
			}
		}
	}
	return nil
}
