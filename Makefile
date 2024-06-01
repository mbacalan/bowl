.PHONY: run test generate

generate:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate

run:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run main.go

test:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go test ./... -cover
