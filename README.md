# snaker

[![Build Status](https://travis-ci.org/godoes/snaker.svg?branch=master)](https://travis-ci.org/godoes/snaker)
[![GoDoc](https://godoc.org/github.com/godoes/snaker?status.svg)](https://godoc.org/github.com/godoes/snaker)

This is a small utility to convert camel cased strings to snake case and back, except some defined words.

## QBS Usage

To replace the original toSnake and back algorithms for [https://github.com/coocood/qbs](https://github.com/coocood/qbs)
you can easily use snaker:

Import snaker
```shell
import (
  github.com/coocood/qbs
  github.com/godoes/snaker
)
```

Register the snaker methods to qbs
```shell
qbs.ColumnNameToFieldName = snaker.SnakeToCamel
qbs.FieldNameToColumnName = snaker.CamelToSnake
```
