// Copyright 2021-2022 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"testing/fstest"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/compile"
	"github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/storage/disk"
	"github.com/cerbos/cerbos/internal/test"
	"github.com/cerbos/cerbos/internal/util"
)

func TestVerify(t *testing.T) {
	eng := mkEngine(t)
	runSuite := func(dir string) (*Result, error) {
		conf := Config{TestsDir: test.PathToDir(t, filepath.Join("verify", dir))}
		return Verify(context.Background(), eng, conf)
	}

	t.Run("valid", func(t *testing.T) {
		result, err := runSuite("valid")
		require.NoError(t, err)
		require.Len(t, result.Suites, 4)

		for _, sr := range result.Suites {
			t.Run(sr.File, func(t *testing.T) {
				is := assert.New(t)
				switch sr.File {
				case "empty_test.yaml":
					is.False(sr.Skipped)
					is.Empty(sr.Principals)
				case "suite_test.yaml", "inline_fixture_test.yaml":
					is.False(sr.Skipped)
					is.Len(sr.Principals, 1)
					for _, action := range sr.Principals[0].Resources[0].Actions {
						is.False(action.Details.Skipped)
						is.False(action.Details.Failed, "Trace:\n%s\n", action.Details.EngineTrace)
						is.Empty(action.Details.Error)
					}
				case "udf_test.yaml":
					is.False(sr.Skipped)
					is.Len(sr.Principals, 2)
					for _, principal := range sr.Principals {
						for _, resource := range principal.Resources {
							for _, action := range resource.Actions {
								is.False(action.Details.Skipped)
								is.False(action.Details.Failed, "Trace:\n%s\n", action.Details.EngineTrace)
								is.Empty(action.Details.Error, fmt.Sprintf("%s | %s | %s", principal.Name, resource.Name, action.Name))
							}
						}
					}
				}
			})
		}
	})

	t.Run("invalid_fixture", func(t *testing.T) {
		result, err := runSuite("invalid_fixture")
		is := require.New(t)
		is.NoError(err)
		is.True(result.Failed)
		is.Len(result.Suites, 1)
		is.True(strings.HasPrefix(result.Suites[0].Status, "failed to load test fixtures from testdata:"))
	})

	t.Run("invalid_test", func(t *testing.T) {
		result, err := runSuite("invalid_test")

		is := require.New(t)
		is.NoError(err)
		is.True(result.Failed)
		is.Len(result.Suites, 2)

		for _, sr := range result.Suites {
			switch sr.File {
			case "invalid_test.yaml":
				is.False(sr.Skipped)
				is.True(sr.Failed)
				is.True(strings.HasPrefix(sr.Status, "failed to load test suite"))
				is.Empty(sr.Principals)
			case "did_not_expected_key_test.yaml":
				is.False(sr.Skipped)
				is.True(sr.Failed)
				is.Equal(sr.Status, "failed to load test suite: invalid TestSuite.Tests: value must contain at least 1 item(s)")
				is.Empty(sr.Principals)
			}
		}
	})
}

const (
	fauxPrincipals = `{"principals":{"harry":{"id":"harry","roles":["user"]}}}`
	principals     = `
---
principals:
  harry:
    id: harry
    policyVersion: '0210210'
    roles:
      - employee
    attr: &harry_attr
      department: marketing
      geography: GB
      team: design
  maggie:
    id: maggie
    policyVersion: '0210210'
    roles:
      - employee
      - manager
    attr:
      << : *harry_attr
      managed_geographies: "GB"
`
	fauxResources = `{"resources":{"draft_leave_request":{"id": "xx11", "kind": "leave_request"}}}`
	resources     = `
---
resources:
  draft_leave_request: &leave_request
    id: xx125
    kind: leave_request
    policyVersion: '20210210'
    attr: &leave_request_attr
      department: marketing
      geography: GB
      id: XX125
      owner: harry
      status: DRAFT
      team: design
  pending_leave_request:
    << : *leave_request
    attr:
      << : *leave_request_attr
      status: PENDING_APPROVAL
`
	testSuiteTemplate = `
---
name: TestSuite
description: Tests for verifying something
tests:
  - name: Harry's draft leave request
    input: &input
      actions:
        - create
        - "view:public"
        - approve
      resources:
        - draft_leave_request
      principals:
        - harry
        - maggie
    expected:
      - principal: harry
        resource: draft_leave_request
        actions:
          create: EFFECT_ALLOW
          "view:public": EFFECT_ALLOW
          approve: EFFECT_DENY
      - principal: maggie
        resource: draft_leave_request
        actions:
          create: EFFECT_DENY
          "view:public": EFFECT_ALLOW
          approve: EFFECT_DENY
  - name: Harry's pending leave request
    input:
      << : *input
      resources:
        - pending_leave_request
    expected:
      - principal: harry
        resource: pending_leave_request
        actions:
          create: EFFECT_ALLOW
          "view:public": EFFECT_ALLOW
          approve: EFFECT_DENY
      - principal: maggie
        resource: pending_leave_request
        actions:
          create: EFFECT_DENY
          "view:public": EFFECT_ALLOW
          approve: EFFECT_ALLOW
{{.Principals}}
{{.Resources}}
`
)

var (
	ts           *template.Template
	initTemplate sync.Once
)

func genTable(t *testing.T, embedResources, embedPrincipals bool) string {
	t.Helper()
	trimSpaceYAML := func(s string) string { // Removes all lines until a first root-level key
		lines := strings.Split(s, "\n")
		i := 0
		for ; i < len(lines); i++ {
			s := strings.TrimSpace(lines[i])
			if s != "" && s != "---" {
				break
			}
		}
		return strings.Join(lines[i:], "\n")
	}

	initTemplate.Do(func() {
		var err error
		ts, err = template.New("suite").Parse(testSuiteTemplate)
		require.NoError(t, err)
	})
	require.NotNil(t, ts)

	data := struct{ Principals, Resources string }{}
	if embedPrincipals {
		data.Principals = trimSpaceYAML(principals)
	}
	if embedResources {
		data.Resources = trimSpaceYAML(resources)
	}

	var sb strings.Builder
	err := ts.Execute(&sb, data)
	require.NoError(t, err)
	return sb.String()
}

func newMapFile(s string) *fstest.MapFile {
	return &fstest.MapFile{Data: []byte(s)}
}

func Test_doVerify(t *testing.T) {
	eng := mkEngine(t)
	const (
		embedded = iota
		external
		mixed
	)
	options := []int{embedded, external, mixed}
	optionTitles := map[int]string{embedded: "EMBEDDED", external: "EXTERNAL", mixed: "MIXED"}
	for _, optionPrincipals := range options {
		for _, optionResources := range options {
			t.Run(fmt.Sprintf("principals = %v, resources = %v", optionTitles[optionPrincipals], optionTitles[optionResources]), func(t *testing.T) {
				fsys := make(fstest.MapFS)
				if optionResources == external {
					fsys[filepath.Join(util.TestDataDirectory, resourcesFileName)+".yaml"] = newMapFile(resources)
				} else if optionResources == mixed {
					fsys[filepath.Join(util.TestDataDirectory, resourcesFileName)+".json"] = newMapFile(fauxResources)
				}
				if optionPrincipals == external {
					fsys[filepath.Join(util.TestDataDirectory, principalsFileName)+".yaml"] = newMapFile(principals)
				} else if optionPrincipals == mixed {
					fsys[filepath.Join(util.TestDataDirectory, principalsFileName)+".json"] = newMapFile(fauxPrincipals)
				}
				table := genTable(t, optionResources != external, optionPrincipals != external)
				fsys["leave_request_test.yaml"] = newMapFile(table)
				result, err := doVerify(context.Background(), fsys, eng, Config{})
				is := require.New(t)
				is.NoError(err)
				is.Len(result.Suites, 1)
				is.False(result.Suites[0].Skipped)
				is.False(result.Failed, "%+v", result.Suites)
			})
		}
	}
	t.Run("Should fail for faux principals", func(t *testing.T) {
		fsys := make(fstest.MapFS)
		fsys[filepath.Join(util.TestDataDirectory, resourcesFileName)+".yaml"] = newMapFile(resources)
		fsys[filepath.Join(util.TestDataDirectory, principalsFileName)+".json"] = newMapFile(fauxPrincipals)

		table := genTable(t, false, false)
		fsys["leave_request_test.yaml"] = newMapFile(table)
		result, err := doVerify(context.Background(), fsys, eng, Config{})
		is := require.New(t)
		is.NoError(err)
		is.Len(result.Suites, 1)
		is.False(result.Suites[0].Skipped)
		is.True(result.Failed)
	})
	t.Run("Should fail for faux resources", func(t *testing.T) {
		fsys := make(fstest.MapFS)
		fsys[filepath.Join(util.TestDataDirectory, resourcesFileName)+".json"] = newMapFile(fauxResources)
		fsys[filepath.Join(util.TestDataDirectory, principalsFileName)+".yaml"] = newMapFile(principals)

		table := genTable(t, false, false)
		fsys["leave_request_test.yaml"] = newMapFile(table)
		result, err := doVerify(context.Background(), fsys, eng, Config{})
		is := require.New(t)
		is.NoError(err)
		is.Len(result.Suites, 1)
		is.False(result.Suites[0].Skipped)
		is.True(result.Failed)
	})
	t.Run("Several subdirectories with test fixtures", func(t *testing.T) {
		fsys := make(fstest.MapFS)
		ts := genTable(t, false, false)
		for _, dir := range []string{"a", "b", "c"} {
			d := filepath.Join(dir, util.TestDataDirectory)
			fsys[d+"/principals.yaml"] = newMapFile(principals)
			fsys[d+"/resources.yaml"] = newMapFile(resources)
			fsys[dir+"/leave_request_test.yaml"] = newMapFile(ts)
		}

		result, err := doVerify(context.Background(), fsys, eng, Config{})
		is := require.New(t)
		is.NoError(err)
		is.Len(result.Suites, 3)
		for i := 0; i < len(result.Suites); i++ {
			is.Len(result.Suites[i].Principals, 2)
			is.Len(result.Suites[i].Principals[0].Resources, 2)
		}
		is.False(result.Failed, "%+v", result.Suites)
	})
	t.Run("Simple test", func(t *testing.T) {
		fsys := make(fstest.MapFS)
		ts := genTable(t, false, false)
		fsys[filepath.Join(util.TestDataDirectory, resourcesFileName)+".yaml"] = newMapFile(resources)
		fsys[filepath.Join(util.TestDataDirectory, principalsFileName)+".yaml"] = newMapFile(principals)
		fsys["leave_request_test.yaml"] = newMapFile(ts)

		result, err := doVerify(context.Background(), fsys, eng, Config{})
		is := require.New(t)
		is.NoError(err)
		is.Len(result.Suites, 1)
		is.False(result.Suites[0].Skipped)
		is.False(result.Failed, "%+v", result.Suites)
	})
}

func mkEngine(t *testing.T) *engine.Engine {
	t.Helper()

	dir := test.PathToDir(t, "store")

	ctx, cancelFunc := context.WithCancel(context.Background())
	t.Cleanup(cancelFunc)

	store, err := disk.NewStore(ctx, &disk.Conf{Directory: dir})
	require.NoError(t, err)

	schemaMgr, err := schema.New(ctx, store)
	require.NoError(t, err)

	eng, err := engine.New(ctx, engine.Components{
		CompileMgr: compile.NewManagerWithDefaultConf(ctx, store, schemaMgr),
		SchemaMgr:  schemaMgr,
		AuditLog:   audit.NewNopLog(),
	})
	require.NoError(t, err)

	return eng
}
