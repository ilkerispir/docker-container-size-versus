cmake_minimum_required(VERSION 3.10)

# Define the project name and where to report bugs.
set(PACKAGE_BUGREPORT
    "https://github.com/GoogleCloudPlatform/cpp-samples/issues")
project(cpp-samples-cloud-run-hello-world CXX C)

# Configure the Compiler options, we will be using C++17 features.
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(Boost 1.66 REQUIRED COMPONENTS program_options)
find_package(Threads)

add_executable(main main.cc)
target_link_libraries(
  main PRIVATE Boost::headers Boost::program_options
                          Threads::Threads)

include(GNUInstallDirs)
install(TARGETS main RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR})