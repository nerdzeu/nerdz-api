package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
)

type config struct {
	Username string
	Password string // optional -> default:
	DbName   string
	Host     string // optional -> default: localhost
	Port     int16  // optional -> default: 5432
	SSLMode  string // optional -> default: disable
}

func Parse(confPath string) (string, error) {
	log.Println("Parsing JSON config file " + confPath)

	contents, e := ioutil.ReadFile(confPath)
	if e != nil {
		return "", e
	}

	var conf config
	e = json.Unmarshal(contents, &conf)

	if e != nil {
		return "", e
	}

	if conf.Username == "" {
		return "", errors.New("Postgresql doesn't support empty username")
	}

	if conf.DbName == "" {
		return "", errors.New("Empty DbName field in " + confPath)
	}

	var ret bytes.Buffer

	ret.WriteString("user=" + conf.Username + " dbname=" + conf.DbName + " host=")

	if conf.Host == "" {
		ret.WriteString("localhost")
	} else {
		ret.WriteString(conf.Host)
	}

	if conf.Password != "" {
		ret.WriteString(" password=" + conf.Password)
	}

	ret.WriteString(" sslmode=")

	if conf.SSLMode == "" {
		ret.WriteString("disable")
	} else {
		ret.WriteString(conf.SSLMode)
	}

	ret.WriteString(" port=")

	if conf.Port == 0 {
		ret.WriteString("5432")
	} else {
		ret.WriteString(strconv.Itoa(int(conf.Port)))
	}

	return ret.String(), nil
}
