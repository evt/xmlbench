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
iMac-2:xmlbench jt$ go version
go version go1.8.1 darwin/amd64
iMac-2:xmlbench jt$ go test -bench=.
BenchmarkParsers/hotels.xml-encoding/xml-4         	      50	  22472369 ns/op
BenchmarkParsers/hotels.xml-cvik/xml-4             	     100	  16518842 ns/op
BenchmarkParsers/hotels.xml-libxml2-4              	     100	  10192672 ns/op
BenchmarkParsers/flights.xml-encoding/xml-4        	     500	   3688593 ns/op
BenchmarkParsers/flights.xml-cvik/xml-4            	    1000	   1998423 ns/op
BenchmarkParsers/flights.xml-libxml2-4             	    1000	   1421344 ns/op
PASS
ok  	xmlbench	9.880s
iMac-2:xmlbench jt$ go test -bench=.
BenchmarkParsers/hotels.xml-encoding/xml-4         	      50	  22240798 ns/op
BenchmarkParsers/hotels.xml-cvik/xml-4             	     100	  15972411 ns/op
BenchmarkParsers/hotels.xml-libxml2-4              	     100	  10096749 ns/op
BenchmarkParsers/flights.xml-encoding/xml-4        	     500	   3689078 ns/op
BenchmarkParsers/flights.xml-cvik/xml-4            	    1000	   2044691 ns/op
BenchmarkParsers/flights.xml-libxml2-4             	    1000	   1439671 ns/op
PASS
ok  	xmlbench	9.912s
```
