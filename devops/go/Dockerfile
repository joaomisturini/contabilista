FROM golang:1.21

ARG USER_UID=1000
ARG USER_GID=1000
ARG USERNAME=go

RUN groupadd --gid $USER_GID $USERNAME && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME
