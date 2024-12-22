package main

import (
	//"github.com/vika-ryt/workspace/proect_calculator_go/pkg/calculation"
	"github.com/vika-ryt/project_go_calculator/internal/application"
)

func main() {
	app := application.New()
	// app.Run()
	app.RunServer()
}