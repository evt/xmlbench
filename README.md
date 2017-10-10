# xmlbench
## go1.9.1
```
iMac-2:xmlbench jt$ go version
go version go1.9.1 darwin/amd64
iMac-2:xmlbench jt$ go test -bench=.
goos: darwin
goarch: amd64
pkg: xmlbench
BenchmarkParsers/hotels.xml-encoding/xml-4         	     100	  18832301 ns/op
BenchmarkParsers/hotels.xml-cvik/xml-4             	     100	  15558379 ns/op
BenchmarkParsers/hotels.xml-libxml2-4              	     200	   9477148 ns/op
BenchmarkParsers/flights.xml-encoding/xml-4        	     500	   3045323 ns/op
BenchmarkParsers/flights.xml-cvik/xml-4            	    1000	   1924364 ns/op
BenchmarkParsers/flights.xml-libxml2-4             	    1000	   1354878 ns/op
PASS
ok  	xmlbench	11.867s
iMac-2:xmlbench jt$ go test -bench=.
goos: darwin
goarch: amd64
pkg: xmlbench
BenchmarkParsers/hotels.xml-encoding/xml-4         	     100	  19223953 ns/op
BenchmarkParsers/hotels.xml-cvik/xml-4             	     100	  15072552 ns/op
BenchmarkParsers/hotels.xml-libxml2-4              	     200	   9497448 ns/op
BenchmarkParsers/flights.xml-encoding/xml-4        	     500	   2997826 ns/op
BenchmarkParsers/flights.xml-cvik/xml-4            	    1000	   1936826 ns/op
BenchmarkParsers/flights.xml-libxml2-4             	    1000	   1337306 ns/op
PASS
ok  	xmlbench	11.833s
```
## go1.8.1
```
```
