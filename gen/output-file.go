// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package gen

import (
	"fmt"
	"strings"
)

const doubleNewline = "\n\n"

type OutputFile struct {
	header          string
	content         []string
	extraData       []string
	generateIfEmpty bool
	overwrite       bool
	format          func(string) (string, error)
	extraDataRender func(file *OutputFile) error
}

func (of *OutputFile) Add(text ...string) {
	of.content = append(of.content, text...)
}

func (of *OutputFile) AddExtraData(text ...string) {
	of.extraData = append(of.extraData, text...)
}

func (of *OutputFile) String() string {
	if !of.generateIfEmpty && len(of.content) == 0 {
		return ""
	}

	b := &strings.Builder{}
	b.WriteString(of.header)
	b.WriteString(doubleNewline)

	for _, s := range of.content {
		b.WriteString(s)
		b.WriteString(doubleNewline)
	}

	return b.String()
}

func (of *OutputFile) Render() (out string, err error) {
	if of.extraDataRender != nil {
		err = of.extraDataRender(of)
		if err != nil {
			return "", fmt.Errorf("failed rendering extra data: %w", err)
		}
	}
	out = of.String()
	if of.format != nil {
		out, err = of.format(out)
		if err != nil {
			return "", fmt.Errorf("failed formatting: %w", err)
		}
	}
	return out, nil
}

func CreateOutputFile(typ string) (ret *OutputFile, err error) {
	def, ok := Types[typ]
	if !ok {
		return nil, fmt.Errorf("unknown file type: %s", typ)
	}
	return &OutputFile{
		header:          def.Header,
		content:         make([]string, 0, 16),
		extraData:       make([]string, 0, 16),
		generateIfEmpty: def.GenerateIfEmpty,
		overwrite:       def.Overwrite,
		format:          def.Format,
		extraDataRender: def.ExtraDataRender,
	}, nil
}
