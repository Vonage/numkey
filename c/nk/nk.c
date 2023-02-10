// NumKey Fast Encoder Command Line Application
//
// nk.c
//
// @category   Tools
// @author     Nicola Asuni <nicola.asuni@vonage.com>
// @copyright  2019 Vonage
// @license    see LICENSE file
// @link       https://github.com/vonage/numkey


#include <inttypes.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../src/numkey/numkey.h"

#ifndef VERSION
#define VERSION "0.0.0-0"
#endif

int main(int argc, char *argv[])
{
    if (argc != 3)
    {
        fprintf(stderr, "NumKey Encoder %s\nUsage: nk COUNTRY NUMBER\n", VERSION);
        return 1;
    }
    fprintf(stdout, "%016" PRIx64, numkey(argv[1], argv[2], strlen(argv[2])));
}
