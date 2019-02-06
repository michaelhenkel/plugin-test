# plugin-test
```
mkdir bar
mkdir foo
go get github.com/michaelhenkel/plugin-test
go build github.com/michaelhenkel/plugin-test
go build -buildmode=plugin -o bar/bar.1.so \
	github.com/michaelhenkel/plugin-test/bar/bar.1.go
go build -buildmode=plugin -o bar.2.so \
	github.com/michaelhenkel/plugin-test/bar/bar.2.go
go build -buildmode=plugin -o foo.1.so \
	github.com/michaelhenkel/plugin-test/foo/foo.1.go
