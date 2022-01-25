package handler

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/slides-gospel/google"
	"github.com/rafaelbmateus/slides-gospel/handler/api"
	"github.com/rafaelbmateus/slides-gospel/handler/web"
	"github.com/rafaelbmateus/slides-gospel/usecase"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

type Handler struct {
	Router  *gin.Engine
	Usecase *usecase.Usecase
	OAuth2  *oauth2.Config
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")
	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/presentations",
	}

	crendential := google.NewCredentials(clientID, clientSecret, redirectURL, scopes)

	router := gin.Default()
	oauth2 := crendential.OAuth([]byte("c00ki3-s3cr3t"))
	router.Use(google.Session("mysession"))
	return &Handler{
		Router:  router,
		Usecase: usecase,
		OAuth2:  oauth2,
	}
}

func (h *Handler) Server() {
	web.NewWeb(h.Router, h.Usecase, h.OAuth2).Routes()
	api.NewAPI(h.Router, h.Usecase, h.OAuth2).Routes()

	if err := h.Router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal().Msg("error on run server")
	}
}
