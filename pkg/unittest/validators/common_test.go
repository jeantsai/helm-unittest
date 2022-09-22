package validators_test

import (
	"github.com/lrills/helm-unittest/pkg/unittest/snapshot"
	"github.com/stretchr/testify/mock"
	yaml "gopkg.in/yaml.v3"
)

func makeManifest(doc string) map[string]interface{} {
	manifest := map[string]interface{}{}
	yaml.Unmarshal([]byte(doc), &manifest)
	return manifest
}

type mockSnapshotComparer struct {
	mock.Mock
}

func (m *mockSnapshotComparer) CompareToSnapshot(content interface{}) *snapshot.CompareResult {
	args := m.Called(content)
	return args.Get(0).(*snapshot.CompareResult)
}
