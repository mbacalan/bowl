.PHONY: run test generate

generate:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate

format:
	gofmt -s -w .

run:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run main.go

test:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go test ./... -cover
