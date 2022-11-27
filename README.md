![Tests](https://github.com/drichelson/gokeyword/actions/workflows/test.yaml/badge.svg)
# gokeyword
A static analysis tool for finding the go keyword in source code
It was designed to be included in [golangci-lint](https://github.com/golangci/golangci-lint)

## Purpose
The motivation for this linter is to detect usages of the go keyword in web servers etc where its direct usage inside a request handler can be an anti-pattern: any panics will cause the whole process to crash.