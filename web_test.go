package utils_test

import (
	"testing"

	"github.com/go-universal/utils"
	"github.com/stretchr/testify/assert"
)

func TestRelativeURL(t *testing.T) {
	tests := []struct {
		root     string
		path     []string
		expected string
	}{
		{"g:/mekramy", []string{"g:/mekramy/utils/web.go"}, "utils/web.go"},
		{"g:/mekramy", []string{"g:/mekramy/utils"}, "utils"},
		{"g:/mekramy", []string{"g:/mekramy"}, ""},
	}

	for _, test := range tests {
		result := utils.RelativeURL(test.root, test.path...)
		assert.Equal(t, test.expected, result, "RelativeURL(%q, %q)", test.root, test.path)
	}
}

func TestAbsoluteURL(t *testing.T) {
	tests := []struct {
		root     string
		path     []string
		expected string
	}{
		{"g:/mekramy", []string{"g:/mekramy/utils/web.go"}, "/utils/web.go"},
		{"g:/mekramy", []string{"g:/mekramy/utils"}, "/utils"},
		{"g:/mekramy", []string{"g:/mekramy"}, "/"},
	}

	for _, test := range tests {
		result := utils.AbsoluteURL(test.root, test.path...)
		assert.Equal(t, test.expected, result, "AbsoluteURL(%q, %q)", test.root, test.path)
	}
}

func TestSanitizeRaw(t *testing.T) {
	tests := []struct {
		data     string
		trim     bool
		expected string
	}{
		{"<script>alert('xss')</script>", false, ""},
		{"<b>bold</b>", true, "bold"},
		{"  <i>italic</i>  ", true, "italic"},
	}

	for _, test := range tests {
		result := utils.SanitizeRaw(test.data, test.trim)
		assert.Equal(t, test.expected, result, "SanitizeRaw(%q, %t)", test.data, test.trim)
	}
}

func TestSanitizeCommon(t *testing.T) {
	tests := []struct {
		data     string
		trim     bool
		expected string
	}{
		{"<script>alert('xss')</script>", false, ""},
		{"<b>bold</b>", true, "<b>bold</b>"},
		{"  <i>italic</i>  ", true, "<i>italic</i>"},
	}

	for _, test := range tests {
		result := utils.SanitizeCommon(test.data, test.trim)
		assert.Equal(t, test.expected, result, "SanitizeCommon(%q, %t)", test.data, test.trim)
	}
}
