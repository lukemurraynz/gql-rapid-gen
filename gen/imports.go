package gen

import (
	"golang.org/x/exp/slices"
	"strings"
)

func goImports(of *OutputFile) (err error) {
	if len(of.extraData) == 0 {
		return nil
	}

	slices.Sort(of.extraData)
	of.extraData = slices.Compact(of.extraData)

	builder := &strings.Builder{}
	builder.WriteString("import (\n")
	for _, v := range of.extraData {
		builder.WriteString("\t\"")
		builder.WriteString(v)
		builder.WriteString("\"\n")
	}
	builder.WriteString(")\n")

	of.content = slices.Insert(of.content, 0, builder.String())

	return nil
}

func tsImports(of *OutputFile) (err error) {
	return nil
}
