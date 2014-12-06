package nerdz

import "path/filepath"

const winDBFileDefinition string = "initdb.bat"

func WinDBCommand() string {
	return Configuration.NERDZDBPath + string(filepath.Separator) +
		linuxDBFileDefinition + " " +
		Configuration.Username + " " + Configuration.DbName + " " + Configuration.Password
}
