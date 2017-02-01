package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Run(t *testing.T) {
	r := require.New(t)
	err := Run("../filetest.json")
	r.NoError(err)
}

func Test_Run_with_Errors(t *testing.T) {
	errs = []error{}
	r := require.New(t)
	err := Run("./examples")
	r.NoError(err)
	r.Len(errs, 3)
	msgs := make([]string, len(errs))
	for _, e := range errs {
		msgs = append(msgs, e.Error())
	}
	r.Contains(msgs, "../cmd/root.go: expected to equal ../cmd/root_test.go")
	r.Contains(msgs, "i/dont/exist.go: does not exist")
	r.Contains(msgs, "../cmd/file.go: does not contain 'i dont exist'")
}
