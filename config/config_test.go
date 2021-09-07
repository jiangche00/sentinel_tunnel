package config

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestFetchTunnelConfigFromConsul(t *testing.T) {

	TestConf = new(TestConfig)
	TestConf.Databases = make([]TestDatabase, 0)
	TestConf.SentinelsAddressList = make([]string, 0)
	TestConsulEndpoint = "http://consul-server.kube-system:8500"

	resp, err := http.Get("http://consul-server.kube-system:8500/v1/kv/config/sentinel_tunnel?raw")
	if err != nil {
		t.Logf("\ncan not fetch tunnel configuration from consul, err: %s\n", err.Error())
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(TestConf)
	if err != nil {
		t.Logf("\nunable to decode configuration, detail: %s\n", err.Error())
	}
	t.Logf("\nTestConf: %v\n", TestConf)
}
