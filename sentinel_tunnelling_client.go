package main

import (
	"sentinel_tunnel/config"
	"sentinel_tunnel/st_logger"
	"sentinel_tunnel/st_sentinel_connection"
	"time"
)

func NewSentinelTunnellingClient() *config.SentinelTunnellingClient {

	Tunnelling_client := config.SentinelTunnellingClient{}
	Tunnelling_client.Configuration = *config.Conf

	var err error

	Tunnelling_client.Sentinel_connection, err =
		st_sentinel_connection.NewSentinelConnection(Tunnelling_client.Configuration.Sentinels_addresses_list)
	if err != nil {
		st_logger.WriteLogMessage(st_logger.FATAL, "an error has occur, ",
			err.Error())
	}

	st_logger.WriteLogMessage(st_logger.INFO, "done initializing Tunnelling")

	return &Tunnelling_client
}

func main() {
	config.InitConfig()
	st_logger.InitializeLogger()
	st_client := NewSentinelTunnellingClient()
	st_client.Start()
	for {
		time.Sleep(1000 * time.Millisecond)
	}
}
