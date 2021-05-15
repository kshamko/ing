.PHONY: gen-cover cover cover-func
SWAGGER := docker run --rm -it -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $(shell pwd)/internal quay.io/goswagger/swagger
LINTER := docker run --rm -v $(shell pwd):/app -v ${GOPATH}/pkg/mod:/go/pkg/mod -w /app golangci/golangci-lint golangci-lint run --enable-all --disable goerr113,cyclop,exhaustivestruct,gci,gofumpt,lll,testpackage,wrapcheck,paralleltest

swagger:
	docker pull quay.io/goswagger/swagger
	cd internal && $(SWAGGER) generate server --spec ../api/ing.swagger.yaml --exclude-main && cd ..

lint:
	docker pull golangci/golangci-lint	
	$(LINTER)

test:
	go test ./...

docker:
	docker build .	