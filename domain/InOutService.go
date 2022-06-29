package domain

type GrandType string

const (
	REFRESH_TOKEN      GrandType = "refresh token"
	AUTHORIZATION_CODE GrandType = "authorization code"
	CLIENT_CREDENTIAL  GrandType = "client credentials"
)

type InputServicesDTO struct {
	GrandType GrandType
}

type OutputServicesDTO struct {
	ResponseServices string
}
