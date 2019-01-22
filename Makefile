PREFIX = 
NAME = go-triangle

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/$(NAME)
clean:
	rm ./bin/$(NAME)
# this one does not needed for now 
# there are no dependanices yet
dep:
	dep ensure

docker-build:
	docker build -t $(PREFIX)$(NAME) .
docker-build-tests:
	docker build -t $(PREFIX)$(NAME)-tests . --target builder
docker-run-unit: docker-build-tests
	docker run --rm $(PREFIX)$(NAME)-tests go test -tags=unit -v

test-unit:
	go test -tags=unit -v

