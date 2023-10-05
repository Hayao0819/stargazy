package utils

import (
	"archive/tar"
	"io"

	errutils "github.com/Hayao0819/stargazy/utils/error"
	executils "github.com/Hayao0819/stargazy/utils/exec"
	fileutils "github.com/Hayao0819/stargazy/utils/file"
	jsonutils "github.com/Hayao0819/stargazy/utils/json"
	pathutils "github.com/Hayao0819/stargazy/utils/path"
	strutils "github.com/Hayao0819/stargazy/utils/str"
	tarutils "github.com/Hayao0819/stargazy/utils/tar"
	"github.com/spf13/cobra"
)

var Err = struct {
	ExitWithErr          func(cmd *cobra.Command, err error)
	RunFuncWithErrorExit func(cmd *cobra.Command, fn func() error)
}{
	ExitWithErr:          errutils.ExitWithErr,
	RunFuncWithErrorExit: errutils.RunFuncWithErrorExit,
}

var Exec = struct {
	Run func(name string, args ...string) error
}{
	Run: executils.Run,
}

var File = struct {
	Copy func(src string, dest string) error
}{
	Copy: fileutils.Copy,
}

var Json = struct {
	Create func(v any, path string) error
}{
	Create: jsonutils.Create,
}

var Path = struct {
	FromSlash           func(path ...string) string
	BaseWithoutExt      func(path string) string
	RenamePathWithoutEx func(path string, new string) string
	Split               func(path string) (string, string, string)
}{
	FromSlash:           pathutils.FromSlash,
	BaseWithoutExt:      pathutils.BaseWithoutExt,
	RenamePathWithoutEx: pathutils.RenamePathWithoutExt,
	Split:               pathutils.Split,
}

var Str = struct {
	Join func(elems ...string) string
}{
	Join: strutils.Join,
}

var Tar = struct {
	AddFile func(tw *tar.Writer, path ...string) error
	Create  func(writer io.Writer, files ...string) (*tar.Writer, error)
}{
	AddFile: tarutils.AddFile,
	Create:  tarutils.Create,
}
