package config

import (
	"encoding/json"
	"os"
	"strconv"
)

type OIDTopicObject struct {
	OID   string `json:"oid"`
	Topic string `json:"topic"`
}

type SNMPEndpointObject struct {
	Endpoint  string           `json:"endpoint"`
	Community string           `json:"community"`
	OIDTopics []OIDTopicObject `json:"oidTopics"`
}

// SNMPConfig basic config
type SNMPMapObject struct {
	SNMPEndpoints []SNMPEndpointObject `json:"snmpEndpoints"`
}

var (
	// SNMPMap is the loaded JSON configuration
	SNMPMap *SNMPMapObject

	// Server is the MQTT server address
	Server string

	// Port is the MQTT server listen port
	Port int

	// ClientID is how the name of the client
	ClientID string

	// TopicPrefix is just that, a prefix for the presented data keys
	TopicPrefix string

	// Interval is the poll interval in seconds
	Interval int
)

// ConnectionString returns the MQTT connection string
func ConnectionString() string {
	return "tcp://" + Server + ":" + strconv.Itoa(Port)
}

// LoadMap loads the file in to the struct
func LoadMap(file string) error {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&SNMPMap)

	if err != nil {
		return err
	}

	return nil
}
