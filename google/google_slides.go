package google

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/rafaelbmateus/slides-gospel/entity"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/slides/v1"
)

func CreatePresentation(apresentation *entity.Presentation, oauth2 *oauth2.Config, token *oauth2.Token) (string, error) {
	slidesService := getService(oauth2.Client(context.Background(), token))
	p := &slides.Presentation{
		Title: apresentation.Title,
	}

	presentation, err := slidesService.Presentations.Create(p).Fields("presentationId").Do()
	if err != nil {
		log.Fatal().Msgf("unable to create presentation. %v", err)
	}
	log.Info().Msgf("created presentation with id: %s", presentation.PresentationId)

	var requests []*slides.Request
	for _, song := range apresentation.Songs {
		requests = append(requests, slideRequest("MAIN_POINT", song.Name)...)
		for _, slide := range song.Content {
			content := strings.Replace(slide, "<br>", "\n", -1)
			requests = append(requests, slideRequest("SECTION_HEADER", content)...)
		}
		requests = append(requests, slideRequest("SECTION_HEADER", "")...)
	}

	for _, slide := range apresentation.Prayer.Content {
		content := strings.Replace(slide, "<br>", "\n", -1)
		requests = append(requests, slideRequest("SECTION_HEADER", content)...)
		requests = append(requests, slideRequest("SECTION_HEADER", "")...)
	}

	body := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}

	response, err := slidesService.Presentations.BatchUpdate(presentation.PresentationId, body).Do()
	if err != nil {
		log.Fatal().Msgf("unable to create slide. %v", err)
	}

	log.Info().Msgf("slides created response: %v", response.PresentationId)

	return presentation.PresentationId, nil
}

func slideRequest(layout, content string) []*slides.Request {
	slideId := uuid.New().String()
	titleId := uuid.New().String()

	return []*slides.Request{
		{
			CreateSlide: &slides.CreateSlideRequest{
				ObjectId: slideId,
				SlideLayoutReference: &slides.LayoutReference{
					PredefinedLayout: layout,
				},
				PlaceholderIdMappings: []*slides.LayoutPlaceholderIdMapping{
					{
						ObjectId: titleId,
						LayoutPlaceholder: &slides.Placeholder{
							Type: "TITLE",
						},
					},
				},
			},
		}, {
			InsertText: &slides.InsertTextRequest{
				ObjectId: titleId,
				Text:     content,
			},
		},
	}
}

func getService(client *http.Client) *slides.Service {
	srv, err := slides.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatal().Msgf("unable to retrieve slides client: %v", err)
	}
	return srv
}
