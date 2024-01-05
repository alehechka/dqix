start:
	gin run start

templ:
	templ generate --watch --path=web/templ

css:
	npm run watch:css
