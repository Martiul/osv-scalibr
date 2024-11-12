// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package mockregistry provides a mock implementation of the registry.Registry interface.
package mockregistry

import (
	"errors"

	"github.com/google/osv-scalibr/common/windows/registry"
)

var (
	errFailedToOpenKey = errors.New("failed to open key")
)

// MockRegistry mocks registry access.
type MockRegistry struct {
	Keys map[string]registry.Key
}

// OpenKey open the requested registry key.
func (o *MockRegistry) OpenKey(path string) (registry.Key, error) {
	if key, ok := o.Keys[path]; ok {
		return key, nil
	}

	return nil, errFailedToOpenKey
}

// Close does nothing when mocking.
func (o *MockRegistry) Close() error {
	return nil
}

// MockKey mocks a registry.Key.
type MockKey struct {
	KName      string
	KClassName string
	KSubkeys   []registry.Key
	KValues    []registry.Value
}

// Name returns the name of the key.
func (o *MockKey) Name() string {
	return o.KName
}

// Close does nothing when mocking.
func (o *MockKey) Close() error {
	return nil
}

// SubkeyNames returns the names of the subkeys of the key.
func (o *MockKey) SubkeyNames() ([]string, error) {
	var names []string
	for _, subkey := range o.KSubkeys {
		names = append(names, subkey.Name())
	}

	return names, nil
}

// Subkeys returns the subkeys of the key.
func (o *MockKey) Subkeys() ([]registry.Key, error) {
	return o.KSubkeys, nil
}

// ClassName returns the class name of the key.
func (o *MockKey) ClassName() ([]byte, error) {
	return []byte(o.KClassName), nil
}

// Values returns the different values contained in the key.
func (o *MockKey) Values() ([]registry.Value, error) {
	return o.KValues, nil
}

// MockValue mocks a registry.Value.
type MockValue struct {
	VName string
	VData []byte
}

// Name returns the name of the value.
func (o *MockValue) Name() string {
	return o.VName
}

// Data returns the data contained in the value.
func (o *MockValue) Data() ([]byte, error) {
	return o.VData, nil
}
