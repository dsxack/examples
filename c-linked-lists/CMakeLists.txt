cmake_minimum_required(VERSION 3.12)
project(c_linked_lists C)

set(CMAKE_C_STANDARD 99)

add_executable(c_linked_lists main.c eprintf.c eprintf.h)