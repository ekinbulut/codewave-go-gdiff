# gdiff

[![Go](https://github.com/ekinbulut/go-http-crawler/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/ekinbulut/go-http-crawler/actions/workflows/go.yml)

A simple crawler application to track any changes on a given website in a interval and sends email.

## Usage

```bash
    #checks google.com each 5 seconds and writes the output to the file.
    go run main.go -u "http://www.google.com" -o "output.html" -i 5 -user "user@email.com"
```


## Author

@ekinbulut

## Contribution

Master branch is the main line. Fork, develop, open PR.


