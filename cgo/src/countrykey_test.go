package numkey

// countrykey_test.go
// @category   Libraries
// @author     Nicola Asuni <nicola.asuni@vonage.com>
// @copyright  2019-2022 Vonage
// @license    see LICENSE file
// @lick       https://github.com/nexmoinc/countrykey

import "testing"

// TCountryKeyData contains test data.
type TCountryKeyData struct {
	country string
	ck      uint16
}

var countrykeyTestData = []TCountryKeyData{
	{"AA", 16705},
	{"AB", 16706},
	{"AC", 16707},
	{"AD", 16708},
	{"AE", 16709},
	{"AF", 16710},
	{"AG", 16711},
	{"AH", 16712},
	{"AI", 16713},
	{"AJ", 16714},
	{"AK", 16715},
	{"AL", 16716},
	{"AM", 16717},
	{"AN", 16718},
	{"AO", 16719},
	{"AP", 16720},
	{"AQ", 16721},
	{"AR", 16722},
	{"AS", 16723},
	{"AT", 16724},
	{"AU", 16725},
	{"AV", 16726},
	{"AW", 16727},
	{"AX", 16728},
	{"AY", 16729},
	{"AZ", 16730},
	{"BA", 16961},
	{"BB", 16962},
	{"BC", 16963},
	{"BD", 16964},
	{"BE", 16965},
	{"BF", 16966},
	{"BG", 16967},
	{"BH", 16968},
	{"BI", 16969},
	{"BJ", 16970},
	{"BK", 16971},
	{"BL", 16972},
	{"BM", 16973},
	{"BN", 16974},
	{"BO", 16975},
	{"BP", 16976},
	{"BQ", 16977},
	{"BR", 16978},
	{"BS", 16979},
	{"BT", 16980},
	{"BU", 16981},
	{"BV", 16982},
	{"BW", 16983},
	{"BX", 16984},
	{"BY", 16985},
	{"BZ", 16986},
	{"CA", 17217},
	{"CB", 17218},
	{"CC", 17219},
	{"CD", 17220},
	{"CE", 17221},
	{"CF", 17222},
	{"CG", 17223},
	{"CH", 17224},
	{"CI", 17225},
	{"CJ", 17226},
	{"CK", 17227},
	{"CL", 17228},
	{"CM", 17229},
	{"CN", 17230},
	{"CO", 17231},
	{"CP", 17232},
	{"CQ", 17233},
	{"CR", 17234},
	{"CS", 17235},
	{"CT", 17236},
	{"CU", 17237},
	{"CV", 17238},
	{"CW", 17239},
	{"CX", 17240},
	{"CY", 17241},
	{"CZ", 17242},
	{"DA", 17473},
	{"DB", 17474},
	{"DC", 17475},
	{"DD", 17476},
	{"DE", 17477},
	{"DF", 17478},
	{"DG", 17479},
	{"DH", 17480},
	{"DI", 17481},
	{"DJ", 17482},
	{"DK", 17483},
	{"DL", 17484},
	{"DM", 17485},
	{"DN", 17486},
	{"DO", 17487},
	{"DP", 17488},
	{"DQ", 17489},
	{"DR", 17490},
	{"DS", 17491},
	{"DT", 17492},
	{"DU", 17493},
	{"DV", 17494},
	{"DW", 17495},
	{"DX", 17496},
	{"DY", 17497},
	{"DZ", 17498},
	{"EA", 17729},
	{"EB", 17730},
	{"EC", 17731},
	{"ED", 17732},
	{"EE", 17733},
	{"EF", 17734},
	{"EG", 17735},
	{"EH", 17736},
	{"EI", 17737},
	{"EJ", 17738},
	{"EK", 17739},
	{"EL", 17740},
	{"EM", 17741},
	{"EN", 17742},
	{"EO", 17743},
	{"EP", 17744},
	{"EQ", 17745},
	{"ER", 17746},
	{"ES", 17747},
	{"ET", 17748},
	{"EU", 17749},
	{"EV", 17750},
	{"EW", 17751},
	{"EX", 17752},
	{"EY", 17753},
	{"EZ", 17754},
	{"FA", 17985},
	{"FB", 17986},
	{"FC", 17987},
	{"FD", 17988},
	{"FE", 17989},
	{"FF", 17990},
	{"FG", 17991},
	{"FH", 17992},
	{"FI", 17993},
	{"FJ", 17994},
	{"FK", 17995},
	{"FL", 17996},
	{"FM", 17997},
	{"FN", 17998},
	{"FO", 17999},
	{"FP", 18000},
	{"FQ", 18001},
	{"FR", 18002},
	{"FS", 18003},
	{"FT", 18004},
	{"FU", 18005},
	{"FV", 18006},
	{"FW", 18007},
	{"FX", 18008},
	{"FY", 18009},
	{"FZ", 18010},
	{"GA", 18241},
	{"GB", 18242},
	{"GC", 18243},
	{"GD", 18244},
	{"GE", 18245},
	{"GF", 18246},
	{"GG", 18247},
	{"GH", 18248},
	{"GI", 18249},
	{"GJ", 18250},
	{"GK", 18251},
	{"GL", 18252},
	{"GM", 18253},
	{"GN", 18254},
	{"GO", 18255},
	{"GP", 18256},
	{"GQ", 18257},
	{"GR", 18258},
	{"GS", 18259},
	{"GT", 18260},
	{"GU", 18261},
	{"GV", 18262},
	{"GW", 18263},
	{"GX", 18264},
	{"GY", 18265},
	{"GZ", 18266},
	{"HA", 18497},
	{"HB", 18498},
	{"HC", 18499},
	{"HD", 18500},
	{"HE", 18501},
	{"HF", 18502},
	{"HG", 18503},
	{"HH", 18504},
	{"HI", 18505},
	{"HJ", 18506},
	{"HK", 18507},
	{"HL", 18508},
	{"HM", 18509},
	{"HN", 18510},
	{"HO", 18511},
	{"HP", 18512},
	{"HQ", 18513},
	{"HR", 18514},
	{"HS", 18515},
	{"HT", 18516},
	{"HU", 18517},
	{"HV", 18518},
	{"HW", 18519},
	{"HX", 18520},
	{"HY", 18521},
	{"HZ", 18522},
	{"IA", 18753},
	{"IB", 18754},
	{"IC", 18755},
	{"ID", 18756},
	{"IE", 18757},
	{"IF", 18758},
	{"IG", 18759},
	{"IH", 18760},
	{"II", 18761},
	{"IJ", 18762},
	{"IK", 18763},
	{"IL", 18764},
	{"IM", 18765},
	{"IN", 18766},
	{"IO", 18767},
	{"IP", 18768},
	{"IQ", 18769},
	{"IR", 18770},
	{"IS", 18771},
	{"IT", 18772},
	{"IU", 18773},
	{"IV", 18774},
	{"IW", 18775},
	{"IX", 18776},
	{"IY", 18777},
	{"IZ", 18778},
	{"JA", 19009},
	{"JB", 19010},
	{"JC", 19011},
	{"JD", 19012},
	{"JE", 19013},
	{"JF", 19014},
	{"JG", 19015},
	{"JH", 19016},
	{"JI", 19017},
	{"JJ", 19018},
	{"JK", 19019},
	{"JL", 19020},
	{"JM", 19021},
	{"JN", 19022},
	{"JO", 19023},
	{"JP", 19024},
	{"JQ", 19025},
	{"JR", 19026},
	{"JS", 19027},
	{"JT", 19028},
	{"JU", 19029},
	{"JV", 19030},
	{"JW", 19031},
	{"JX", 19032},
	{"JY", 19033},
	{"JZ", 19034},
	{"KA", 19265},
	{"KB", 19266},
	{"KC", 19267},
	{"KD", 19268},
	{"KE", 19269},
	{"KF", 19270},
	{"KG", 19271},
	{"KH", 19272},
	{"KI", 19273},
	{"KJ", 19274},
	{"KK", 19275},
	{"KL", 19276},
	{"KM", 19277},
	{"KN", 19278},
	{"KO", 19279},
	{"KP", 19280},
	{"KQ", 19281},
	{"KR", 19282},
	{"KS", 19283},
	{"KT", 19284},
	{"KU", 19285},
	{"KV", 19286},
	{"KW", 19287},
	{"KX", 19288},
	{"KY", 19289},
	{"KZ", 19290},
	{"LA", 19521},
	{"LB", 19522},
	{"LC", 19523},
	{"LD", 19524},
	{"LE", 19525},
	{"LF", 19526},
	{"LG", 19527},
	{"LH", 19528},
	{"LI", 19529},
	{"LJ", 19530},
	{"LK", 19531},
	{"LL", 19532},
	{"LM", 19533},
	{"LN", 19534},
	{"LO", 19535},
	{"LP", 19536},
	{"LQ", 19537},
	{"LR", 19538},
	{"LS", 19539},
	{"LT", 19540},
	{"LU", 19541},
	{"LV", 19542},
	{"LW", 19543},
	{"LX", 19544},
	{"LY", 19545},
	{"LZ", 19546},
	{"MA", 19777},
	{"MB", 19778},
	{"MC", 19779},
	{"MD", 19780},
	{"ME", 19781},
	{"MF", 19782},
	{"MG", 19783},
	{"MH", 19784},
	{"MI", 19785},
	{"MJ", 19786},
	{"MK", 19787},
	{"ML", 19788},
	{"MM", 19789},
	{"MN", 19790},
	{"MO", 19791},
	{"MP", 19792},
	{"MQ", 19793},
	{"MR", 19794},
	{"MS", 19795},
	{"MT", 19796},
	{"MU", 19797},
	{"MV", 19798},
	{"MW", 19799},
	{"MX", 19800},
	{"MY", 19801},
	{"MZ", 19802},
	{"NA", 20033},
	{"NB", 20034},
	{"NC", 20035},
	{"ND", 20036},
	{"NE", 20037},
	{"NF", 20038},
	{"NG", 20039},
	{"NH", 20040},
	{"NI", 20041},
	{"NJ", 20042},
	{"NK", 20043},
	{"NL", 20044},
	{"NM", 20045},
	{"NN", 20046},
	{"NO", 20047},
	{"NP", 20048},
	{"NQ", 20049},
	{"NR", 20050},
	{"NS", 20051},
	{"NT", 20052},
	{"NU", 20053},
	{"NV", 20054},
	{"NW", 20055},
	{"NX", 20056},
	{"NY", 20057},
	{"NZ", 20058},
	{"OA", 20289},
	{"OB", 20290},
	{"OC", 20291},
	{"OD", 20292},
	{"OE", 20293},
	{"OF", 20294},
	{"OG", 20295},
	{"OH", 20296},
	{"OI", 20297},
	{"OJ", 20298},
	{"OK", 20299},
	{"OL", 20300},
	{"OM", 20301},
	{"ON", 20302},
	{"OO", 20303},
	{"OP", 20304},
	{"OQ", 20305},
	{"OR", 20306},
	{"OS", 20307},
	{"OT", 20308},
	{"OU", 20309},
	{"OV", 20310},
	{"OW", 20311},
	{"OX", 20312},
	{"OY", 20313},
	{"OZ", 20314},
	{"PA", 20545},
	{"PB", 20546},
	{"PC", 20547},
	{"PD", 20548},
	{"PE", 20549},
	{"PF", 20550},
	{"PG", 20551},
	{"PH", 20552},
	{"PI", 20553},
	{"PJ", 20554},
	{"PK", 20555},
	{"PL", 20556},
	{"PM", 20557},
	{"PN", 20558},
	{"PO", 20559},
	{"PP", 20560},
	{"PQ", 20561},
	{"PR", 20562},
	{"PS", 20563},
	{"PT", 20564},
	{"PU", 20565},
	{"PV", 20566},
	{"PW", 20567},
	{"PX", 20568},
	{"PY", 20569},
	{"PZ", 20570},
	{"QA", 20801},
	{"QB", 20802},
	{"QC", 20803},
	{"QD", 20804},
	{"QE", 20805},
	{"QF", 20806},
	{"QG", 20807},
	{"QH", 20808},
	{"QI", 20809},
	{"QJ", 20810},
	{"QK", 20811},
	{"QL", 20812},
	{"QM", 20813},
	{"QN", 20814},
	{"QO", 20815},
	{"QP", 20816},
	{"QQ", 20817},
	{"QR", 20818},
	{"QS", 20819},
	{"QT", 20820},
	{"QU", 20821},
	{"QV", 20822},
	{"QW", 20823},
	{"QX", 20824},
	{"QY", 20825},
	{"QZ", 20826},
	{"RA", 21057},
	{"RB", 21058},
	{"RC", 21059},
	{"RD", 21060},
	{"RE", 21061},
	{"RF", 21062},
	{"RG", 21063},
	{"RH", 21064},
	{"RI", 21065},
	{"RJ", 21066},
	{"RK", 21067},
	{"RL", 21068},
	{"RM", 21069},
	{"RN", 21070},
	{"RO", 21071},
	{"RP", 21072},
	{"RQ", 21073},
	{"RR", 21074},
	{"RS", 21075},
	{"RT", 21076},
	{"RU", 21077},
	{"RV", 21078},
	{"RW", 21079},
	{"RX", 21080},
	{"RY", 21081},
	{"RZ", 21082},
	{"SA", 21313},
	{"SB", 21314},
	{"SC", 21315},
	{"SD", 21316},
	{"SE", 21317},
	{"SF", 21318},
	{"SG", 21319},
	{"SH", 21320},
	{"SI", 21321},
	{"SJ", 21322},
	{"SK", 21323},
	{"SL", 21324},
	{"SM", 21325},
	{"SN", 21326},
	{"SO", 21327},
	{"SP", 21328},
	{"SQ", 21329},
	{"SR", 21330},
	{"SS", 21331},
	{"ST", 21332},
	{"SU", 21333},
	{"SV", 21334},
	{"SW", 21335},
	{"SX", 21336},
	{"SY", 21337},
	{"SZ", 21338},
	{"TA", 21569},
	{"TB", 21570},
	{"TC", 21571},
	{"TD", 21572},
	{"TE", 21573},
	{"TF", 21574},
	{"TG", 21575},
	{"TH", 21576},
	{"TI", 21577},
	{"TJ", 21578},
	{"TK", 21579},
	{"TL", 21580},
	{"TM", 21581},
	{"TN", 21582},
	{"TO", 21583},
	{"TP", 21584},
	{"TQ", 21585},
	{"TR", 21586},
	{"TS", 21587},
	{"TT", 21588},
	{"TU", 21589},
	{"TV", 21590},
	{"TW", 21591},
	{"TX", 21592},
	{"TY", 21593},
	{"TZ", 21594},
	{"UA", 21825},
	{"UB", 21826},
	{"UC", 21827},
	{"UD", 21828},
	{"UE", 21829},
	{"UF", 21830},
	{"UG", 21831},
	{"UH", 21832},
	{"UI", 21833},
	{"UJ", 21834},
	{"UK", 21835},
	{"UL", 21836},
	{"UM", 21837},
	{"UN", 21838},
	{"UO", 21839},
	{"UP", 21840},
	{"UQ", 21841},
	{"UR", 21842},
	{"US", 21843},
	{"UT", 21844},
	{"UU", 21845},
	{"UV", 21846},
	{"UW", 21847},
	{"UX", 21848},
	{"UY", 21849},
	{"UZ", 21850},
	{"VA", 22081},
	{"VB", 22082},
	{"VC", 22083},
	{"VD", 22084},
	{"VE", 22085},
	{"VF", 22086},
	{"VG", 22087},
	{"VH", 22088},
	{"VI", 22089},
	{"VJ", 22090},
	{"VK", 22091},
	{"VL", 22092},
	{"VM", 22093},
	{"VN", 22094},
	{"VO", 22095},
	{"VP", 22096},
	{"VQ", 22097},
	{"VR", 22098},
	{"VS", 22099},
	{"VT", 22100},
	{"VU", 22101},
	{"VV", 22102},
	{"VW", 22103},
	{"VX", 22104},
	{"VY", 22105},
	{"VZ", 22106},
	{"WA", 22337},
	{"WB", 22338},
	{"WC", 22339},
	{"WD", 22340},
	{"WE", 22341},
	{"WF", 22342},
	{"WG", 22343},
	{"WH", 22344},
	{"WI", 22345},
	{"WJ", 22346},
	{"WK", 22347},
	{"WL", 22348},
	{"WM", 22349},
	{"WN", 22350},
	{"WO", 22351},
	{"WP", 22352},
	{"WQ", 22353},
	{"WR", 22354},
	{"WS", 22355},
	{"WT", 22356},
	{"WU", 22357},
	{"WV", 22358},
	{"WW", 22359},
	{"WX", 22360},
	{"WY", 22361},
	{"WZ", 22362},
	{"XA", 22593},
	{"XB", 22594},
	{"XC", 22595},
	{"XD", 22596},
	{"XE", 22597},
	{"XF", 22598},
	{"XG", 22599},
	{"XH", 22600},
	{"XI", 22601},
	{"XJ", 22602},
	{"XK", 22603},
	{"XL", 22604},
	{"XM", 22605},
	{"XN", 22606},
	{"XO", 22607},
	{"XP", 22608},
	{"XQ", 22609},
	{"XR", 22610},
	{"XS", 22611},
	{"XT", 22612},
	{"XU", 22613},
	{"XV", 22614},
	{"XW", 22615},
	{"XX", 22616},
	{"XY", 22617},
	{"XZ", 22618},
	{"YA", 22849},
	{"YB", 22850},
	{"YC", 22851},
	{"YD", 22852},
	{"YE", 22853},
	{"YF", 22854},
	{"YG", 22855},
	{"YH", 22856},
	{"YI", 22857},
	{"YJ", 22858},
	{"YK", 22859},
	{"YL", 22860},
	{"YM", 22861},
	{"YN", 22862},
	{"YO", 22863},
	{"YP", 22864},
	{"YQ", 22865},
	{"YR", 22866},
	{"YS", 22867},
	{"YT", 22868},
	{"YU", 22869},
	{"YV", 22870},
	{"YW", 22871},
	{"YX", 22872},
	{"YY", 22873},
	{"YZ", 22874},
	{"ZA", 23105},
	{"ZB", 23106},
	{"ZC", 23107},
	{"ZD", 23108},
	{"ZE", 23109},
	{"ZF", 23110},
	{"ZG", 23111},
	{"ZH", 23112},
	{"ZI", 23113},
	{"ZJ", 23114},
	{"ZK", 23115},
	{"ZL", 23116},
	{"ZM", 23117},
	{"ZN", 23118},
	{"ZO", 23119},
	{"ZP", 23120},
	{"ZQ", 23121},
	{"ZR", 23122},
	{"ZS", 23123},
	{"ZT", 23124},
	{"ZU", 23125},
	{"ZV", 23126},
	{"ZW", 23127},
	{"ZX", 23128},
	{"ZY", 23129},
	{"ZZ", 23130},
}

var countrykeyTestDataError = []TCountryKeyData{
	{"", 0},
}

func TestCountryKeyError(t *testing.T) {
	for _, v := range countrykeyTestDataError {
		v := v
		t.Run("", func(t *testing.T) {
			t.Parallel()
			ck := CountryKey(v.country)
			if ck != v.ck {
				t.Errorf("The code value is different, expected %#v got %#v", v.ck, ck)
			}
		})
	}
}

func TestCountryKey(t *testing.T) {
	for _, v := range countrykeyTestData {
		v := v
		t.Run("", func(t *testing.T) {
			t.Parallel()
			ck := CountryKey(v.country)
			if ck != v.ck {
				t.Errorf("The code value is different, expected %#v got %#v", v.ck, ck)
			}
		})
	}
}

func BenchmarkCountryKey(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountryKey("ZZ")
	}
}

func TestDecodeCountryKey(t *testing.T) {
	for _, v := range countrykeyTestData {
		v := v
		t.Run("", func(t *testing.T) {
			t.Parallel()
			country := DecodeCountryKey(v.ck)
			if country != v.country {
				t.Errorf("The country hash value is different, expected %#v got: %#v", v.country, country)
			}
		})
	}
}

func BenchmarkDecodeCountryKey(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DecodeCountryKey(23130)
	}
}
