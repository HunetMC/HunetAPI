package database

import (
	"encoding/json"
    "io/ioutil"
)

// Conf Constructor for config.json
type Conf struct {
	MySQL string `json:"mysql"`
}

// Reading configuration
func readConf() (*Conf, error) {
	const file = "config.json"
	conf := new(Conf) // Initialize constructor
	
	// Read config.json
	cValue, err := ioutil.ReadFile(file)
	if err != nil {
		return conf, err
	}
	
	// Decode json and mapping to constructor
	err = json.Unmarshal([]byte(cValue), conf)
	if err != nil {
		return conf, err
	}
	
	return conf, err
}

func GetSQLURI() string {
	conf, _ := readConf()
	return conf.MySQL
}