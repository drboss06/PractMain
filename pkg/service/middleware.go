package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Service) AuthInterceptor(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "token is not provided")
	}
	//VerifyToken(token[0])
	userID, err := s.Authorization.ParseToken(token[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	ctx = context.WithValue(ctx, "user_id", userID)

	return ctx, nil
}

func (s *Service) AuthUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		token, err := s.Authorization.GenerateToken("test", "test")
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

//func (s *Service) userIdentity(c *context.Context) {
//	header := c.GetHeader(authorizationHeader)
//	if header == "" {
//		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
//		return
//	}
//
//	headerParts := strings.Split(header, " ")
//	if len(headerParts) != 2 {
//		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
//		return
//	}
//
//	userId, err := h.services.Authorization.ParseToken(headerParts[1])
//	if err != nil {
//		newErrorResponse(c, http.StatusUnauthorized, err.Error())
//	}
//	c.Set("userId", userId)
//}
