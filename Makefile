build:
	go build -o serial cmd/serial/main.go

run: build
	@chmod a+x serial
	./serial ${ARGS}

clean:
	rm serial

test: deps
	@echo "Make files for test at test directory."
	@./scripts/setTestEnv.sh
	@echo "--------------------------------------------------------------------"
	-go test -cover ./...
	@echo "Remove files for test at test directory."
	@echo "--------------------------------------------------------------------"
	rm -rf test/*
	touch test/.gitkeep

deps:
	dep ensure
	go mod vendor