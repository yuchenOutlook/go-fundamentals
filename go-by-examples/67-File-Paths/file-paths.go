package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	prt := fmt.Println
	p := filepath.Join("dir1", "dir2", "filename")
	prt("p:", p)

	prt(filepath.Join("dir1//", "filename"))
	prt(filepath.Join("dir1/.../dir1", "filename"))
	prt("p:", p)

	prt("Dir(p): ", filepath.Dir(p))
	prt("Base(p): ", filepath.Base(p))

	prt(filepath.IsAbs("dir/file"))
	prt(filepath.IsAbs("/dir/file"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	prt(ext)

	prt(strings.TrimSuffix(filename, ext))

	// Find a relative path between a base and a target, it returns an error if the target
	// cannot be made relative to base.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	prt(rel)

	rel, err = filepath.Rel("a/b", "c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}