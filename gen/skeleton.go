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

// templateFiles includes core and plugin templates
//
//go:embed skeleton/*
var skeletonFiles embed.FS

func WriteSkeleton(outputDir string) (err error) {
	err = fs.WalkDir(skeletonFiles, ".", func(path string, d fs.DirEntry, err error) error {
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
		if ext == ".replace" {
			target = strings.Replace(target, ".replace", "", 1)
			ext = filepath.Ext(target)
		}
		if base == "gitkeep" {
			target = strings.Replace(target, "gitkeep", ".gitkeep", 1)
		}

		srcFile, err := fs.ReadFile(skeletonFiles, path)
		if err != nil {
			return fmt.Errorf("failed reading skeleton file '%s': %w", path, err)
		}

		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("failed getting info for skeleton file '%s': %w", path, err)
		}

		err = os.WriteFile(target, srcFile, info.Mode())
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
