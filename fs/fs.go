package fs

import (
	"bazil.org/fuse/fs"
	_ "bazil.org/fuse/fs/fstestutil"
)

// FS implements the hello world file system.
type FS struct {
	CgroupDir string
}

func (fs FS) Root() (fs.Node, error) {
	return Dir{fs.CgroupDir}, nil
}
