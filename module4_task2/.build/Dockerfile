FROM golang:1.15.8-buster

WORKDIR /module4_task2
COPY ./ /module4_task2

RUN apt-get update && apt-get install -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

RUN wget --progress=dot:giga https://github.com/gohugoio/hugo/releases/download/v0.79.0/hugo_0.79.0_Linux-64bit.deb \
    && dpkg -i hugo_0.79.0_Linux-64bit.deb \
    && rm hugo_0.79.0_Linux-64bit.deb

RUN npm install -g markdownlint-cli@0.29.0

ENTRYPOINT ["/module4_task2/ls"]
