project("main")
include_directories(${CMAKE_SOURCE_DIR}/../../include)
add_executable(main main.c)
target_link_libraries(main ${CMAKE_SOURCE_DIR}/../../lib/librustscryptbindings.a)
