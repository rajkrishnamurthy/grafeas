// Copyright 2018 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provenance

import (
	"testing"

	ppb "github.com/grafeas/grafeas/proto/v1/provenance_go_proto"
)

func TestValidateBuildProvenance(t *testing.T) {
	tests := []struct {
		desc     string
		p        *ppb.BuildProvenance
		wantErrs bool
	}{
		{
			desc:     "missing ID, want error(s)",
			p:        &ppb.BuildProvenance{},
			wantErrs: true,
		},
		{
			desc: "nil command, want error(s)",
			p: &ppb.BuildProvenance{
				Id:       "8c0b1847-f78b-4bf7-8b2e-38e1bb48b125",
				Commands: []*ppb.Command{nil},
			},
			wantErrs: true,
		},
		{
			desc: "invalid command, want error(s)",
			p: &ppb.BuildProvenance{
				Id: "8c0b1847-f78b-4bf7-8b2e-38e1bb48b125",
				Commands: []*ppb.Command{
					{},
				},
			},
			wantErrs: true,
		},
		{
			desc: "nil built artifact, want error(s)",
			p: &ppb.BuildProvenance{
				Id:             "8c0b1847-f78b-4bf7-8b2e-38e1bb48b125",
				BuiltArtifacts: []*ppb.Artifact{nil},
			},
			wantErrs: true,
		},
		{
			desc: "invalid source provenance, want error(s)",
			p: &ppb.BuildProvenance{
				Id: "8c0b1847-f78b-4bf7-8b2e-38e1bb48b125",
				SourceProvenance: &ppb.Source{
					FileHashes: map[string]*ppb.FileHashes{"foo/bar": nil},
				},
			},
			wantErrs: true,
		},
		{
			desc: "valid build provenance, want success",
			p: &ppb.BuildProvenance{
				Id: "8c0b1847-f78b-4bf7-8b2e-38e1bb48b125",
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateBuildProvenance(tt.p)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateBuildProvenance(%+v): got success, want error(s)", tt.desc, tt.p)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateBuildProvenance(%+v): got error(s) %v, want success", tt.desc, tt.p, errs)
		}
	}
}

func TestValidateCommand(t *testing.T) {
	tests := []struct {
		desc     string
		c        *ppb.Command
		wantErrs bool
	}{
		{
			desc:     "missing name, want error(s)",
			c:        &ppb.Command{},
			wantErrs: true,
		},
		{
			desc: "valid command, want success",
			c: &ppb.Command{
				Name: "wc",
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateCommand(tt.c)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateCommand(%+v): got success, want error(s)", tt.desc, tt.c)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateCommand(%+v): got error(s) %v, want success", tt.desc, tt.c, errs)
		}
	}
}

func TestValidateArtifact(t *testing.T) {
	tests := []struct {
		desc     string
		a        *ppb.Artifact
		wantErrs bool
	}{
		{
			desc:     "valid artifact, want success",
			a:        &ppb.Artifact{},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateArtifact(tt.a)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateArtifact(%+v): got success, want error(s)", tt.desc, tt.a)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateArtifact(%+v): got error(s) %v, want success", tt.desc, tt.a, errs)
		}
	}
}

func TestValidateSource(t *testing.T) {
	tests := []struct {
		desc     string
		s        *ppb.Source
		wantErrs bool
	}{
		{
			desc: "nil file hashes, want error(s)",
			s: &ppb.Source{
				FileHashes: map[string]*ppb.FileHashes{"foo/bar": nil},
			},
			wantErrs: true,
		},
		{
			desc: "invalid file hashes, want error(s)",
			s: &ppb.Source{
				FileHashes: map[string]*ppb.FileHashes{
					"foo/bar": &ppb.FileHashes{},
				},
			},
			wantErrs: true,
		},
		{
			desc:     "valid source, want success",
			s:        &ppb.Source{},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateSource(tt.s)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateSource(%+v): got success, want error(s)", tt.desc, tt.s)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateSource(%+v): got error(s) %v, want success", tt.desc, tt.s, errs)
		}
	}
}

func TestValidateFileHashes(t *testing.T) {
	tests := []struct {
		desc     string
		f        *ppb.FileHashes
		wantErrs bool
	}{
		{
			desc:     "missing file hash, want error(s)",
			f:        &ppb.FileHashes{},
			wantErrs: true,
		},
		{
			desc: "empty file hash, want error(s)",
			f: &ppb.FileHashes{
				FileHash: []*ppb.Hash{},
			},
			wantErrs: true,
		},
		{
			desc: "nil file hash element, want error(s)",
			f: &ppb.FileHashes{
				FileHash: []*ppb.Hash{nil},
			},
			wantErrs: true,
		},
		{
			desc: "invalid file hash element, want error(s)",
			f: &ppb.FileHashes{
				FileHash: []*ppb.Hash{
					{},
				},
			},
			wantErrs: true,
		},
		{
			desc: "valid file hashes, want success",
			f: &ppb.FileHashes{
				FileHash: []*ppb.Hash{
					{
						Type:  ppb.Hash_SHA256,
						Value: []byte("foobar"),
					},
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateFileHashes(tt.f)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateFileHashes(%+v): got success, want error(s)", tt.desc, tt.f)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateFileHashes(%+v): got error(s) %v, want success", tt.desc, tt.f, errs)
		}
	}
}

func TestValidateHash(t *testing.T) {
	tests := []struct {
		desc     string
		h        *ppb.Hash
		wantErrs bool
	}{
		{
			desc:     "missing type, want error(s)",
			h:        &ppb.Hash{},
			wantErrs: true,
		},
		{
			desc: "missing value, want error(s)",
			h: &ppb.Hash{
				Type: ppb.Hash_SHA256,
			},
			wantErrs: true,
		},
		{
			desc: "valid hash, want success",
			h: &ppb.Hash{
				Type:  ppb.Hash_SHA256,
				Value: []byte("foobar"),
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateHash(tt.h)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateHash(%+v): got success, want error(s)", tt.desc, tt.h)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateHash(%+v): got error(s) %v, want success", tt.desc, tt.h, errs)
		}
	}
}
