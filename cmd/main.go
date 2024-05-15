package main

import (
	"monify/internal"
	"monify/internal/infra"
)

func main() {
	s := internal.NewServer(infra.NewProductionConfig())
	s.Start("8080")
}
