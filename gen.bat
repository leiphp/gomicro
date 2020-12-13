cd datamodels/protos
protoc --micro_out=../ --go_out=../ prods.proto
cd .. && cd ..
