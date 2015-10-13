package areacode

import "encoding/json"

type Code struct {
	CountryCode int64 `json:"country_code"`
	Name 		string `json:"name"`
	Example 	string `json:"example"`
}

const Codes = `[  
    {
        "country_code": 52,
        "name": "Mexico",
        "example": "5212221234567"
    },
    {
        "country_code": 1,
        "name": "UnitedStates",
        "example": "12015550123"
    },
    {
        "country_code": 247,
        "name": "AscensionIsland",
        "example": "2476889"
    },
    {
        "country_code": 93,
        "name": "Afghanistan",
        "example": "93701234567"
    },
    {
        "country_code": 355,
        "name": "Albania",
        "example": "355661234567"
    },
    {
        "country_code": 213,
        "name": "Algeria",
        "example": "213551234567"
    },
    {
        "country_code": 1,
        "name": "AmericanSamoa",
        "example": "16847331234"
    },
    {
        "country_code": 376,
        "name": "Andorra",
        "example": "376312345"
    },
    {
        "country_code": 244,
        "name": "Angola",
        "example": "244923123456"
    },
    {
        "country_code": 1,
        "name": "Anguilla",
        "example": "12642351234"
    },
    {
        "country_code": 1,
        "name": "AntiguaandBarbuda",
        "example": "12684641234"
    },
    {
        "country_code": 54,
        "name": "Argentina",
        "example": "5491123456789"
    },
    {
        "country_code": 374,
        "name": "Armenia",
        "example": "37477123456"
    },
    {
        "country_code": 297,
        "name": "Aruba",
        "example": "2975601234"
    },
    {
        "country_code": 61,
        "name": "Australia",
        "example": "61412345678"
    },
    {
        "country_code": 43,
        "name": "Austria",
        "example": "43644123456"
    },
    {
        "country_code": 994,
        "name": "Azerbaijan",
        "example": "994401234567"
    },
    {
        "country_code": 1,
        "name": "Bahamas",
        "example": "12423591234"
    },
    {
        "country_code": 973,
        "name": "Bahrain",
        "example": "9733600XXXX"
    },
    {
        "country_code": 880,
        "name": "Bangladesh",
        "example": "8801812345678"
    },
    {
        "country_code": 1,
        "name": "Barbados",
        "example": "12462501234"
    },
    {
        "country_code": 375,
        "name": "Belarus",
        "example": "375294911911"
    },
    {
        "country_code": 32,
        "name": "Belgium",
        "example": "32470123456"
    },
    {
        "country_code": 501,
        "name": "Belize",
        "example": "5016221234"
    },
    {
        "country_code": 229,
        "name": "Benin",
        "example": "22990011234"
    },
    {
        "country_code": 1,
        "name": "Bermuda",
        "example": "14413701234"
    },
    {
        "country_code": 975,
        "name": "Bhutan",
        "example": "97517XXXXXX"
    },
    {
        "country_code": 591,
        "name": "Bolivia",
        "example": "5917XXXXXXX"
    },
    {
        "country_code": 387,
        "name": "BosniaandHerzegovina",
        "example": "38761123456"
    },
    {
        "country_code": 267,
        "name": "Botswana",
        "example": "26771123456"
    },
    {
        "country_code": 55,
        "name": "Brazil",
        "example": "55(XX)XXXX-XXXX,+55(XX)9XXXX-XXXX"
    },
    {
        "country_code": 1,
        "name": "BritishVirginIslands",
        "example": "12843001234"
    },
    {
        "country_code": 673,
        "name": "Brunei",
        "example": "6737XXXXXX"
    },
    {
        "country_code": 359,
        "name": "Bulgaria",
        "example": "35948123456"
    },
    {
        "country_code": 226,
        "name": "BurkinaFaso",
        "example": "22670123456"
    },
    {
        "country_code": 257,
        "name": "Burundi",
        "example": "25779561234"
    },
    {
        "country_code": 855,
        "name": "Cambodia",
        "example": "85591234567"
    },
    {
        "country_code": 237,
        "name": "Cameroon",
        "example": "23771234567"
    },
    {
        "country_code": 1,
        "name": "Canada",
        "example": "12042345678"
    },
    {
        "country_code": 238,
        "name": "CapeVerde",
        "example": "2389911234"
    },
    {
        "country_code": 1,
        "name": "CaymanIslands",
        "example": "1345XXXXXXX"
    },
    {
        "country_code": 236,
        "name": "CentralAfricanRepublic",
        "example": "23670012345"
    },
    {
        "country_code": 235,
        "name": "Chad",
        "example": "23563012345"
    },
    {
        "country_code": 56,
        "name": "Chile",
        "example": "56961234567"
    },
    {
        "country_code": 86,
        "name": "China",
        "example": "8613123456789"
    },
    {
        "country_code": 61,
        "name": "ChristmasIsland",
        "example": "61412345678"
    },
    {
        "country_code": 61,
        "name": "Cocos(Keeling)Islands",
        "example": "61412345678"
    },
    {
        "country_code": 57,
        "name": "Colombia",
        "example": "573211234567"
    },
    {
        "country_code": 269,
        "name": "Comoros",
        "example": "2693212345"
    },
    {
        "country_code": 242,
        "name": "Congo-Brazzaville",
        "example": "242061234567"
    },
    {
        "country_code": 243,
        "name": "Congo-Kinshasa",
        "example": "243991234567"
    },
    {
        "country_code": 682,
        "name": "CookIslands",
        "example": "68271234"
    },
    {
        "country_code": 506,
        "name": "CostaRica",
        "example": "50683123456"
    },
    {
        "country_code": 385,
        "name": "Croatia",
        "example": "3859XXXXXXXX"
    },
    {
        "country_code": 53,
        "name": "Cuba",
        "example": "5351234567"
    },
    {
        "country_code": 357,
        "name": "Cyprus",
        "example": "35796123456"
    },
    {
        "country_code": 420,
        "name": "CzechRepublic",
        "example": "420601123456"
    },
    {
        "country_code": 225,
        "name": "C\u00f4ted\u2019Ivoire",
        "example": "225XXXXXXX"
    },
    {
        "country_code": 45,
        "name": "Denmark",
        "example": "4520123456"
    },
    {
        "country_code": 253,
        "name": "Djibouti",
        "example": "25377831001"
    },
    {
        "country_code": 1,
        "name": "Dominica",
        "example": "1767225XXX"
    },
    {
        "country_code": 1,
        "name": "DominicanRepublic",
        "example": "18092345678"
    },
    {
        "country_code": 593,
        "name": "Ecuador",
        "example": "593991234567"
    },
    {
        "country_code": 20,
        "name": "Egypt",
        "example": "20100XXXXXXX"
    },
    {
        "country_code": 503,
        "name": "ElSalvador",
        "example": "50370123456"
    },
    {
        "country_code": 240,
        "name": "EquatorialGuinea",
        "example": "240222XXXXXX"
    },
    {
        "country_code": 291,
        "name": "Eritrea",
        "example": "2917123456"
    },
    {
        "country_code": 372,
        "name": "Estonia",
        "example": "37251234567"
    },
    {
        "country_code": 251,
        "name": "Ethiopia",
        "example": "251911234567"
    },
    {
        "country_code": 500,
        "name": "FalklandIslands",
        "example": 50051234
    },
    {
        "country_code": 298,
        "name": "FaroeIslands",
        "example": "29821XXXX"
    },
    {
        "country_code": 679,
        "name": "Fiji",
        "example": "6797012345"
    },
    {
        "country_code": 358,
        "name": "Finland",
        "example": "358412345678"
    },
    {
        "country_code": 33,
        "name": "France",
        "example": "33655551122"
    },
    {
        "country_code": 594,
        "name": "FrenchGuiana",
        "example": "5946942XXXXX"
    },
    {
        "country_code": 689,
        "name": "FrenchPolynesia",
        "example": "689212345"
    },
    {
        "country_code": 241,
        "name": "Gabon",
        "example": "24106031234"
    },
    {
        "country_code": 220,
        "name": "Gambia",
        "example": "2203012345"
    },
    {
        "country_code": 995,
        "name": "Georgia",
        "example": "995555123456"
    },
    {
        "country_code": 49,
        "name": "Germany",
        "example": "49178XXXXXXXX"
    },
    {
        "country_code": 233,
        "name": "Ghana",
        "example": "233231234567"
    },
    {
        "country_code": 350,
        "name": "Gibraltar",
        "example": "35057123456"
    },
    {
        "country_code": 30,
        "name": "Greece",
        "example": "306912345678"
    },
    {
        "country_code": 299,
        "name": "Greenland",
        "example": "299221234"
    },
    {
        "country_code": 1,
        "name": "Grenada",
        "example": "14734031234"
    },
    {
        "country_code": 590,
        "name": "Guadeloupe",
        "example": "590690301234"
    },
    {
        "country_code": 1,
        "name": "Guam",
        "example": "16713001234"
    },
    {
        "country_code": 502,
        "name": "Guatemala",
        "example": "50251234567"
    },
    {
        "country_code": 44,
        "name": "Guernsey",
        "example": "447781123456"
    },
    {
        "country_code": 224,
        "name": "Guinea",
        "example": "22460201234"
    },
    {
        "country_code": 245,
        "name": "Guinea-Bissau",
        "example": "2455012345"
    },
    {
        "country_code": 592,
        "name": "Guyana",
        "example": "5926091234"
    },
    {
        "country_code": 509,
        "name": "Haiti",
        "example": "50934101234"
    },
    {
        "country_code": 504,
        "name": "Honduras",
        "example": "50491234567"
    },
    {
        "country_code": 852,
        "name": "HongKongSARChina",
        "example": "85251234567"
    },
    {
        "country_code": 36,
        "name": "Hungary",
        "example": "36201234567"
    },
    {
        "country_code": 354,
        "name": "Iceland",
        "example": "3546101234"
    },
    {
        "country_code": 91,
        "name": "India",
        "example": "919XXXXXXXXX"
    },
    {
        "country_code": 62,
        "name": "Indonesia",
        "example": "62812345678"
    },
    {
        "country_code": 98,
        "name": "Iran",
        "example": "989123456789"
    },
    {
        "country_code": 964,
        "name": "Iraq",
        "example": "9647912345678"
    },
    {
        "country_code": 353,
        "name": "Ireland",
        "example": "353850123456"
    },
    {
        "country_code": 44,
        "name": "IsleofMan",
        "example": "447924123456"
    },
    {
        "country_code": 972,
        "name": "Israel",
        "example": "972501234567"
    },
    {
        "country_code": 39,
        "name": "Italy",
        "example": "393123456789"
    },
    {
        "country_code": 1,
        "name": "Jamaica",
        "example": "18762101234"
    },
    {
        "country_code": 81,
        "name": "Japan",
        "example": "817012345678"
    },
    {
        "country_code": 44,
        "name": "Jersey",
        "example": "447797XXXXXX"
    },
    {
        "country_code": 962,
        "name": "Jordan",
        "example": "962790123456"
    },
    {
        "country_code": 7,
        "name": "Kazakhstan",
        "example": "77710009998"
    },
    {
        "country_code": 254,
        "name": "Kenya",
        "example": "254712XXXXXX"
    },
    {
        "country_code": 686,
        "name": "Kiribati",
        "example": "68661234"
    },
    {
        "country_code": 965,
        "name": "Kuwait",
        "example": "965500XXXXX"
    },
    {
        "country_code": 996,
        "name": "Kyrgyzstan",
        "example": "996700123456"
    },
    {
        "country_code": 856,
        "name": "Laos",
        "example": "8562023123456"
    },
    {
        "country_code": 371,
        "name": "Latvia",
        "example": "37121234567"
    },
    {
        "country_code": 961,
        "name": "Lebanon",
        "example": "96171XXXXXX"
    },
    {
        "country_code": 266,
        "name": "Lesotho",
        "example": "26650123456"
    },
    {
        "country_code": 231,
        "name": "Liberia",
        "example": "2314612345"
    },
    {
        "country_code": 218,
        "name": "Libya",
        "example": "218912345678"
    },
    {
        "country_code": 423,
        "name": "Liechtenstein",
        "example": "423661234567"
    },
    {
        "country_code": 370,
        "name": "Lithuania",
        "example": "37061234567"
    },
    {
        "country_code": 352,
        "name": "Luxembourg",
        "example": "352628123456"
    },
    {
        "country_code": 853,
        "name": "MacauSARChina",
        "example": "85366123456"
    },
    {
        "country_code": 389,
        "name": "Macedonia",
        "example": "38972345678"
    },
    {
        "country_code": 261,
        "name": "Madagascar",
        "example": "261301234567"
    },
    {
        "country_code": 265,
        "name": "Malawi",
        "example": "26599XXXXXXX"
    },
    {
        "country_code": 60,
        "name": "Malaysia",
        "example": "60123456789"
    },
    {
        "country_code": 960,
        "name": "Maldives",
        "example": "9607712345"
    },
    {
        "country_code": 223,
        "name": "Mali",
        "example": "22365012345"
    },
    {
        "country_code": 356,
        "name": "Malta",
        "example": "35696961234"
    },
    {
        "country_code": 692,
        "name": "MarshallIslands",
        "example": "6922351234"
    },
    {
        "country_code": 596,
        "name": "Martinique",
        "example": "596696201234"
    },
    {
        "country_code": 222,
        "name": "Mauritania",
        "example": "22222123456"
    },
    {
        "country_code": 230,
        "name": "Mauritius",
        "example": "2302512345"
    },
    {
        "country_code": 262,
        "name": "Mayotte",
        "example": "262639XXXXXX"
    },
    {
        "country_code": 52,
        "name": "Mexico",
        "example": "5212221234567"
    },
    {
        "country_code": 691,
        "name": "Micronesia",
        "example": "6913501234"
    },
    {
        "country_code": 373,
        "name": "Moldova",
        "example": "37365012345"
    },
    {
        "country_code": 377,
        "name": "Monaco",
        "example": "377612345678"
    },
    {
        "country_code": 976,
        "name": "Mongolia",
        "example": "97688123456"
    },
    {
        "country_code": 382,
        "name": "Montenegro",
        "example": "38267622901"
    },
    {
        "country_code": 1,
        "name": "Montserrat",
        "example": "16644923456"
    },
    {
        "country_code": 212,
        "name": "Morocco",
        "example": "212650XXXXXX"
    },
    {
        "country_code": 258,
        "name": "Mozambique",
        "example": "25882XXXXXXX"
    },
    {
        "country_code": 95,
        "name": "Myanmar(Burma)",
        "example": "9592123456"
    },
    {
        "country_code": 264,
        "name": "Namibia",
        "example": "264811234567"
    },
    {
        "country_code": 674,
        "name": "Nauru",
        "example": "6745551234"
    },
    {
        "country_code": 977,
        "name": "Nepal",
        "example": "9779841234567"
    },
    {
        "country_code": 31,
        "name": "Netherlands",
        "example": "31612345678"
    },
    {
        "country_code": 687,
        "name": "NewCaledonia",
        "example": "687751234"
    },
    {
        "country_code": 64,
        "name": "NewZealand",
        "example": "64211234567"
    },
    {
        "country_code": 505,
        "name": "Nicaragua",
        "example": "50581234567"
    },
    {
        "country_code": 227,
        "name": "Niger",
        "example": "22793123456"
    },
    {
        "country_code": 234,
        "name": "Nigeria",
        "example": "2348021234567"
    },
    {
        "country_code": 683,
        "name": "Niue",
        "example": "6831234"
    },
    {
        "country_code": 672,
        "name": "NorfolkIsland",
        "example": "672381234"
    },
    {
        "country_code": 850,
        "name": "NorthKorea",
        "example": "8501921234567"
    },
    {
        "country_code": 1,
        "name": "NorthernMarianaIslands",
        "example": "16702345678"
    },
    {
        "country_code": 47,
        "name": "Norway",
        "example": "4741234567"
    },
    {
        "country_code": 968,
        "name": "Oman",
        "example": "96892XXXXXX"
    },
    {
        "country_code": 92,
        "name": "Pakistan",
        "example": "923012345678"
    },
    {
        "country_code": 680,
        "name": "Palau",
        "example": "6806201234"
    },
    {
        "country_code": 970,
        "name": "PalestinianTerritories",
        "example": "970599123456"
    },
    {
        "country_code": 507,
        "name": "Panama",
        "example": "507600XXXXX"
    },
    {
        "country_code": 675,
        "name": "PapuaNewGuinea",
        "example": "6756812345"
    },
    {
        "country_code": 595,
        "name": "Paraguay",
        "example": "595961XXXXXX"
    },
    {
        "country_code": 51,
        "name": "Peru",
        "example": "51912345678"
    },
    {
        "country_code": 63,
        "name": "Philippines",
        "example": "639051234567"
    },
    {
        "country_code": 48,
        "name": "Poland",
        "example": "48512345678"
    },
    {
        "country_code": 351,
        "name": "Portugal",
        "example": "351912345678"
    },
    {
        "country_code": 1,
        "name": "PuertoRico",
        "example": "1787XXXXXX"
    },
    {
        "country_code": 974,
        "name": "Qatar",
        "example": "97433123456"
    },
    {
        "country_code": 40,
        "name": "Romania",
        "example": "407XXXXXXXX"
    },
    {
        "country_code": 7,
        "name": "Russia",
        "example": "79123456789"
    },
    {
        "country_code": 250,
        "name": "Rwanda",
        "example": "250720123456"
    },
    {
        "country_code": 262,
        "name": "R\u00e9union",
        "example": "262692123456"
    },
    {
        "country_code": 211,
        "name": "SouthSudan",
        "example": "211977123456"
    },
    {
        "country_code": 590,
        "name": "SaintBarth\u00e9lemy",
        "example": "590690221234"
    },
    {
        "country_code": 290,
        "name": "SaintHelena",
        "example": "2902158"
    },
    {
        "country_code": 1,
        "name": "SaintKittsandNevis",
        "example": "18695561234"
    },
    {
        "country_code": 1,
        "name": "SaintLucia",
        "example": "17582845678"
    },
    {
        "country_code": 590,
        "name": "SaintMartin",
        "example": "590690221234"
    },
    {
        "country_code": 508,
        "name": "SaintPierreandMiquelon",
        "example": "508551234"
    },
    {
        "country_code": 1,
        "name": "SaintVincentandtheGrenadines",
        "example": "17844301234"
    },
    {
        "country_code": 685,
        "name": "Samoa",
        "example": "685601234"
    },
    {
        "country_code": 378,
        "name": "SanMarino",
        "example": "37866661212"
    },
    {
        "country_code": 966,
        "name": "SaudiArabia",
        "example": "966512345678"
    },
    {
        "country_code": 221,
        "name": "Senegal",
        "example": "221701012345"
    },
    {
        "country_code": 381,
        "name": "Serbia",
        "example": "381601234567"
    },
    {
        "country_code": 248,
        "name": "Seychelles",
        "example": "2482510123"
    },
    {
        "country_code": 232,
        "name": "SierraLeone",
        "example": "23225123456"
    },
    {
        "country_code": 65,
        "name": "Singapore",
        "example": "658XXXXXXX"
    },
    {
        "country_code": 421,
        "name": "Slovakia",
        "example": "421912123456"
    },
    {
        "country_code": 386,
        "name": "Slovenia",
        "example": "38631234567"
    },
    {
        "country_code": 677,
        "name": "SolomonIslands",
        "example": "6777421234"
    },
    {
        "country_code": 252,
        "name": "Somalia",
        "example": "25290792024"
    },
    {
        "country_code": 27,
        "name": "SouthAfrica",
        "example": "27711234567"
    },
    {
        "country_code": 82,
        "name": "SouthKorea",
        "example": "821023456789"
    },
    {
        "country_code": 34,
        "name": "Spain",
        "example": "34612345678"
    },
    {
        "country_code": 94,
        "name": "SriLanka",
        "example": "94712345678"
    },
    {
        "country_code": 249,
        "name": "Sudan",
        "example": "249911231234"
    },
    {
        "country_code": 597,
        "name": "Suri\"name\"",
        "example": "5977412345"
    },
    {
        "country_code": 47,
        "name": "SvalbardandJanMayen",
        "example": "4741234567"
    },
    {
        "country_code": 268,
        "name": "Swaziland",
        "example": "26876XXXXXX"
    },
    {
        "country_code": 46,
        "name": "Sweden",
        "example": "46701234567"
    },
    {
        "country_code": 41,
        "name": "Switzerland",
        "example": "41741234567"
    },
    {
        "country_code": 963,
        "name": "Syria",
        "example": "963944567890"
    },
    {
        "country_code": 239,
        "name": "S\u00e3oTom\u00e9andPr\u00edncipe",
        "example": "2399812345"
    },
    {
        "country_code": 886,
        "name": "Taiwan",
        "example": "886912345678"
    },
    {
        "country_code": 992,
        "name": "Tajikistan",
        "example": "992917123456"
    },
    {
        "country_code": 255,
        "name": "Tanzania",
        "example": "255612345678"
    },
    {
        "country_code": 66,
        "name": "Thailand",
        "example": "668XXXXXXXX"
    },
    {
        "country_code": 670,
        "name": "Timor-Leste",
        "example": "67077212345"
    },
    {
        "country_code": 228,
        "name": "Togo",
        "example": "22890112345"
    },
    {
        "country_code": 690,
        "name": "Tokelau",
        "example": "6905190"
    },
    {
        "country_code": 676,
        "name": "Tonga",
        "example": "6767715123"
    },
    {
        "country_code": 1,
        "name": "TrinidadandTobago",
        "example": "18682911234"
    },
    {
        "country_code": 216,
        "name": "Tunisia",
        "example": "21620123456"
    },
    {
        "country_code": 90,
        "name": "Turkey",
        "example": "905012345678"
    },
    {
        "country_code": 993,
        "name": "Turkmenistan",
        "example": "99366123456"
    },
    {
        "country_code": 1,
        "name": "TurksandCaicosIslands",
        "example": "16492311234"
    },
    {
        "country_code": 688,
        "name": "Tuvalu",
        "example": "688901234"
    },
    {
        "country_code": 1,
        "name": "U.S.VirginIslands",
        "example": "13406421234"
    },
    {
        "country_code": 256,
        "name": "Uganda",
        "example": "256712345678"
    },
    {
        "country_code": 380,
        "name": "Ukraine",
        "example": "380391234567"
    },
    {
        "country_code": 971,
        "name": "UnitedArabEmirates",
        "example": "97150XXXXXXX"
    },
    {
        "country_code": 44,
        "name": "UnitedKingdom",
        "example": "447400XXXXXX"
    },
    {
        "country_code": 1,
        "name": "UnitedStates",
        "example": "12015550123"
    },
    {
        "country_code": 598,
        "name": "Uruguay",
        "example": "59894231234"
    },
    {
        "country_code": 998,
        "name": "Uzbekistan",
        "example": "998912345678"
    },
    {
        "country_code": 678,
        "name": "Vanuatu",
        "example": "6785912345"
    },
    {
        "country_code": 379,
        "name": "VaticanCity",
        "example": "3790669812345"
    },
    {
        "country_code": 58,
        "name": "Venezuela",
        "example": "584121234567"
    },
    {
        "country_code": 84,
        "name": "Vietnam",
        "example": "84912345678"
    },
    {
        "country_code": 681,
        "name": "WallisandFutuna",
        "example": "681501234"
    },
    {
        "country_code": 967,
        "name": "Yemen",
        "example": "967712345678"
    },
    {
        "country_code": 260,
        "name": "Zambia",
        "example": "260955123456"
    },
    {
        "country_code": 263,
        "name": "Zimbabwe",
        "example": "263711234567"
    },
    {
        "country_code": 358,
        "name": "\u00c5landIslands",
        "example": "358412345678"
    },
    {
        "name": "USA(America)+1",
        "country_code": 1,
        "example": "12015550123"
    }
]`

func (self *Code) GetAreaCodes() (data []Code) {
    codes := []byte(Codes)
    json.Unmarshal(codes,&data)

    return data
}


