package utils

import (
    "encoding/json"
    "log"
    "io/ioutil"
    "errors"
)

type config struct {
    Username string
    Password string
    DbName   string
    SSLMode  string
}

func parse(confPath string) (conf *config, e error) {
    log.Println("Parsing JSON config file " + confPath);

    contents, e := ioutil.ReadFile(confPath)
    if e != nil {
        return nil, e
    }

    //WIP
}
