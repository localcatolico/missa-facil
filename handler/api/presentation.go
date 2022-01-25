package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/slides-gospel/entity"
	"golang.org/x/oauth2"
)

type presentationRequest struct {
	Title    string   `json:"title"`
	PrayerID string   `json:"prayer"`
	SongsID  []string `json:"songs"`
}

func (me *API) CreatePresentation(c *gin.Context) {
	var p presentationRequest
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error on bind json - " + err.Error()})
		return
	}

	var songs []entity.Song
	for _, s := range p.SongsID {
		if s != "" {
			song, err := me.Usecase.GetSong(s)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error on search song " + s + " - " + err.Error()})
				return
			}
			songs = append(songs, song)
		}
	}

	prayer, err := me.Usecase.GetPrayer(p.PrayerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error on search prayer - " + err.Error()})
		return
	}

	token := new(oauth2.Token)
	session := sessions.Default(c)
	token.AccessToken = session.Get("AccessToken").(string)
	token.RefreshToken = session.Get("RefreshToken").(string)
	token.RefreshToken = session.Get("RefreshToken").(string)
	t := session.Get("Expiry").(string)
	token.Expiry, _ = time.Parse(time.RFC3339, t)
	token.TokenType = session.Get("TokenType").(string)

	presentation, err := me.Usecase.CreatePresentation(me.OAuth2, token, p.Title, prayer, songs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error on create apresentation - " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, presentation)
}
