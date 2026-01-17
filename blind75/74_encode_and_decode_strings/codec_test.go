package codec

import (
	"reflect"
	"testing"
)

func TestCodec(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		{
			name: "example 1",
			strs: []string{"Hello", "World"},
		},
		{
			name: "example 2 - empty string",
			strs: []string{""},
		},
		{
			name: "empty list",
			strs: []string{},
		},
		{
			name: "single string",
			strs: []string{"test"},
		},
		{
			name: "string with delimiter",
			strs: []string{"#", "##", "a#b"},
		},
		{
			name: "string with numbers",
			strs: []string{"123", "456"},
		},
		{
			name: "mixed strings",
			strs: []string{"abc", "", "def", "ghij"},
		},
		{
			name: "special characters",
			strs: []string{"!@#$%", "^&*()", "{}[]"},
		},
		{
			name: "long strings",
			strs: []string{
				"This is a long string with many words",
				"Another long string",
			},
		},
		{
			name: "unicode characters",
			strs: []string{"Hello 世界", "🎉🎊"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := &Codec{}

			// Test encode then decode
			encoded := codec.Encode(tt.strs)
			decoded := codec.Decode(encoded)

			if !reflect.DeepEqual(decoded, tt.strs) {
				t.Errorf("Encode/Decode mismatch: got %v, want %v", decoded, tt.strs)
			}
		})
	}
}

func TestCodecEdgeCases(t *testing.T) {
	codec := &Codec{}

	t.Run("empty string in middle", func(t *testing.T) {
		strs := []string{"before", "", "after"}
		encoded := codec.Encode(strs)
		decoded := codec.Decode(encoded)
		if !reflect.DeepEqual(decoded, strs) {
			t.Errorf("got %v, want %v", decoded, strs)
		}
	})

	t.Run("multiple empty strings", func(t *testing.T) {
		strs := []string{"", "", ""}
		encoded := codec.Encode(strs)
		decoded := codec.Decode(encoded)
		if !reflect.DeepEqual(decoded, strs) {
			t.Errorf("got %v, want %v", decoded, strs)
		}
	})

	t.Run("strings that look like encoding", func(t *testing.T) {
		strs := []string{"5#hello", "10#world"}
		encoded := codec.Encode(strs)
		decoded := codec.Decode(encoded)
		if !reflect.DeepEqual(decoded, strs) {
			t.Errorf("got %v, want %v", decoded, strs)
		}
	})
}
