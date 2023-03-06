package main

import (
	"erp-web/initialize"
)

func main() {
	initialize.InitLogger()
	initialize.Initdb()
	g := initialize.Routers()
	g.Run("127.0.0.1:8094")
}
