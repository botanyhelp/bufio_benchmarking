package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	// there are usually gzip files in /usr/lib/ on linux machines.  
	// these were found on my GCP cloudshell machine with this command:
	// find /usr/lib/ -name "*.gz" -print > /var/tmp/gz_files.txt
	gzipFilePaths := []string{
		"/usr/lib/x86_64-linux-gnu/perl/debian-config-data-5.38.2/config.sh.debug.gz",
		"/usr/lib/x86_64-linux-gnu/perl/debian-config-data-5.38.2/config.sh.static.gz",
		"/usr/lib/x86_64-linux-gnu/perl/debian-config-data-5.38.2/config.sh.shared.gz",
		"/usr/lib/google-cloud-sdk/platform/gsutil/gslib/tests/test_data/favicon.ico.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/unpack200.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/javadoc.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jstat.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/rmid.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/rmic.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/keytool.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jdb.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/rmiregistry.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jconsole.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/serialver.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/javac.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jinfo.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jdeps.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jarsigner.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jmap.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jjs.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/java.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jps.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/pack200.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jstack.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jar.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jrunscript.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jcmd.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/jstatd.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/ja_JP.UTF-8/man1/javap.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/unpack200.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/javadoc.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jstat.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/rmid.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/rmic.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/keytool.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jdb.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/rmiregistry.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jconsole.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/serialver.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/javac.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jinfo.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jdeps.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jarsigner.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jmap.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jjs.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/java.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jps.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/pack200.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jstack.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jar.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jrunscript.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jcmd.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/jstatd.1.gz",
		"/usr/lib/jvm/java-11-openjdk-amd64/man/man1/javap.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/javadoc.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jwebserver.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jstat.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jdeprscan.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jfr.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/keytool.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jdb.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/rmiregistry.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jconsole.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/serialver.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/javac.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jmod.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jinfo.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jdeps.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jarsigner.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jmap.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/java.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jps.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jhsdb.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jstack.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jshell.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jar.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jrunscript.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jcmd.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jstatd.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jpackage.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/javap.1.gz",
		"/usr/lib/jvm/java-21-openjdk-amd64/man/man1/jlink.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/javadoc.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jstat.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jdeprscan.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jfr.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/keytool.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jdb.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/rmiregistry.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jconsole.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/serialver.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/javac.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jmod.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jinfo.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jdeps.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jarsigner.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jmap.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/java.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jps.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jhsdb.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jstack.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jshell.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jar.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jrunscript.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jcmd.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jstatd.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jpackage.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/javap.1.gz",
		"/usr/lib/jvm/java-17-openjdk-amd64/man/man1/jlink.1.gz",
	}
	for _, gzipFilePath := range gzipFilePaths {
		// Open .gz file
		//file, err := os.Open("/usr/lib/jvm/java-17-openjdk-amd64/man/man1/java.1.gz")
		file, err := os.Open(gzipFilePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()
	
		// Create gzip.Reader from opened file
		gzipReader, err := gzip.NewReader(file)
		if err != nil {
			fmt.Printf("Error creating gzip reader: %v\n", err)
			return
		}
		defer gzipReader.Close()
	
		// Read uncompressed data using a chosen buffer
		buffer := make([]byte, 1024)
		totalBytesRead :=0
		for {
			n, err := gzipReader.Read(buffer) // Read into the buffer
			if n > 0 {
				//fmt.Print(string(buffer[:n]))
	                        totalBytesRead += n
			}
			if err == io.EOF { // Check for end of file
				break
			}
			if err != nil { // Handle other errors
				fmt.Printf("Error reading from gzip reader: %v\n", err)
				return
			}
		}
		fmt.Printf("totalBytesRead: %d\nfile: %s\n\n", totalBytesRead, gzipFilePath)
	}
}
