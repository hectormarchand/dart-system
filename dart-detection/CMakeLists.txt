cmake_minimum_required(VERSION 3.28.3)
project(dart-detection)

# Find OpenCV using pkg-config
find_package(PkgConfig REQUIRED)
find_package(CURL REQUIRED)
pkg_check_modules(OPENCV REQUIRED opencv4)

include_directories(${OPENCV_INCLUDE_DIRS} ${CURL_INCLUDE_DIR} include)
link_directories(${OPENCV_LIBRARY_DIRS})
add_definitions(${OPENCV_CFLAGS_OTHER})

set(CMAKE_CXX_STANDARD 17)

# Glob all .cpp files in src/
file(GLOB SRC_FILES src/*.cpp)

# Add executable with all source files
add_executable(dart_detection ${SRC_FILES})

target_link_libraries(dart_detection ${OPENCV_LIBRARIES} ${CURL_LIBRARIES})
