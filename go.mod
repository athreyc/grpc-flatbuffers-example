module github.com/athreyc/grpc-flatbuffers-example

go 1.17

require (
	github.com/google/flatbuffers v23.5.26+incompatible
	github.com/athreyc/grpc-flatbuffers-example v0.0.0-20211215101029-a94887c57c64
	google.golang.org/grpc v1.61.0
)

replace  github.com/safeie/grpc-flatbuffers-example => github.com/athreyc/grpc-flatbuffers-example v0.0.0-20211215101029-a94887c57c64

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
