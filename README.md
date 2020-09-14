Task
----
There is a flow of input data on stdin (paths to local files and urls). Program should count the number of words "Go" in all sources and print it to stdout. The computing should me done with specified limited number of goroutines. If there are less data sources than specified number of goroutines, redundant ones should not be created.


How to use
----------

```bash
echo -e 'https://golang.org\n/etc/passwd\nhttps://golang.org\nhttps://golang.org\nhttps://golang.org' | go run main.go
```