# cmake_minimum_required(VERSION 3.5)
# project(org_chart_test CXX)

# # Find Drogon test framework
# find_package(Drogon REQUIRED)

# # Add the source files for the controllers, models, utils, filters, and plugins
# set(PROJECT_SOURCES
#     ${CMAKE_SOURCE_DIR}/controllers/DepartmentsController.cc
#     ${CMAKE_SOURCE_DIR}/models/Department.cc
#     ${CMAKE_SOURCE_DIR}/models/Person.cc
#     ${CMAKE_SOURCE_DIR}/models/Job.cc
#     ${CMAKE_SOURCE_DIR}/utils/utils.cc
#     ${CMAKE_SOURCE_DIR}/filters/LoginFilter.cc
#     ${CMAKE_SOURCE_DIR}/plugins/Jwt.cc
#     ${CMAKE_SOURCE_DIR}/plugins/JwtPlugin.cc
# )

# # Add the test executable
# add_executable(${PROJECT_NAME} test_main.cc test_controllers.cc ${PROJECT_SOURCES})

# # Include directories for the test executable
# target_include_directories(${PROJECT_NAME} PRIVATE
#     ${CMAKE_SOURCE_DIR}
#     ${CMAKE_SOURCE_DIR}/third_party/drogon/lib/inc
#     ${CMAKE_SOURCE_DIR}/third_party/drogon/orm_lib/inc
#     ${CMAKE_SOURCE_DIR}/third_party/jwt-cpp/include
#     ${CMAKE_SOURCE_DIR}/controllers
#     ${CMAKE_SOURCE_DIR}/models
#     ${CMAKE_SOURCE_DIR}/utils
#     ${CMAKE_SOURCE_DIR}/filters
#     ${CMAKE_SOURCE_DIR}/plugins
# )

# # Link against Drogon libraries
# target_link_libraries(${PROJECT_NAME} PRIVATE drogon)

# # Enable test discovery for CTest
# enable_testing()
# add_test(NAME ${PROJECT_NAME} COMMAND ${PROJECT_NAME})

# # If you're using Drogon's ParseAndAddDrogonTests function
# # include(${CMAKE_SOURCE_DIR}/third_party/drogon/cmake_modules/DrogonTest.cmake)
# ParseAndAddDrogonTests(${PROJECT_NAME})

# Current Working
# =======================================================
cmake_minimum_required(VERSION 3.5)
project(org_chart_test CXX)

# Find Drogon test framework
find_package(Drogon REQUIRED)

# --- Automatic Source File Discovery ---
# Use file(GLOB_RECURSE) to automatically find all application source files.
# This prevents linker errors when you add new files to the project.
# ${CMAKE_SOURCE_DIR} refers to the root directory of the project.
file(GLOB_RECURSE APP_SOURCES
    "${CMAKE_SOURCE_DIR}/controllers/*.cc"
    "${CMAKE_SOURCE_DIR}/models/*.cc"
    "${CMAKE_SOURCE_DIR}/filters/*.cc"
    "${CMAKE_SOURCE_DIR}/plugins/*.cc"
    "${CMAKE_SOURCE_DIR}/utils/*.cc"
)

# --- Test Executable ---
# Add the test executable, including its own files and all the discovered
# application source files.
add_executable(${PROJECT_NAME}
    test_main.cc
    test_controllers.cc
    ${APP_SOURCES}
)

# --- Include Directories ---
# Add all directories where header files might be located.
target_include_directories(${PROJECT_NAME} PRIVATE
    ${CMAKE_SOURCE_DIR}
    ${CMAKE_SOURCE_DIR}/third_party/drogon/lib/inc
    ${CMAKE_SOURCE_DIR}/third_party/drogon/orm_lib/inc
    ${CMAKE_SOURCE_DIR}/third_party/jwt-cpp/include
    ${CMAKE_SOURCE_DIR}/controllers
    ${CMAKE_SOURCE_DIR}/models
    ${CMAKE_SOURCE_DIR}/utils
    ${CMAKE_SOURCE_DIR}/filters
    ${CMAKE_SOURCE_DIR}/plugins
)

# --- Link Libraries ---
# Link against Drogon and the bcrypt library for password hashing.
target_link_libraries(${PROJECT_NAME} PRIVATE drogon bcrypt)

# --- Test Setup ---
# Enable test discovery for CTest
enable_testing()
add_test(NAME ${PROJECT_NAME} COMMAND ${PROJECT_NAME})
ParseAndAddDrogonTests(${PROJECT_NAME})