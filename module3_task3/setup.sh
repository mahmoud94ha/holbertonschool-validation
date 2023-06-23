#!/bin/bash

sudo apt-get update && sudo apt-get install -y make git

#download hugo
sudo curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
sudo apt install ./hugo.deb
#clean
sudo rm hugo.deb

# linter
sudo go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
sudo npm install -g markdownlint-cli
sudo npm install -g markdown-link-check
