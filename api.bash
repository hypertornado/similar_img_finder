curl -XGET 'localhost:9201/images/img/_search' -d '{
    "query" : {
      "bool": {
        should: {
          "terms": {
            "f": [6,7,9,17,22,23,27,28,30,39,43,44,45,46,54,60,61,63,64,65,67,69,72,75,80,90,91,92,97,99,100,106,107,108,109,112,113,115,117,127,129,134,135,137,140,142,143,146,151,155,157,158,159,161,164,166,168,169,173,177,185,192,195,199,202,204,207,208,212,214,215,219,222,223,224,226,228,233,234,236,237,243,244,251,257,265,269,271,273,275,279,280,281,283,292,293,299,306,309,311,313,314,317,320,322,328,331,334,335,336,339,342,344,346,347,349,357,360,362,365,366,370,375,377,378,380,381,383,384,391,393,394,395,397,398,399,400,401,402,407,410,415,418,421,424,426,428,435,445,448,449,453,455,459,462,467,468,472,474,475,478,484,489,491,494,496,501,502,504,505,506,509,510,513,514,521,523,524,526,528,530,537,538,541,545,552,553,554,557,561,569,577,581,591,592,593,594,596,613,615,616,623,631,633,638,644,645,647,648,653,656,657,662,663,664,665,668,669,676,686,694,695,698,702,719,720,724,729,731,732,736,745,750,753,755,758,761,763,764,765,767,771,772,774,777,778,787,789,790,791,792,793,796,797,798,799,801,803,805,808,811,814,816,822,824,828,832,835,837,848,850,851,852,859,861,863,865,867,872,874,875,877,880,883,886,887,891,894,896,898,906,909,911,912,916,917,918,921,928,929,931,932,940,942,943,945,946,947,950,951,952,954,955,957,959,960,961,962,963,965,968,971,984,985,987,990,995,998,1004,1007,1009,1012,1013,1019,1020,1023,1024,1026,1029,1036,1038,1040,1041,1046,1047,1048,1049,1052,1058,1061,1062,1072,1077,1084,1094,1096,1097,1100,1101,1102,1105,1107,1109,1110,1111,1112,1113,1115,1120,1124,1127,1128,1129,1132,1144,1154,1155,1156,1160,1163,1169,1173,1175,1177,1180,1182,1188,1191,1200,1204,1206,1210,1211,1212,1220,1224,1226,1229,1235,1236,1237,1243,1244,1246,1247,1248,1250,1251,1252,1257,1259,1262,1263,1264,1268,1275,1279,1280,1283,1285,1286,1288,1289,1291,1292,1293,1295,1297,1300,1304,1305,1310,1311,1312,1314,1323,1325,1327,1330,1331,1332,1333,1334,1336,1337,1339,1340,1344,1348,1350,1354,1360,1364,1369,1376,1383,1384,1389,1394,1395,1399,1403,1404,1407,1413,1414,1418,1419,1421,1422,1425,1426,1429,1432,1433,1435,1437,1440,1441,1444,1446,1452,1454,1455,1456,1459,1463,1466,1471,1478,1481,1482,1485,1487,1491,1492,1493,1494,1496,1497,1498,1499,1500,1503,1515,1516,1518,1520,1521,1523,1524,1525,1527,1538,1539,1542,1544,1546,1553,1557,1559,1563,1564,1566,1572,1573,1577,1581,1584,1591,1598,1601,1606,1615,1616,1617,1618,1621,1624,1625,1627,1628,1640,1646,1650,1652,1660,1661,1664,1667,1668,1669,1674,1678,1679,1681,1691,1694,1695,1699,1702,1703,1707,1709,1711,1716,1717,1718,1729,1732,1735,1736,1737,1740,1743,1747,1748,1751,1753,1755,1756,1761,1765,1768,1770,1772,1773,1782,1784,1789,1790,1792,1795,1800,1805,1807,1808,1816,1818,1822,1823,1828,1829,1834,1837,1841,1844,1845,1858,1860,1865,1869,1871,1872,1875,1885,1889,1891,1894,1897,1900,1903,1904,1905,1912,1914,1915,1919,1922,1924,1928,1933,1936,1937,1944,1945,1946,1949,1952,1953,1954,1956,1965,1970,1971,1972,1973,1976,1978,1979,1981,1988,1991,1992,1996,1997,1998,2000,2003,2007,2010,2012,2016,2022,2026,2029,2030,2033,2036,2037,2039,2041,2042,2044,2046,2057,2063,2064,2068,2069,2072,2073,2074,2079,2082,2085,2086,2094,2096,2098,2101,2102,2105,2107,2117,2120,2123,2125,2126,2128,2129,2130,2133,2137,2139,2142,2143,2145,2149,2153,2157,2160,2166,2168,2170,2176,2177,2180,2182,2184,2185,2186,2194,2196,2198,2199,2200,2203,2204,2207,2208,2209,2211,2214,2217,2222,2223,2224,2226,2227,2233,2234,2241,2244,2246,2247,2250,2251,2252,2254,2258,2261,2264,2267,2271,2274,2275,2276,2279,2282,2283,2286,2297,2305,2306,2307,2311,2316,2320,2321,2323,2324,2326,2329,2340,2342,2343,2345,2346,2348,2352,2358,2363,2369,2373,2374,2376,2377,2378,2380,2382,2383,2384,2386,2387,2389,2394,2395,2396,2397,2398,2400,2405,2406,2414,2422,2423,2424,2426,2427,2428,2429,2434,2436,2437,2439,2440,2441,2442,2443,2447,2448,2449,2450,2451,2456,2457,2461,2462,2465,2466,2467,2468,2472,2482,2483,2490,2491,2494,2495,2496,2497,2503,2506,2509,2528,2529,2531,2535,2536,2540,2541,2543,2544,2546,2548,2550,2554,2559,2562,2564,2568,2570,2571,2573,2574,2577,2582,2583,2584,2593,2601,2603,2604,2605,2608,2611,2616,2620,2621,2622,2624,2627,2629,2634,2636,2640,2641,2642,2643,2646,2656,2658,2661,2662,2663,2664,2666,2667,2668,2670,2671,2674,2680,2683,2691,2693,2695,2697,2698,2701,2705,2707,2714,2719,2723,2727,2731,2736,2737,2739,2740,2741,2754,2756,2760,2761,2762,2763,2764,2768,2769,2772,2773,2774,2776,2779,2786,2788,2793,2795,2797,2803,2807,2810,2812,2815,2818,2829,2831,2842,2846,2847,2851,2852,2855,2856,2859,2862,2864,2866,2869,2890,2893,2894,2896,2899,2901,2903,2904,2909,2910,2911,2914,2917,2924,2925,2934,2935,2937,2940,2941,2943,2945,2949,2950,2951,2953,2964,2965,2966,2967,2972,2973,2976,2979,2984,2985,2986,2989,2991,2993,3000,3002,3008,3009,3011,3014,3017,3024,3026,3032,3038,3040,3041,3043,3048,3050,3052,3053,3064,3069,3073,3076,3078,3080,3081,3086,3087,3090,3091,3093,3094,3095,3101,3109,3111,3122,3129,3136,3138,3147,3148,3152,3155,3158,3159,3160,3165,3173,3178,3179,3182,3185,3189,3194,3201,3202,3205,3208,3209,3210,3211,3213,3215,3218,3220,3221,3231,3232,3238,3242,3244,3247,3251,3255,3261,3263,3272,3273,3285,3286,3287,3290,3292,3298,3300,3302,3304,3305,3309,3310,3319,3321,3324,3326,3333,3335,3340,3341,3345,3350,3354,3356,3362,3364,3365,3367,3368,3379,3380,3382,3385,3388,3389,3391,3392,3399,3403,3407,3409,3417,3420,3421,3426,3430,3441,3444,3447,3449,3453,3455,3457,3462,3466,3467,3469,3473,3475,3476,3480,3481,3485,3486,3487,3490,3493,3494,3495,3498,3499,3501,3504,3505,3508,3510,3511,3517,3521,3528,3532,3534,3535,3538,3540,3542,3549,3559,3564,3565,3566,3567,3569,3573,3574,3576,3583,3585,3586,3588,3589,3596,3599,3603,3606,3607,3608,3609,3610,3614,3617,3620,3624,3626,3628,3629,3632,3635,3639,3641,3643,3645,3651,3656,3661,3664,3666,3668,3671,3676,3677,3678,3679,3682,3687,3688,3690,3691,3692,3699,3700,3702,3704,3720,3729,3730,3731,3735,3737,3739,3744,3745,3748,3755,3763,3764,3765,3766,3769,3772,3774,3775,3776,3778,3780,3785,3786,3789,3793,3794,3795,3802,3808,3809,3815,3817,3822,3825,3828,3829,3830,3834,3836,3841,3848,3850,3853,3854,3856,3857,3859,3861,3862,3864,3870,3877,3878,3880,3881,3886,3890,3892,3899,3902,3903,3904,3905,3906,3912,3913,3923,3924,3925,3928,3929,3930,3933,3934,3935,3939,3943,3960,3962,3965,3969,3973,3976,3981,3983,3990,3991,3992,3993,4002,4003,4005,4007,4008,4009,4018,4019,4022,4023,4025,4038,4039,4041,4042,4043,4045,4046,4047,4048,4049,4051,4054,4055,4057,4058,4059,4063,4070,4071,4073,4074,4081,4082,4083,4085,4086,4089,4095],
            minimum_match: 0
          }
        }
      }
    }
}'
