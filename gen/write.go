// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package gen

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func (o *Output) Write(outputDir string) (err error) {
	for name, f := range o.files {
		filename := filepath.Join(outputDir, name)

		content := f.String()

		if len(content) == 0 && !f.generateIfEmpty {
			log.Printf("Not generating %s as empty file", name)
			continue
		}

		err = os.MkdirAll(filepath.Dir(filename), fs.ModePerm)

		var flags = os.O_RDWR | os.O_TRUNC | os.O_CREATE

		if !f.overwrite {
			flags |= os.O_EXCL
		}

		ref, err := os.OpenFile(filename, flags, fs.ModePerm)
		if err != nil {
			if !f.overwrite && errors.Is(err, fs.ErrExist) {
				// We don't overwrite these files, so silently ignore
				return nil
			}
			return fmt.Errorf("failed opening file %s: %w", name, err)
		}

		_, err = ref.WriteString(content)
		if err != nil {
			_ = ref.Close()
			return fmt.Errorf("failed writing file %s: %w", name, err)
		}

		err = ref.Close()
		if err != nil {
			return fmt.Errorf("failed closing file %s: %w", name, err)
		}
		log.Printf("Successfully wrote %s", name)
	}

	return nil
}
