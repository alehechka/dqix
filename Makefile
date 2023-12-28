start:
	gin run start

build-static:
	./scripts/build-static.sh

templ:
	templ generate --watch --path=web/templ

css:
	npm run watch:css
