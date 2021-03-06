build: deps
	go build -o serial cmd/serial/main.go

run: build
	@chmod a+x serial
	./serial ${ARGS}

clean:
	-rm serial
	-rm -rf test/*
	-rm cover.*
	-rm -rf release
	-rm ./docs/man/en/serial.1.gz
	-rm ./docs/man/ja/serial.1.gz
	-touch test/.gitkeep

doc:
	pandoc ./docs/man/en/serial.1.md -s -t man > ./docs/man/en/serial.1
	pandoc ./docs/man/ja/serial.1.md -s -t man > ./docs/man/ja/serial.1
	gzip -f ./docs/man/en/serial.1
	gzip -f ./docs/man/ja/serial.1

install:
	install -m 0755 -D ./serial /usr/local/bin/.
	install -m 0644 -D ./docs/man/en/serial.1.gz /usr/share/man/man1/serial.1.gz
	install -m 0644 -D ./docs/man/ja/serial.1.gz /usr/share/man/ja/man1/serial.1.gz

pre_test:
	@echo "Clean test directory."
	-@rm -rf test/*
	-@rm -rf cmd/serial/*.txt
	@echo "Make files for test at test directory."
	@./scripts/setTestEnv.sh
	@echo "--------------------------------------------------------------------"

test: deps pre_test
	-@go test -cover ./... -v -coverprofile=cover.out
	-@go tool cover -html=cover.out -o cover.html
	@echo "--------------------------------------------------------------------"
	-@rm -rf test/*
	-@rm -rf cmd/serial/*.txt
	@touch test/.gitkeep
	@echo "The tool saved the coverage information in an HTML. See cover.html"

deps:
	dep ensure
	go mod vendor

release:
	./scripts/Release.sh
