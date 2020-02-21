DBNAME:=[[.AppName]]
ENV:=development

run:
	go run src/main.go

proto-build:
	protoc -I ./../proto user.proto --go_out=plugins=grpc:./src/interface/rpc
	protoc -I ./../proto room.proto --go_out=plugins=grpc:./src/interface/rpc

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/up:
	sql-migrate up -env=$(ENV)
