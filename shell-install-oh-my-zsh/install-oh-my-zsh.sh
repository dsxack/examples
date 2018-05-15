#!/usr/bin/env bash

sudo apt update
sudo apt install -y zsh
sh -c "\$(wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)"
