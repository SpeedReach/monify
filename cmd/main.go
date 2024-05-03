package main

import "monify/internal"

func main() {
	s := internal.NewProduction()
	s.Start("8080")
}
