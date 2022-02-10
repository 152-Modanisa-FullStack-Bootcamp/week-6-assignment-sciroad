package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// "initialBalanceAmount": 0,
//     "minumumBalanceAmount": -100

type Config struct {
	InitialBalanceAmount float64 `json:"initialBalanceAmount"`
	MinumumBalanceAmount float64 `json:"minumumBalanceAmount"`
}

var config = &Config{}

func init() {
	file, err := os.Open(".config/" + env + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reader, config)
	if err != nil {
		panic(err)
	}

}
func Get() *Config {
	return config
}
