package gomod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		data := []byte(`
module example.com/hello/world

go 1.12

require (
	example.com/foo/bar v1.0.0 
	example.com/indirect v0.0.1 // indirect
)

replace (
	example.com/old v1.0.0 => example.com/new v2.0.0
)

exclude (
	example.com/exclude v0.2.3
)
`)
		got, err := Parse(data)

		want := &GoMod{
			Module: Module{Path: "example.com/hello/world"},
			Go:     "1.12",
			Require: []Require{
				{Path: "example.com/foo/bar", Version: "v1.0.0", Indirect: false},
				{Path: "example.com/indirect", Version: "v0.0.1", Indirect: true},
			},
			Exclude: []Module{
				{Path: "example.com/exclude", Version: "v0.2.3"},
			},
			Replace: []Replace{
				{
					Old: Module{Path: "example.com/old", Version: "v1.0.0"},
					New: Module{Path: "example.com/new", Version: "v2.0.0"},
				},
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("fail", func(t *testing.T) {
		data := []byte(`err data`)
		_, err := Parse(data)
		assert.Error(t, err)
	})
}
