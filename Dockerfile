# build go package in ./pocket-base as builder
FROM golang:1.21.6-alpine3.19 AS builder

WORKDIR /app

COPY ./pocketbase/go.mod ./pocketbase/go.sum ./

RUN go env -w GO111MODULE=on

RUN go mod download

COPY ./pocketbase/*.go .

COPY ./pocketbase/migrations ./migrations

RUN go build -o pocketbase .

# build vue client in ./vue-client as node-builder
FROM node:20.11.0-alpine3.19 AS node-builder

WORKDIR /app

COPY ./vue-client/package.json .

RUN npm install

COPY ./vue-client .

RUN npm run build:docker-single

# build final image
FROM golang:1.21.6-alpine3.19 AS final

WORKDIR /app

COPY --from=builder /app/pocketbase ./

COPY --from=node-builder /app/dist ./dist

# Set to true to disable the PocketBase UI if not using Docker Compose
ENV POCKETBASE_DISABLE_UI=false

EXPOSE 8090

CMD ["/app/pocketbase", "serve", "--http=0.0.0.0:80"]
