# filetest

Do a lot of file/code generation? Yeah, me too! And you know what I've noticed about it? It's a huge pain in the ass test! So how do I propose we go about it? Well, with this tool!

```text
$ go get -u github.com/markbates/filetest
```

## Usage

The most basic usage is to create a file called `filetest.json`. Really imaginative, isn't? Then fill it with a array of stuff you want to test.

```json
[{
  "path": "cmd/file.go",
  "contains": [
    "type File struct"
  ]
}, {
  "path": "cmd/root.go",
  "contains": ["pwd, _ = os.Getwd()"]
}]
```

Then just run:

```
$ filetest
```

That's it!

## More Complex Usage

Ok, so you want to make things more complex? OK, I hear you!

Using the `-c` flag you can point "the tool" at either a particalar `.json` file you want to use, or at a directory of `.json` files that it will use to run it's tests.

```text
$ filetest -c some/specific/file.json
$ filetest -c some/rando/directory
```

## Fail Fast

Ashamed of all the failures you're getting and only want to see them 1 at a time? I get it, we all feel overwhelmed sometimes. The `-f` flag has your back.

## Use Go and Want to Add These Tests to Your Test Suite?

Well, you sure are a pushy one, but OK. You're in luck. This is a Go tool, so it's not hard. In fact, the tests for this tool USE THIS TOOL!! OMG!! INCEPTION!!!

```go
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
```

## Match the Whole Darn Thing!

OK, so by now you're saying, what if I want to make the entire file against a known file? What am I supposed to do copy and paste the whole darn thing into a JSON file? C'mon? First, let's take the tone down a bit, there's really no need for that type of langage. I'm trying to help, remember? All we need to do in that case is use the `equals_path` attribute in our JSON file to point at another file. Did I confuse you? Yeah, that's what I thought.

Here's an example:

```json
[{
  "path": "../cmd/errors.go",
  "equals_path": "../cmd/errors.go"
}, {
  "path": "../cmd/root.go",
  "equals_path": "../cmd/root_test.go"
}]
```
