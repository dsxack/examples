#!/usr/bin/env bash

fallocate -l 4G /swapfile
chmod 600 /swapfile
mkswap /swapfile
swapon /swapfile
echo "/swapfile   none    swap    sw    0   0" | sudo tee -a /etc/fstab
sysctl vm.swappiness=10
echo "vm.swappiness=10" | sudo tee -a /etc/sysctl.conf
sysctl vm.vfs_cache_pressure=50
echo "vm.vfs_cache_pressure=50" | tee -a /etc/sysctl.conf
