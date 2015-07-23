package nerdz

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	DbUsername string
	DbPassword string // optional -> default:
	DbName     string
	DbHost     string // optional -> default: localhost
	DbPort     int16  // optional -> default: 5432
	DbSSLMode  string // optional -> default: disable
	NERDZPath  string
	NERDZUrl   string
	NERDZURL   *url.URL `json:"-"`
	Languages  []string
	Templates  map[uint8]string
	EnableLog  bool  //optional: default: false
	Port       int16 // API port, optional -> default: 7536
}

var Configuration *Config

// initConfiguration initialize the API parsing the configuration file
func initConfiguration(path string) error {
	log.Println("Parsing JSON config file " + path)

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	Configuration = new(Config)
	if err = json.Unmarshal(contents, Configuration); err != nil {
		return err
	}

	var dirs []os.FileInfo
	if dirs, err = ioutil.ReadDir(Configuration.NERDZPath + "/data/langs/"); err != nil || len(dirs) == 0 {
		return errors.New("Check your NERDZPath: " + Configuration.NERDZPath)
	}

	for _, language := range dirs {
		if language.Name() != "index.html" {
			Configuration.Languages = append(Configuration.Languages, language.Name())
		}
	}

	if dirs, err = ioutil.ReadDir(Configuration.NERDZPath + "/tpl/"); err != nil {
		return err
	}

	Configuration.Templates = make(map[uint8]string)
	for _, tpl := range dirs {
		if tpl.Name() != "index.html" {
			var tplNumber int
			if tplNumber, err = strconv.Atoi(tpl.Name()); err != nil {
				return err
			}

			var byteName []byte
			if byteName, err = ioutil.ReadFile(Configuration.NERDZPath + "/tpl/" + tpl.Name() + "/NAME"); err != nil {
				return err
			}
			Configuration.Templates[uint8(tplNumber)] = string(byteName)
		}
	}

	if Configuration.Port == 0 {
		Configuration.Port = 7536
	}

	if Configuration.NERDZUrl != "" {
		if url, e := url.Parse(Configuration.NERDZUrl); e == nil {
			Configuration.NERDZURL = url
		} else {
			return e
		}
	} else {
		return errors.New("NERDZUrl is a required field")
	}

	return nil
}

// It returns a valid connection string on success, Error otherwise
func (conf *Config) ConnectionString() (string, error) {
	if Configuration.DbUsername == "" {
		return "", errors.New("Postgresql doesn't support empty username")
	}

	if Configuration.DbName == "" {
		return "", errors.New("Empty DbName field")
	}

	var ret bytes.Buffer
	ret.WriteString("user=" + Configuration.DbUsername + " dbname=" + Configuration.DbName + " host=")

	if Configuration.DbHost == "" {
		ret.WriteString("localhost")
	} else {
		ret.WriteString(Configuration.DbHost)
	}

	if Configuration.DbPassword != "" {
		ret.WriteString(" password=" + Configuration.DbPassword)
	}

	ret.WriteString(" sslmode=")

	if Configuration.DbSSLMode == "" {
		ret.WriteString("disable")
	} else {
		ret.WriteString(Configuration.DbSSLMode)
	}

	ret.WriteString(" port=")

	if Configuration.DbPort == 0 {
		ret.WriteString("5432")
	} else {
		ret.WriteString(strconv.Itoa(int(Configuration.DbPort)))
	}

	return ret.String(), nil
}
