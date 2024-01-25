// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package gen

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// skeletonFiles includes core skeleton files
//
//go:embed skeleton/*
var skeletonFiles embed.FS

func WriteSkeleton(outputDir string) (err error) {
	files, err := fs.Sub(skeletonFiles, "skeleton")
	if err != nil {
		panic(err) // Compilation fault
	}
	err = fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		target := filepath.Join(outputDir, path)

		if d.IsDir() {
			err = os.MkdirAll(target, fs.ModePerm)
			if err != nil {
				return fmt.Errorf("failed creating directory '%s': %w", path, err)
			}
			return nil
		}

		base := filepath.Base(target)
		ext := filepath.Ext(target)
		if base == "gomod" && ext == "" {
			target = strings.Replace(target, "gomod", "gomod", 1)
			base = "go"
			ext = ".mod"
		}
		if ext == ".replace" {
			target = strings.Replace(target, ".replace", "", 1)
			ext = filepath.Ext(target)
		}
		if base == "gitignore" {
			target = strings.Replace(target, "gitignore", ".gitignore", 1)
		}
		if base == "gitkeep" {
			target = strings.Replace(target, "gitkeep", ".gitkeep", 1)
		}

		srcFile, err := fs.ReadFile(files, path)
		if err != nil {
			return fmt.Errorf("failed reading skeleton file '%s': %w", path, err)
		}

		err = os.WriteFile(target, srcFile, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed writing skeleton file '%s': %w", path, err)
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func WritePluginSkeleton(files fs.FS, output *Output) (err error) {
	files, err = fs.Sub(files, "skeleton")
	if err != nil {
		panic(err) // Compilation fault
	}

	err = fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		target := path

		base := filepath.Base(target)
		ext := filepath.Ext(target)
		if base == "gomod" && ext == "" {
			target = strings.Replace(target, "gomod", "go.mod", 1)
			base = "go"
			ext = ".mod"
		}
		if ext == ".replace" {
			target = strings.Replace(target, ".replace", "", 1)
			ext = filepath.Ext(target)
		}
		if base == "gitignore" {
			target = strings.Replace(target, "gitignore", ".gitignore", 1)
		}
		if base == "gitkeep" {
			target = strings.Replace(target, "gitkeep", ".gitkeep", 1)
		}

		srcFile, err := fs.ReadFile(files, path)
		if err != nil {
			return fmt.Errorf("failed reading skeleton file '%s': %w", path, err)
		}

		_, err = output.Create(RAW_SKEL, target, string(srcFile))
		if err != nil {
			return fmt.Errorf("failed writing skeleton file '%s': %w", path, err)
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
