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
    && apt-get remove --purge --auto-remove -y \
    && rm -rf /var/lib/apt/lists/*

ENV GO111MODULE on
WORKDIR /go/src/
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]
