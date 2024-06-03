# tfmod-list

tfmod-list is a small utility written in Go that allows you to collect and list all the modules in a Terraform project.

This utility takes a directory as an argument, loads the configuration, and parses the module calls to collect all the module sources. It prints out a list of unique modules.

## Prerequisites
* Go version 1.22

## Restrictions

* Only support local module

## How to Use

```shell
tfmod-list example/
```

## Author

Kazunori KOJIMA

## License
MPL