# autogen

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/posener/autogen)

Package autogen automatically generates files from templates.

## Usage

Create a main package somewhere in your project. In that package, locate templates for generating
files. These templates should have the generated file names suffixed with ".gotmpl". The
generated files will be located by default in the parent directory, but this location can be
customized with the `Location` option. The main package should define the variables to be used
in the templates and run the `autogen.Execute` function with this variable.
The main package should include a `//go:generate go run .` comment.

See example at the [./example](./example) directory.

## Sub Packages

* [example](./example): Package example is a simple example that uses autogen.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
