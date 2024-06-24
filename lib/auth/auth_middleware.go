package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"monify/lib"
	"net/http"
	"strings"
)

type Middleware struct {
	JwtSecret string
}

func (m Middleware) HttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "" {
			token, err := validateBearerToken(auth, m.JwtSecret)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), lib.UserIdContextKey{}, token)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}

func (m Middleware) GrpcExtractUserId(ctx context.Context, req any, info *grpc.UnaryServerInfo) (uuid.UUID, error) {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return uuid.Nil, nil
	}
	auths := md.Get("authorization")
	if len(auths) == 0 {
		return uuid.Nil, nil
	}
	auth := auths[0]
	token, err := validateBearerToken(auth, m.JwtSecret)
	if err != nil {
		return uuid.Nil, status.Error(codes.Unauthenticated, err.Error())
	}
	return token, nil

}

func validateBearerToken(token string, secret string) (uuid.UUID, error) {
	if strings.HasPrefix(token, "Bearer ") {
		tokenStr := token[7:]
		token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			return uuid.Nil, err
		}
		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
			userId, err := uuid.Parse(claims.Subject)
			if err != nil {
				return uuid.Nil, errors.New("invalid user id in token")
			}
			return userId, nil
		} else {
			switch err {
			case nil:
				return uuid.Nil, errors.New("invalid token")
			default:
				return uuid.Nil, err
			}
		}
	}
	return uuid.Nil, errors.New("invalid token")
}
