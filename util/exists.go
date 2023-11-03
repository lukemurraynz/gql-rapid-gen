package util

import "os"

func FileExists(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	if s.IsDir() {
		return false
	}
	return true
}
