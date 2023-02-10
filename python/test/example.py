#!/usr/bin/env python3

# Usage example for NumKey Python wrapper
# https://github.com/vonage/numkey

import numkey as nk


# BASIC NUMKEY FUNCTIONS
# --------------------------

nk.decode_numkey(14027409114588157055)
# (b'XJ', b'762942138198343')

nk.numkey(country=b'XJ', number=b'762942138198343')
# 14027409114588157055

nk.compare_numkey_country(0xd6a23089b8e15cdf, 0xd6a2300000000000)
# 0

nk.numkey_hex(14027409114588157055)
# b'c2ab5e44f21a947f'

nk.parse_numkey_hex(b'c2ab5e44f21a947f')
# 14027409114588157055
