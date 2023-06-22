FROM golang:1.15.8-alpine

# Install wget and container-structure-test
RUN apk update && \
    apk add --no-cache wget && \
    wget -O /usr/local/bin/container-structure-test https://storage.googleapis.com/container-structure-test/latest/container-structure-test-linux-amd64 && \
    chmod +x /usr/local/bin/container-structure-test

WORKDIR /module4_task2
COPY . .

# Set the command to execute the markdownlint command
CMD ["markdownlint", "-f", "DEPLOY.md"]
