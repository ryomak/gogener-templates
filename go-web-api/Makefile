
init: install init-db generate-mysql

## generate
generate-mysql:
	mysqldef -u[[.AppName]] -p[[.AppName]] [[.AppName]] < schema.sql
	sqlboiler mysql
generate-di:
	wire ./di/wire.go
generate-swagger:
	swag init

login:
	gcloud auth login
prod:
	gcloud config set project [[.AppName]]
prod-cloudsql:
	cloud_sql_proxy -instances=[[.AppName]]:asia-northeast1:[[.AppName]]=tcp:3306
prod-migrate:
	@$(eval PASSWORD=$(shell stty -echo; read -p "Password: " pwd; stty echo; echo $$pwd))
	mysqldef -uroot -p$(PASSWORD) [[.AppName]] < schema.sql
prod-deploy: prod
	go mod vendor
	GO111MODULE=off gcloud --project [[.AppName]] --quiet app deploy app.yaml

## install
install: install-sqlboiler install-gcloud install-sqldef install-wire install-direnv install-swag

install-gcloud:
	brew cask install google-cloud-sdk
	brew cask info google-cloud-sdk
install-sqldef:
	curl -L -O https://github.com/k0kubun/sqldef/releases/download/v0.5.12/mysqldef_darwin_amd64.zip
	unzip mysqldef_darwin_amd64.zip
	rm mysqldef_darwin_amd64.zip
	mv mysqldef /usr/local/bin/mysqldef
install-sqlboiler:
	go get -u github.com/volatiletech/sqlboiler
	go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
install-cloud-sql-proxy:
	curl -o cloud_sql_proxy https://dl.google.com/cloudsql/cloud_sql_proxy.darwin.amd64
	chmod +x cloud_sql_proxy
	mv cloud_sql_proxy /usr/local/bin/cloud_sql_proxy
install-wire:
	go get github.com/google/wire/cmd/wire
install-direnv: 
	brew install direnv
install-swag:
	go get github.com/swaggo/swag/cmd/swag
init-db:
	sh ./script/db/create_user.sh 
