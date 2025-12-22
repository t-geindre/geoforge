package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(zerolog.NewConsoleWriter())

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal().Err(err).Msg("Game crashed")
	}

}
