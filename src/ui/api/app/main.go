package main

import (
	"log"
	"rms-backend/src/core/utils"
	"rms-backend/src/ui/api"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	godotenv.Load(".env-for-local")
	api := api.NewAPI(getServerHostAndPort())
	api.Serve()
}

func getServerHostAndPort() (string, int) {
	host := utils.GetenvWithDefault("SERVER_HOST", "0.0.0.0")
	portStr := utils.GetenvWithDefault("SERVER_PORT", "8000")
	var port int
	if v, err := strconv.Atoi(portStr); err != nil {
		log.Fatal("The server port env variable must be a number (e.g 8000)")
	} else {
		port = v
	}
	return host, port
}
