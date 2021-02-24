package webdata

import (
	"bytes"
	"io"
	"io/fs"
)

// TODO: minify

func readFs(s fs.FS) []byte {
	var buff bytes.Buffer
	fs.WalkDir(s, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		f, err := s.Open(path)
		if err != nil {
			panic(err)
		}
		io.Copy(&buff, f)
		return nil
	})
	return buff.Bytes()
}
