# build stage
FROM golang:alpine AS build-env
#RUN apk --no-cache add build-base git bzr mercurial gcc
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /app/goapp /app/
ENTRYPOINT ./goapp