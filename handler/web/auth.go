package web

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/rs/zerolog/log"
)

func (me *Web) Login(c *gin.Context) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	if retrievedState != c.Query("state") {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid session state: %s", retrievedState))
		return
	}

	token, err := me.OAuth2.Exchange(c.Request.Context(), c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := me.OAuth2.Client(c.Request.Context(), token)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer email.Body.Close()
	data, err := ioutil.ReadAll(email.Body)
	if err != nil {
		log.Fatal().Msgf("could not read body: %s", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user entity.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Fatal().Msgf("unmarshal userinfo failed: %s", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session.Set("AccessToken", token.AccessToken)
	session.Set("RefreshToken", token.RefreshToken)
	session.Set("TokenType", token.TokenType)
	session.Set("Expiry", token.Expiry.Format(time.RFC3339))
	session.Set("Email", user.Email)
	session.Set("Name", user.Name)
	err = session.Save()
	if err != nil {
		log.Fatal().Msgf("session save failed: %s", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Info().Msgf("user logged with success: %v", user)
	c.Redirect(http.StatusFound, "/")
}

func (me *Web) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func UserToken(c *gin.Context) interface{} {
	session := sessions.Default(c)
	log.Info().Msgf("user token from default session: %v", session.Get("AccessToken"))
	return session.Get("AccessToken")
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
