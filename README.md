# gomicro-samples


## Installation
install **protoc** binary at https://github.com/protocolbuffers/protobuf/releases
copy executable file protoc to ***%GOPATH%/bin*** 

Install protoc-gen-go and go-micro (v2)
```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
$ go get https://github.com/micro/protoc-gen-micro/v2
```
that automatically create executable file at ***%GOPATH%/bin*** 
run this command in your shell
```sh
$ protoc --micro_out=helloworld/proto --go_out=helloworld/proto --proto_path=helloworld helloworld/proto/hello.proto
```
hopefully that success generated files in ***helloworld/proto*** dir


## Samples
### HelloWorld
run server with 
```sh
$ go run helloworld/server.go
```
then test with client 
```sh
$ go run helloworld/client.go
```

License
----
MIT


**100% Free Hell Yeah ! royadv**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)