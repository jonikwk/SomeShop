package configuration

import (
	"encoding/xml"
	"os"

	"github.com/golang/glog"
)

type Configuration struct {
	XMLName  xml.Name `xml:"configuration"`
	Connect  *Connect
	User     *User
	Settings *Settings
}

type Connect struct {
	XMLName xml.Name `xml:"connect"`
	DBHost  string   `xml:"dbhost"`
	DBName  string   `xml:"dbname"`
	DBType  string   `xml:"dbtype"`
	Port    int      `xml:"port"`
}

type User struct {
	XMLName  xml.Name `xml:"user"`
	Login    string   `xml:"login"`
	Password string   `xml:"password"`
}

type Settings struct {
	XMLName       xml.Name `xml:"settings"`
	BotToken      string   `xml:"botToken"`
	UpdateOfSet   int      `xml:"updateOfSet"`
	UpdateTimeout int      `xml:"updateTimeOut"`
	MapsApiKey    string   `xml:"mapsApiKey"`
}

func (this *Configuration) ParseConfigurationFile() {
	file, err := os.Open("config.xml")
	if err != nil {
		glog.Exit(err)
	}
	defer file.Close()
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(this)
	if err != nil {
		glog.Exit(err)
	}
}
