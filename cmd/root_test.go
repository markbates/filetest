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
	r.Len(errs, 2)
}
