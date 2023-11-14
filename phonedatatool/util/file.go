package util

import (
	"fmt"
	"os"
)

func AssureFileNotExist(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return fmt.Errorf("check file existence %v failed: %v", path, err)
		}
	} else {
		return fmt.Errorf("file %v already exists", path)
	}
}

func AssureAllFileNotExist(paths ...string) error {
	for _, p := range paths {
		if err := AssureFileNotExist(p); err != nil {
			return err
		}
	}
	return nil
}
