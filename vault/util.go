package vault

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetDefaultToken() (string, error) {
	if env := os.Getenv("VAULT_TOKEN"); env != "" {
		return env, nil
	}

	rawToken, err := ioutil.ReadFile(fmt.Sprintf("%s/.vault-token", os.Getenv("HOME")))
	if err != nil {
		return "", err
	}

	return string(rawToken), nil
}

func GetDefaultAddr() string {
	return os.Getenv("VAULT_ADDR")
}
