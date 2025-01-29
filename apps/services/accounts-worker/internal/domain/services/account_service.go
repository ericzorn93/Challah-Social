package services

import (
	"apps/services/accounts-worker/internal/app/ports"
	"context"
	"libs/backend/auth/m2m"
	boot "libs/backend/boot"
	userEntities "libs/backend/domain/user/entities"
	"libs/backend/httpauth"
	accountsapiv1 "libs/backend/proto-gen/go/accounts/accountsapi/v1"
	"libs/backend/proto-gen/go/accounts/accountsapi/v1/accountsapiv1connect"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
)

// AccountService handles generic interactions
type AccountService struct {
	Logger                    boot.Logger
	AccountConsumer           boot.AMQPConsumer
	RegistrationServiceClient accountsapiv1connect.AccountServiceClient
	M2MClient                 m2m.M2MGenerator
}

// AccountServiceParams is a struct to hold the parameters for the AccountService
type AccountServiceParams struct {
	Logger          boot.Logger
	AccountConsumer boot.AMQPConsumer
	AccountsAPIURI  string
}

// NewAccountService will construct the auth service
func NewAccountService(params AccountServiceParams) AccountService {
	registrationServiceClient := accountsapiv1connect.NewAccountServiceClient(http.DefaultClient, params.AccountsAPIURI)

	return AccountService{
		Logger:                    params.Logger,
		AccountConsumer:           params.AccountConsumer,
		RegistrationServiceClient: registrationServiceClient,
	}
}

// RegisterUser is an application interface method to handle user registration
// webhooks
func (s AccountService) CreateAccount(ctx context.Context, user userEntities.User) error {
	// Call the accounts-api to create the account
	req := connect.NewRequest(&accountsapiv1.CreateAccountRequest{
		ClerkUserId:  user.ClerkUserID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		Gender:       user.Gender,
		EmailAddress: user.EmailAddress.String(),
	})
	req.Header().Add(httpauth.M2MHeaderKey, httpauth.M2MHeaderValue)
	account, err := s.RegistrationServiceClient.CreateAccount(ctx, req)

	if err != nil {
		s.Logger.Error("Cannot create account in Accounts API", slog.Any("error", err))
		return ports.ErrUserNotCreated
	}

	s.Logger.Info("Account created in Accounts API", slog.Any("isSuccess", account.Msg.GetIsSuccess()))
	return nil
}
