package httpauth

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
)

// Validator ensure handling of access token by service
type Validator interface {
	EnsureValidToken(ctx context.Context, accessToken string) (*clerk.SessionClaims, error)
}
