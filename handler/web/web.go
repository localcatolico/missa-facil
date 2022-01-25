package web

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/slides-gospel/usecase"
	"golang.org/x/oauth2"
)

type Web struct {
	Router  *gin.Engine
	Usecase *usecase.Usecase
	OAuth2  *oauth2.Config
}

func NewWeb(router *gin.Engine, usecase *usecase.Usecase, oauth2 *oauth2.Config) *Web {
	return &Web{
		Router:  router,
		Usecase: usecase,
		OAuth2:  oauth2,
	}
}

func (me *Web) Home(c *gin.Context) {
	state := randToken()
	session := sessions.Default(c)
	email := session.Get("Email")
	if email == nil {
		session.Set("state", state)
		session.Save()
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":    "Página Principal",
		"loginURL": me.OAuth2.AuthCodeURL(state),
		"email":    email,
		"prayers":  me.Usecase.GetPrayers(),
		"songs":    me.Usecase.GetSongs(),
	})
}

func (me *Web) Help(c *gin.Context) {
	c.HTML(http.StatusOK, "help.tmpl", gin.H{
		"title": "Ajuda",
	})
}

func (me *Web) Done(c *gin.Context) {
	c.HTML(http.StatusOK, "done.tmpl", gin.H{
		"title": "Apresentação Finalizada",
	})
}
