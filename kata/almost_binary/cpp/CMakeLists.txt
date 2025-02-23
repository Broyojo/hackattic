set(PROJECT_NAME "almost_binary")

cmake_minimum_required(VERSION 3.15)
project(${PROJECT_NAME})

set(SOURCES
    src/main.cpp
)

add_executable(${PROJECT_NAME} ${SOURCES})

target_include_directories(${PROJECT_NAME} PUBLIC include)
