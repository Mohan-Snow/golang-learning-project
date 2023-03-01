.PHONY: run
run:
	@export EXTERNAL_API_TOKEN=.token && go run http-server/*.go

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
		-e EXTERNAL_API_TOKEN=/secret/.token \
		go-learning-proj

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
	@docker container rm b4821e029bf9 de345bac9d43 e6c89cbb238d cc930b83e04c fc173217ce52 f9654ecc94fd bab394c001ae 77984ca5912f c4a62adddcba 6259fa28747a 09d3cadc10a6