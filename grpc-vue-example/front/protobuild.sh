mkdir -p ./src/assets
protoc -I ./../proto user.proto --js_out=import_style=commonjs,binary:./src/assets --grpc-web_out=import_style=commonjs,mode=grpcweb:./src/assets
protoc -I ./../proto room.proto --js_out=import_style=commonjs,binary:./src/assets --grpc-web_out=import_style=commonjs,mode=grpcweb:./src/assets
