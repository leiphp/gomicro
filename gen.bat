cd Services/protos
protoc --go_out=../ Models.proto
protoc --micro_out=../ --go_out=../ ProdService.proto
protoc-go-inject-tag -input=../Models.pb.go
protoc-go-inject-tag -input=../ProdService.pb.go
cd .. && cd ..
