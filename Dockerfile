# build go package in ./pocket-base as builder
FROM golang:1.20.0-alpine3.17 AS builder

WORKDIR /app

COPY ../pocket-base/go.mod ../pocket-base/go.sum ./

RUN go env -w GO111MODULE=on

RUN go mod download

COPY ../pocket-base/*.go .

COPY ../pocket-base/migrations ./migrations

RUN go build -o pocketbase .

# build vue client in ./vue-client as node-builder
FROM node:18.14.0-alpine3.17 AS node-builder

WORKDIR /app

COPY ../vue-client/package.json .

RUN npm install

COPY ../vue-client .

RUN npm run build:docker-single

# build final image
FROM golang:1.20.0-alpine3.17 AS final

WORKDIR /app

COPY --from=builder /app/pocketbase ./

COPY --from=node-builder /app/dist ./dist

EXPOSE 8090

RUN ls /app

CMD ["/app/pocketbase", "serve", "--http=0.0.0.0:8090"]