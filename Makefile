
unit_test:
	docker-compose down
	docker-compose up -d
	go clean -testcache && go test ./... -v
	docker-compose down

provider_test:
	docker-compose down
	docker-compose up -d
	go test ./... -tags="pact_test" -run TestProvider -v
	docker-compose down

all_test:unit_test provider_test
	
docker_build:
	docker build . -t service

docker_run:
	docker run --publish 5050:5050 service

	
