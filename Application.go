package main

import (
	log "github.com/sirupsen/logrus"
	"yellowgreenorgreenyellow/api"
)

var version string = "dev"

func main() {
	log.Info("current Version " + version)
	api.Serve()
}
