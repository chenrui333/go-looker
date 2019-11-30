build:
	go get -d -t -v ./...
	go build ./...

generate: swagger.json
	$(RM) -r client models
	docker run -it --rm \
		-v $(PWD):/go/src/github.com/chenrui333/go-looker \
		-w /go/src/github.com/chenrui333/go-looker \
		quay.io/goswagger/swagger generate client --name looker --spec $<

swagger.yaml:
	curl -fSsL https://demo.looker.com:19999/api/3.1/swagger.json -o $@
