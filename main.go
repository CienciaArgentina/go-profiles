package main

import (
	"github.com/CienciaArgentina/go-profiles/cmd"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	cmd.Execute()
}
