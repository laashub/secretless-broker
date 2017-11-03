package authenticator

import (
  "encoding/base64"
  "fmt"
  "net/http"
  "os"

  "github.com/kgilpin/secretless/config"
  "github.com/kgilpin/secretless/conjur"
)

// TODO: cleanup redundancy with pg/backend
var HostUsername = os.Getenv("CONJUR_AUTHN_LOGIN")
var HostAPIKey = os.Getenv("CONJUR_AUTHN_apiKey")

type ConjurAuthenticator struct {
  Config config.ListenerConfig
}

func (self ConjurAuthenticator) Authenticate(values map[string]string, r *http.Request) error {
	username := values["username"]
	if username == "" {
		return fmt.Errorf("Conjur connection parameter 'username' is not available")
	}
	apiKey := values["api_key"]
	if apiKey == "" {
		return fmt.Errorf("Conjur connection parameter 'api_key' is not available")
	}

  if token, err := conjur.Authenticate(username, apiKey); err != nil {
    return err
  } else {
	  r.Header.Set("Authorization", fmt.Sprintf("Token token=\"%s\"", base64.StdEncoding.EncodeToString([]byte(token.Token))))
	  return nil
  }
}