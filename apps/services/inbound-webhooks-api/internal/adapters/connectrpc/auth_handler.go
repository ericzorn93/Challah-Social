package connectrpc

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"

	"apps/services/inbound-webhooks-api/internal/app"
	boot "libs/backend/boot"
	userEntities "libs/backend/domain/user/entities"
	userValueObjects "libs/backend/domain/user/valueobjects"
	commonv1 "libs/backend/proto-gen/go/common/v1"
	inboundwebhooksapiv1 "libs/backend/proto-gen/go/webhooks/inboundwebhooksapi/v1"
	inboundwebhooksapiv1connect "libs/backend/proto-gen/go/webhooks/inboundwebhooksapi/v1/inboundwebhooksapiv1connect"
)

// AuthHandler handles all gRPC endpoints for inbound webhooks
type AuthHandler struct {
	inboundwebhooksapiv1connect.UnimplementedInboundWebhooksAuthServiceHandler
	Logger      boot.Logger
	Application app.Application
}

// NewAuthHandler will return a pointer to the inbound webhooks API server
func NewAuthHandler(logger boot.Logger, application app.Application) *AuthHandler {
	return &AuthHandler{
		Logger:      logger,
		Application: application,
	}
}

// ClerkAuthUserEvent handles incoming Webhooks from Auth0 and will attach the message
// to an exchange within the message broker
func (h *AuthHandler) ClerkAuthUserEvent(
	ctx context.Context,
	req *connect.Request[inboundwebhooksapiv1.ClerkUserAuthEventRequest],
) (*connect.Response[commonv1.Empty], error) {
	eventType := req.Msg.GetType()

	switch eventType {
	case inboundwebhooksapiv1.ClerkUserEventType_CLERK_USER_EVENT_TYPE_USER_CREATED:
		data := req.Msg.GetData()
		h.Logger.Info("User printed", slog.Any("user", data))

		// Parse Info
		emailAddresses := data.GetEmailAddresses()
		if len(emailAddresses) == 0 {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("no email addresses provided"))
		}
		emailAddress := userValueObjects.NewEmailAddress(emailAddresses[0].EmailAddress)
		userName := data.GetUsername()

		if err := h.Application.AuthService.RegisterUser(
			userEntities.NewUser(
				userEntities.WithEmailAddress(emailAddress),
				userEntities.WithUserUsername(userName),
				userEntities.WithMetadata(make(map[string]any, 0)),
			),
		); err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		return connect.NewResponse(&commonv1.Empty{}), nil
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid event type"))
	}
}
