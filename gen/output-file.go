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
	generateIfEmpty bool
}

func (of *OutputFile) Add(text string) {
	of.content = append(of.content, text)
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

func CreateOutputFile(typ string) (ret *OutputFile, err error) {
	def, ok := Types[typ]
	if !ok {
		return nil, fmt.Errorf("unknown file type: %s", typ)
	}
	return &OutputFile{
		header:          def.Header,
		content:         make([]string, 0, 16),
		generateIfEmpty: def.GenerateIfEmpty,
	}, nil
}
