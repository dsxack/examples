#!/usr/bin/env bash

echo "0  1    * * *   root apt-get upgrade -y 2>&1 >> /root/apt-get-upgrade.log" | sudo tee -a /etc/crontab
sudo service cron reload
