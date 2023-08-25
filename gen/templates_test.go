package gen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TemplateEmbed(t *testing.T) {
	for k, _ := range templateCache {
		t.Log(k)
	}
	assert.True(t, len(templateCache) >= 2)
}
