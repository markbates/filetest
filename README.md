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
