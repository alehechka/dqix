# BASE IMAGES
FROM --platform=$BUILDPLATFORM golang:1.21 as go-base
FROM --platform=$BUILDPLATFORM node:20.7.0-bullseye as node-base

# BUILD WEB

FROM node-base as web-builder

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm install --prefer-offline

COPY Makefile .
COPY tailwind.config.js .
COPY web/templ web/templ
COPY web/static web/static

RUN npm run build:css

# GENERATE TEMPL

FROM go-base as templ-builder

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY web/templ web/templ

RUN templ generate --path=web/templ

# BUILD SERVER

FROM go-base as go-builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
COPY --from=templ-builder app/web/templ web/templ

ENV CGO_ENABLED=0
ARG RELEASE_VERSION=latest
ARG TARGETOS TARGETARCH

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o server -ldflags "-X 'main.Version=${RELEASE_VERSION}'" main.go 

# SERVE

FROM scratch

WORKDIR /app

COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /app/server .

COPY --from=web-builder app/web/static web/static
COPY web/data web/data

ENTRYPOINT [ "/app/server" ]
