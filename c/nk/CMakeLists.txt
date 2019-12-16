set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/lib)
set(RUNTIME_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/cmd)

# Add the binary tree directory to the search path for linking and include files
link_directories(${PROJECT_BINARY_DIR}/src/numkey)
include_directories(${CMAKE_CURRENT_BINARY_DIR} ${PROJECT_BINARY_DIR}/src/numkey)

file(COPY DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
add_executable(nk nk.c)
target_link_libraries(nk numkey)

# --- PACKAGING ---

install(TARGETS "nk" DESTINATION "bin" COMPONENT "nk")
install(FILES "../../README.md" DESTINATION "share/doc/nk" COMPONENT "nk")
install(FILES "../../VERSION" DESTINATION "share/doc/nk" COMPONENT "nk")
install(FILES "../../RELEASE" DESTINATION "share/doc/nk" COMPONENT "nk")
install(FILES "../../LICENSE" DESTINATION "share/doc/nk" COMPONENT "nk" RENAME copyright)

set(CPACK_PACKAGE_NAME "nk")
set(CPACK_PACKAGE_DESCRIPTION "NumKey CLI encoder")
set(CPACK_PACKAGE_DESCRIPTION_SUMMARY "NumKey Fast Encoder Command Line Application")

include(CPack)
