.PHONY: run
run:
	go run main.go

.PHONY: doc
doc:
	pandoc -t beamer -o doc/design.pdf doc/design.md
