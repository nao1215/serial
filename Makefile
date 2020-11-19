build:
	go build -o serial cmd/serial/main.go

run: build
	@chmod a+x serial
	./serial ${ARGS}

clean:
	-rm serial
	-rm -rf test/*
	-rm cover.*
	-touch test/.gitkeep

pre_test:
	@echo "Make files for test at test directory."
	@./scripts/setTestEnv.sh
	@echo "--------------------------------------------------------------------"

test: deps pre_test
	-@go test -cover ./... -v -coverprofile=cover.out
	-@go tool cover -html=cover.out -o cover.html
	@echo "--------------------------------------------------------------------"
	@rm -rf test/*
	@touch test/.gitkeep
	# Rewrite the main() test result of the coverage file:)
	@sed -i -e 's/func main() <span class="cov0" title="0">{/func main() <span class="cov8" title="1">{/g' cover.html
	@echo "The tool saved the coverage information in an HTML. See cover.html"

deps:
	dep ensure
	go mod vendor