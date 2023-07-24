module github.com/openpds/connectord

go 1.20

require (
	github.com/openpds/connector-sdk v0.0.0-20230724155932-02aad21fe979
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)

//replace github.com/openpds/connector-sdk => ../connector-sdk
