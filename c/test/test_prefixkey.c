// NumKey
//
// test_prefixkey.c
//
// @category   Tools
// @author     Nicola Asuni
// @license    see LICENSE file
// @link       https://github.com/Vonage/numkey

// Test for prefixey

#include <inttypes.h>
#include <stdio.h>
#include <string.h>
#include <time.h>
#include "../src/numkey/prefixkey.h"


static const int k_prefixkey_test_size = 14;

typedef struct test_prefixkey_data_t
{
    const char* number;
    uint64_t    pk;
} test_prefixkey_data_t;

static const test_prefixkey_data_t test_prefixkey_data[] =
{
    {"", 0},
    {"0", 0},
    {"00", 0},
    {"00000000000000", 0},
    {"000000000000000", 0},
    {"0000000000000000", 0},
    {"000000000000001", 1},
    {"0000000000000019", 1},
    {"1", 100000000000000L},
    {"10", 100000000000000L},
    {"10000000000000", 100000000000000L},
    {"100000000000000", 100000000000000L},
    {"1000000000000000", 100000000000000L},
    {"999999999999999", 999999999999999L},
};

// returns current time in nanoseconds
uint64_t get_time()
{
    struct timespec t;
    (void) timespec_get(&t, TIME_UTC);
    return (((uint64_t)t.tv_sec * 1000000000) + (uint64_t)t.tv_nsec);
}

int test_prefixkey()
{
    int errors = 0;
    int i = 0;
    uint64_t pk = 0;
    for (i=0 ; i < k_prefixkey_test_size; i++)
    {
        pk = prefixkey(test_prefixkey_data[i].number, strlen(test_prefixkey_data[i].number));
        if (pk != test_prefixkey_data[i].pk)
        {
            (void) fprintf(stderr, "%s (%d): Unexpected prefixkey: expected 0x%016" PRIx64 ", got 0x%016" PRIx64 "\n", __func__, i, test_prefixkey_data[i].pk, pk);
            ++errors;
        }
    }
    return errors;
}

void benchmark_prefixkey()
{
    uint64_t tstart = 0, tend = 0;
    int i = 0;
    int size = 100000;
    tstart = get_time();
    for (i=0 ; i < size; i++)
    {
        prefixkey("123456789012345", 15);
    }
    tend = get_time();
    (void) fprintf(stdout, " * %s : %lu ns/op\n", __func__, (tend - tstart)/size);
}

int main()
{
    static int errors = 0;

    errors += test_prefixkey();

    benchmark_prefixkey();

    return errors;
}
