# Tests

#set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}/lib)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/lib)
set(RUNTIME_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/lib)

# Add the binary tree directory to the search path for linking and include files
link_directories(${PROJECT_BINARY_DIR}/src/numkey)
include_directories(${CMAKE_CURRENT_BINARY_DIR} ${PROJECT_BINARY_DIR}/src/numkey)

#define macro to simplify adding tests
macro (do_test arg)
  add_test (${arg} ${TARGET_SYSTEM_EMULATOR} ${arg}${CMAKE_EXECUTABLE_SUFFIX})
endmacro ()

# create a smoke test
function(SMOKE_TEST test_name test_file dependencies)
  add_executable(${test_name} ${test_file})
  # Link the executable to the libraries
  target_link_libraries(${test_name} ${dependencies})
  # run test
  do_test (${test_name})
endfunction(SMOKE_TEST)

file(GLOB TEST_BIN_FILES "data/*.bin")
file (COPY ${TEST_BIN_FILES} DESTINATION ${CMAKE_CURRENT_BINARY_DIR})

SMOKE_TEST (test_binsearch test_binsearch.c numkey)
SMOKE_TEST (test_binsearch_col test_binsearch_col.c numkey)
SMOKE_TEST (test_binsearch_file test_binsearch_file.c numkey)
SMOKE_TEST (test_hex test_hex.c numkey)
SMOKE_TEST (test_set test_set.c numkey)
SMOKE_TEST (test_example test_example.c numkey)
SMOKE_TEST (test_test_numkey test_numkey.c numkey)
SMOKE_TEST (test_test_prefixkey test_prefixkey.c numkey)
SMOKE_TEST (test_test_countrykey test_countrykey.c numkey)
