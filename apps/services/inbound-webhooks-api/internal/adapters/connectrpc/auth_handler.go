package connectrpc

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"apps/services/inbound-webhooks-api/internal/app"
	"apps/services/inbound-webhooks-api/internal/utils"
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
	// Validate the Clerk Webhook
	jsonBody, err := h.marshalProtobufMessageToJSON(req.Msg)
	if err != nil {
		h.Logger.Error("Error marshaling protobuf message to JSON", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	if err := utils.ValidateClerkAuthWebhook(req.Header(), jsonBody); err != nil {
		h.Logger.Error("Error validating clerk auth user.created webhook", slog.Any("error", err))
		// TODO: Add back when auth is implemented properly with tghe correct request body
		// return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// Convert event type to enum value
	finalEventType := h.convertClerkAuthUserEventTypeToEnum(req.Msg.GetType())

	switch finalEventType {
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
		clerkUserId := data.GetId()

		if err := h.Application.AuthService.RegisterUser(
			userEntities.NewUser(
				userEntities.WithClerkUserID(clerkUserId),
				userEntities.WithUserFirstName(data.GetFirstName()),
				userEntities.WithUserLastName(data.GetLastName()),
				userEntities.WithEmailAddress(emailAddress),
				userEntities.WithUserUsername(userName),
				userEntities.WithUserGender(data.GetGender()),
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

// convertClerkAuthUserEventTypeToEnum will convert the user event type to an enum value for the ingestion of
// the clerk user events
func (*AuthHandler) convertClerkAuthUserEventTypeToEnum(eventType string) inboundwebhooksapiv1.ClerkUserEventType {
	const (
		userCreatedEventType string = "user.created"
	)

	var finalEventType inboundwebhooksapiv1.ClerkUserEventType
	switch eventType {
	case userCreatedEventType:
		finalEventType = inboundwebhooksapiv1.ClerkUserEventType_CLERK_USER_EVENT_TYPE_USER_CREATED
	default:
		finalEventType = inboundwebhooksapiv1.ClerkUserEventType_CLERK_USER_EVENT_TYPE_UNKNOWN
	}

	return finalEventType
}

// marshalProtobufMessageToJSON will marshal a protobuf message to a byte slice for JSON
func (h AuthHandler) marshalProtobufMessageToJSON(msg proto.Message) ([]byte, error) {
	marshaler := protojson.MarshalOptions{
		UseProtoNames:   false,
		EmitUnpopulated: true,
		Indent:          "  ",
	}

	return marshaler.Marshal(msg)
}
