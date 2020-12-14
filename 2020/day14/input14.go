package main

var inputTest string = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

var inputTest2 string = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

var input string = `mask = 110000011XX0000X101000X10X01XX001011
mem[49397] = 468472
mem[50029] = 23224119
mem[39033] = 191252712
mem[37738] = 25669
mem[45831] = 238647542
mem[55749] = 1020
mem[29592] = 57996
mask = 100X10XXX101011X10X0110111X01X0X0010
mem[10526] = 1843
mem[2144] = 177500
mem[33967] = 5833292
mem[58979] = 25707732
mask = 100010X011XX00X11011010011101100XXX1
mem[1729] = 1042
mem[30433] = 366890
mem[7726] = 2862
mem[19747] = 52273994
mask = 11001X0011010110X01X011X001X0XX01010
mem[40528] = 32637378
mem[16008] = 30888145
mask = X11X1X0X10X10110011X0001X01001X100X0
mem[27746] = 14986812
mem[45873] = 4381392
mem[26216] = 538203
mask = 1100101011X00010101111001001XX1X0011
mem[30777] = 84408647
mem[6931] = 133210956
mem[5173] = 7497
mem[65147] = 912575421
mem[12597] = 55281597
mem[20417] = 909474
mem[65270] = 1914920
mask = X100XX10XX010X10X110000000X0X1100100
mem[50768] = 3383
mem[59421] = 111147
mem[33900] = 427465715
mem[33084] = 14313354
mem[12648] = 17983288
mask = 11X0100011X011100X00100X01111000XX11
mem[17710] = 60
mem[30013] = 296
mem[48130] = 31469003
mem[45585] = 3231589
mask = X1XX1010110001X0XX000010X0101010X01X
mem[20502] = 15059188
mem[29762] = 375
mem[24169] = 594
mem[24197] = 64508559
mem[8424] = 108440
mem[20424] = 21436372
mask = X10010001XX0X1100000X00000X010X00001
mem[18190] = 448461
mem[37090] = 5353
mem[39942] = 5084619
mem[18325] = 1962539
mask = 10101110110000X010100X1X10XX1X1X1101
mem[9299] = 6164
mem[8421] = 990
mem[23905] = 34526767
mem[44233] = 39766571
mask = 1110X1X01010X1111X0XX1X01110X011001X
mem[53340] = 16503076
mem[59433] = 378862
mem[18190] = 1792792
mem[56498] = 227
mask = 1100100X11000X1X0100X00010X01X010101
mem[65168] = 265913
mem[40500] = 18368848
mem[39558] = 1810777
mem[24300] = 911
mem[47807] = 3491
mem[6201] = 267177
mem[17369] = 21952
mask = 1111101010100X0111001X10011XX1110100
mem[32283] = 17550
mem[55129] = 56452456
mem[7945] = 2961
mask = 1X00101X0101001011101000010XX1001100
mem[1120] = 7335
mem[65276] = 493090
mem[17104] = 220
mask = 11001X101101101110111XX01001110X0000
mem[15933] = 859
mem[50326] = 3145522
mem[48794] = 367683
mem[24561] = 57849668
mem[43526] = 103212
mem[33478] = 20703997
mask = 11001010111X01111001100X100X110X0110
mem[718] = 1589870
mem[8424] = 1123972
mem[966] = 7551
mask = 11X01010110001X00000X1X0101X10000000
mem[16160] = 26953
mem[16417] = 419431373
mem[54811] = 430477
mem[4340] = 180411
mask = 10X0X00011X100101X1X1010X1X111X00X10
mem[37425] = 922346
mem[289] = 810051
mem[58526] = 86518
mem[374] = 92968
mem[37165] = 6023
mem[61397] = 8223350
mask = 1X001000X1X11X100000100X1011111X1110
mem[43693] = 743
mem[9418] = 1128022
mem[11571] = 47294995
mem[449] = 52713877
mask = X1XX1X1011000110XX0001X01XX001000000
mem[29924] = 1125544
mem[10782] = 342783
mem[15523] = 218611
mem[8009] = 1866
mask = 10XX011X11000X001X100110100111110100
mem[40200] = 54187
mem[19587] = 45108
mem[50857] = 1309
mem[18658] = 11992852
mask = 1X001000XX1001101100001X010X00001001
mem[21333] = 7608315
mem[9746] = 259920
mem[63211] = 126262747
mem[59768] = 65880460
mask = 11X11X100X1000X01X00X011110011111001
mem[59121] = 293545
mem[14925] = 17664197
mem[60673] = 1663
mem[45765] = 195645
mem[33094] = 58807
mask = 1X0010X011X1001X1001X0110XX0000000X0
mem[32288] = 20128
mem[50857] = 1189904
mem[18918] = 913
mem[7726] = 50248226
mem[22429] = 18716
mem[7848] = 272580
mask = 01XX100010010X1X0X1X00X1X110X1100000
mem[40002] = 72763964
mem[20337] = 36642182
mem[19538] = 230553
mem[11992] = 8409
mask = 11001000X11X111000XXX011X0111000111X
mem[63876] = 969
mem[1336] = 5375872
mem[31377] = 5165
mem[41185] = 161434
mem[38292] = 634
mask = 1X0010101X00011010X1X10101X011XX1010
mem[59768] = 10746
mem[27445] = 2335
mem[26812] = 58960
mem[40116] = 104178572
mem[40702] = 48107383
mask = 00000001X0X1011XX011X00X01111100X11X
mem[18702] = 150975
mem[62270] = 502767513
mem[6931] = 15732227
mem[12320] = 3799
mem[29975] = 99827
mask = 1100100X1100XX1000001X00000010X00110
mem[17011] = 11786404
mem[25382] = 98379404
mem[35946] = 791341
mem[49767] = 719
mem[11664] = 738
mask = 000XX0011011011000111010X11X111001X1
mem[53375] = 513
mem[776] = 31438875
mem[26228] = 6566431
mem[62653] = 352
mem[8883] = 13700386
mem[17292] = 66198210
mask = 1110XX001110X110000X1000010001100100
mem[65123] = 23447
mem[53419] = 1784255
mem[32201] = 472209
mask = 10X0X0001101X01000100X000011001001X1
mem[45831] = 4941253
mem[17666] = 7
mem[52211] = 250885474
mem[33711] = 38546733
mem[54654] = 108397257
mem[54577] = 7660097
mask = 110010001100X11010000100010XX110X010
mem[48263] = 203073
mem[46274] = 329424784
mask = XX00000XXX0X0X101011X11001001100X111
mem[46639] = 245946590
mem[24300] = 769
mem[54106] = 23763
mem[35221] = 970549
mem[23333] = 322574122
mem[32283] = 9651
mem[38047] = 804
mask = 01X00XXXX101011010110010X10001010X01
mem[52675] = 50846938
mem[43900] = 69746023
mem[54409] = 1786723
mem[30815] = 4286
mem[37] = 4678667
mask = 1X0X1X1010000101110XXX0X0001011X11X0
mem[40133] = 158160
mem[13432] = 984
mask = 1110X000011X0110001X10010100X0001000
mem[28551] = 97731716
mem[21298] = 1506013
mask = 110XX01X10100110100X01X001001111X001
mem[5461] = 26227
mem[4650] = 1623
mask = 110110X0110XX110010000001000101X0001
mem[18167] = 5899011
mem[45492] = 18393
mem[13148] = 171228654
mem[59109] = 52915776
mem[37] = 1212
mask = 111X1000110X01X01000110X0X00110X1011
mem[13148] = 11483926
mem[33841] = 22637
mem[60690] = 16733
mem[35555] = 125444
mem[19999] = 10615
mem[49083] = 57306580
mem[2958] = 113424903
mask = 1X00X0X011X010100110011XX0X110000X10
mem[16044] = 2922
mem[58981] = 99
mem[17754] = 41326186
mem[57873] = 767731
mask = 0000110011011X10101110X001X0X1101X00
mem[53194] = 54243360
mem[15023] = 258913
mem[37425] = 678
mem[36057] = 2068683
mem[6540] = 145235
mem[46515] = 5824196
mask = 1XX0X00X1X0000X010101X00001000101011
mem[42985] = 2821
mem[17666] = 178146480
mem[35891] = 111717
mem[37731] = 280009
mem[45606] = 27440
mem[14991] = 26844935
mask = 01X0X010X10101101011001001X001110100
mem[45084] = 377769619
mem[58867] = 3974659
mem[48117] = 374339883
mem[1141] = 1632150
mask = 1010X1X0X10000X01010X01010011111XX0X
mem[45122] = 3222
mem[2300] = 16240
mem[58035] = 6201
mem[40871] = 16257123
mem[24285] = 12751
mem[57579] = 24679
mask = X1XX000X11X01110010X0000000011110001
mem[10424] = 280052
mem[36995] = 398570435
mem[160] = 6920
mem[42829] = 3609
mem[49083] = 76851
mask = 11001XX011011X100X00000011011X0X011X
mem[24655] = 976
mem[56929] = 23232
mem[63878] = 63802677
mem[19968] = 15946871
mask = 1100101011X1XX1110X100X0X0X11100X100
mem[29216] = 2636405
mem[3744] = 344561
mem[60039] = 11290842
mem[45769] = 9817
mem[52361] = 250607
mem[43526] = 6568339
mem[28084] = 47601
mask = X10010101101101X1001X001X00011001000
mem[33294] = 65108649
mem[39245] = 1562390
mem[18702] = 880826
mask = X110X00000100100110000X1X01111X1X011
mem[62194] = 21047
mem[56498] = 8195045
mem[19165] = 7369328
mem[13257] = 536577153
mask = XX00100X0X00111000X1X000X0X011110010
mem[6133] = 795
mem[40702] = 1159
mem[49254] = 936358
mem[20224] = 33223599
mask = 10001000111100101X11100XX101111000X1
mem[12938] = 250757561
mem[8424] = 795011162
mem[6681] = 444240
mask = XX001000111011100000X00100X0XX00000X
mem[34480] = 317
mem[642] = 6967048
mem[27203] = 3233
mask = X100100011X011X0000X11X0XX1X10010X11
mem[9519] = 6889363
mem[48618] = 56235450
mem[45084] = 3643761
mem[22351] = 128696
mask = 1X00101011XX001X1000101100X01X011XX0
mem[43960] = 1039599408
mem[29626] = 8360561
mem[31260] = 256268877
mem[50373] = 1706687
mem[24558] = 753
mask = 111011X000X0XX1010000X01011XX0XX1101
mem[37425] = 562
mem[32022] = 231573
mem[52827] = 36198
mem[1203] = 187184
mask = 11X0101011110010100X00XX10001010001X
mem[27236] = 50136301
mem[36499] = 18610469
mem[23179] = 193
mem[2602] = 520829
mask = 1X1011000010X1100000X10110X00X1001X1
mem[58650] = 17011909
mem[30325] = 1792
mem[21629] = 146235659
mask = 1X000010110X011001X0000X10000X10X1X1
mem[56201] = 65276
mem[45769] = 27536
mem[63677] = 76310013
mem[32288] = 38391157
mem[2732] = 553
mem[21153] = 674
mask = 110010001100X1X0000X00X111X001001010
mem[20650] = 1639
mem[37394] = 2020484
mem[10598] = 46526712
mem[18167] = 18124530
mask = 1100X00XX100X1100X0X000000X0111X000X
mem[49767] = 503
mem[23201] = 170673423
mem[37394] = 2873290
mask = 11001010X10X00X01011X00XX00101110001
mem[12597] = 4852003
mem[45585] = 241
mem[6816] = 252644
mem[55923] = 3191
mem[59547] = 165517
mem[10853] = 1769226
mem[37991] = 238
mask = 11001000110X11100X0001XX100110X0010X
mem[22590] = 60452
mem[59590] = 18099
mem[50198] = 21070930
mem[5308] = 5434548
mem[7675] = 6165055
mask = 11XX1001010011100X01010X011111010010
mem[1312] = 30936
mem[48263] = 2432189
mem[58137] = 3014
mask = 1000XX101X0100000100X01XX01000000100
mem[27203] = 610377
mem[11538] = 1967996
mem[32288] = 26776
mem[7745] = 330
mem[43272] = 1383
mem[18399] = 6837
mask = 111X1000X11001X0X0XX11010100X0100X1X
mem[17790] = 7714503
mem[54074] = 32718129
mem[5352] = 1054
mask = 11001110001001011100X00101X000X10X00
mem[18972] = 783671072
mem[59100] = 54416
mem[59256] = 621566
mem[31471] = 591
mem[2884] = 2615461
mem[51] = 790
mask = 11101110XX1X0XX11000X1110X1011011000
mem[20222] = 882
mem[27763] = 7914
mem[32294] = 145898791
mem[33294] = 254866534
mem[24498] = 96614215
mem[45811] = 59795025
mask = 1100X0X0110001100X0001XX000X1X100X01
mem[31950] = 1352
mem[10853] = 766
mem[3709] = 5103902
mask = 110000X0110001100100011X10X001000XX1
mem[30788] = 426
mem[19168] = 42816
mem[27236] = 45039961
mem[21448] = 8723202
mem[48744] = 11100131
mem[37] = 3152
mask = 1X0X1010110X0XX001000X00101XX1000101
mem[25916] = 52795821
mem[1763] = 5368864
mem[13148] = 378742711
mem[10853] = 4345777
mem[64644] = 8348080
mask = X11011000010011110001X011X100X00000X
mem[45572] = 172063
mem[39527] = 19012657
mem[24187] = 758186
mem[65360] = 97
mem[37394] = 2174365
mem[22260] = 170639258
mem[11465] = 45577
mask = 11011X10XXXX01X111000111111X001X1001
mem[33046] = 40550135
mem[55128] = 487381
mem[48068] = 7496218
mem[24391] = 15110
mask = 11XX1010X0X00XXX11000011110X111X1001
mem[56260] = 2566
mem[40500] = 11350955
mem[16482] = 470
mask = 110110X01100011001X0010X101X001011X0
mem[11839] = 1035
mem[27964] = 455
mem[21803] = 109558713
mem[20663] = 1163
mem[12474] = 36111
mask = X10010XX1100X11X0010010101001001111X
mem[15464] = 51852071
mem[59553] = 620
mem[28798] = 248109182
mask = 11X000001X0X0110000X01111X111X110X01
mem[22073] = 3262
mem[17070] = 33580553
mem[11911] = 2692
mask = 1X100000X1110110001011001X011X110X0X
mem[10155] = 747210936
mem[57352] = 1286964
mem[12621] = 3237187
mem[58650] = 17477
mem[13702] = 759723
mask = 11X010X01XX00110XX00X1X001X011101011
mem[38922] = 205
mem[45585] = 99912
mem[53888] = 48069
mem[44233] = 1788
mask = 1110110X00X00110000010X00000100X0XX1
mem[11817] = 4458
mem[58578] = 4618
mem[27624] = 173091087
mask = XX0110X011010100010X00101010010111X1
mem[14010] = 3227436
mem[492] = 6881522
mem[5687] = 2478716
mem[12673] = 14623351
mem[53812] = 140355
mask = 110010X0110X0X10X0XX011100X01X10001X
mem[3709] = 6604
mem[19531] = 29597
mem[38507] = 2150917
mem[59768] = 56061470
mem[54074] = 4058
mask = 11101XX0X010011XX000X0XX01100X111010
mem[60451] = 1612
mem[42190] = 37042
mem[20069] = 96923
mem[21689] = 592
mem[1247] = 8651172
mem[48777] = 40334782
mask = XX1010X011100110X1100X10010X1000X010
mem[18179] = 45826
mem[33139] = 838529759
mask = 01X010001X0X011X00X01001111XX1100X1X
mem[33377] = 1739
mem[35840] = 6769704
mem[14441] = 22736868
mem[22630] = 1700619
mask = X110100X1010X11011001110010X0XX0101X
mem[39736] = 41854026
mem[5320] = 172367335
mem[24297] = 10252548
mask = X11X101011000110000XX010000001101010
mem[46338] = 393890
mem[55364] = 969778
mem[32531] = 267024186
mem[704] = 3741
mem[50527] = 218631
mask = 11X0X00XX0X001X011000XX0000011110011
mem[374] = 216
mem[30607] = 788
mem[17248] = 1204
mem[21290] = 356140
mem[11719] = 12908630
mem[5338] = 98892
mask = XX00100011010110X0001X001000X0100X1X
mem[65147] = 16521590
mem[50886] = 4725
mem[29082] = 846562
mem[26065] = 24418411
mem[56929] = 403301
mem[489] = 2168
mask = 1000100011010110XXX01010X10010XX0010
mem[36629] = 2579
mem[20122] = 2088646
mem[2798] = 2730
mem[20062] = 232728360
mem[27203] = 3015
mem[47864] = 10789801
mask = 1XX00X001100X0100110X1000XX110X1X011
mem[25357] = 1792
mem[25872] = 56296
mem[1964] = 26399389
mask = 1X00100011010X1X101X0X1X01000100000X
mem[61985] = 42984
mem[19168] = 394494472
mem[30890] = 213
mem[58650] = 581887
mem[50658] = 2763
mask = 10X0XX00110XX010X010000X0X010X100010
mem[42533] = 153956
mem[58867] = 470369
mem[59441] = 176314
mem[53867] = 1949039
mem[59547] = 2730
mask = 1010X00011XX0X10XX100110X0011X110011
mem[62586] = 2420073
mem[56548] = 7379
mem[50515] = 2405893
mask = 10001X00110110101010XX1X010X001X0101
mem[33606] = 211
mem[3055] = 106121132
mem[12465] = 823
mask = 11001X10X0XX010111001011X10XX1X0110X
mem[40544] = 476165
mem[23184] = 280716800
mem[12930] = 63529
mem[46092] = 2274568
mem[38292] = 1051815696
mem[48873] = 1125500
mask = XX00XX10110101100X1X11101000001101X1
mem[41185] = 228856274
mem[20806] = 6455676
mem[10598] = 9012
mem[18273] = 3452904
mem[43960] = 117914
mem[8412] = 16428888
mem[56401] = 15927
mask = 110010XX110101100X000011X0010X010X10
mem[15287] = 1639969
mem[53222] = 60401483
mem[21266] = 5960
mem[32861] = 7234007
mem[61866] = 36199944
mem[19264] = 550701
mask = X0X0X100110110X01X1X001X110XX0110001
mem[13148] = 3209260
mem[49522] = 22692520
mem[45544] = 532538
mem[38922] = 127394
mem[53475] = 850137
mem[41422] = 762248838`
