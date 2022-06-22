FROM ubuntu:20.04

# As root
ARG DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -y git curl wget gcc

RUN wget --quiet https://go.dev/dl/go1.18.3.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz

# As vscode
RUN useradd --shell /bin/bash -m vscode
USER vscode

ENV PATH=$PATH:/usr/local/go/bin

RUN go install -v golang.org/x/tools/gopls@latest && \
go install -v github.com/ramya-rao-a/go-outline@latest && \
go install -v github.com/stamblerre/gocode@latest

