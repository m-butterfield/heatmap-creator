cloudrunbasecommand := gcloud run deploy --region=us-central1
deployservercommand := $(cloudrunbasecommand) --image=gcr.io/mattbutterfield/heatmap-creator heatmap-creator

terraformbasecommand := cd infra && terraform
terraformvarsarg := -var-file=secrets.tfvars

export DB_SOCKET=host=localhost dbname=heatmap-creator

deploy: docker-build docker-push
	$(deployservercommand)

docker-build: run-webpack-prod
	docker-compose build

docker-push:
	docker-compose push

fmt:
	go fmt ./...
	staticcheck ./...
	yarn prettier --loglevel warn --write .
	cd infra/ && terraform fmt

generate:
	go run github.com/99designs/gqlgen generate
	yarn graphql-codegen

reset-db:
	dropdb --if-exists heatmap-creator
	createdb heatmap-creator
	psql -f dump.sql heatmap-creator

migrate:
	go run cmd/migrate/main.go

run-server: export USE_LOCAL_FS=true
run-server: export SQL_LOGS=true
run-server: export GQL_PLAYGROUND=true
run-server:
	go run cmd/server/main.go

run-webpack:
	yarn run webpack --mode development --watch

run-webpack-prod:
	rm -rf app/static/js/dist
	yarn run webpack --mode production

tf-plan:
	$(terraformbasecommand) plan $(terraformvarsarg)

tf-apply:
	$(terraformbasecommand) apply $(terraformvarsarg)

tf-refresh:
	$(terraformbasecommand) apply $(terraformvarsarg) -refresh-only

test: export DB_SOCKET=host=localhost dbname=heatmap-creator_test
test:
	dropdb --if-exists heatmap-creator_test
	createdb heatmap-creator_test
	psql heatmap-creator_test -c 'CREATE EXTENSION citext'
	psql heatmap-creator_test -c 'CREATE EXTENSION "uuid-ossp"'
	go run cmd/migrate/main.go
	go test -v ./app/...

update-deps:
	go get -u ./...
	go mod tidy
	yarn upgrade
	cd infra && terraform init -upgrade && cd -
