package nerdz

import "path/filepath"

const linuxDBFileDefinition string = "initdb.sh"

func LinuxDBCommand() string {
	return "sh " + Configuration.NERDZDBPath + string(filepath.Separator) +
		linuxDBFileDefinition + " " +
		Configuration.Username + " " + Configuration.DbName + " " + Configuration.Password
}
