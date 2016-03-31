/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
	"strings"
)

// Config represents the configuration file structure
type Config struct {
	DbUsername string
	DbPassword string // optional -> default:
	DbName     string
	DbHost     string // optional -> default: localhost
	DbPort     int16  // optional -> default: 5432
	DbSSLMode  string // optional -> default: disable
	NERDZPath  string
	NERDZHost  string
	Languages  []string
	Scopes     []string
	Templates  map[uint8]string
	EnableLog  bool  //optional: default: false
	Port       int16 // API port, optional -> default: 7536
	Host       string
	Scheme     string
}

// Configuration represent the parsed configuration file
var Configuration *Config

var scopes = []string{
	"profile",
	"projects",
	"pms",
	"notifications",
	"messages",         // implies profile_messages and project_messages
	"profile_messages", // implies "profile_comments"
	"project_messages", // implies "project_comments"
	"followers",
	"following",
	"friends",
	"profile_comments",
	"project_comments",
	"base", // read only access to every scope above
}

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

	Configuration.Scopes = scopes

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

	if Configuration.NERDZHost != "" {
		if _, e := url.Parse(Configuration.NERDZHost); e != nil {
			return e
		}
	} else {
		return errors.New("NERDZHost is a required field")
	}

	if Configuration.Host != "" {
		if _, e := url.Parse(Configuration.Host); e != nil {
			return e
		}
	} else {
		return errors.New("Host is a required field")
	}

	if !strings.HasPrefix(Configuration.Scheme, "http") {
		return errors.New("Scheme shoud be http or https only. Https is mandatory in production environment")
	}

	return nil
}

// APIURL returns the the API host:port URL
func (conf *Config) APIURL() *url.URL {
	host := Configuration.Host
	if Configuration.Port != 80 && Configuration.Port != 443 {
		host += ":" + strconv.Itoa(int(Configuration.Port))
	}
	return &url.URL{
		Scheme: Configuration.Scheme,
		Host:   host,
	}
}

// ConnectionString returns a valid connection string on success, Error otherwise
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
