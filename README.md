# xmlbench

```
iMac-2:xml jt$ go test -bench=.
goos: darwin
goarch: amd64
pkg: xmlbench
BenchmarkParser1/1.xml-encoding/xml-4         	     100	  20722206 ns/op
BenchmarkParser1/1.xml-cvik/xml-4             	     100	  15585583 ns/op
BenchmarkParser1/1.xml-libxml2-4              	     100	  11215866 ns/op
PASS
ok  	xmlbench	4.835s
```
