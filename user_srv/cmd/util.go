package cmd

import (
	"fmt"

	"volcano.user_srv/utils"
)

func CheckPath(path string) error {
	if path == "" {
		return fmt.Errorf("the configuration file doesn't specify: %s", path)
	}

	if !utils.PathExists(path) {
		return fmt.Errorf("the configuration file doesn't exist: %s", path)
	}

	
	if !utils.PathExists(fmt.Sprintf("%s/%s", path, mode)) {
		return fmt.Errorf("the configuration file doesn't exist in this mode: %s/%s", path, mode)
	}

	return nil
}