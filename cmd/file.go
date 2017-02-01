package cmd

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"golang.org/x/sync/errgroup"
)

type File struct {
	Path     string
	Contains []string
}

func (f File) Test() error {
	_, err := os.Stat(f.Path)
	if err != nil {
		return Add(errors.Errorf("%s: does not exist", f.Path))
	}
	b, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}
	for _, s := range f.Contains {
		if !bytes.Contains(b, []byte(s)) {
			err = Add(errors.Errorf("%s: does not contain '%s'", f.Path, s))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Files []File

func (ff Files) Test() error {
	var g errgroup.Group
	for _, f := range ff {
		g.Go(f.Test)
	}
	return g.Wait()
}
