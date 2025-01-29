package httpauth

import (
	"context"
	"errors"
	"libs/backend/boot"
	"strings"

	"connectrpc.com/connect"
)

const (
	tokenValueLength = 2
)

type AuthInerceptor struct {
	logger               boot.Logger
	accessTokenValidator Validator
}

// NewAuthInterceptor will intercept connectRPC requests and handle authentication
func NewAuthInterceptor(logger boot.Logger) AuthInerceptor {
	return AuthInerceptor{logger: logger, accessTokenValidator: NewAccessTokenValidator(logger)}
}

// Incoming will handle auth for incoming server requests
func (a AuthInerceptor) Incoming() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			m2mTokenHeader := req.Header().Get(M2MHeaderKey)
			authTokenHeaderValue := req.Header().Get(AuthorizationHeaderKey)

			// Check if M2M token is present and don't authenticate
			if m2mTokenHeader == M2MHeaderValue {
				return next(ctx, req)
			}

			// Check token in handlers.
			if authTokenHeaderValue == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}

			// Validate token length
			authTokenValues := strings.Split(authTokenHeaderValue, " ")

			if len(authTokenValues) != tokenValueLength {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("token invalid"),
				)
			}

			// Validate incoming token
			accessToken := authTokenValues[1]
			claims, err := a.accessTokenValidator.EnsureValidToken(ctx, accessToken)

			// Validate claims
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("token claims invalid"),
				)
			}

			// Set the validated custom claims to the context
			ctx = SetClaimsToContext(ctx, claims)

			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
