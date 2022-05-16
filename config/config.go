package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sentinel_tunnel/st_logger"
	"sentinel_tunnel/st_sentinel_connection"

	flag "github.com/spf13/pflag"
)

type get_db_address_by_name_function func(db_name string) (string, error)

var (
	VERSION        string
	ConsulEndpoint string
	Conf           *SentinelTunnellingConfiguration
)

type SentinelTunnellingDbConfig struct {
	Name       string `json:"Name"`
	Local_port string `json:"Local_Port"`
}

type SentinelTunnellingConfiguration struct {
	Sentinels_addresses_list []string                     `json:"Sentinels_addresses_list"`
	Databases                []SentinelTunnellingDbConfig `json:"Databases"`
}

type SentinelTunnellingClient struct {
	Configuration       SentinelTunnellingConfiguration
	Sentinel_connection *st_sentinel_connection.Sentinel_connection
}

func InitConfig() {
	Conf = new(SentinelTunnellingConfiguration)
	Conf.Databases = make([]SentinelTunnellingDbConfig, 0)
	Conf.Sentinels_addresses_list = make([]string, 0)
	flag.StringVar(&ConsulEndpoint, "consul-endpoint", "http://consul-server.kube-system:8500",
		"default consul server endpoint: http://consul-server.kube-system:8500")
	FetchTunnelConfigFromConsul(Conf)
	flag.Parse()
}

func FetchTunnelConfigFromConsul(conf *SentinelTunnellingConfiguration) {
	resp, err := http.Get(ConsulEndpoint + "/v1/kv/config/sentinel_tunnel?raw")
	if err != nil {
		log.Fatalf("can not fetch tunnel configuration from consul, err: %s\n", err.Error())
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(conf)
	if err != nil {
		log.Fatalf("unable to decode configuration, detail: %s\n", err.Error())
	}
}

func createTunnelling(conn1 net.Conn, conn2 net.Conn) {
	io.Copy(conn1, conn2)
	conn1.Close()
	conn2.Close()
}

func handleConnection(c net.Conn, db_name string,
	get_db_address_by_name get_db_address_by_name_function) {
	db_address, err := get_db_address_by_name(db_name)
	if err != nil {
		st_logger.WriteLogMessage(st_logger.ERROR, "cannot get db address for ", db_name,
			",", err.Error())
		c.Close()
		return
	}
	db_conn, err := net.Dial("tcp", db_address)
	if err != nil {
		st_logger.WriteLogMessage(st_logger.ERROR, "cannot connect to db ", db_name,
			",", err.Error())
		c.Close()
		return
	}
	go createTunnelling(c, db_conn)
	go createTunnelling(db_conn, c)
}

func handleSigleDbConnections(listening_port string, db_name string,
	get_db_address_by_name get_db_address_by_name_function) {

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", listening_port))
	if err != nil {
		st_logger.WriteLogMessage(st_logger.FATAL, "cannot listen to port ",
			listening_port, err.Error())
	}

	st_logger.WriteLogMessage(st_logger.INFO, "listening on port ", listening_port,
		" for connections to database: ", db_name)
	for {
		conn, err := listener.Accept()
		if err != nil {
			st_logger.WriteLogMessage(st_logger.FATAL, "cannot accept connections on port ",
				listening_port, err.Error())
		}
		st_logger.WriteLogMessage(st_logger.INFO, fmt.Sprintf("client %v connected", conn.RemoteAddr().String()))
		go handleConnection(conn, db_name, get_db_address_by_name)
	}
}

func (st_client *SentinelTunnellingClient) Start() {
	for _, db_conf := range st_client.Configuration.Databases {
		go handleSigleDbConnections(db_conf.Local_port, db_conf.Name,
			st_client.Sentinel_connection.GetAddressByDbName)
	}
}
