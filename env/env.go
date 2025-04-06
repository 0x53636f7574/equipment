package env

import (
	"bufio"
	"os"
	"strings"
)

func GetVar(name string) string {
	return os.Getenv(name)
}

func GetVarOrDefault(name, def string) string {
	value, present := os.LookupEnv(name)

	if !present {
		return def
	}

	return value
}

func SetVar(name, value string) error {
	return os.Setenv(name, value)
}

func SetupEnv(envFilePath string) error {
	envFile, err := os.Open(envFilePath)
	if err != nil {
		return err
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		params := strings.Split(line, "=")
		if len(params) != 2 {
			continue
		}
		_ = SetVar(params[0], params[1])
	}
	return nil
}
