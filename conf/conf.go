// config project conf.go
package conf

import (
	"encoding/json"
	"io/ioutil"
	"kino/site"
	"log"
)

const ConfFile = "./config.json"

// Sites------------------------------------------------------------------------
type config struct {
	Sites []site
}

func (c config) String() string {
	result := "Config\r\n\tSites"
	for _, v := range c.Sites {
		result += "\r\n\t\t" + v.String()
	}
	return result
}

func Read() *config {

	var result config

	rawDataIn, err := ioutil.ReadFile(ConfFile)
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}

	err = json.Unmarshal(rawDataIn, &result)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}

	return &result
}

func (c config) Write() {
	rawDataOut, err := json.MarshalIndent(&c, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(ConfFile, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}
