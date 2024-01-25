package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//go:generate go run github.com/mjdrgn/gql-rapid-gen generate
//go:generate pnpm run generate
//go:generate go run .

func main() {
	backendDirs, err := ListSubdirectories("./backend/")
	if err != nil {
		panic(fmt.Errorf("failed reading backend dirs: %w", err))
	}
	for _, d := range backendDirs {
		lambdaPath := filepath.Join("./backend/", d, "lambda")
		if DirExists(lambdaPath) {
			err := UpdateLambdas(lambdaPath)
			if err != nil {
				panic(fmt.Errorf("update lambdas %s: %w", d, err))
			}
		}

		err = CmdInDir("./backend/"+d+"/", "terraform", "fmt")
		if err != nil {
			panic(fmt.Errorf("terraform fmt backend/%s: %w", d, err))
		}
	}

	err = CmdInDir("./backend", "terraform", "fmt")
	if err != nil {
		panic(fmt.Errorf("terraform fmt backend: %w", err))
	}
}

func UpdateLambdas(path string) error {
	entries, err := fs.ReadDir(os.DirFS(path), ".")
	if err != nil {
		return fmt.Errorf("readdir err: %w", err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if e.Name() == "build" {
			continue
		}

		err = CmdInDir(filepath.Join(path, e.Name()), "go", "get")
		if err != nil {
			return fmt.Errorf("go get err: %w", err)
		}
		err = CmdInDir(filepath.Join(path, e.Name()), "go", "mod", "tidy")
		if err != nil {
			return fmt.Errorf("go mod tidy err: %w", err)
		}
	}
	return nil
}

func ListSubdirectories(dir string) (resp []string, err error) {
	list, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed reading: %w", err)
	}
	resp = make([]string, 0, len(list))
	for _, de := range list {
		if de.IsDir() {
			resp = append(resp, de.Name())
		}
	}
	return resp, nil
}

func CmdInDir(dir, cmd string, args ...string) error {
	cmdObj := exec.Command(cmd, args...)
	abs, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("abs error: %w", err)
	}
	log.Printf("Running '%s' in '%s'", cmd, abs)
	cmdObj.Dir = abs
	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr

	err = cmdObj.Run()
	if err != nil {
		return fmt.Errorf("exec error: %w", err)
	}

	return nil
}

func DirExists(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	if s.IsDir() {
		return true
	}
	return false
}
