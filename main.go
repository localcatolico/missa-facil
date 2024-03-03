package main

import (
	"github.com/rafaelbmateus/slides-gospel/database/memory"
	"github.com/rafaelbmateus/slides-gospel/handler"
	"github.com/rafaelbmateus/slides-gospel/usecase"
	"github.com/rs/zerolog/log"
)

var (
	Version = "no version provided"
	Commit  = "no commit hash provided"
)

func main() {
	log.Info().Msgf("starting using version %s and commit %s", Version, Commit)
	prayersURL := "https://raw.githubusercontent.com/localcatolico/localcatolico-oracoes/main/prayers.json"
	musicsURL := "https://raw.githubusercontent.com/localcatolico/localcatolico-musicas/main/data/musics.json"
	memory, err := memory.NewMemory(prayersURL, musicsURL)
	if err != nil {
		log.Error().Msgf("error on create memory %q", err)

		return
	}

	handler.NewHandler(usecase.NewUsecase(memory)).Server()
}
