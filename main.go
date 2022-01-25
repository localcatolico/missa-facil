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
	memory, _ := memory.NewMemory("data/prayers.json", "data/songs.json")
	handler.NewHandler(usecase.NewUsecase(memory)).Server()
}
