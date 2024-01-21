start:
	gin --buildArgs "--buildvcs=false" run start 

install-templ:
	go install github.com/a-h/templ/cmd/templ@$(shell go list -m -f '{{ .Version }}' github.com/a-h/templ)

templ:
	templ generate --watch --path=web/templ

css:
	npm run watch:css

parse:
	go run main.go parse wikidot
