// Usage example for NumKey
// https://github.com/vonage/numkey

#include <inttypes.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../src/numkey/numkey.h"

int main()
{
    // BASIC NUMKEY FUNCTIONS
    // --------------------------

    numkey_t dnk = {0};
    decode_numkey(0xc2ab5e44f21a947f, &dnk);
    fprintf(stdout, "%s %s\n", dnk.country, dnk.number);
    // XJ 762942138198343

    uint64_t nk = numkey("XJ", "762942138198343", 15);
    fprintf(stdout, "%" PRIx64 "\n", nk);
    // c2ab5e44f21a947f

    int cmp = compare_numkey_country(0xd6a23089b8e15cdf, 0xd6a2300000000000);
    fprintf(stdout, "%d\n", cmp);
    // 0

    char ns[17] = "";
    numkey_hex(0xc2ab5e44f21a947f, ns);
    fprintf(stdout, "%s\n", ns);
    // c2ab5e44f21a947f

    nk = parse_numkey_hex("c2ab5e44f21a947f");
    fprintf(stdout, "%" PRIx64 "\n", nk);
    // c2ab5e44f21a947f

    // ============================================================================

    return 0;
}
