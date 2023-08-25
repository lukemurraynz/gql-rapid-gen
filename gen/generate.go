package gen

import (
	"fmt"
	"golang.org/x/exp/maps"
)

type Output struct {
	files map[string]*OutputFile
}

func (o *Output) EnsureHasFile(typ string, filename string) (of *OutputFile, err error) {
	typObj := Types[typ]
	fullName := typObj.DirectoryPrefix + filename + typObj.Extension

	of, pres := o.files[fullName]
	if !pres {
		of, err = CreateOutputFile(typ)
		if err != nil {
			return nil, fmt.Errorf("failed creating output file: %w", err)
		}
		if o.files == nil {
			o.files = make(map[string]*OutputFile, 256)
		}
		o.files[fullName] = of
	}

	of.generateIfEmpty = true

	return of, nil
}

func (o *Output) AppendOrCreate(typ string, filename string, content string) (of *OutputFile, err error) {
	typObj := Types[typ]
	fullName := typObj.DirectoryPrefix + filename + typObj.Extension

	of, pres := o.files[fullName]
	if !pres {
		of, err = CreateOutputFile(typ)
		if err != nil {
			return nil, fmt.Errorf("failed creating output file: %w", err)
		}
		if o.files == nil {
			o.files = make(map[string]*OutputFile, 256)
		}
		o.files[fullName] = of
	}

	of.Add(content)
	return of, nil
}

func (o *Output) HasFile(typ string, filename string) bool {
	typObj := Types[typ]
	fullName := typObj.DirectoryPrefix + filename + typObj.Extension

	_, pres := o.files[fullName]
	return pres
}

func (o *Output) GetFiles() map[string]*OutputFile {
	return maps.Clone(o.files)
}
