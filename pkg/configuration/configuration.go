package configuration

import (
	"encoding/xml"
	"os"

	"github.com/golang/glog"
)

//Configuration -
type Configuration struct {
	XMLName  xml.Name `xml:"configuration"`
	Connect  *Connect
	User     *User
	Settings *Settings
}

//Connect -
type Connect struct {
	XMLName xml.Name `xml:"connect"`
	DBHost  string   `xml:"dbhost"`
	DBName  string   `xml:"dbname"`
	DBType  string   `xml:"dbtype"`
	Port    int      `xml:"port"`
}

//User -
type User struct {
	XMLName  xml.Name `xml:"user"`
	Login    string   `xml:"login"`
	Password string   `xml:"password"`
}

//Settings -
type Settings struct {
	XMLName       xml.Name `xml:"settings"`
	BotToken      string   `xml:"botToken"`
	UpdateOfSet   int      `xml:"updateOfSet"`
	UpdateTimeout int      `xml:"updateTimeOut"`
	MapsAPIKey    string   `xml:"mapsAPIKey"`
}

//ParseConfigurationFile -
func (c *Configuration) ParseConfigurationFile() {
	file, err := os.Open("config.xml")
	if err != nil {
		glog.Exit(err)
	}
	defer file.Close()
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		glog.Exit(err)
	}
}
