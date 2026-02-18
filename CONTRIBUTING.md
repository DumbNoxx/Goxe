# contributing to goxe

thanks for your interest in contributing.

goxe aims to stay small, fast and maintainable.
please keep changes focused and minimal.

## development setup

requirements:

* go 1.25.5+

clone and build:

```bash
git clone git@github.com:DumbNoxx/goxe.git
cd goxe
Task build 
```

this executes:

```bash
go build -o bin/goxe ./cmd/goxe 
```

if you don't use task, run the command directly.


run locally:

```bash
./cmd/goxe
```

## coding guidelines

* prefer simple, readable code over clever abstractions
* avoid unnecessary dependencies
* keep allocations and performance in mind
* document exported functions
* follow standard go formatting:

``` bash
go fmt ./...
```

## tests

we use task as the primary development runner.

run benchmarks:

```bash
Task test
```

this executes:

```bash
go test -bench=. -benchmem ./cmd/goxe
```

if you don't use task, run the command directly.

guidelines:

* include benchmarks for performance-related changes
* include unit tests for bug fixes when possible
* avoid flaky timing-based tests
* keep tests fast and deterministic

## commit style

use clear, descriptive commit messages:

```
feat: add udp buffer optimization
fix: prevent race condition in aggregator
docs: update readme
```

small commits are preferred over large, mixed changes.

## issues

before opening a new issue:

* check if it already exists
* provide a clear description
* include logs or reproduction steps when relevant

feature requests should explain the use case.

## pull requests

1. fork the repository
2. create a feature branch
3. implement your change
4. add tests if applicable
5. open a pull request

please keep pull requests focused on a single change.

large refactors should be discussed in an issue first.

## philosophy

goxe prioritizes:

* performance
* simplicity
* maintainability

if a change conflicts with these goals,
it may not be accepted.

thanks for contributing.

