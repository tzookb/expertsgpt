package main

import (
	"flag"
	"tzook/gogpt/internal/runner"
)

func main() {
	openaiToken := flag.String("openaiToken", "", "open ai token")
	flag.Parse()

	app := runner.NewApp(*openaiToken)
	app.Boom()
}
