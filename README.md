# csvManipulate

A CLI to read a csv, remove all but one column, that is passed via flag.

# Requirements

[Go 1.18+](https://go.dev/dl/)

# Run

Execute the `go run` command

    go run main.go

You can also build and execute the code with:

OS | Command
--- | --- 
Windows | `go build main.go && main.exe`
Linux | `go build main.go && ./main`

It accepts command-line arguments, listed by: 

    go run main.go --help
    
# Features

### csv

Defines a path to the csv that will be parsed (default `example.csv`)

    -csv="input.csv"
    
### column

The column number that will remain in the csv (at the moment only one column alowed, default `0`)

    -column=2
    
# Usage

With the following file configuration and the following inputs, it will output (on the same exe location as "output.csv") the value of the "colum2".

```csv
column0,column1,column2
0000000,1111111,2222222
0000000,1111111,2222222
0000000,1111111,2222222
```

    go run main.go -csv="input.csv" -column=2
    
```csv
column2
2222222
2222222
2222222
```
