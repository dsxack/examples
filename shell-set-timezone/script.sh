#!/usr/bin/env bash

echo "Europe/Moscow" | sudo tee /etc/timezone
sudo dpkg-reconfigure -f noninteractive tzdata
