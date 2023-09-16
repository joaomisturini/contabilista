FROM golang:1.21

ARG USER_UID=1000
ARG USER_GID=1000
ARG USERNAME=go

RUN groupadd --gid $USER_GID $USERNAME && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME
USER $USERNAME

RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
