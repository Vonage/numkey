/*
 * @(#)NumKey.java
 */

package com.nexmo.numkey;

public class NumKey {

    private final static long NKBMASK_NUMBER = Long.parseUnsignedLong("003FFFFFFFFFFFF0", 16);  //!< Bit mask for the short code or E.164 number (max 15 digits)  [ 00000000 00111111 11111111 11111111 11111111 11111111 11111111 11110000 ].
    private final static long NKBMASK_LENGTH = 0x0F; //!< Bit mask for the number length [ 00001111 ].
    private final static byte NKBMASK_COUNTRY = 0x1F; //!< Bit mask for country five bit [ 00011111 ].
    private final static byte NKBSHIFT_COUNTRY_FL = 59; //!< COUNTRY first letter LSB position from the NumKey LSB.
    private final static byte NKBSHIFT_COUNTRY_SL = 54; //!< COUNTRY second letter LSB position from the NumKey LSB.
    private final static byte NKBSHIFT_NUMBER = 4; //!< NUMBER LSB position from the NumKey LSB.
    private final static byte NKCSHIFT_CHAR = 64; //!< Value shift to encode characters to numbers (A=1, ..., Z=26).
    private final static byte NKNUMMAXLEN = 15; //!< Maximum number length for E.164 and key reversibility.
    private final static byte PKNUMMAXLEN = 15; //!< Maximum number of digits to store for the prefixkey.

    public static class NumData {
        public byte[] country;
        public byte[] number;
    }

    public static class NumDataStr {
        public String country;
        public String number;
    }

    /**
     * Encodes the input byte character to a numeric value.
     * NOTE: This is safe to be used only with A-Z characters.
     *
     * @param chr Character to encode.
     * @return Encoded character.
     */
    private static long encodeChar(byte chr) {
        return ((long) chr - NKCSHIFT_CHAR);
    }

    /**
     * Encode country code.
     *
     * @param country ISO 3166 alpha-2 country code.
     * @return Encoded country code.
     */
    private static long encodeCountry(byte[] country) {
        return ((encodeChar(country[0]) << NKBSHIFT_COUNTRY_FL) | (encodeChar(country[1]) << NKBSHIFT_COUNTRY_SL));
    }

    /**
     * Decode country code into 2-char byte array.
     *
     * @param nk NumKey.
     * @return decoded ISO 3166 alpha-2 country code.
     */
    private static byte[] decodeCountry(long nk) {
        return new byte[]{
                (byte) (((nk >>> NKBSHIFT_COUNTRY_FL) & NKBMASK_COUNTRY) + NKCSHIFT_CHAR),
                (byte) (((nk >>> NKBSHIFT_COUNTRY_SL) & NKBMASK_COUNTRY) + NKCSHIFT_CHAR)
        };
    }

    /**
     * Encode number string.
     *
     * @param number byte array containing the Short code or E.164 LVN number.
     * @return Encoded number.
     */
    private static long encodeNumber(byte[] number) {
        long num = 0;
        int b;
        int i, j = 0;
        int len = number.length;
        if (len > NKNUMMAXLEN) {
            j = (len - NKNUMMAXLEN); // last 15 digits
            len = 0;                 // flag non-revesible encoding
        }
        for (i = j; i < number.length; i++) {
            b = (int) number[i] - '0';
            num = (num * 10) + b;
        }
        return ((num << NKBSHIFT_NUMBER) | (len & NKBMASK_LENGTH));
    }

    /**
     * Decode number into byte array.
     *
     * @param nk NumKey.
     * @return Number byte array.
     */
    private static byte[] decodeNumber(long nk) {
        int size = (int) (nk & NKBMASK_LENGTH);
        long num = ((nk & NKBMASK_NUMBER) >>> NKBSHIFT_NUMBER);
        int rem;
        int i;
        byte[] number = new byte[size];
        for (i = (size - 1); i >= 0; i--) {
            rem = (int) (num % 10);
            number[i] = (byte) (rem + '0');
            num = Long.divideUnsigned(num, 10);
        }
        return number;
    }

    /**
     * Encode numkey.
     *
     * @param country ISO 3166 alpha-2 country code.
     * @param number  Short code or LVN number.
     * @return NumKey 64 bit code.
     */
    public static long numkey(byte[] country, byte[] number) {
        return (encodeCountry(country) | encodeNumber(number));
    }

    /**
     * Encode numkey.
     *
     * @param country ISO 3166 alpha-2 country code.
     * @param number  Short code or LVN number.
     * @return NumKey 64 bit code.
     */
    public static long numkey(String country, String number) {
        return (encodeCountry(country.getBytes()) | encodeNumber(number.getBytes()));
    }

    /**
     * Decode a NumKey code to get the individual components.
     *
     * @param nk NumKey code.
     * @return country and number data as byte arrays.
     */
    public static NumData decodeNumkey(long nk) {
        NumData ret = new NumData();
        ret.country = decodeCountry(nk);
        ret.number = decodeNumber(nk);
        return ret;
    }

    /**
     * Decode a NumKey code to get the individual components
     *
     * @param nk NumKey code.
     * @return country and number data as strings.
     */
    public static NumDataStr decodeNumkeyStr(long nk) {
        NumDataStr ret = new NumDataStr();
        ret.country = new String(decodeCountry(nk));
        ret.number = new String(decodeNumber(nk));
        return ret;
    }

    /**
     * Compares two NumKeys by country only.
     *
     * @param nka The first NumKey to be compared.
     * @param nkb The second NumKey to be compared.
     * @return -1 if the first country is smaller than the second, 0 if they are equal and 1 if the first is greater than the second.
     */
    public static int compareNumkeyCountry(long nka, long nkb) {
        return Long.compareUnsigned((nka >>> NKBSHIFT_COUNTRY_SL), (nkb >>> NKBSHIFT_COUNTRY_SL));
    }

    /**
     * Returns NumKey hexadecimal string (16 characters).
     *
     * @param nk NumKey code.
     * @return hexadecimal string.
     */
    public static String numkeyHexString(long nk) {
        String s = Long.toUnsignedString(nk, 16);
        return new String(new char[16 - s.length()]).replace('\0', '0') + s;
    }

    /**
     * Returns NumKey hexadecimal string (16 characters).
     *
     * @param nk NumKey code.
     * @return hexadecimal byte array.
     */
    public static byte[] numkeyHex(long nk) {
        return numkeyHexString(nk).getBytes();
    }

    /**
     * Parses a NumKey hexadecimal string and returns the code.
     *
     * @param ns NumKey hexadecimal string (it must contain 16 hexadecimal characters).
     * @return A NumKey code.
     */
    public static long parseNumkeyHex(String ns) {
        return Long.parseUnsignedLong(ns, 16);
    }

    /**
     * Parses a NumKey hexadecimal string and returns the code.
     *
     * @param ns NumKey hexadecimal string (it must contain 16 hexadecimal characters).
     * @return A NumKey code.
     */
    public static long parseNumkeyHex(byte[] ns) {
        String s = new String(ns);
        return parseNumkeyHex(s);
    }

    /**
     * Encode a number string into long.
     * The encoded number is always 15 digits long as it is either right-padded with zeros or truncated.
     *
     * @param number byte array containing the E.164 number or prefix (max 18 digits or it will be truncated).
     * @return Encoded number.
     */
    public static long prefixkey(String number) {
        return prefixkey(number.getBytes());
    }

    /**
     * Encode a number string into long.
     * The encoded number is always 15 digits long as it is either right-padded with zeros or truncated.
     *
     * @param number byte array containing the E.164 number or prefix (max 18 digits or it will be truncated).
     * @return Encoded number.
     */
    public static long prefixkey(byte[] number) {
        long num = 0;
        int b;
        int i = 0;
        byte[] bnum = number;
        int len = bnum.length;
        if (len > PKNUMMAXLEN) {
            len = PKNUMMAXLEN; // truncate number
        }
        for (i = 0; i < len; i++) {
            b = (int) bnum[i] - '0';
            num = (num * 10) + b;
        }
        for (i = len; i < PKNUMMAXLEN; i++) {
            num = (num * 10); // zero right-padding
        }
        return num;
    }

    /**
     * Encode country code into a short number.
     *
     * @param country ISO 3166 alpha-2 country code.
     * @return Encoded country code.
     */
    public static short countrykey(String country) {
        return countrykey(country.getBytes());
    }

    /**
     * Encode country code into a short number.
     *
     * @param country ISO 3166 alpha-2 country code.
     * @return Encoded country code.
     */
    public static short countrykey(byte[] country) {
        return (short) (((country[0]) << 8) | (country[1]));
    }

    /**
     * Decode countrykey into 2-char byte array.
     *
     * @param ck CountryKey.
     * @return decoded ISO 3166 alpha-2 country code.
     */
    public static String decodeCountrykeyString(short ck) {
        return new String(decodeCountrykey(ck));
    }

    /**
     * Decode countrykey into 2-char byte array.
     *
     * @param ck CountryKey.
     * @return decoded ISO 3166 alpha-2 country code.
     */
    public static byte[] decodeCountrykey(short ck) {
        return new byte[]{
                (byte) ((ck & 0xFF00) >>> 8),
                (byte) (ck & 0x00FF),
        };
    }
}
