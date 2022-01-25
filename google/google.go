// Package google provides you access to Google's OAuth2
// infrastructure. The implementation is based on this blog post:
// http://skarlso.github.io/2016/06/12/google-signin-with-go/
package google

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var store sessions.CookieStore

// Credentials stores google client-ids.
type Credentials struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

func NewCredentials(clientID, secret, redirectURL string, scopes []string) *Credentials {
	return &Credentials{
		ClientID:     clientID,
		ClientSecret: secret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
	}
}

// OAuth the authorization path.
func (c Credentials) OAuth(cookieSecret []byte) *oauth2.Config {
	store = sessions.NewCookieStore(cookieSecret)

	return &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Scopes:       c.Scopes,
		Endpoint:     google.Endpoint,
	}
}

func Session(name string) gin.HandlerFunc {
	return sessions.Sessions(name, store)
}
