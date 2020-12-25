package main

import "github.com/asatake/trophies-management/api/trophy/infra"

func main() {
	infra.Router.Run()
}
