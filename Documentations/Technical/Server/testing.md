# Testing & Quality Assurance

## Overview

This document describes how to run tests, check code quality, and ensure the reliability of the backend server.

## Running Tests

The backend uses Go's built-in testing framework. To run all tests:

```bash
make test
```

> Test files are located alongside their respective code files and are named with the `_test.go` suffix. In these files, functions are prefixed with `Test` to be recognized as tests.

## Test Coverage

The server covers more than 80% of the codebase with tests. The coverage report will be generated in the `coverage.out` file, which can be viewed in a browser using:

```bash
make coverage
```

## Linting

The project uses `golangci-lint` for static code analysis. To run the linter:

```bash
make lint
```
