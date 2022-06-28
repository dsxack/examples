# Examples

Many different pieces of code

- [Examples](#examples)
  * [Administration](#administration)
    + [Ansible](#ansible)
    + [Docker](#docker)
    + [Scripts](#scripts)
  * [Languages](#languages)
    + [C language](#c-language)
    + [Go language](#go-language)
    + [Java language](#java-language)
    + [JavaScript language](#javascript-language)
    + [PHP language](#php-language)
    + [Python language](#python-language)
    + [Rust language](#rust-language)
    + [Inter language communication](#inter-language-communication)
  * [Algorithms](#algorithms)
    + [Binary search](#binary-search)
    + [Binary tree](#binary-tree)
    + [Extensible array](#extensible-array)
    + [Hashmap](#hashmap)
    + [Linked lists](#linked-lists)
    + [Markov chain](#markov-chain)
    + [Quick sort](#quick-sort)
    + [Fibonacci](#fibonacci)

## Administration

### Ansible

* [ansible-autoupdate](ansible-autoupdate) - Ansible role which adds daily updates cron job
* [ansible-elastic-logstash-kibana](ansible-elastic-logstash-kibanaansible-mariadb) - Ansible role which deploy ELK stack docker containers
* [ansible-mariadb](ansible-mariadb]) - Ansible role which deploy mariadb docker container
* [ansible-minecraft](ansible-minecraft) - Ansible role which deploy minecraft server docker container
* [ansible-oh-my-zsh](ansible-oh-my-zsh) - Ansible role which install oh-my-zsh for current user
* [ansible-openvpn](ansible-openvpn) - Ansible role which deploy openvpn docker container
* [ansible-portainer](ansible-portainer) - Ansible role which deploy portainer docker container
* [ansible-prometheus-grafana](ansible-prometheus-grafana) - Ansible role which deploy prometheus and grafana docker containers
* [ansible-registry](ansible-registry) - Ansible role which install docker registry docker container
* [ansible-set-locale](ansible-set-locale) - Ansible role which sets english locale on the server
* [ansible-set-timezone](ansible-set-timezone) - Ansible role which sets Europe/Moscow timezone on the server
* [ansible-socks5](ansible-socks5) - Ansible role which deploy socks5 proxy docker container
* [ansible-swap](ansible-swap) - Ansible role which sets swap on the server
* [ansible-traefik](ansible-traefik) - Ansible role which deploy traefik docker container
* [ansible-whoami](ansible-whoami) - Ansible role which deploy whoami docker container

### Docker

* [docker-compose-apache-php](docker-compose-apache-php) - Docker compose project with apache + php services
* [docker-compose-php70](docker-compose-php70) - Docker compose project with php-fpm service
* [docker-golang](docker-golang) - Dockerfile for golang
* [docker-nginx-consul](docker-nginx-consul) - Nginx+consul+consultemplate docker image
    
### Scripts 

* [shell-install-docker](shell-install-docker) - Shell script which installs docker
* [shell-install-nvm](shell-install-nvm) - Shell script which installs NVM
* [shell-install-oh-my-zsh](shell-install-oh-my-zsh) - Shell script which installs Oh-my-zsh
* [shell-install-zsh-pure](shell-install-zsh-pure) - Shell script which installs zsh-pure
* [shell-set-daily-updates](shell-set-daily-updates) - Shell script which adds daily updates cron job
* [shell-set-swap](shell-set-swap) - Shell script which installs swap
* [shell-set-timezone](shell-set-timezone) - Shell script which sets Europe/Moscow timezone

## Languages

### C language
    
* [c-bsearch](c-bsearch) - Binary search algorithm written in C language
* [c-btree](c-btree) - Binary tree structure written in C language
* [c-extensible-array](c-extensible-array) - Dynamicaly extensible array written in C language
* [c-hashmap](c-hashmap) - HashMap structure written in C language
* [c-linked-lists](c-linked-lists) - Linked list structure written in C language
* [c-markov-chain](c-markov-chain) - Markov chain algorithm written in C language
* [c-qsort](c-qsort) - Quick sort algorithm written in C language
    
### Go language

* [go-benchmark](go-benchmark) - Benchmark of sum in many goroutines written in Go
* [go-cgo-benchmark](go-cgo-benchmark) - Fibonacci processing benchmark with cgo and without cgo
* [go-cgo-php-extension](go-cgo-php-extension) - PHP extension written in Go 
* [go-cgo-python-cdll](go-cgo-python-cdll) - Python shared library written in Go loading 
* [go-cgo-python-extension](go-cgo-python-extension) - Python extension written in Go
* [go-cgo-shared-lib](go-cgo-shared-lib) - Shared library written in Go
* [go-code-generation](go-code-generation) - Code generalization with code generation
* [go-consul](go-consul) - Consul communication with Go
* [go-http-echo-over-tcp](go-http-echo-over-tcp) - HTTP echo server over tcp written in Go 
* [go-http-echo-over-tcp-thread-pooled](go-http-echo-over-tcp-thread-pooled) - HTTP echo server over tcp written in Go 
    with thread pool
* [go-http-mirror](go-http-mirror) - HTTP mirror server written in Go 
* [go-http-pipe](go-http-pipe) - HTTP pipe server written in Go
* [go-philosopher-dinner](go-philosopher-dinner) - Philosopher dinner app from Rustbook written in Go
* [go-plus-rust](go-plus-rust) - Go plus Rust languages intercommunication
* [go-plus-scrypt-rust-c](go-plus-scrypt-rust-c) - Rust scrypt bindings for go
* [go-tcp-echo](go-tcp-echo) - TCP echo server written in Go
* [go-telegrambot](go-telegrambot) - Telegram bot written in Go
* [go-webassembly](go-webassembly) - Webassembly app written in Go
* [go-websocket](go-websocket) - Websocket app written in Go
* [go-x509-certs](go-x509-certs) - x509 certs manipulation with Go
* [go-exec-from-memory](go-exec-from-memory) - Running binary app from memory without saving on the disk
* [go-https-attack-proxy](go-https-attack-proxy) - HTTPS/HTTP attack proxy that sniffing requests and responses
* [go-plugins-cgo-with-same-symbol-names](go-plugins-cgo-with-same-symbol-names) - Example of using go plugins that statically linked C libraries with CGO with same symbol names
* [go-net-with-custom-codec](go-net-with-custom-codec) – Example of usage gonet library with custom codec

### Java language

* [java-http-echo-over-tcp](java-http-echo-over-tcp) - HTTP echo server over tcp written in Java
* [java-telegrambot](java-telegrambot) - Telegram bot written in Java
    
### JavaScript language

* [js-gulp-plugin](js-gulp-plugin) - Gulp plugin template

### PHP language

* [php-chmod777](php-chmod777) - Recursive files changing mod PHP script

### Python language 

* [python-http-echo-over-tcp](python-http-echo-over-tcp) - HTTP echo server over tcp written in Python 

### Rust language 
   
* [rust-benchmark](rust-benchmark) - Benchmark of sum in many threads written in Rust
* [rust-convert-temperatures](rust-convert-temperatures) - Converting from celsius to fahrenheit in Rust 
* [rust-fibonacci](rust-fibonacci) - Fibonacci algorithm written in Rust language
* [rust-http-echo-over-tcp](rust-http-echo-over-tcp) - HTTP echo server over tcp written in Rust
* [rust-http-echo-over-tcp-with-tokio](rust-http-echo-over-tcp-with-tokio) - HTTP echo server over tcp written in Rust
    with Tokio library
* [rust-http-mirror](rust-http-mirror) - HTTP mirror server written in Rust
* [rust-http-simple-rest](rust-http-simple-rest) - HTTP simple REST written in Rust
* [rust-philosopher-dinner](rust-philosopher-dinner) - Philosopher dinner app from Rustbook written in Rust
* [rust-secret-number](rust-secret-number) - Secret number app from Rustbook written in Rust
* [rust-telegrambot](rust-telegrambot) - Telegram bot written in Rust
* [rust-time](rust-time) - Time manipulation in Rust
* [rust-embed](rust-embed) - Shared library written in Rust

### Inter language communication

* [go-cgo-python-cdll](go-cgo-python-cdll) - Python shared library written in Go loading 
* [go-cgo-python-extension](go-cgo-python-extension) - Python extension written in Go
* [go-cgo-shared-lib](go-cgo-shared-lib) - Shared library written in Go
* [go-cgo-php-extension](go-cgo-php-extension) - PHP extension written in Go 
* [go-plus-rust](go-plus-rust) - Go plus Rust languages intercommunication
* [rust-embed](rust-embed) - Shared library written in Rust 
 
## Algorithms

### Binary search

* [c-bsearch](c-bsearch) - Binary search algorithm written in C language

### Binary tree

* [c-btree](c-btree) - Binary tree structure written in C language
    
### Extensible array
    
* [c-extensible-array](c-extensible-array) - Dynamicaly extensible array written in C language
    
### Hashmap

* [c-hashmap](c-hashmap) - HashMap structure written in C language
    
### Linked lists

* [c-linked-lists](c-linked-lists) - Linked list structure written in C language
    
### Markov chain

* [c-markov-chain](c-markov-chain) - Markov chain algorithm written in C language
    
### Quick sort

* [c-qsort](c-qsort) - Quick sort algorithm written in C language

### Fibonacci

* [rust-fibonacci](rust-fibonacci) - Fibonacci algorithm written in Rust language
