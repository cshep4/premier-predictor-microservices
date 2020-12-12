package integration

import "fmt"

const (
	grpcHostname = "localhost:50051"
	httpHostname = "http://localhost:8080/%s"

	jwtSecret     = "some-jwt-secret"
	password      = "plaintextPassword123"
	createdUserId = "created user id"

	notFoundErrorEmail            = "notfound@test.com"
	getUserErrorEmail             = "getusererror@test.com"
	failUpdatePasswordErrorEmail  = "failupdatepassword@test.com"
	failUpdateSignatureErrorEmail = "failupdatesignature@test.com"
	failCreateErrorEmail          = "failcreate@test.com"
	failEmailErrorEmail           = "failemail@test.com"
	validEmail                    = "valid@test.com"
)

type (
	errorResponse struct {
		Message string `json:"message"`
	}

	loginResponse struct {
		Id    string `json:"id"`
		Token string `json:"token"`
	}

	registerResponse struct {
		Id    string `json:"id"`
		Token string `json:"token"`
	}
)

func buildUrl(path string) string {
	return fmt.Sprintf(httpHostname, path)
}
