package filefinder

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindFile(t *testing.T) {
	type args struct {
		ctx       context.Context
		finder    FileFinder
		startPath string
		filename  string
		subDir    string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "File exists",
			args: args{
				ctx:       context.Background(),
				finder:    NewSequential(),
				startPath: "", // to be replaced with tmpDir in the test case
				filename:  "testfile1.txt",
				subDir:    "subdir1",
			},
			want:    "subdir1/testfile1.txt", // This path will be joined with tmpDir
			wantErr: assert.NoError,
		},
		{
			name: "File does not exist",
			args: args{
				ctx:       context.Background(),
				finder:    NewSequential(),
				startPath: "", // to be replaced with tmpDir in the test case
				filename:  "nonexistent.txt",
			},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name: "File exists",
			args: args{
				ctx:       context.Background(),
				finder:    NewConcurrent(),
				startPath: "", // to be replaced with tmpDir in the test case
				filename:  "testfile1.txt",
				subDir:    "subdir1",
			},
			want:    "subdir1/testfile1.txt", // This path will be joined with tmpDir
			wantErr: assert.NoError,
		},
		{
			name: "File does not exist",
			args: args{
				ctx:       context.Background(),
				finder:    NewConcurrent(),
				startPath: "", // to be replaced with tmpDir in the test case
				filename:  "nonexistent.txt",
			},
			want:    "",
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			defer func() {
				assert.NoError(t, os.RemoveAll(tmpDir)) // Ensure cleanup after test is done
			}()

			tt.args.startPath = tmpDir // Set the temporary directory for the test

			if tt.args.subDir != "" {
				createTempFile(t, tmpDir, tt.args.subDir, tt.args.filename)
				tt.want = filepath.Join(tmpDir, tt.want) // Update the want path to include tmpDir
			}

			got, err := tt.args.finder.FindFile(tt.args.ctx, tt.args.startPath, tt.args.filename)
			if !tt.wantErr(t, err, fmt.Sprintf("FindFile(%v, %v, %v)", tt.args.ctx, tt.args.startPath, tt.args.filename)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindFile(%v, %v, %v)", tt.args.ctx, tt.args.startPath, tt.args.filename)
		})
	}
}

func BenchmarkSequentialFileFinder(b *testing.B) {
	benchmarkFileFinder(b, &sequential{})
}

func BenchmarkConcurrentFileFinder(b *testing.B) {
	benchmarkFileFinder(b, &concurrent{})
}

func benchmarkFileFinder(b *testing.B, finder FileFinder) {
	ctx := context.Background()
	rootPath := getProjectRootPath(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := finder.FindFile(ctx, rootPath, "finder.go")
		require.NoError(b, err)
	}
}

func getProjectRootPath(t require.TestingT) string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	output, err := cmd.Output()
	require.NoError(t, err)
	// The output will contain a newline character at the end, so we trim space to clean it up.
	return string(bytes.TrimSuffix(output, []byte("\n")))
}

func createTempFile(t *testing.T, tmpDir string, subDir string, filename string) {
	subDirPath := filepath.Join(tmpDir, subDir)
	require.NoError(t, os.Mkdir(subDirPath, 0755))
	filePath := filepath.Join(subDirPath, filename)
	require.NoError(t, os.WriteFile(filePath, []byte{}, 0644))
}
