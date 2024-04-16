# Build stage
FROM golang:latest AS builder
ARG package=./cmd/thumbnailer
WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN --mount=type=cache,target=/go,id=go-path go mod download

COPY . .

ENV CGO_ENABLED 1

# Build the application
RUN --mount=type=cache,target=/go,id=go-path go build -a -ldflags '-w -s -linkmode external -extldflags "-static"' -o main ${package}

# Deploy stage
FROM alpine:latest
ARG USER=deploy
ARG config=config.yaml
ENV HOME /home/$USER

RUN apk add --update sudo
RUN adduser -D $USER \
    && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
    && chmod 0440 /etc/sudoers.d/$USER

USER $USER
WORKDIR $HOME

# Copy Swagger docs from the build stage
COPY ./config/config*.yaml ./config/
COPY --from=builder /build/main .

ENV APP_CONFIG_FILE ${config}
ENV GIN_MODE release
EXPOSE 8080

ENTRYPOINT ./main --config=${APP_CONFIG_FILE}