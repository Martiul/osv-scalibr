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

//go:build linux

package regpatchlevel

import (
	"context"
	"fmt"

	"github.com/google/osv-scalibr/extractor"
	"github.com/google/osv-scalibr/extractor/standalone"
	"github.com/google/osv-scalibr/purl"
)

// Name of the extractor
const Name = "windows/regpatchlevel"

// Extractor implements the regpatchlevel extractor.
type Extractor struct{}

// Name of the extractor.
func (e Extractor) Name() string { return Name }

// Version of the extractor.
func (e Extractor) Version() int { return 0 }

// Extract is a no-op for Linux.
func (e *Extractor) Extract(ctx context.Context, input *standalone.ScanInput) ([]*extractor.Inventory, error) {
	return nil, fmt.Errorf("only supported on Windows")
}

// ToPURL converts an inventory created by this extractor into a PURL.
func (e *Extractor) ToPURL(i *extractor.Inventory) (*purl.PackageURL, error) {
	return nil, fmt.Errorf("only supported on Windows")
}

// ToCPEs converts an inventory created by this extractor into CPEs, if supported.
func (e *Extractor) ToCPEs(i *extractor.Inventory) ([]string, error) {
	return nil, fmt.Errorf("only supported on Windows")
}
