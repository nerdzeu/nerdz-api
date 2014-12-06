package nerdz

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Starts the API configuration from the JSON configuration file
// An error is returned if something wrong happens
func ExecuteAPIConfiguration(configurationFilePath string) error {
	if err := InitConfiguration(configurationFilePath); err != nil {
		return err
	}

	localPath, errPath := os.Getwd()
	// get the current working directory
	if errPath != nil {
		return errPath
	}

	// starts the configuration of the dbms
	errDB := startsDBConfiguration()
	if errDB != nil {
		return errDB
	}

	InitDB()

	// get all the test files that will be executed
	testFiles, errTestFile := retrieveTestFile(localPath)
	if errTestFile != nil {
		return errTestFile
	}

	for _, testFile := range testFiles {
		cmd := exec.Command("go", "test", "-compiler", Configuration.Compiler, testFile)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			fmt.Errorf("%s - %s", fmt.Sprint(err), stderr.String())
			return err
		}

		fmt.Println(out.String())

	}

	return nil
}

func retrieveTestFile(localPath string) ([]string, error) {
	testFiles := make([]string, 0, 5)

	var visit filepath.WalkFunc = func(path string, f os.FileInfo, err error) error {
		if !strings.Contains(path, "_test.go") {
			return nil
		}

		testFiles = append(testFiles, path)

		return nil
	}

	filepath.Walk(localPath, visit)

	if len(testFiles) == 0 {
		return nil, errors.New("There aren't test files to execute...")
	}

	return testFiles, nil
}

func startsDBConfiguration() error {
	dbCommand := ""
	switch runtime.GOOS {
	case "linux":
		dbCommand = LinuxDBCommand()
	case "windows":
		// Unable to use WinDBCommand() - wrong definition (?)
		//dbCommand = WinDBCommand()
	}

	parts := strings.Fields(dbCommand)

	return exec.Command(parts[0], parts[1], parts[2], parts[3]).Run()

}
