all: clean update test build-all

clean: 
	@echo '- cleaning...'
	@rm -rf bin

update: 
	@echo '- update dependencies...'
	@go get -u
	@echo '- removing not used dependencies...'
	@go mod tidy

mod-vendor: 
	@echo 'vendoring...'
	@go mod vendor

build-all: build-linux build-mac build-win

build-linux:
	@echo '- building linux...'
	@GOOS=linux go build ${LD_FLAGS} -o bin/$(APP_NAME)

build-arm:
	@echo '- building arm...'
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ${LD_FLAGS} -o bin/$(APP_NAME).arm

build-mac:
	@echo '- building mac...'
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build ${LD_FLAGS} -o bin/$(APP_NAME).app

build-win:
	@echo '- building windows...'
	@GOOS=windows GOARCH=386 CGO_ENABLED=0 go build ${LD_FLAGS} -o bin/$(APP_NAME).exe

test:
	@echo '- testing with coverage...'
	@go test -v ./...
