FROM golang:1.15.8-buster

WORKDIR /module4_task2
COPY ./ /module4_task2

RUN apt-get update && apt-get install -y --no-install-recommends \
    npm \
    wget \
    curl \
    zip \
    && rm -rf /var/lib/apt/lists/*

RUN wget --progress=dot:giga https://github.com/gohugoio/hugo/releases/download/v0.79.0/hugo_0.79.0_Linux-64bit.deb \
    && dpkg -i hugo_0.79.0_Linux-64bit.deb \
    && rm hugo_0.79.0_Linux-64bit.deb

RUN curl -s https://deb.nodesource.com/setup_16.x | bash \
    && apt-get install -y nodejs \
    && rm -rf /var/lib/apt/lists/*

RUN npm install -g markdownlint-cli

ENTRYPOINT ["/module4_task2/ls"]
