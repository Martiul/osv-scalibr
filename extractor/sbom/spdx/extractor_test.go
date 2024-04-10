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

package spdx_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/osv-scalibr/extractor"
	"github.com/google/osv-scalibr/extractor/sbom/spdx"
	"github.com/google/osv-scalibr/purl"
)

func TestFileRequired(t *testing.T) {
	var e extractor.InventoryExtractor = spdx.Extractor{}

	tests := []struct {
		name           string
		path           string
		wantIsRequired bool
	}{
		{
			name:           "sbom.spdx",
			path:           "testdata/sbom.spdx",
			wantIsRequired: true,
		},
		{
			name:           "sbom.SPDX",
			path:           "testdata/sbom.SPDX",
			wantIsRequired: true,
		},
		{
			name:           "sbom.SpDx",
			path:           "testdata/sbom.SpDx",
			wantIsRequired: true,
		},
		{
			name:           "sbom.spdx.json",
			path:           "testdata/sbom.spdx.json",
			wantIsRequired: true,
		},
		{
			name:           "sbom.spdx.yml",
			path:           "testdata/sbom.spdx.yml",
			wantIsRequired: true,
		},
		{
			name:           "sbom.spdx.rdf",
			path:           "testdata/sbom.spdx.rdf",
			wantIsRequired: true,
		},
		{
			name:           "random_file.ext",
			path:           "testdata/random_file.ext",
			wantIsRequired: false,
		},
		{
			name:           "sbom.spdx.foo.ext",
			path:           "testdata/sbom.spdx.foo.ext",
			wantIsRequired: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := e.FileRequired(tt.path, 0); got != tt.wantIsRequired {
				t.Fatalf("FileRequired(%s): got %v, want %v", tt.path, got, tt.wantIsRequired)
			}
		})
	}
}

func TestExtract(t *testing.T) {
	var e extractor.InventoryExtractor = spdx.Extractor{}

	tests := []struct {
		name          string
		path          string
		wantErr       error
		wantInventory []*extractor.Inventory
	}{
		{
			name: "sbom.spdx.json",
			path: "testdata/sbom.spdx.json",
			wantInventory: []*extractor.Inventory{
				{
					Name: "cpe:2.3:a:nginx:nginx:1.21.1",
					Metadata: &spdx.Metadata{
						CPEs: []string{"cpe:2.3:a:nginx:nginx:1.21.1"},
					},
					Locations: []string{"testdata/sbom.spdx.json"},
					Extractor: e.Name(),
				},
				{
					Name: "openssl",
					Metadata: &spdx.Metadata{
						PURL: getPURL("openssl", "1.1.1l"),
					},
					Locations: []string{"testdata/sbom.spdx.json"},
					Extractor: e.Name(),
				},
			},
		},
		{
			name: "purl_and_cpe.spdx.json",
			path: "testdata/purl_and_cpe.spdx.json",
			wantInventory: []*extractor.Inventory{
				{
					Name: "nginx",
					Metadata: &spdx.Metadata{
						CPEs: []string{"cpe:2.3:a:nginx:nginx:1.21.1"},
						PURL: getPURL("nginx", "1.21.1"),
					},
					Locations: []string{"testdata/purl_and_cpe.spdx.json"},
					Extractor: e.Name(),
				},
				{
					Name: "openssl",
					Metadata: &spdx.Metadata{
						PURL: getPURL("openssl", "1.1.1l"),
					},
					Locations: []string{"testdata/purl_and_cpe.spdx.json"},
					Extractor: e.Name(),
				},
			},
		},
		{
			name: "sbom.spdx",
			path: "testdata/sbom.spdx",
			wantInventory: []*extractor.Inventory{
				{
					Name: "cpe:2.3:a:nginx:nginx:1.21.1",
					Metadata: &spdx.Metadata{
						CPEs: []string{"cpe:2.3:a:nginx:nginx:1.21.1"},
					},
					Locations: []string{"testdata/sbom.spdx"},
					Extractor: e.Name(),
				},
				{
					Name: "openssl",
					Metadata: &spdx.Metadata{
						PURL: getPURL("openssl", "1.1.1l"),
					},
					Locations: []string{"testdata/sbom.spdx"},
					Extractor: e.Name(),
				},
			},
		},
		{
			name: "sbom.spdx.yml",
			path: "testdata/sbom.spdx.yml",
			wantInventory: []*extractor.Inventory{
				{
					Name: "cpe:2.3:a:nginx:nginx:1.21.1",
					Metadata: &spdx.Metadata{
						CPEs: []string{"cpe:2.3:a:nginx:nginx:1.21.1"},
					},
					Locations: []string{"testdata/sbom.spdx.yml"},
					Extractor: e.Name(),
				},
				{
					Name: "openssl",
					Metadata: &spdx.Metadata{
						PURL: getPURL("openssl", "1.1.1l"),
					},
					Locations: []string{"testdata/sbom.spdx.yml"},
					Extractor: e.Name(),
				},
			},
		},
		{
			name: "sbom.spdx.rdf",
			path: "testdata/sbom.spdx.rdf",
			wantInventory: []*extractor.Inventory{
				{
					Name: "cpe:2.3:a:nginx:nginx:1.21.1",
					Metadata: &spdx.Metadata{
						CPEs: []string{"cpe:2.3:a:nginx:nginx:1.21.1"},
					},
					Locations: []string{"testdata/sbom.spdx.rdf"},
					Extractor: e.Name(),
				},
				{
					Name: "openssl",
					Metadata: &spdx.Metadata{
						PURL: getPURL("openssl", "1.1.1l"),
					},
					Locations: []string{"testdata/sbom.spdx.rdf"},
					Extractor: e.Name(),
				},
			},
		},
		{
			name:    "invalid_sbom.spdx",
			path:    "testdata/invalid_sbom.spdx",
			wantErr: cmpopts.AnyError,
		},
		{
			name:    "sbom.spdx.foo.ext",
			path:    "testdata/sbom.spdx.foo.ext",
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tt := range tests {
		// Note the subtest here
		t.Run(tt.name, func(t *testing.T) {
			r, err := os.Open(tt.path)
			defer func() {
				if err = r.Close(); err != nil {
					t.Errorf("Close(): %v", err)
				}
			}()
			if err != nil {
				t.Fatal(err)
			}

			input := &extractor.ScanInput{Path: tt.path, Reader: r}
			got, err := e.Extract(context.Background(), input)
			if diff := cmp.Diff(tt.wantErr, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("Extract(%s) unexpected error (-want +got):\n%s", tt.path, diff)
			}

			want := tt.wantInventory

			if diff := cmp.Diff(want, got, cmpopts.SortSlices(invLess)); diff != "" {
				t.Errorf("Extract(%s) (-want +got):\n%s", tt.path, diff)
			}
		})
	}
}

func TestToPURL(t *testing.T) {
	e := spdx.Extractor{}
	want := &purl.PackageURL{
		Type:      purl.TypePyPi,
		Name:      "name",
		Namespace: "namespace",
		Version:   "1.2.3",
	}
	i := &extractor.Inventory{
		Name: "name",
		Metadata: &spdx.Metadata{
			PURL: want,
			CPEs: []string{},
		},
		Locations: []string{"location"},
	}
	got, err := e.ToPURL(i)
	if err != nil {
		t.Fatalf("ToPURL(%v): %v", i, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ToPURL(%v) (-want +got):\n%s", i, diff)
	}
}

func TestToCPEs(t *testing.T) {
	e := spdx.Extractor{}
	want := []string{"cpe1", "cpe2"}
	i := &extractor.Inventory{
		Name: "name",
		Metadata: &spdx.Metadata{
			CPEs: want,
		},
		Locations: []string{"location"},
	}
	got, err := e.ToCPEs(i)
	if err != nil {
		t.Fatalf("ToCPEs(%v): %v", i, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ToCPEs(%v) (-want +got):\n%s", i, diff)
	}
}

func invLess(i1, i2 *extractor.Inventory) bool {
	return i1.Name < i2.Name
}

func getPURL(name, version string) *purl.PackageURL {
	return &purl.PackageURL{
		Type:       purl.TypeGeneric,
		Name:       name,
		Version:    version,
		Qualifiers: purl.Qualifiers{},
	}
}
