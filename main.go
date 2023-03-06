package main

import "shop-app-API/app"

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
