.PHONY: run
run:
	@export GENERATION_QUERY_PARAM=secret/.token && export DB_CONNECTION=secret/.db_conn && go run http-server/*.go

.PHONY: build
build:
	@go build -o ./app http-server/*.go

.PHONY: test
test:
	@go test -v ./tests

.PHONY: docker-build
docker-build:
	@docker build -t go-learning-proj .

.PHONY: docker-run
docker-run:
	@docker run --name go-learning-proj \
		-d \
		--rm \
		-p 8080:8080 \
		-v `pwd`/secret:/secret \
		-e PORT="8081" \
		-e GENERATION_QUERY_PARAM=/secret/.token \
		-e DB_CONNECTION=/secret/.db_conn \
		go-learning-proj

#fix command to run postgres in docker container
.PHONY: docker-run-postgres
docker-run-postgres:
	@docker run -d \
    	--name postgres \
    	-p 5455:5432 \
    	-e POSTGRES_USER=postgres \
    	-e POSTGRES_PASSWORD=postgres \
    	-e POSTGRES_DB=backend \
    	-d \
    	postgres

.PHONY: restart
restart:
	@docker restart defc84bc55a9

.PHONY: docker-stop
docker-stop:
	@docker stop go-learning-proj

.PHONY: images
images:
	@docker images

.PHONY: all-containers
all-containers:
	@docker ps -a

.PHONY: remove-image
remove-image:
	@docker image rm -f go-learning-proj

.PHONY: remove-container
remove-container:
	@docker container rm postgres