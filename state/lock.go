package state

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

const LockVersion = "0.9"

type LockFile struct {
	Version string
	Plugins []string
	Files   []string
}

func (lf *LockFile) Save(path string) (err error) {

	lf.Version = LockVersion
	slices.Sort(lf.Plugins)
	slices.Sort(lf.Files)

	raw, err := json.MarshalIndent(lf, "", "\t")
	if err != nil {
		return fmt.Errorf("failed marshalling lock file: %w", err)
	}
	err = os.WriteFile(path, raw, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed writing lock file: %w", err)
	}
	return nil
}

func LoadLockFile(path string) (lf *LockFile, err error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed reading lock file: %w", err)
	}

	lf = &LockFile{}
	err = json.Unmarshal(raw, lf)
	if err != nil {
		return nil, fmt.Errorf("failed parsing lock file: %w", err)
	}
	return lf, nil
}
