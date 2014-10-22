package nerdz

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
    "os"
)

type Config struct {
	Username  string
	Password  string // optional -> default:
	DbName    string
	Host      string // optional -> default: localhost
	Port      int16  // optional -> default: 5432
	SSLMode   string // optional -> default: disable
	NERDZPath string
    Languages []string
    Templates map[uint8]string
}

var Configuration *Config

func InitConfiguration(path string) error {
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

	return nil
}

// It returns a valid connection string on success, Error otherwise
func (conf *Config) ConnectionString() (string, error) {
	if Configuration.Username == "" {
		return "", errors.New("Postgresql doesn't support empty username")
	}

	if Configuration.DbName == "" {
		return "", errors.New("Empty DbName field")
	}

	var ret bytes.Buffer
	ret.WriteString("user=" + Configuration.Username + " dbname=" + Configuration.DbName + " host=")

	if Configuration.Host == "" {
		ret.WriteString("localhost")
	} else {
		ret.WriteString(Configuration.Host)
	}

	if Configuration.Password != "" {
		ret.WriteString(" password=" + Configuration.Password)
	}

	ret.WriteString(" sslmode=")

	if Configuration.SSLMode == "" {
		ret.WriteString("disable")
	} else {
		ret.WriteString(Configuration.SSLMode)
	}

	ret.WriteString(" port=")

	if Configuration.Port == 0 {
		ret.WriteString("5432")
	} else {
		ret.WriteString(strconv.Itoa(int(Configuration.Port)))
	}

	return ret.String(), nil
}
