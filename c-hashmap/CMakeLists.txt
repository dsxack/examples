cmake_minimum_required(VERSION 3.12)
project(c_hashmap C)

set(CMAKE_C_STANDARD 99)

add_executable(c_hashmap main.c eprintf.c eprintf.h)