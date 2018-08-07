cmake_minimum_required(VERSION 3.12)
project(cpp_qsort)

set(CMAKE_CXX_STANDARD 14)

include_directories(.)

add_executable(c_qsort
        main.c)
