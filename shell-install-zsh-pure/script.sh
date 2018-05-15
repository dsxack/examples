#!/usr/bin/env bash

npm install --global pure-prompt
cat ~/.zshrc | sed s/ZSH_THEME.*/ZSH_THEME=\"pure\"/ > ~/.zshrc
