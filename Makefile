.PHONY: run
run: dep
	go run cmd/api/main.go

.PHONY: doc
doc:
	pandoc -t beamer -o doc/design.pdf doc/design.md

.PHONY: dep
dep:
	dep ensure

.PHONY: dep
test:
	go test ./...
