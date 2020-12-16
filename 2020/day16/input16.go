package main

var inputTest = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

var inputTest2 = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`

var input = `departure location: 37-594 or 615-952
departure station: 50-562 or 573-968
departure platform: 49-584 or 592-971
departure track: 28-727 or 744-957
departure date: 35-930 or 943-965
departure time: 38-811 or 829-962
arrival location: 35-446 or 467-950
arrival station: 29-234 or 245-969
arrival platform: 47-416 or 431-970
arrival track: 38-134 or 160-962
class: 30-493 or 506-953
duration: 43-335 or 346-949
price: 33-635 or 654-953
route: 43-399 or 410-974
row: 32-848 or 854-951
seat: 36-777 or 788-965
train: 35-109 or 122-969
type: 38-673 or 694-960
wagon: 50-168 or 193-971
zone: 46-215 or 232-954

your ticket:
131,103,109,67,127,97,89,79,163,59,73,83,61,107,53,193,167,101,71,197

nearby tickets:
631,272,574,809,804,282,84,288,999,232,831,323,625,629,777,703,279,893,841,360
349,891,70,212,74,889,438,455,842,759,474,438,744,355,522,411,483,835,903,378
201,71,247,245,317,622,362,983,283,899,917,95,789,286,493,412,671,873,399,313
197,382,664,876,212,617,84,491,803,838,95,764,632,460,259,446,160,915,764,468
539,915,378,319,561,206,522,83,468,639,164,658,549,305,488,885,416,164,56,772
326,368,374,481,811,329,777,935,364,594,896,623,867,125,273,792,304,69,199,469
803,748,848,444,700,358,481,357,266,792,658,794,362,314,284,753,773,24,811,374
76,196,666,719,275,557,506,478,805,539,860,80,720,799,850,846,333,394,904,351
760,436,623,475,903,842,808,558,894,675,471,757,328,919,699,873,912,888,105,109
754,616,705,514,876,854,516,585,627,762,321,804,128,750,631,319,561,476,550,857
836,87,367,53,930,659,364,683,385,65,582,359,308,901,58,291,432,664,811,659
943,207,10,274,856,774,836,797,363,694,920,900,891,315,327,777,97,257,929,252
310,489,124,258,847,252,84,255,64,351,355,659,87,924,702,883,84,860,408,833
793,885,298,635,721,814,131,695,210,374,594,901,871,895,553,799,367,477,871,657
262,789,398,531,903,755,886,705,234,866,835,987,127,767,841,377,270,276,922,630
712,878,486,758,708,74,254,634,648,245,473,106,365,667,55,663,512,799,755,268
639,267,476,259,904,258,370,551,911,304,842,694,233,802,256,521,555,755,748,539
134,623,527,358,242,548,50,60,880,882,276,485,509,724,369,577,794,833,916,476
886,809,490,292,362,390,929,64,201,726,468,578,468,871,122,989,897,127,760,628
890,917,475,700,830,544,758,548,271,949,288,509,791,988,249,441,322,212,255,467
527,101,521,711,302,196,849,703,442,807,751,490,522,889,319,702,163,326,807,621
330,206,295,319,384,211,759,382,870,725,723,722,808,836,94,480,11,83,836,250
929,437,50,362,791,370,410,819,53,166,626,274,665,655,471,85,718,475,839,74
471,263,514,55,758,776,568,532,805,871,593,524,261,615,870,896,917,892,579,251
324,627,295,129,513,380,300,523,864,364,868,653,700,310,208,295,859,577,510,96
436,706,863,266,252,359,165,76,463,269,196,251,666,93,232,661,264,477,431,562
548,65,257,611,251,532,798,842,440,553,833,380,83,93,720,829,847,485,756,766
328,720,325,626,271,547,446,295,413,569,488,918,256,581,756,860,309,870,507,792
924,943,264,553,197,672,663,615,334,232,585,658,535,523,947,548,251,272,716,699
716,2,876,256,51,541,86,391,438,106,480,896,870,537,703,107,109,376,549,798
403,805,895,913,444,233,574,927,555,945,839,107,546,547,799,831,899,333,632,906
199,929,297,98,122,771,594,353,303,264,918,935,858,284,854,438,554,362,520,766
18,578,196,762,534,713,62,287,791,279,704,127,309,621,128,375,883,556,393,761
486,83,799,623,877,588,528,581,899,66,888,772,486,811,550,470,89,88,522,323
388,621,619,491,212,295,308,644,127,133,163,549,211,320,871,163,266,747,868,68
855,364,277,122,465,483,254,527,362,292,753,134,943,752,871,809,543,304,234,577
841,533,627,751,408,506,375,667,195,490,514,215,163,947,63,945,289,299,263,357
503,510,209,802,664,87,232,440,70,582,746,268,748,892,868,776,917,920,200,897
735,245,870,386,246,387,749,525,312,73,926,282,287,844,633,545,255,493,549,517
488,446,700,316,916,478,50,792,948,326,553,754,980,763,858,520,470,471,210,658
620,633,432,513,772,730,490,868,349,831,893,94,808,335,62,300,275,726,261,861
771,355,721,860,85,842,906,798,370,914,285,294,302,214,457,718,670,697,626,656
794,437,322,479,990,920,762,312,575,773,764,798,352,753,318,539,793,528,274,616
579,76,352,775,472,534,534,258,533,851,885,758,899,437,252,664,758,233,864,546
925,489,108,549,809,267,701,635,858,53,380,810,162,805,87,650,776,881,724,747
800,838,315,272,214,363,123,570,105,831,713,315,350,54,209,474,477,89,759,303
352,657,539,718,194,765,63,698,394,760,897,131,372,97,557,126,756,857,287,20
69,805,62,406,197,60,279,385,389,348,948,698,831,51,433,623,390,248,718,712
509,256,618,920,211,333,709,323,662,601,771,713,573,773,720,53,98,469,660,316
162,515,211,639,888,523,332,889,928,543,330,809,471,626,625,353,92,626,372,527
538,928,860,383,490,394,584,862,438,718,596,762,665,699,811,777,354,517,867,550
352,560,862,379,193,904,670,580,290,920,654,260,7,536,805,848,301,323,528,215
101,108,288,257,788,106,859,120,835,358,711,306,832,704,64,126,384,559,375,83
858,760,745,302,704,483,86,58,895,783,437,913,87,74,203,288,444,560,80,435
123,590,754,889,270,63,202,478,369,557,517,536,335,663,277,392,746,54,98,479
201,492,97,293,76,380,842,299,488,316,841,769,633,552,309,58,442,841,462,482
270,323,592,517,632,576,555,912,936,310,795,286,100,810,831,161,574,274,288,808
349,943,626,711,196,550,546,51,290,97,541,325,663,867,370,619,838,535,541,849
379,830,210,235,435,699,620,725,621,633,258,131,662,261,807,930,547,837,319,544
716,197,522,671,481,471,251,65,85,517,354,168,547,204,648,510,258,213,841,310
923,670,895,871,749,474,877,97,750,75,892,623,673,90,395,163,707,878,720,3
809,535,70,750,354,921,702,514,199,270,379,284,745,981,708,655,477,480,901,748
865,705,760,332,280,565,555,383,66,710,846,439,751,838,917,276,776,324,201,273
297,84,920,878,720,574,871,884,93,393,807,667,878,847,572,794,908,831,124,71
124,316,911,755,577,233,84,128,93,984,330,803,397,233,442,489,626,664,750,105
802,359,299,398,325,872,762,127,275,775,900,200,834,330,408,911,915,431,672,877
410,648,576,277,806,268,593,101,594,160,474,717,724,890,443,246,701,210,397,905
66,732,373,485,205,804,410,61,660,201,209,834,367,209,334,581,196,532,271,508
276,77,511,861,861,176,884,265,381,206,435,541,328,493,918,796,881,133,91,100
353,518,835,699,348,844,896,718,259,709,161,298,272,885,82,262,331,876,105,243
196,122,328,299,64,252,55,849,658,901,308,561,479,772,797,486,895,410,383,211
703,350,310,392,720,878,99,55,441,23,922,255,708,368,492,796,126,878,659,765
55,558,577,205,92,792,352,830,907,96,945,899,756,873,128,246,213,550,630,636
318,131,922,864,901,483,752,659,401,296,79,293,798,310,925,101,436,898,130,351
869,445,879,508,328,943,727,918,708,903,370,309,163,520,620,86,630,320,478,692
119,315,506,806,126,531,279,664,838,106,885,322,701,319,909,878,486,365,59,291
410,829,281,541,196,476,662,873,899,421,671,368,94,895,747,334,277,511,509,769
628,58,308,828,308,927,844,248,617,946,370,354,160,507,714,416,914,296,577,574
320,755,947,934,443,322,671,318,858,916,751,901,123,82,386,232,889,132,467,493
543,561,241,866,73,921,287,857,715,807,92,161,366,392,375,548,842,892,626,252
313,476,837,437,50,252,762,890,774,0,232,69,476,473,559,131,698,758,129,661
872,889,287,106,752,884,351,262,81,319,454,718,389,493,291,920,486,664,907,633
379,593,376,433,366,871,614,196,715,381,516,305,671,316,581,161,197,383,862,478
322,77,882,347,536,353,97,741,389,867,214,874,288,490,830,98,593,80,867,761
527,713,753,502,234,617,391,298,90,629,634,902,757,389,547,367,511,859,537,524
583,130,286,805,437,266,778,204,632,433,293,702,699,901,328,52,658,536,697,215
888,773,626,724,301,298,256,578,514,888,754,833,431,845,788,12,103,478,506,414
895,326,545,909,668,474,924,836,7,483,413,51,380,539,445,943,772,893,946,870
256,390,87,543,716,73,914,311,373,411,670,949,370,898,347,257,199,627,976,719
201,474,873,796,975,920,385,706,863,380,415,335,766,843,199,193,377,513,618,58
264,293,551,529,579,298,589,198,582,297,838,51,480,268,247,766,908,720,234,902
545,881,316,478,325,560,211,297,901,164,749,316,909,472,125,194,727,883,388,405
585,291,592,533,758,525,918,929,848,261,248,91,86,900,872,271,669,878,296,411
105,384,82,739,870,414,204,926,633,53,467,76,837,133,918,79,166,695,199,79
757,635,379,919,594,356,386,260,60,497,209,889,331,764,774,744,479,809,103,329
940,251,493,302,888,575,870,390,900,388,669,768,879,106,788,900,514,248,232,548
179,633,872,835,906,331,252,764,324,281,713,371,894,923,561,758,943,122,900,512
765,80,78,542,512,721,655,263,889,861,376,504,323,296,552,435,624,467,320,856
756,922,384,209,392,80,315,96,697,903,567,101,251,253,432,100,532,667,846,390
365,916,437,701,442,555,685,332,616,520,318,625,809,435,522,286,574,673,533,709
900,528,3,304,124,209,845,508,705,727,513,930,195,194,874,128,869,319,395,248
284,382,633,992,856,300,352,625,727,791,304,322,882,861,744,902,307,255,792,918
922,106,284,745,323,247,655,890,285,308,862,558,208,124,821,892,61,131,712,362
887,882,201,360,294,256,655,760,299,240,873,632,92,72,62,370,877,714,274,265
921,844,574,791,363,517,388,716,204,767,619,324,363,666,5,267,923,897,311,319
770,316,867,708,890,85,398,323,377,23,879,902,697,793,473,473,896,247,431,108
396,880,348,808,869,377,418,766,88,658,854,747,724,94,260,307,270,632,294,528
382,694,200,579,846,94,949,376,545,233,902,79,194,323,294,923,334,382,98,187
445,846,755,272,717,529,536,244,334,898,777,392,697,78,722,286,920,663,619,399
868,829,774,869,263,63,915,377,752,322,236,701,69,205,916,69,68,711,385,811
802,657,203,51,379,109,912,208,624,769,302,480,700,399,356,271,278,405,750,577
519,583,71,592,63,487,66,580,735,927,256,53,859,319,52,326,517,476,581,892
271,204,128,252,160,384,327,701,775,367,405,920,360,618,331,467,74,617,555,720
840,314,122,97,629,103,774,756,81,412,23,255,838,711,488,948,537,59,439,615
132,328,170,367,103,829,702,896,707,377,270,715,300,412,296,215,923,389,373,280
546,473,883,754,809,509,634,183,593,863,540,299,703,898,370,60,866,260,357,794
631,366,946,574,927,289,655,210,272,380,366,435,487,658,832,178,628,54,354,555
163,307,717,431,546,353,878,984,856,201,910,861,862,575,768,926,334,133,269,788
926,930,727,874,698,664,258,378,188,759,247,283,888,469,557,802,749,562,800,289
781,913,631,53,511,844,347,61,413,126,70,353,673,857,296,126,708,308,93,666
285,295,506,284,617,107,815,437,925,545,615,835,715,128,287,108,357,269,539,354
491,273,256,367,493,892,83,157,253,204,375,130,914,908,195,300,882,762,389,98
78,477,431,84,275,975,562,665,573,792,895,320,70,901,714,559,207,945,667,416
811,861,131,109,249,582,235,839,549,289,584,395,910,105,372,232,717,576,389,123
267,94,764,791,594,276,547,903,709,389,848,829,348,81,107,855,838,801,511,564
72,924,415,260,198,372,434,866,888,765,726,560,368,512,551,444,551,529,519,464
348,899,725,395,52,438,720,891,261,830,440,461,866,920,884,330,73,331,529,699
323,402,944,859,664,283,291,919,165,626,195,762,372,126,318,322,868,395,919,837
122,277,362,523,887,94,668,762,550,371,125,862,842,239,592,123,346,76,398,890
198,52,91,315,100,18,747,249,387,717,724,860,801,199,580,711,870,100,84,792
699,191,486,267,71,411,615,491,844,948,309,624,702,923,630,132,803,355,488,83
252,77,307,250,86,847,482,838,519,575,482,843,574,97,160,777,551,928,593,779
69,871,763,314,122,592,399,510,793,292,301,531,710,872,560,453,672,51,293,912
912,556,694,829,290,99,194,480,210,295,881,375,659,374,123,401,267,130,352,619
522,551,280,803,431,620,323,230,433,550,472,357,721,99,709,755,310,927,804,915
860,61,901,530,386,375,797,279,716,831,278,275,488,922,103,371,234,248,244,352
70,777,482,472,203,554,873,7,368,758,517,413,61,269,584,312,667,623,377,899
351,498,519,133,809,795,204,51,532,829,522,547,416,396,623,574,671,85,246,635
309,524,215,713,511,907,211,103,124,57,826,199,384,799,378,254,214,803,81,766
382,493,872,887,168,994,209,76,657,477,372,562,547,840,886,561,75,803,62,848
248,854,369,445,376,927,844,233,213,881,631,888,353,96,902,538,5,660,833,89
943,475,757,81,515,387,372,630,202,59,165,77,493,281,937,360,760,398,261,930
204,714,281,74,96,273,393,389,503,387,164,809,865,248,854,947,912,79,360,919
831,713,873,59,287,540,118,902,296,534,436,289,164,306,292,900,625,486,89,480
846,168,662,357,346,747,710,401,880,839,696,746,831,855,524,702,210,161,214,525
929,213,627,511,439,362,233,212,830,624,170,797,486,283,395,574,750,618,917,719
322,246,92,616,483,320,386,392,902,661,59,260,523,299,762,492,864,943,533,585
245,378,474,667,665,854,895,573,809,330,690,61,545,315,319,578,670,844,948,74
193,817,533,438,617,624,365,519,626,385,580,287,832,705,245,894,278,755,856,470
255,193,709,84,478,914,328,755,438,352,289,89,873,998,314,770,356,309,776,335
696,832,579,796,669,506,897,57,842,945,213,727,211,751,769,382,259,23,250,864
887,870,713,519,168,831,874,270,270,754,939,103,620,665,201,582,132,254,319,512
410,443,376,324,164,357,83,352,410,256,353,774,239,535,618,194,283,85,527,271
318,209,920,471,487,248,189,96,615,108,468,267,553,909,923,443,617,371,399,346
467,750,68,415,414,773,406,545,790,668,412,512,283,246,624,62,305,528,745,707
844,134,413,168,662,796,507,673,770,742,107,552,622,416,831,556,434,293,128,746
758,81,910,530,367,323,363,483,944,703,639,127,84,869,556,517,480,522,620,324
915,662,906,124,381,331,305,706,522,119,76,477,467,706,697,901,865,107,772,558
886,204,348,164,538,503,894,763,443,507,273,391,622,416,543,891,58,329,895,871
703,243,792,258,592,382,200,394,766,129,725,352,583,304,322,203,895,884,85,163
346,316,290,635,880,767,459,413,763,78,531,83,922,873,165,484,760,919,720,106
749,391,299,834,164,666,658,860,327,699,413,510,909,789,239,287,562,291,774,915
251,257,489,460,442,788,197,619,357,263,205,582,55,129,100,89,246,537,374,659
254,440,60,584,434,73,86,334,888,169,944,232,752,554,727,71,671,530,442,548
769,808,314,280,316,402,873,361,390,312,292,261,538,431,482,106,704,793,286,332
889,576,916,329,873,132,805,886,374,74,539,406,659,352,94,848,266,697,876,909
831,517,615,287,903,880,559,599,97,842,839,316,210,702,857,437,392,295,538,711
624,927,293,861,209,508,260,391,431,843,501,894,168,72,393,473,535,196,923,578
853,797,283,311,303,922,725,105,878,576,203,834,162,334,756,915,553,861,537,799
906,706,767,759,556,373,839,915,61,549,129,769,631,134,596,948,490,726,484,469
165,843,346,489,92,520,772,427,434,907,314,536,384,166,727,492,411,593,311,829
583,489,459,87,443,390,768,71,543,831,286,160,764,777,272,799,166,492,377,526
310,511,270,761,899,656,480,445,716,793,842,796,943,549,741,863,914,304,195,774
884,754,636,347,365,63,760,582,560,513,854,773,442,254,714,908,205,723,672,205
726,313,913,837,467,467,222,386,855,348,329,53,303,298,131,468,898,410,673,665
126,263,578,566,534,302,847,793,64,253,291,486,273,332,919,722,353,68,519,543
884,808,845,346,208,546,829,252,883,298,530,204,875,593,554,771,486,768,177,857
884,379,562,775,252,935,211,273,904,865,410,858,881,763,366,833,92,289,706,866
289,323,626,860,862,876,713,555,103,931,489,544,314,309,899,807,491,83,354,362
463,209,797,878,628,324,621,761,699,908,491,896,214,211,360,263,546,509,715,324
210,749,55,925,277,88,442,722,473,590,433,71,520,847,63,366,261,357,489,273
871,942,214,750,214,720,294,751,269,477,830,472,300,877,383,627,946,745,910,898
892,943,210,71,481,777,932,380,555,577,847,123,131,94,435,856,198,631,83,322
84,864,554,904,351,654,261,767,555,450,790,949,492,531,524,489,846,305,762,353
897,917,7,281,416,320,764,709,69,544,924,762,616,902,877,521,872,263,365,215
842,301,351,487,78,753,947,500,746,399,745,256,727,294,473,64,370,323,351,257
247,80,514,470,473,530,580,758,276,209,906,295,764,512,329,502,946,204,858,576
672,706,788,197,874,886,63,858,886,287,59,360,267,667,283,875,329,355,483,825
631,768,704,715,718,948,89,330,287,394,89,764,245,921,97,314,374,768,202,500
792,703,50,330,876,102,76,664,550,705,332,194,252,924,371,410,583,880,512,612
372,654,505,894,525,839,527,551,803,584,251,859,881,860,318,516,864,387,274,309
360,754,59,301,325,289,349,291,833,313,232,275,547,164,911,102,914,99,501,943
836,346,51,719,773,394,483,259,708,194,765,753,294,408,272,578,759,900,664,559
560,943,895,282,108,277,755,655,660,94,673,370,106,330,76,889,318,835,989,394
443,470,701,869,545,699,618,477,795,167,562,190,530,390,532,382,871,839,944,560
606,165,267,520,833,708,65,281,885,81,886,506,862,543,324,274,949,480,355,882
772,195,95,911,261,19,621,632,573,52,625,663,331,525,432,446,929,537,855,327
632,758,562,544,716,245,100,484,211,276,123,485,263,769,501,324,160,360,350,831
130,256,322,68,620,560,880,763,292,196,866,306,54,4,233,347,347,544,860,886
901,129,618,753,866,594,576,770,352,471,908,890,379,415,801,241,887,539,438,708
755,533,70,801,886,752,755,363,297,825,761,279,332,319,436,889,386,788,707,631
655,868,81,834,365,638,897,87,773,196,717,255,313,722,592,927,671,324,490,252
165,808,755,215,276,581,758,70,486,362,862,10,593,757,748,944,754,263,761,886
674,292,625,579,322,383,68,94,198,839,307,948,252,65,296,548,575,490,541,319
213,719,360,53,906,532,502,667,555,710,380,841,775,357,561,800,59,63,303,232
796,719,838,907,884,714,704,782,547,107,749,103,124,133,775,882,234,254,361,834
847,654,249,573,378,762,833,94,304,944,514,98,440,479,776,733,381,470,246,331
571,194,247,471,272,258,323,513,521,71,906,866,442,385,756,56,716,806,522,365
374,253,695,416,562,301,506,720,668,334,435,838,397,382,923,788,507,478,317,733
56,317,862,296,130,664,791,544,197,232,66,470,246,410,475,105,833,259,234,942
694,788,288,252,269,769,891,912,656,194,862,236,250,297,839,701,479,102,857,762
237,305,695,807,753,949,659,480,61,708,78,96,511,320,839,774,284,710,234,859
381,838,318,747,658,291,370,79,261,881,738,52,269,700,803,833,310,912,908,304
894,193,60,326,102,299,884,835,762,490,840,741,313,331,902,949,333,165,618,520
59,487,285,168,626,594,398,129,412,767,482,574,196,484,674,108,543,355,98,581
882,90,313,697,664,726,582,208,416,839,803,860,20,412,909,87,377,74,296,70
525,377,715,197,726,917,735,64,245,211,361,304,435,655,810,802,671,578,197,211
520,764,384,322,440,2,923,562,797,794,396,594,706,211,87,320,511,263,511,431
635,460,802,896,557,760,553,866,389,857,889,554,124,443,205,756,234,892,389,562
313,753,913,313,842,455,469,273,754,881,166,125,911,303,298,856,372,259,621,659
437,747,671,716,708,660,593,948,668,804,517,81,756,949,272,701,664,837,415,2
929,707,892,904,545,249,287,331,467,621,898,82,866,721,181,946,517,769,286,126
530,317,409,727,759,124,584,552,762,777,88,373,128,272,266,576,865,88,266,347
550,56,839,811,477,388,357,628,50,557,892,857,745,673,586,862,537,318,165,758
540,261,659,492,484,134,993,381,547,767,898,706,76,441,768,286,710,525,411,889
802,320,353,318,749,902,131,899,861,752,521,468,695,789,103,273,541,393,548,997
918,558,203,396,282,235,293,160,535,319,510,59,398,389,711,60,767,86,706,200
533,233,912,301,745,70,259,763,891,298,289,669,392,558,872,126,390,249,335,998
358,55,710,471,53,161,642,90,269,129,709,298,70,97,838,87,372,133,631,554
234,388,540,334,182,492,910,519,89,671,312,397,434,531,629,256,201,379,745,560
68,833,669,589,895,726,266,353,386,770,533,660,664,370,101,249,913,129,534,507
538,930,262,865,513,326,897,233,554,130,901,56,203,700,807,562,480,239,508,556
368,750,713,476,310,293,829,507,241,545,700,788,68,93,126,57,750,282,476,443
827,330,308,916,882,280,874,93,361,292,754,254,889,578,949,93,929,943,96,902
746,381,766,171,101,517,762,673,522,490,524,860,257,581,210,524,397,166,706,276
903,519,894,79,382,927,90,98,824,804,748,81,102,922,665,776,335,294,756,700
875,389,790,59,928,55,289,854,319,407,304,373,320,763,582,748,761,514,878,846`
