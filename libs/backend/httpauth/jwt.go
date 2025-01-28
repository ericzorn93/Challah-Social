package httpauth

import (
	"context"
	"libs/backend/boot"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
)

// AccessTokenValidator will implement the auth.Validator interface
// for access token validation
type AccessTokenValidator struct {
	logger boot.Logger
}

// NewAccessTokenValidator handles validator construction
func NewAccessTokenValidator(logger boot.Logger) AccessTokenValidator {
	return AccessTokenValidator{logger}
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func (v AccessTokenValidator) EnsureValidToken(ctx context.Context, accessToken string) (*clerk.SessionClaims, error) {
	return jwt.Verify(ctx, &jwt.VerifyParams{
		Token: accessToken,
	})
}
