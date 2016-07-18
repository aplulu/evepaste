gogo:
	protoc --proto_path=/usr/local/include/:${GOPATH}/src/:. --gofast_out=. *.proto

go:
	protoc --proto_path=/usr/local/include/:${GOPATH}/src/:. --go_out=. *.proto
clean:
	rm ./*.pb.go

