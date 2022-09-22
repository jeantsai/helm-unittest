package common

import (
	// "bytes"

	yaml "gopkg.in/yaml.v3"
)

// TrustedMarshalYAML marshal yaml without error returned, if an error happens it panics
func TrustedMarshalYAML(d interface{}) string {
	// b := bytes.Buffer{}
	// encoder := yaml.NewEncoder(&b)
	// encoder.SetIndent(2)
	// err := encoder.Encode(d)
	s, err := yaml.Marshal(d)
	if err != nil {
		panic(err)
	}
	// s := b.String()
	return string(s)
}
