# xmlbench

```
iMac-2:xmlbench jt$ go version
go version go1.9 darwin/amd64
iMac-2:xmlbench jt$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkParser1/1.xml-encoding/xml-4         	     100	  20701932 ns/op
BenchmarkParser1/1.xml-cvik/xml-4             	     100	  16271892 ns/op
BenchmarkParser1/1.xml-libxml2-4              	     100	  10072970 ns/op
PASS
ok  	_/Users/jt/xmlbench/src/xmlbench	5.628s
iMac-2:xmlbench jt$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkParser1/1.xml-encoding/xml-4         	     100	  20719905 ns/op
BenchmarkParser1/1.xml-cvik/xml-4             	     100	  16337585 ns/op
BenchmarkParser1/1.xml-libxml2-4              	     100	  10274935 ns/op
PASS
ok  	_/Users/jt/xmlbench/src/xmlbench	4.825s
iMac-2:xmlbench jt$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkParser1/1.xml-encoding/xml-4         	      50	  20157980 ns/op
BenchmarkParser1/1.xml-cvik/xml-4             	     100	  17606331 ns/op
BenchmarkParser1/1.xml-libxml2-4              	     100	  11324089 ns/op
PASS
ok  	_/Users/jt/xmlbench/src/xmlbench	3.996s
```
