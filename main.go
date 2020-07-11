package main

import (
	"github.com/CienciaArgentina/go-profiles/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	cmd.Execute()
}
