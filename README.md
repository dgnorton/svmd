# svmd
Converts CSV, TSV, and other separated values formats to a markdown table

## Install

```
go install github.com/dgnorton/svmd
```

## Usage

CSV, TSV, etc. can either be piped in through `stdin` or read from a file using the `-f` flag. Default output is to `stdout` but an output file can specified using the `-o` flag. For other options, run `svmd -h`.

```
cat some.csv | svmd
```
```
svmd -f some.csv
```
