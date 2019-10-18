package filesystem

import (
	"context"
	"testing"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/internal/server/filesystem/internal"
)

type fakeFileSystemAPI struct{}

func (fakeFileSystemAPI) PathExists(path string) (bool, error) {
	return true, nil
}
func (fakeFileSystemAPI) Mkdir(path string) error {
	return nil
}
func (fakeFileSystemAPI) Rmdir(path string) error {
	return nil
}
func (fakeFileSystemAPI) LinkPath(tgt string, src string) error {
	return nil
}

func TestMkdirWindows(t *testing.T) {
	v1alpha1, nil := apiversion.NewVersion("v1alpha1")
	testCases := []struct {
		name string
		path string
		pathCtx	internal.PathContext
		version apiversion.Version
		expectError bool
	}{
		{
			name: "absolute path outside of container context with container context set",
			path: `C:\foo\bar`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "absolute path inside container context with container context set",
			path: `C:\var\lib\kubelet\pods\pv1`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: false,
		},
		{
			name: "absolute path outside of plugin context with plugin context set",
			path: `C:\foo\bar`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "absolute path inside plugin context with plugin context set",
			path: `C:\var\lib\kubelet\plugins\pv1`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: false,
		},
		{
			name: "relative path with invalid character `:` beyond drive letter prefix",
			path: `csi-plugin\pv1:foo`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid character `/`",
			path: `csi-plugin\pv1/foo`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid character `*`",
			path: `csi-plugin\pv1*foo`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid character `?`",
			path: `csi-plugin?pv1\foo`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid character `|`",
			path: `csi-plugin|pv1\foo`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid characters `..`",
			path: `csi-plugin\..\..\..\system32`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "absolute path with invalid prefix `\\`",
			path: `\\csi-plugin\..\..\..\system32`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
	}
	srv, err := NewServer("windows", `C:\var\lib\kubelet\plugins`, `C:\var\lib\kubelet\pods`, &fakeFileSystemAPI{})
	if err != nil {
		t.Fatalf("FileSystem Server could not be initialized for testing: %v", err)
	}
	for _, tc := range testCases {
		t.Logf("test case: %s", tc.name)
		req := &internal.MkdirRequest {
			Path: tc.path,
			Context: tc.pathCtx,
		}
		mkdirResponse, _ := srv.Mkdir(context.TODO(), req, tc.version)
		if tc.expectError && mkdirResponse.Error == "" {
			t.Errorf("Expected error but Mkdir returned a nil error")
		}
		if !tc.expectError && mkdirResponse.Error != "" {
			t.Errorf("Expected no errors but Mkdir returned error: %s", mkdirResponse.Error)
		}
	}
}

func TestRmdirWindows(t *testing.T) {
	v1alpha1, nil := apiversion.NewVersion("v1alpha1")
	testCases := []struct {
		name string
		path string
		pathCtx	internal.PathContext
		version apiversion.Version
		expectError bool
	}{
		{
			name: "absolute path outside of container context with container context set",
			path: `C:\foo\bar`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "absolute path inside container context with container context set",
			path: `C:\var\lib\kubelet\pods\pv1`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: false,
		},
		{
			name: "absolute path outside of plugin context with plugin context set",
			path: `C:\foo\bar`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "absolute path inside plugin context with plugin context set",
			path: `C:\var\lib\kubelet\plugins\pv1`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: false,
		},
		{
			name: "relative path with invalid character `:` beyond drive letter prefix",
			path: `csi-plugin\pv1:foo`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid character `/`",
			path: `csi-plugin\pv1/foo`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid character `*`",
			path: `csi-plugin\pv1*foo`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid character `?`",
			path: `csi-plugin?pv1\foo`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid character `|`",
			path: `csi-plugin|pv1\foo`,
			pathCtx: internal.PLUGIN,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "relative path with invalid characters `..`",
			path: `csi-plugin\..\..\..\system32`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
		{
			name: "absolute path with invalid prefix `\\`",
			path: `\\csi-plugin\..\..\..\system32`,
			pathCtx: internal.CONTAINER,
			version: v1alpha1,
			expectError: true,
		},
	}
	srv, err := NewServer("windows", `C:\var\lib\kubelet\plugins`, `C:\var\lib\kubelet\pods`, &fakeFileSystemAPI{})
	if err != nil {
		t.Fatalf("FileSystem Server could not be initialized for testing: %v", err)
	}
	for _, tc := range testCases {
		t.Logf("test case: %s", tc.name)
		req := &internal.RmdirRequest {
			Path: tc.path,
			Context: tc.pathCtx,
		}
		rmdirResponse, _ := srv.Rmdir(context.TODO(), req, tc.version)
		if tc.expectError && rmdirResponse.Error == "" {
			t.Errorf("Expected error but Rmdir returned a nil error")
		}
		if !tc.expectError && rmdirResponse.Error != "" {
			t.Errorf("Expected no errors but Rmdir returned error: %s", rmdirResponse.Error)
		}
	}
}
