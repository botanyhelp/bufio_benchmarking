## bufio_benchmarking golang scripts to benchmark performance 

* **bufio_benchmarking** repo is a pile of golang scripts to benchmark performance 
* typically we might alter buffer sizes and file sizes under various hardware/OS conditions to test and understand expected performance

### go run it

```
go run main.py
```

* there might be a hardcoded file path in the script that we must change since we do not have a file at this PATH:
    * **/usr/lib/jvm/java-17-openjdk-amd64/man/man1/java.1.gz**
    * I do see that file on GCP cloudshell, so that hardcoded path did work on GCP cloudshell
* ..so we fix by changing source code or putting a gzip file at that PATH

```
go run main.go 
Error opening file: open /usr/lib/jvm/java-17-openjdk-amd64/man/man1/java.1.gz: no such file or directory
```
 
* we can create another file with 1 million random string characters
* ..and gzip it:

```
mkdir /var/tmp/bufio_benchmarking/
< /dev/urandom tr -dc 'a-zA-Z0-9' | head -c 1000000 | fold -w 20 |gzip > /var/tmp/bufio_benchmarking/randomstrings.txt.gz
```

* ..and then add that full path to hardcoded file path in main.go
* then we can run it and see how long it takes to read 1MB of compressed text
* ..with time:

```
$ time go run main.go 
totalBytesRead: 189086
real    0m8.916s
user    0m19.289s
sys     0m2.844s
```

* some of the hardcoded filepath things are commented out so that code may run in docker container which will not have a ton of gz files laying around
* the dummy gz files are created on the fly and gzipped
