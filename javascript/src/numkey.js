/** NumKey Javascript Library
 * 
 * numkey.js
 * 
 * @category   Libraries
 * @author     Nicola Asuni <nicola.asuni@vonage.com>
 * @copyright  2019-2020 Vonage
 * @license    see LICENSE file
 * @link       https://github.com/nexmoinc/numkey
 */

// NOTE: Javascript numbers are 64 bit floats with a 53 bit precision.

var NKBMASK_COUNTRY_FL = 0xF8000000 //!< Bit mask for the ISO 3166 alpha-2 country code first letter  [ 11111000 00000000 00000000 00000000 ]
var NKBMASK_COUNTRY_SL = 0x07C00000 //!< Bit mask for the ISO 3166 alpha-2 country code second letter [ 00000111 11000000 00000000 00000000 ]

var NKBMASK_NUMBER_HI = 0x003FFFFF //!< Bit mask for the hi part of short code or E.164 number (max 15 digits)  [ 00000000 00111111 11111111 11111111 ]
var NKBMASK_NUMBER_LO = 0xFFFFFFF0 //!< Bit mask for the lo part of short code or E.164 number (max 15 digits)  [ 11111111 11111111 11111111 11110000 ]

var NKBMASK_32 = 0xFFFFFFFF //!< 32-Bit mask [ 11111111 11111111 11111111 11111111 ]

var NKBMASK_LENGTH = 0x0000000F //!< Bit mask for the number length [ 00000000 00000000 00000000 00001111 ]

var NKBSHIFT_COUNTRY_FL = 27 //!< COUNTRY first letter LSB position from the NumKey LSB
var NKBSHIFT_COUNTRY_SL = 22 //!< COUNTRY second letter LSB position from the NumKey LSB

var NKBSHIFT_NUMBER_HI = 10 //!< NUMBER MSB position from the NumKey MSB
var NKBSHIFT_NUMBER_LO = 4 //!< NUMBER LSB position from the NumKey LSB

var NKSLENGTH_COUNTRY = 3 //!< Number of characters in the country code + NULL terminator
var NKSLENGTH_NUMBER = 16 //!< Number of characters in the number code + NULL terminator

var NKCSHIFT_CHAR = 64 //!< Value shift to encode characters to numbers (A=1, ..., Z=26)

var NKNUMDIV = 0x10000000 //!< = [2^28] divider to get the number HI bits
var NKNUMMUL = 0x100000000 //!< = [2^32] multiplier to set the number HI bits

var NKZEROSHIFT = 48 //!< ASCII code of the '0' character

var NKNUMMAXLEN = 15 //!< Maximum number length for E.164 and key reversibility

var PKNUMMAXLEN = 15 //!< Maximum number of digits to store for the prefixkey.

function encodeChar(c) {
    return ((c - NKCSHIFT_CHAR) >>> 0)
}

function encodeCountry(country) {
    return (((encodeChar(country.charCodeAt(0)) << NKBSHIFT_COUNTRY_FL) | (encodeChar(country.charCodeAt(1)) << NKBSHIFT_COUNTRY_SL)) >>> 0);
}

function decodeCountry(nk) {
    return String.fromCharCode((((nk.hi & NKBMASK_COUNTRY_FL) >>> NKBSHIFT_COUNTRY_FL) + NKCSHIFT_CHAR), (((nk.hi & NKBMASK_COUNTRY_SL) >>> NKBSHIFT_COUNTRY_SL) + NKCSHIFT_CHAR));
}

function encodeNumber(number) {
    var size = number.length;
    if (size < 1) {
        return {
            "hi": 0,
            "lo": 0
        };
    }
    var num = 0; // 53 bit precision
    var b;
    var i, j = 0;
    var len = size;
    if (size > NKNUMMAXLEN) {
        j = (size - NKNUMMAXLEN); // last 15 digits
        len = 0; // flag non-revesible encoding
    }
    for (i = j; i < size; i++) {
        b = (number.charCodeAt(i) - NKZEROSHIFT);
        num = (num * 10) + b;
    }
    return {
        "hi": ((((num / NKNUMDIV) >>> 0) & NKBMASK_NUMBER_HI) >>> 0),
        "lo": ((((num << NKBSHIFT_NUMBER_LO) >>> 0) | ((len & NKBMASK_LENGTH) >>> 0)) >>> 0)
    };
}

function decodeNumber(nk) {
    var size = (nk.lo & NKBMASK_LENGTH) >>> 0;
    if (size == 0) {
        return "";
    }
    var number = new Array(size);
    var numhi = ((nk.hi & NKBMASK_NUMBER_HI) >>> 0);
    var numlo = ((nk.lo & NKBMASK_NUMBER_LO) >>> 0);
    var num = ((numhi >>> NKBSHIFT_NUMBER_LO) * NKNUMMUL) + (((numhi << (32 - NKBSHIFT_NUMBER_LO)) >>> 0) + (numlo >>> NKBSHIFT_NUMBER_LO));
    var rem = 0;
    var i;
    for (i = (size - 1); i >= 0; i--) {
        rem = ((num % 10) >>> 0);
        number[i] = String.fromCharCode(rem + NKZEROSHIFT);
        num = (num / 10);
    }
    return number.join("");
}

function numKey(country, number) {
    var ne = encodeNumber(number);
    return {
        "hi": ((encodeCountry(country) | ne.hi) >>> 0),
        "lo": ne.lo
    };
}

function decodeNumKey(nk) {
    return {
        "country": decodeCountry(nk),
        "number": decodeNumber(nk)
    };
}

function compare(a, b) {
    return ((a < b) ? -1 : ((a > b) ? 1 : 0));
}

function compareNumKeyCountry(nka, nkb) {
    return compare((nka.hi >>> NKBSHIFT_COUNTRY_SL), (nkb.hi >>> NKBSHIFT_COUNTRY_SL));
}

function padL08(s) {
    return ("00000000" + s).slice(-8);
}

function numKeyString(nk) {
    return padL08(nk.hi.toString(16)) + padL08(nk.lo.toString(16));
}

function parseHex(ns) {
    return {
        "hi": parseInt(ns.substring(0, 8), 16) >>> 0,
        "lo": parseInt(ns.substring(8, 16), 16) >>> 0,
    };
}

function prefixKey(number) {
    var size = number.length;
    if (size < 1) {
        return {
            "hi": 0,
            "lo": 0
        };
    }
    var num = 0;
    var b;
    var i = 0;
    if (size > PKNUMMAXLEN) {
        size = PKNUMMAXLEN; // truncate number
    }
    for (i = 0; i < size; i++) {
        b = (number.charCodeAt(i) - NKZEROSHIFT);
        num = (num * 10) + b;
    }
    for (i = size; i < PKNUMMAXLEN; i++) {
        num = (num * 10); // zero right-padding
    }
    return {
        "hi": ((((num / NKNUMMUL) >>> 0) & NKBMASK_32) >>> 0),
        "lo": ((num & NKBMASK_32) >>> 0)
    };
}

function countryKey(country) {
    return (country.charCodeAt(0) << 8) | (country.charCodeAt(1));
}

function decodeCountryKey(ck) {
    return String.fromCharCode((ck & 0xFF00) >> 8, (ck & 0x00FF));
}

if (typeof(module) !== 'undefined') {
    module.exports = {
        numKey: numKey,
        decodeNumKey: decodeNumKey,
        compareNumKeyCountry: compareNumKeyCountry,
        numKeyString: numKeyString,
        parseHex: parseHex,
        prefixKey: prefixKey,
        countryKey: countryKey,
        decodeCountryKey: decodeCountryKey,
    }
}