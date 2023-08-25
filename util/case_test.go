package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitSnake(t *testing.T) {
	in := "my_test_string"
	expect := []string{"my", "test", "string"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitDash(t *testing.T) {
	in := "my-test-string"
	expect := []string{"my", "test", "string"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitSpace(t *testing.T) {
	in := "my test string"
	expect := []string{"my", "test", "string"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitTitle(t *testing.T) {
	in := "MyTestString"
	expect := []string{"my", "test", "string"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitCamel(t *testing.T) {
	in := "myTestString"
	expect := []string{"my", "test", "string"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitLower(t *testing.T) {
	in := "test"
	expect := []string{"test"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitCamelEndUpper(t *testing.T) {
	in := "myTestStringS"
	expect := []string{"my", "test", "string", "s"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitUpper(t *testing.T) {
	in := "TEST"
	expect := []string{"test"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitNumbers(t *testing.T) {
	in := "1234"
	expect := []string{"1234"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitGreek(t *testing.T) {
	in := "ΔαααΘθθθ"
	expect := []string{"δααα", "θθθθ"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitCamelNum(t *testing.T) {
	in := "myTestString32"
	expect := []string{"my", "test", "string32"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitID(t *testing.T) {
	in := "ID"
	expect := []string{"id"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitSingle(t *testing.T) {
	in := "A"
	expect := []string{"a"}
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_splitEmpty(t *testing.T) {
	in := ""
	var expect []string
	out := split(in)

	assert.Equal(t, expect, out)
}

func Test_TitleCase(t *testing.T) {
	in := "my-test-string"
	expect := "MyTestString"
	out := TitleCase(in)

	assert.Equal(t, expect, out)
}

func Test_CamelCase(t *testing.T) {
	in := "my-test-string"
	expect := "myTestString"
	out := CamelCase(in)

	assert.Equal(t, expect, out)
}

func Test_DashCase(t *testing.T) {
	in := "myTestString"
	expect := "my-test-string"
	out := DashCase(in)

	assert.Equal(t, expect, out)
}

func Test_UnderCase(t *testing.T) {
	in := "myTestString"
	expect := "my_test_string"
	out := UnderCase(in)

	assert.Equal(t, expect, out)
}
