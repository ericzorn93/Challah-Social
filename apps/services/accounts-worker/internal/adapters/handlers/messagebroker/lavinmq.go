package messagebroker

import (
	"apps/services/accounts-worker/internal/app"
	"context"
	"libs/backend/auth/m2m"
	"libs/backend/boot"
	userEntities "libs/backend/domain/user/entities"
	userValueObjects "libs/backend/domain/user/valueobjects"
	accountseventsv1 "libs/backend/proto-gen/go/accounts/accountsevents/v1"
	"log/slog"

	"google.golang.org/protobuf/proto"
)

// LavinMQHandler handles all incoming events from LavinMQ
type LavinMQHandler struct {
	Logger   boot.Logger
	Consumer boot.AMQPConsumer
	M2M      m2m.M2MGenerator
	App      app.App
}

// LavinMQHandler is the constructor for LavinMQHandler
func NewLavinMQHandler(logger boot.Logger, consumer boot.AMQPConsumer, app app.App) LavinMQHandler {
	return LavinMQHandler{
		Logger:   logger,
		Consumer: consumer,
		App:      app,
	}
}

// HandleUserRegisteredEvent handles the user registered event
func (h LavinMQHandler) HandleUserRegisteredEvent(ctx context.Context, queueName string) error {
	msgs, err := h.Consumer.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		h.Logger.Error("Cannot consume messages", slog.Any("error", err))
		return err
	}

	// Loop through all messages in queue
	for msg := range msgs {
		// Unmarshal the message to userRegisteredEvent
		var userRegisteredEvent accountseventsv1.UserRegistered
		proto.Unmarshal(msg.Body, &userRegisteredEvent)
		h.Logger.Info("User printed", slog.Any("userRegisteredEvent", &userRegisteredEvent))

		// Parse CommonID
		emailAddress := userValueObjects.NewEmailAddress(userRegisteredEvent.EmailAddress)

		// Create User in Accounts API
		user := userEntities.NewUser(
			userEntities.WithEmailAddress(emailAddress),
			userEntities.WithUserUsername(userRegisteredEvent.Username),
		)

		// Create Account in Accounts API
		go func() {
			if err := h.App.AccountService.CreateAccount(ctx, user); err != nil {
				h.Logger.Error("Cannot create account", slog.Any("error", err))
			}
		}()
	}

	return nil
}
