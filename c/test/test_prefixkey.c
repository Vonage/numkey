// NumKey
//
// test_prefixkey.c
//
// @category   Tools
// @author     Nicola Asuni <nicola.asuni@vonage.com>
// @copyright  2019-2022 Vonage
// @license    see LICENSE file
// @link       https://github.com/Vonage/numkey

// Test for prefixey

#if __STDC_VERSION__ >= 199901L
#define _XOPEN_SOURCE 600
#else
#define _XOPEN_SOURCE 500
#endif

#include <stdio.h>
#include <inttypes.h>
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
    clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &t);
    return (((uint64_t)t.tv_sec * 1000000000) + (uint64_t)t.tv_nsec);
}

int test_prefixkey()
{
    int errors = 0;
    int i;
    uint64_t pk;
    for (i=0 ; i < k_prefixkey_test_size; i++)
    {
        pk = prefixkey(test_prefixkey_data[i].number, strlen(test_prefixkey_data[i].number));
        if (pk != test_prefixkey_data[i].pk)
        {
            fprintf(stderr, "%s (%d): Unexpected prefixkey: expected 0x%016" PRIx64 ", got 0x%016" PRIx64 "\n", __func__, i, test_prefixkey_data[i].pk, pk);
            ++errors;
        }
    }
    return errors;
}

void benchmark_prefixkey()
{
    uint64_t tstart, tend;
    int i;
    int size = 100000;
    tstart = get_time();
    for (i=0 ; i < size; i++)
    {
        prefixkey("123456789012345", 15);
    }
    tend = get_time();
    fprintf(stdout, " * %s : %lu ns/op\n", __func__, (tend - tstart)/size);
}

int main()
{
    static int errors = 0;

    errors += test_prefixkey();

    benchmark_prefixkey();

    return errors;
}
