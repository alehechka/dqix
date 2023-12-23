start:
	gin run start

build-static:
	./scripts/build-static.sh

templ:
	templ generate --watch --path=internal/templ

css:
	npm run watch:css
