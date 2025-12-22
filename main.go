package main

import (
	"fmt"
	"net/http"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog"

	_ "net/http/pprof"
)

func main() {
	log := zerolog.New(zerolog.NewConsoleWriter())

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal().Err(err).Msg("Game crashed")
	}

}

// pprofPort override it at compilation with -ldflags "-X main.pprofPort=XXXX"
var pprofPort = "6060"

func init() {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("localhost:%s", pprofPort), nil)
		if err != nil {
			panic(err)
		}
	}()
}
