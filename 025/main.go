package main

import "github.com/narie-monarie/config"

func main() {
	err := config.Migrate()
	if err != nil {
		panic(err.Error())
	}
	err = config.Connect()
	if err != nil {
		panic(err.Error())
	}
}
