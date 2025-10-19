package main

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/app"
)

func main() {
	// TODO: Will add config here
	router := app.SetupRouter()
	router.Run()
}
