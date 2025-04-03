package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/go-universal/utils"
)

func TestExtractNumbers(t *testing.T) {
	input := "abc123def456"
	expected := "123456"
	result := utils.ExtractNumbers(input)
	assert.Equal(t, expected, result, "ExtractNumbers(%q)", input)
}

func TestExtractAlphaNum(t *testing.T) {
	input := "abc123!@_def456"
	expected := "abc123def456"
	result := utils.ExtractAlphaNum(input)
	assert.Equal(t, expected, result, "ExtractAlphaNum(%q)", input)
}

func TestExtractAlphaNumPersian(t *testing.T) {
	input := "abc123!@_def456سلام گچپژ"
	expected := "abc123def456سلامگچپژ"
	result := utils.ExtractAlphaNumPersian(input)
	assert.Equal(t, expected, result, "ExtractAlphaNumPersian(%q)", input)
}

func TestRandomNumeric(t *testing.T) {
	length := uint(10)
	result := utils.RandomNumeric(length)
	assert.Equal(t, int(length), len(result), "RandomNumeric(%d) length mismatch", length)
	assert.Regexp(t, `^\d+$`, result, "RandomNumeric(%d) should return numeric string", length)
}

func TestRandomAlphaNum(t *testing.T) {
	length := uint(10)
	result := utils.RandomAlphaNum(length)
	assert.Equal(t, int(length), len(result), "RandomAlphaNum(%d) length mismatch", length)
	assert.Regexp(t, `^[A-Z0-9]+$`, result, "RandomAlphaNum(%d) should return alphanumeric string", length)
}

func TestSlugify(t *testing.T) {
	input := []string{"Hello-- ", "  World!"}
	expected := "hello-world"
	result := utils.Slugify(input...)
	assert.Equal(t, expected, result, "Slugify(%q)", input)
}

func TestSlugifyPersian(t *testing.T) {
	input := []string{"سلام", "دنیا!"}
	expected := "سلام-دنیا"
	result := utils.SlugifyPersian(input...)
	assert.Equal(t, expected, result, "SlugifyPersian(%q)", input)
}

func TestConcat(t *testing.T) {
	input := []string{"Hello", "", "      ", "World"}
	sep := " "
	expected := "Hello World"
	result := utils.Concat(sep, input...)
	assert.Equal(t, expected, result, "Concat(%q, %q)", sep, input)
}

func TestFormatNumber(t *testing.T) {
	layout := "%d Dollars"
	value := 100000
	expected := "100,000 Dollars"
	result := utils.FormatNumber(layout, value)
	assert.Equal(t, expected, result, "FormatNumber(%q, %d)", layout, value)
}

func TestFormatRx(t *testing.T) {
	data := "123456"
	pattern := `(\d{3})(\d{2})(\d{1})`
	repl := "($1) $2-$3"
	expected := "(123) 45-6"
	result, err := utils.FormatRx(data, pattern, repl)
	assert.NoError(t, err, "FormatRx(%q, %q, %q) returned error", data, pattern, repl)
	assert.Equal(t, expected, result, "FormatRx(%q, %q, %q)", data, pattern, repl)
}
