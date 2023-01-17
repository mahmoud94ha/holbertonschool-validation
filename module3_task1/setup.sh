#!/bin/bash
sudo apt-get update
sudo apt-get install -y make wget
sudo wget https://github.com/gohugoio/hugo/releases/download/v0.109.0/hugo_extended_0.109.0_linux-amd64.deb
sudo dpkg -i hugo_extended_0.109.0_linux-amd64.deb