FROM golang:1.19.3

ARG USERNAME=AI1411
# hadolint ignore=DL3008

ENV TZ Asia/Tokyo
RUN useradd -s /bin/bash -m ${USERNAME} \
    && apt-get update \
    && apt-get install --no-install-recommends --no-install-suggests -y \
        less \
        unzip \
        zsh \
    && sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" \
    && apt-get install --no-install-recommends -y \
        iputils-ping \
        jq \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1 \
    && go install mvdan.cc/gofumpt@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && go install github.com/kyoh86/richgo@latest \
    && go install -v google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0 \
    && go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 \
    && go install -v github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.4.1 \
    && go install -v github.com/stormcat24/protodep@0.1.3 \
    && go install -v github.com/golang/mock/mockgen@v1.5.0 \
    && go install gotest.tools/gotestsum@latest \
    && apt-get remove --purge --auto-remove -y \
    && rm -rf /var/lib/apt/lists/*

ENV GO111MODULE on
WORKDIR /go/src/
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]
