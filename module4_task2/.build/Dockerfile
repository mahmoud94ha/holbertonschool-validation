FROM golang:1.15.8-buster

WORKDIR /module4_task2
COPY ./ /module4_task2

RUN apt-get update \
    && wget --progress=dot:giga https://github.com/gohugoio/hugo/releases/download/v0.79.0/hugo_0.79.0_Linux-64bit.deb \
    && dpkg -i hugo_0.79.0_Linux-64bit.deb \
    && rm -rf hugo_0.79.0_Linux-64bit.deb \
    && apt-get purge -y wget \
    && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/module4_task2/ls"]
