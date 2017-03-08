package cmd

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"golang.org/x/sync/errgroup"
)

//File represents a file inside the test json.
type File struct {
	Path        string   `json:"path"`
	Contains    []string `json:"contains"`
	NotContains []string `json:"!contains"`
	EqualsPath  string   `json:"equals_path"`
	Absent      bool     `json:"absent"`
}

func (f File) Test() error {
	_, err := os.Stat(f.Path)
	if err != nil && f.Absent == false {
		return Add(errors.Errorf("%s: does not exist", f.Path))
	}

	if f.Absent {
		return Add(errors.Errorf("%s: should not be present", f.Path))
	}

	b, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return errors.Errorf("%s: %s", f.Path, err.Error())
	}
	for _, s := range f.Contains {
		if !bytes.Contains(b, []byte(s)) {
			err = Add(errors.Errorf("%s: does not contain '%s'", f.Path, s))
			if err != nil {
				return err
			}
		}
	}
	for _, s := range f.NotContains {
		if bytes.Contains(b, []byte(s)) {
			err = Add(errors.Errorf("%s: should not contain '%s'", f.Path, s))
			if err != nil {
				return err
			}
		}
	}
	if f.EqualsPath != "" {
		_, err = os.Stat(f.EqualsPath)
		if err != nil {
			return Add(errors.Errorf("%s: %s does not exist", f.Path, f.EqualsPath))
		}
		bb, err := ioutil.ReadFile(f.EqualsPath)
		if err != nil {
			return errors.Errorf("%s: %s", f.Path, err.Error())
		}
		if string(bb) != string(b) {
			return Add(errors.Errorf("%s: expected to equal %s", f.Path, f.EqualsPath))
		}
	}
	return nil
}

//Files represents a slice of File
type Files []File

//Test calls Test for each of the Files in the File slice.
func (ff Files) Test() error {
	var g errgroup.Group
	for _, f := range ff {
		g.Go(f.Test)
	}
	return g.Wait()
}
