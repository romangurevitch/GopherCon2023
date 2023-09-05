package main

import (
	"github.com/romangurevitch/gophercon2023/internal/challenge/implme/intermediate/poke/app"
	"github.com/romangurevitch/gophercon2023/internal/challenge/implme/intermediate/poke/client"
)

func main() {
	pokeAPP := app.NewPokeApp(client.New())
	pokeAPP.Start()
}
