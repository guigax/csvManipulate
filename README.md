# csvManipulate

A CLI to read a csv, remove all but one column, that is parameterized via flag.

# Requirements

[Go 1.18+](https://go.dev/dl/)

# Build

    go build

It accepts command-line arguments, listed by: 

    csvManipulate --help
    
# Features

### csv

Defines a path to the csv that will be parsed (default `example.csv`)

    -csv="input.csv"
    
### column

The column number that will remain in the csv (at the moment only one column alowed, default `0`)

    -column=2
    
# Usage

With the following file configuration and the following inputs, it will output (in the same the working directory as `output.csv`) the value of the `column2`.

**Input**

```csv
column0,column1,column2
0000000,1111111,2222222
0000000,1111111,2222222
0000000,1111111,2222222
```

**Command**

    csvManipulate -csv="input.csv" -column=2

**Output**

```csv
column2
2222222
2222222
2222222
```
