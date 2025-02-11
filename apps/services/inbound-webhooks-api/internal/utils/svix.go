package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	svix "github.com/svix/svix-webhooks/go"
)

// ValidateClerkAuthWebhook validates the headers of a webhook request
// coming from Clerk Auth
func ValidateClerkAuthWebhook(headers http.Header, body []byte) error {
	removedBodyMap := make(map[string]interface{})
	json.Unmarshal(body, &removedBodyMap)

	removedBodyData, err := json.Marshal(removedBodyMap)
	if err != nil {
		return fmt.Errorf("error marshalling clerk auth webhook data: %w", err)
	}

	signingSecret, found := os.LookupEnv("CLERK_WEBHOOK_SIGNING_SECRET")
	if !found {
		return fmt.Errorf("error getting webhook signing secret")
	}

	wh, err := svix.NewWebhook(signingSecret)
	if err != nil {
		return fmt.Errorf("error creating webhook: %w", err)
	}

	if err := wh.Verify(removedBodyData, headers); err != nil {
		return fmt.Errorf("error verifying webhook: %w", err)
	}

	return nil
}
