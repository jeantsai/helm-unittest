package valueutils_test

import (
	"testing"

	. "github.com/lrills/helm-unittest/pkg/unittest/valueutils"
	"github.com/stretchr/testify/assert"
)

func TestGetValueOfSetPath(t *testing.T) {
	a := assert.New(t)
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"b":   []interface{}{"_", map[string]interface{}{"c": "yes"}},
			"d":   "no",
			"e.f": "false",
			"g":   map[string]interface{}{"h": "\"quotes\""},
		},
	}

	var expectionsMapping = map[string]interface{}{
		"a.b[1].c":                "yes",
		"a.b[0]":                  "_",
		"a.b":                     []interface{}{"_", map[string]interface{}{"c": "yes"}},
		"a['d']":                  "no",
		"a[\"e.f\"]":              "false",
		"a.g.h":                   "\"quotes\"",
		"x":                       nil,
		"a.x":                     nil,
		`$.a.b[? @.c=="yes"]`:     []interface{}([]interface{}{map[string]interface{}{"c": "yes"}}),
		"$.a['d']":                "no",
	}

	for path, expect := range expectionsMapping {
		actual, err := GetValueOfSetPath(data, path)
		a.Equal(expect, actual)
		a.Nil(err)
	}
}

func TestGetValueOfSetPathError(t *testing.T) {
	a := assert.New(t)
	data := map[string]interface{}{
		"a": map[interface{}]interface{}{
			"b":   []interface{}{"_"},
			"c.d": "no",
		},
	}

	var expectionsMapping = map[string]string{
		"a.b[0].c": "unknown parameter a.b.0.c",
		// jsonpath return nil from a[ on_existing_index ]
		// "a[0]":     "can't get [0] from a non array type:\nb:\n- _\nc.d: \"no\"\n",
		",":        "parsing error: ,\t:1:1 - 1:2 unexpected \",\" while scanning extensions",
		"a.b[0[]]": "parsing error: a.b[0[]]\t:1:6 - 1:7 unexpected \"[\" while scanning array key expected \"]\"",
		"a[c[0]]":  "unknown parameter c.0",
		// "x":        "unknown parameter x",
		"x.x":      "unknown parameter x.x",
	}

	for path, expect := range expectionsMapping {
		actual, err := GetValueOfSetPath(data, path)
		a.Nil(actual)
		a.EqualError(err, expect)
	}
}

func TestBuildValueOfSetPath(t *testing.T) {
	a := assert.New(t)
	data := map[interface{}]interface{}{"foo": "bar"}

	var expectionsMapping = map[string]interface{}{
		"a.b":    map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": data}},
		"a[1]":   map[interface{}]interface{}{"a": []interface{}{nil, data}},
		"a[1].b": map[interface{}]interface{}{"a": []interface{}{nil, map[interface{}]interface{}{"b": data}}},
	}

	for path, expected := range expectionsMapping {
		actual, err := BuildValueOfSetPath(data, path)
		a.Equal(actual, expected)
		a.Nil(err)
	}
}

func TestBuildValueSetPathError(t *testing.T) {
	a := assert.New(t)
	data := map[interface{}]interface{}{"foo": "bar"}

	actual, err := BuildValueOfSetPath(data, "")

	a.Nil(actual)
	a.NotNil(err)
	a.EqualError(err, "set path is empty")
}

func TestMergeValues(t *testing.T) {
	a := assert.New(t)
	dest := map[interface{}]interface{}{
		"a": map[interface{}]interface{}{
			"b":   []interface{}{"_", map[interface{}]interface{}{"c": "yes"}},
			"e.f": "false",
		},
	}
	src := map[interface{}]interface{}{
		"a": map[interface{}]interface{}{
			"b":   []interface{}{"_", map[interface{}]interface{}{"c": "no"}, "a"},
			"d":   "no",
			"e.f": "yes",
		},
	}
	expected := map[interface{}]interface{}{
		"a": map[interface{}]interface{}{
			"b":   []interface{}{"_", map[interface{}]interface{}{"c": "no"}, "a"},
			"d":   "no",
			"e.f": "yes",
		},
	}
	actual := MergeValues(dest, src)
	a.Equal(expected, actual)
}
