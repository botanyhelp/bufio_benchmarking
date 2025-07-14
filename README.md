## bufio_benchmarking golang scripts to benchmark performance 

* **bufio_benchmarking** repo is a pile of golang scripts to benchmark performance 
* typically we might alter buffer sizes and file sizes under various hardware/OS conditions to test and understand expected performance

### go run it

```
go run main.py
```

* there might be a hardcoded file path in the script that we must change since we do not have a file at this PATH:
* **/var/tmp/bufio_benchmarking/xaa30.csv.gz**
* ..so we fix by changing source code or putting a gzip file at that PATH

```
go run main.go 
Error opening file: open 
```
