package httpauth

const (
	// Used for headers to be sent and recieved via HTTP and ConnectRPC
	AuthorizationHeaderKey = "Authorization"

	// Machine to Machine request header to ensure validation is not required
	M2MHeaderKey = "X-M2M-Token"

	// Machine to Machine request header value
	M2MHeaderValue = "@challah-social-m2m@"
)
