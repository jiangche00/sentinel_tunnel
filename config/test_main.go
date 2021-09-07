package config

import (
	"os"
	"testing"
)

type TestConfig struct {
	SentinelsAddressList []string       `json:"Sentinels_addresses_list"`
	Databases            []TestDatabase `json:"Databases"`
}

type TestDatabase struct {
	Name      string `json:"Name"`
	LocalPort string `json:"Local_Port"`
}

var (
	TestConsulEndpoint string
	TestConf           *TestConfig
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
