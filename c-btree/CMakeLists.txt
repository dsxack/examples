cmake_minimum_required(VERSION 3.12)
project(c_btree C)

set(CMAKE_C_STANDARD 99)

add_executable(c_btree main.c eprintf.c eprintf.h)