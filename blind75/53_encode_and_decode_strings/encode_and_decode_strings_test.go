package encode_and_decode_strings

import (
	"reflect"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "example 1 - basic strings",
			input: []string{"Hello", "World"},
		},
		{
			name:  "example 2 - empty string",
			input: []string{""},
		},
		{
			name:  "example 3 - empty array",
			input: []string{},
		},
		{
			name:  "example 4 - strings with special characters",
			input: []string{"a#b", "c", "de#f"},
		},
		{
			name:  "single string",
			input: []string{"test"},
		},
		{
			name:  "multiple empty strings",
			input: []string{"", "", ""},
		},
		{
			name:  "strings with spaces",
			input: []string{"hello world", "foo bar"},
		},
		{
			name:  "strings with numbers",
			input: []string{"123", "456#789"},
		},
		{
			name:  "long strings",
			input: []string{"this is a very long string with many characters", "another long one"},
		},
		{
			name:  "unicode characters",
			input: []string{"Hello 世界", "🚀 emoji"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := Encode(tt.input)
			decoded := Decode(encoded)
			if !reflect.DeepEqual(decoded, tt.input) {
				t.Errorf("Encode/Decode failed: got %v, want %v", decoded, tt.input)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "basic example",
			input: []string{"Hello", "World"},
			want:  "5#Hello5#World",
		},
		{
			name:  "empty string",
			input: []string{""},
			want:  "0#",
		},
		{
			name:  "empty array",
			input: []string{},
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.input)
			if got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "basic example",
			input: "5#Hello5#World",
			want:  []string{"Hello", "World"},
		},
		{
			name:  "empty string",
			input: "0#",
			want:  []string{""},
		},
		{
			name:  "empty input",
			input: "",
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Decode(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
