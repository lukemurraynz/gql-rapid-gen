// Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD

package gen

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os/exec"
	"strings"
)

func formatGo(in string) (string, error) {

	cmd := exec.Command("goimports")

	cmd.Stdin = strings.NewReader(in)
	//wc, err := cmd.StdinPipe()
	//if err != nil {
	//	return "", fmt.Errorf("failed setting goimports stdin: %w", err)
	//}

	buf := bytes.NewBuffer(make([]byte, 0, len(in)))
	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		log.Printf("failed running goimports: %s", err)
		formatted, err := format.Source([]byte(in))
		if err != nil {
			log.Printf(in)
			return in, fmt.Errorf("failed formatting Go code: %w", err)
		}
		return string(formatted), nil
	}

	return buf.String(), nil
}
