# plugin-test
## Build
```
mkdir bar
mkdir foo
go get github.com/michaelhenkel/plugin-test
go build github.com/michaelhenkel/plugin-test
go build -buildmode=plugin -o bar/bar.1.so \
	go/src/github.com/michaelhenkel/plugin-test/bar/bar.1.go
go build -buildmode=plugin -o bar/bar.2.so \
	go/src/github.com/michaelhenkel/plugin-test/bar/bar.2.go
go build -buildmode=plugin -o foo/foo.1.so \
	go/src/github.com/michaelhenkel/plugin-test/foo/foo.1.go
```
## Run
```
./plugin-test
```
## Start client
```
nc localhost 3333
```
## Load resource definitions
```
{"plugin":"loadIRD","name":"foo","version":"1"}
{"plugin":"loadIRD","name":"bar","version":"1"}
```
## Create resources
```
{"plugin":"runIR","name":"foo","action":"create","data":{"name":"foo1"}}
{"plugin":"runIR","name":"bar","action":"create","data":{"name":"bar1","properties":"prop1"}}
```
## Load v2 resource definition
```
{"plugin":"loadIRD","name":"bar","version":"2"}
```
## Create v2 resource
```
{"plugin":"runIR","name":"bar","action":"create","data":{"name":"bar2","properties":"prop1","references":"ref1"}}
```
