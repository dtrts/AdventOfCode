package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	log.Printf("Starting...")

	testInputPtr := flag.Bool("testInput", false, "a bool")
	flag.Parse()

	if *testInputPtr {
		fmt.Println("Using test Input")
		input = inputTest
	}

	signals := parseInput(input)
	fmt.Println("Part 1:", countOutputsPart1(signals))

	// Looking at inputs, you can determin a list of condidates
	numberLengthMap := map[string]int{"0": 6, "1": 2, "2": 5, "3": 5, "4": 4, "5": 5, "6": 6, "7": 3, "8": 7, "9": 6}
	// lengthNumbersMap := map[int][]string{2: {"1"}, 3: {"7"}, 5: {"2", "3", "5"}, 6: {"0", "6", "9"}, 7: {"8"}}
	fmt.Println(numberLengthMap)

	part2 := 0

	for _, signal := range signals {
		fmt.Println(signal)
		candidates := generateCandidates(signal.inputs)
		foundPattern := map[string]string{"1": candidates["1"][0], "4": candidates["4"][0], "7": candidates["7"][0], "8": candidates["8"][0]}
		// candidates = removeFound(foundPattern, candidates)

		// Top is letter in 7 but not in 1
		// top := findTop(foundPattern["7"], foundPattern["1"])

		// For all numbers len5 only num3 contains num7
		foundPattern["3"] = findNum(candidates["3"], foundPattern["7"])
		candidates["3"] = []string{}
		candidates = removeFound(foundPattern, candidates)

		// For all numbers of len6, only num9 contains num4
		foundPattern["9"] = findNum(candidates["9"], foundPattern["4"])
		candidates["9"] = []string{}
		candidates = removeFound(foundPattern, candidates)

		// For remaining 0 candidates (0 and 6) only 0 contains 1
		foundPattern["0"] = findNum(candidates["0"], foundPattern["1"])
		candidates["0"] = []string{}
		candidates = removeFound(foundPattern, candidates)

		foundPattern["6"] = candidates["6"][0]
		candidates["0"] = []string{}
		candidates = removeFound(foundPattern, candidates)

		topRight := findTop(foundPattern["8"], foundPattern["6"])
		// 2 contains top right, so top right will have something from 2 in it.
		if findTop(topRight, candidates["2"][0]) == "x" {
			foundPattern["5"] = candidates["2"][1]
			foundPattern["2"] = candidates["2"][0]
		} else {
			foundPattern["5"] = candidates["2"][0]
			foundPattern["2"] = candidates["2"][1]
		}

		candidates = removeFound(foundPattern, candidates)

		outputString := ""
		for _, output := range signal.outputs {
			for k, v := range foundPattern {
				if sortString(v) == sortString(output) {
					outputString += k
				}
			}
		}
		// fmt.Println(signal.outputs, foundPattern)
		outputNum, err := strconv.Atoi(outputString)
		if err != nil {
			panic(err)
		}
		// fmt.Println(foundPattern)
		// For all numbers of len6, num6 is missing num7
		fmt.Println(candidates, foundPattern, signal.outputs, outputNum)
		part2 += outputNum
	}
	fmt.Println(part2)
	log.Printf("Duration: %s", time.Since(start))
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func findNum(options []string, num string) string {
	// find a string which contains all of the characters of 7
	for _, option := range options {
		charCount := 0
		for _, numChar := range num {
			for _, optionChar := range option {
				if numChar == optionChar {
					charCount++
				}
			}
		}
		if charCount == len(num) {
			return option
		}
	}
	return ""
}

func findThree(threes []string, seven string) string {
	// find a string which contains all of the characters of 7
	for _, three := range threes {
		charCount := 0
		for _, sevenChar := range seven {
			for _, threeChar := range three {
				if sevenChar == threeChar {
					charCount++
				}
			}
		}
		if charCount == len(seven) {
			return three
		}
	}
	return ""
}

func findTop(seven, one string) string {
	for _, sevenRune := range seven {
		found := false
		for _, oneRune := range one {
			if sevenRune == oneRune {
				found = true
			}
		}
		if !found {
			return string(sevenRune)
		}
	}
	return "x"
}

type signal struct {
	inputs  [10]string
	outputs [4]string
}

func countOutputsPart1(signals []signal) (count int) {
	for _, signal := range signals {
		for _, output := range signal.outputs {
			if len(output) == 7 || len(output) == 4 || len(output) == 3 || len(output) == 2 {
				count++
			}
		}
	}
	return count
}

func parseInput(input string) []signal {
	lines := strings.Split(input, "\n")
	signals := make([]signal, len(lines))
	for i, line := range lines {
		inputOutputSplit := strings.Split(line, " | ")
		inputsSlice := strings.Split(inputOutputSplit[0], " ")
		outputsSlice := strings.Split(inputOutputSplit[1], " ")

		for j := 0; j < 10; j++ {
			signals[i].inputs[j] = sortString(inputsSlice[j])
		}
		for j := 0; j < 4; j++ {
			signals[i].outputs[j] = sortString(outputsSlice[j])
		}
	}
	return signals
}

func generateCandidates(inputs [10]string) map[string][]string {
	lengthNumbersMap := map[int][]string{2: {"1"}, 3: {"7"}, 4: {"4"}, 5: {"2", "3", "5"}, 6: {"0", "6", "9"}, 7: {"8"}}
	candidates := map[string][]string{}
	for i := 0; i < 10; i++ {
		stringNum := strconv.Itoa(i)
		candidates[stringNum] = make([]string, 0)
	}
	for _, input := range inputs {
		for _, number := range lengthNumbersMap[len(input)] {
			candidates[number] = append(candidates[number], input)
		}
	}
	return candidates
}

func tidyCandidates(foundPattern map[string]string, candidates map[string][]string) map[string][]string {
	for candidateSolve(candidates) {
		for num, options := range candidates {
			if len(options) == 1 {
				foundPattern[num] = options[0]
			}
		}
		candidates = removeFound(foundPattern, candidates)
	}
	candidates = removeFound(foundPattern, candidates)
	return candidates
}

func removeFound(foundPattern map[string]string, candidates map[string][]string) map[string][]string {
	for _, v := range foundPattern {
		for k2, v2 := range candidates {
			candidates[k2] = deleteMatch(v2, v)
		}
	}
	return candidates
}

func candidateSolve(candidates map[string][]string) bool {
	for _, options := range candidates {
		if len(options) == 1 {
			return true
		}
	}
	return false
}

func deleteMatch(s []string, c string) []string {
	for i, v := range s {
		if v == c {
			return remove(s, i)
		}
	}
	return s
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

/*
lengthNumbersMap := map[int][]string{2: {"1"}, 3: {"7"}, 4: {"4"}, 5: {"2", "3", "5"}, 6: {"0", "6", "9"}, 7: {"8"}}
"Correct" Assignment:
0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

 5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg


Test Assignment:
 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc

1,4,7,8

top is in 7 not in 1



*/

type digit struct {
	top         bool
	mid         bool
	bottom      bool
	topLeft     bool
	topRight    bool
	bottomLeft  bool
	bottomRight bool
}

type digitMap struct {
	top         string
	mid         string
	bottom      string
	topLeft     string
	topRight    string
	bottomLeft  string
	bottomRight string
}

var inputTest string = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

var input string = `dg debg edgfc afbgcd efdbgc gdc bfdceag bfdec febcad gfaec | dcg bdaegfc egbd dcgfe
ebafgc feadcb deagf gdacf gae edbg ge bcefagd dgbefa ebfda | eg bgfaed fgdaecb agcdbfe
gfcbe bfg fg efbca bcdgea dfge gdcfba dfbgcae ebcgd cgbedf | fg cegabdf cebdg cgebdf
bdafg gbdfea ge fegacd gabcdf eadbg egd cgafdbe acdbe gfbe | cebda bfgda bfdgace adcbefg
fdb fdbce cbade fcdagb fcgdbe cgbfea fd dgfe afbdecg fgecb | fedg cdebf fd dafcbg
dcgefa fcgbde bdecg cde edfb cfdeagb feagbc de efbgc bdcga | ed febgc fabgec cegadf
fbcaeg cfgdae egc edgcf fdcag ebdcf dgae bagcfd eg eagfcbd | gdcfeab adge cbgfad cedfb
edagf bfg dbcfge cgeab aebfg dcgeaf fb daebfg eafcbdg bfad | ebafg gfedac bfg ecagb
bafgce agcdb gaefc agcfed defbag cbgfa afb fb ecfdgba cebf | bf bf dacbg cfgab
egcdbaf begdaf acgfdb fcba cfbdg cbd gdabf fcdge bc ecgdba | dcb dgefc dcbgf ebcfgad
cgbef faedc acfbgd fcebdg ebfcag febdacg ba gaeb ebcfa abf | cefagb ba dceaf bfa
gbaed cbafeg ebc bcgedf gbcaedf becga geafc cfba adgfec bc | gaedb aegbcf cb abfc
dgbec fd acgdfbe bfaced efad cagbef dcgabf bcfed cafbe bdf | dgcbefa ecdgafb fdb fbd
dcagfeb daf decgf fa dcfag gdacfe dbgcef cbdga aedcbf fega | gafcd fad afd fdaceb
bcga gfcad cgdbf caf dbcfage ca cbdafe efdga egdcfb dbcagf | fdecbga bfdcg fca efabdc
eafd adfgecb cgfdeb egfab ea cgfba gea dfgeb debacg eafbgd | ae efda fdae fade
dce acdfg gbdace ecbgdf gadbe eacb dfgceab agedc gbafed ec | fgedacb dec bdgfcae ced
ebfda fabcde adgcfe fdc cedb eagbfcd bfacd gedfba cd cbfga | bgdfcea dcf afcgb ebfacd
cebadf ecdag gadb bcage deabcg eda fbaecg fdceg fagebdc da | edbcag badg ad gbad
dbafegc fgdaec acegd aedfbg fcdba bdg bgacde cdbag ebcg bg | gecb bg bdagc cedag
gb bgf bagfce geacf acdgef fbecgd bacg geabf cdfageb defba | cbgfae egbfdc afebgc fbdea
da fda bedcaf cbfde egbaf ecgdbf aefdb adbc afgdce dcegafb | gafeb gadbfce ad ad
cedf bdagf dc ebafc dca dfbac cadfgbe bcfaeg acdgeb bcfead | ebdgfca bcdfa adc dabgf
dcbga daf efga dgcbfe fgdca af efbdac fgbdeac efdcg dgfeac | fbdegac fdagc bdeagcf agfe
abg gb ecfdab afegcb afbed dgfab cadfg gdeb dagfeb degafcb | bg egbd bdafg deafb
bge ecadg fgecdab dgfecb egdcba egdba dcfaeg gbdaf bcae eb | be aecb be ecdag
cgade dbgec debgcf cdbeag ea gafcd fgabedc edba cae befgac | eac ea daeb eac
bca cb acfgde gbdea adcbe acgbef cebgafd fdcb fcade ebcadf | acgfde dabfceg cb bagedcf
cdfgeab acd cbdfg dfcga gedbac fdcgba gdefa gfdecb ac fbca | ac afgecdb gcefbd ac
aec afcdegb gceafd fbdagc ea aebd agebc adcbg gcebad gbcef | cdgbfa aec fecgb adfbgc
efcda cgefa cag dceafb acdfgb gc agefcdb adcfge agbef edcg | cg cdaef agc gca
dgbeaf bdg eabd cadegf edbfg cgfbe cabfdg dgefbca daefg db | gbd edbgcaf eagfdcb edab
gdbca fdgecb fabceg geadbfc cdafbg ga gca cdfbg bdace fgad | fdcbg ag fdga agc
cdfgaeb ca dcbfe cabg eca befga cdfega bcgaef ebacf gbefad | aec bcga cea eabfgd
efgd deagc aegcdf dgcaf gdacefb ecafbg gfa ebacgd cfabd gf | dagbfec bcfad fag fag
dcbage beagcf dgfaeb eac bcaf ebfcadg cfegd ac beafg gfcae | eca cae acegf abfc
cfeag agbdfc bc fdagb cbdf cab faebgdc fdbega bcdage abfcg | cbfgda cgbaed gbafd fcbag
facgbd begda bf fecb adefc bfdecag fbeadc fab afdbe egafdc | dfcgea edbacfg fb bcef
dbf dfac cebfa egdcb bfecd df afbdeg fedcba fcebga aecdgfb | ceafbd cbegd gacdebf bdf
dgfae agfbedc badgfc cagb ca fdgca dca dbcgfe cfgbd dacebf | cad ac gdbaecf gacb
fcgad dbgac ebcfdag fcagde feda gcaef fdc cdfgbe df gbacef | cgaebf cfd fgecda gdecfa
cadfbe gbcaed geac ecb begcd dbega gdbcf ec acbdegf agbfde | cdbgea ecb fdabegc efdgba
adgfb afdegb dbgcfae cgdbef aedg bedafc dab ad gbfac ebfgd | bad da adb geda
dcgbe aecd gcfadb ad dag gabed efcbagd bafeg ecbdag bcdgef | bdecg abgde gbfceda ad
dcfgbe fedcgab cba gcfaeb abcgd dgfac ab dgecab aebd ebcdg | fcgeab deab baed cab
acbef gbfca fg bfge ecdagfb fadbec fgc ecabfg fgacde dacgb | fegb cgf adfgceb efbg
eaf fbde dcabgfe cdgea ef gfbdac bcafd facbde ecadf cbaegf | fbde fe fbed fegabc
cdbfe gfcd edbfga cgdfaeb dfb gbfcae df ebdac gfbedc cbgef | bcedf dbf eagbfc cgdf
edf bade ecgfdb gacfeb fegbda eagfb daegf gacdf egdbafc ed | fedgba deab dbaegf agfdc
bdef aed dgeab de ecabgdf agefb afebdg fecdga gfecba bagcd | aed fbed bedf dfgaeb
dgcefa gfdecb ae aec dagcfeb ecabf cebfg cfbeag afdbc beag | gabe ace dgeafc cgbef
dcfeb eadfbg ecad ed afbdec bcfda acfdgb ecfbgad dfe cegbf | adec acedbf ed ed
abd bfcd adceg bgefad dfceba db decba ecabf efagcb cdfbega | egdbfa deagc cdbf befdacg
dbacfeg bcdfa defbg aedfb gaed eba gacefb ae cedbfg dfaebg | abe ae deag cbeafdg
acbfdg abcdef gcdb gedfab dcagf dfg gcbdfea dabfc eacgf dg | dgcb afceg cbgd bdgc
dgbace bgaed dbgacf fb bfea fgdcaeb gbf dcefg dgefb gbeadf | bcgfead dgafbc fabe dbfge
egdf fbagdc gbd dg ecbfga egbfdc gcbef bcged egbdfca adbce | dg gecfdb adcfgeb efdbcag
gcfdb gacbd fc ceabgd gacbef egbdf acbdgf adcgbfe dcaf fbc | dcaf fdcgb bgcafe cdaf
edbacf dgfaec feb feadc bcfed cdefbga bdegc cfba dfaebg bf | ecbfd ecbafgd adgbcfe fcedb
ba cgafeb fdbecg dbag dfcae ebdgf bgedacf abfde fdgeab afb | ebgadf bcdgafe egbdfa abf
defbag dcgf acbdgf edcba gecfba fdgba cg gac adcgb ecadfgb | gbacd cgfd gfcdba gc
fbd dfecag dacgbf dcebgf bd cbfae cdgfa dbag dbcfa cgbfead | dbga bfd gadcef cgfaed
daceg ecfdgb ab gab bcdgf cdafegb cgadbf bdaefg gdbac afbc | ba bfac gbdcf cbdegaf
ef gbfcd caebg cdfe gafdbc bcgfe feg fcdgeab efdgcb dgfbae | gfe begfdc bgefc cefgadb
eafc fecdab fcebd bdegcf dfeacbg cab bdegac ac gadbf bfacd | bcedf bcdega dfbac fbdcgea
bef aecbd edgbcf ebcadf adefg dbecfga bf eadbf cfab gabecd | fb cabf cafb efb
agef fecda cdfbga cgefda efcbad fg gdf cdbge ecgdabf cdgef | cadegfb ecfdg dfgabce gf
ebcf gcbed fcgbdea gef gecdbf baecgd ef fabdg egdbf gdfcea | gfedac cgdfeab fe ef
adebf fecg gedbc fg bgfdeca dfg fagdbc dgefb gbdcea cfdbge | dfg dfg gdf ebgfadc
gacb dfabg abfdc gfa fadgcb ag gdefac ebcdaf bfdcega gfbed | ceabdf fadcb bfdca bfegd
ecbagf bdfae gfbd adebcf gdfae agf ebadfg abdcgef fg egdac | gdbf afg gf gdefa
bfde cbafdge agced dbfacg gcbfd bfdgce ef gdcef feagbc feg | dgcfe gefcdb aecdg cegad
dc cead fdc adgcf dgfae dfgaec cfgab gbdeaf gebdcf edgfabc | bafcg cdf cfdegb dace
bf cbedg fgeca fgeacb bcafdeg fgdaec badfgc gcbef fbc baef | fb fbea bf fbc
db dcfbe cgfeb daefc efdgcb edabfg bed cgbd facbdeg eacfbg | baefcg fabgec gfbced egbafc
baedcf cadeb beg cgeadfb cbdage edfag eabgd bg cabg cebdgf | cgab fcbegd gb bg
egbfd ecdafg gfbdec eda ad fabd acbge becfgad adgbe dabegf | ebfgd da gbdfce cfgeadb
caegf dafbcg cbef dfgae bdegca fc cebagf gfc cbage cafgbed | cf abcedg cf cf
gcaebdf bc adcfe cbaefd bcd gfdace decab cbfa bdcgef adegb | ecgfdba abdge cb bcd
egcbfda gfda efacbg edfgca cdeab cgfdbe dg ceagd cgfae cgd | dg bacde geacd dafg
baf efbagd bgaecdf gadcf ecbf baegcd afdecb cedab fb adbcf | bafecd fbcda fagedb bf
efadbc ebdac dca dbefa bdfgae gfabdc agbfedc cefa cgedb ac | cafe dgebfa cefabdg bdeac
cegabd ebdfc dbcefa cgdfba fc cebda efac fbegd gebcfda dfc | bdfge ebacd becgadf fdc
gcbea bd febgac fdbeca ceadgb fcgde ebd cedgb gdeabfc abdg | gafbedc gcfde bed gdbce
dac cbgd cbdea ceafdg gadbe dcbaeg fbcae cd gadbef gdafceb | adc cbfaged ecgfad dc
dfgeac dfage gdfc gc febgad dbgcae gce fcbae cbdafeg cfaeg | dbfegca gec gce afdceg
fdbaceg cadfge gbedc edfgb agecb dc dcaebg dcg acdb bafcge | gafbdec gbcae fgdebac cegbfad
fcabed gbaefc dgeb dgfaecb bcfde bcfged gefdc adgfc ge gce | gbde ecfdb cge gebd
gdebc cdab decag ecb fgebd ebcfagd cgbade gceadf abcfge bc | ecb dbac cdab ecb
dgcbfa caebd gcbfdae dgbfc efdbgc ecg dfeagc eg bfeg gbced | cgfdeb eg ebgfcd gbdce
ba cdfgae bdcafge dagecb cegfb dcbfae gabd eba gaecb cegad | gadb ab feadcb efadbgc
bed dcbage dgce de bgdfca ebdca bcefa agcdb gacfdeb abdegf | fdcgaeb ceafdgb cdge cagdbef
beacgfd dgebaf agefbc fagbc fgcdeb cabgd cfg fc gebaf eacf | afce fgc fc gebafc
gfcadeb deg fbdec egafb dg fdegb aefcdb edacbg fgdc dcbfge | deg dbeagcf deg afdecbg
egacf aeg fbcag ea gcfdae ecda bagdef baegdfc edfcg fcedbg | dace dgbefc ae ae
degbf fbgdae bdcfae decgbfa fgcd dbc bgdefc bdgce gabce dc | cd dbc dc abgdfe
fda df bafceg cfgeadb abgefd fdgb gdaef eagbf cgaed abedcf | cfadbge bcdefa fad cgead
acd gfebca agdfbc debcf agcef bagedfc fdecga aedg da dcfae | gbcdfa gaed feadc agfce
adbfgce dfgeab cd ecdb dcgbfa fcgae fdegb egcdf fgdcbe gcd | cfdbeg efdabg caefg cafgdbe
fb dbgcf cfbe gdbce eabfcdg gadefb dgcebf bfd gdcfa cedbga | bcef aedgfbc cfbe gfadc
cdgefa gcdaf baedfgc fag dgef edafc adbgc abfegc fg bcdeaf | fdcbae facedb fg bacgd
cebfg fd fdecag gdbcef gbead def fdcb fcaebg fedbg gafecdb | dgbcfe fdbeg gdfcbe cfegba
deabgf cdaf dabfg fgbca deagcb bac fbdcga ac cgefb caefdbg | degbca ac egbadc cdfa
bafcge dfec bacdgf ecbgd bgfedc dcbgeaf eagdb dfgcb ce ecb | bdgacfe aefbcdg bce dfgcb
edcbgf agfedc ebafgd cgfad fdcba cega febcgad gc fcg dgfea | cg gcea fcg bgefda
efgab gced gc gfc abedcf bdcefg febcg bgafcd acbdfge befdc | gcf fgc bedfc becfd
ceagbd adbf gfdec eadcf fabce agdbcef abcfge dac bfcdae da | bfegacd fcead dabegc dac
af fga gcbea dcgeaf cgdef fbgdce fdea febgcda fbgdca faegc | fcebgd fdegcb gfecbad eagfc
fegcad decg daegf badecgf cdbafe fgacd dbegaf dc gcbaf adc | gdaebf dgfea acd dcefag
decfb eg dbfcag egd adbcefg gefbda bdgef cdfeag geab dbgaf | fdgba gdcefab eg fabdg
gfbaced adcbg ecbdga gfceda cbgfe acgbdf ced edab de ebdcg | gaebfcd adcgbf ebacgd ced
ba eacgfd beadcg dfgac dbgcaf fdgeb bcaf bga fecgdba dfbga | dacgf acedgb agbdf dbfegca
cdbgfa dfgbe gdabf bdfage ef gedbc fde bcafed bcagdef gafe | efd edf fe ebcdg
befdg bagced dcb fgaecb fbced cfda dcbfea cfagebd cfeba cd | cd cd fdgbe fdac
cfead cafb gcdaeb ca adfecbg degaf fecbd ecabdf aec bgdefc | cbedfa bacfedg eac ceafd
dgacbe dfcbgea decbgf cagdf daceg cde adeb de fcbgae cbgea | bdae ed gbdcae cfbaeg
fbedc fgceabd eg cge adcbg baeg dgecb fdcgea badceg gbcadf | dabgec gbae cbdfga aebg
fbgd aedcfgb baf abgecf fcabd fb dcbae afgcde dcfag adgcbf | gafbdc fab fb dgbf
dgbacfe bafgd ceda gfabec cd fdc cfebgd acfbe fcbad cedbfa | cdaefb fbecga fceab cd
dfceba aefg ebagc cfbgad afdgcbe bafegc gcedb ea aeb acfbg | gfaecb eabfdc bagcf geaf
dfcb bfedg agedf defgcb db fgcbdea dbe fgceb gefacb dgcabe | fcdb db db acdbge
fgbcad dabce fdbe dfaec dcb cafdbe db gecba efcdag bgeafcd | cegdfba faebcd cgeadf fabgdc
baegcf dgcb ecgafdb bgdeca cd fegad caegb cedafb cde egdca | efgad cbegfa egbdfca aefbdc
bcdaf cgb bafge afcbeg fgcebda bfacg dgfebc cg eabgfd geac | cbg bcfdgea gbc fcdab
fg dbegcaf decabf gfd gdaebf decga cfbg cdgef ebgfcd fbdce | gfd fegbda cdage dfceg
fdeca cb gbedac gecfab bac edacfbg agfbe efadgb cbgf fcabe | acb dfgaeb cfdae bfgeac
befcag cdgfa acg befcda fbagd cg cfead bgdafec edcg dgfaec | cag dcefabg gca gecd
decab cga acgefb agbef adcgefb cg faegbd cfbg acbge fdcega | abgdcef cg cg cbdae
abc fcbge cdafbe aegc gbefcd ca gbface bdfga cbafg acdebfg | gcbfde cgafbed geca fgbeac
afbegdc edacb dabeg edcfa cebfad bc ceb dbfc cegafd gbeafc | efbgca becda bce adbeg
aedcgbf ab ecadg gdeba fgab dba abgdfe fcebda fecbgd dgefb | decga bda cgafedb afcdbe
dg degbc aebcdf cegfb gbcafd daceb dcg cegdab adeg egcabdf | gd bcged gcbdafe cegbd
cbgd acbefg db dbaecg dab afecdbg dfage cfebda aegbd agceb | dcbg ecdgba db dbcg
fedgc agf ga fcagbe dfgcae dega fdbceg fbcda dacfg efcdbga | ag fga ga ga
ca fdgbac bcedf afgdeb adfgecb cdbfa cefagd afc bgfad cgba | gacb ac bcgdfae cdagef
acdbfe abdge gca gc gcfdba aebcg faebc bgefdca cefg cfbeag | bcdagf dgeba gfec ecfg
fdbga afd dbgaefc abfc af fgbcda ebcadg dcgaef fgdbe cdgab | fdgab gefbacd af ebdgf
gcafdbe efd ef cdeba abdfcg gdfaeb egcf bedcfg gfbcd cfdbe | dceab gdcabf ef fe
bdca cdfgeb begfa bdf bafgdc db fcgade ecadfgb dbfag gcfad | fbd fbd bgdaf gafdc
efbacg bgedc abcdge bgd dcba db agbce gcfde cgbdeaf eafdbg | bfcega cgdbe bd bdg
fbe fcgbeda fb cebfa ebgac egacdb gebadf ebgcaf adcfe cbfg | bdcafge afgbdec bdegcaf fb
bdcgae fcdbag cdeafg defa fcgea af gbfce cegad bcefdga acf | acf cgdeab dafbcg aefd
ebcag eda ecfabd fedcgb cgfed eagcdf da acedg gcdfbea agfd | dgace agfecd egcbfda cbdeaf
dbgec gebfdc cba bdgfaec acebd ac acdbeg agcd afbed bgaefc | cegdab eadcgb feabd acb
afcg cgadb bedcga efgdb fab gafcbed dabgf afcdbg dcfbae af | baf fbegacd abf fab
ae ebagcd gacdebf gdfeac afdge febgd faec ade dgacfb cadgf | gfdea ae eda bgfde
gbdacfe faedbg be cbgfa ebac cbafgd efb gbfce fecbag dgcfe | fcegb acgfb bafedg ecab
dbefac fdegbca ec ced cdafb cdeab dgface abegd fcadbg cefb | baegd dce edc gdbea
gadecb dcge fcabd ec afegbc gdbae bcdfeag aecbd fbgdea ace | eac geabdf cae bacedgf
efabd ga ebadcf cgdfeba adfceg fadbge degbc baged bfag eag | abgf bfga aeg aeg
eafdbc acgdbe dfae deb gfcbe fdbcag ed fbced cfabd fadegcb | cfgeb ed ebd eadf
egd caedgf de gbcea edafbg eadgb fadgbec bedf bgfda dgabcf | de efdb aecgb gadbf
ceg cade dgfbac cfdeg agcfd fgdeb ce gaecfb fcdageb defacg | fdabgc aced edfgb cdae
fadebc fa fade dbgfca ecfbg ebdcfga dbeac egcadb afb ceafb | fa daef cegfadb dfae
eadgbc fbaecg gdb db fcebdg cdgfa gbcad egbca abde gefcabd | gbd db cbdga acgbe
cab dabgef defbcga ca aecgfb aecf dcegab bfaeg bacfg gfbcd | gfdcb bfaedg ca egabf
bace egbafd fceabg beafg afc ca bafecgd afegcd dgcbf bgfac | bdcefag dfbcg cebgfad fagdebc
afc gcadf dacbfe ca gdbacf befcagd gcba cbdgfe cfbdg fdage | agfcedb febgcd ca bcga
cdgbfa gbdac bda gdebc fcebagd cbeagf gfda efbdca ad fagbc | cbfga gafd fgad gadbc
eb edfagb agecdfb agdfbc gebc cagbfe edfac efb acfbg acfbe | febdag egcb bdgfca eb
cadgebf dgbf bf baf gdbea gbdace gdbefa bagfec ebdfa cafde | ebdfa ebdafgc gfdbea gfdbae
fca cafdg bdacfe bgfa cbdga dcbage af gfcde bfgcda bcegfad | caf caf bafg afc
ecdfa fegda cedgfb ge gcdfae deg efcdab fgdab cgafdeb cgae | fdgeacb gabfcde afgdecb dfgceb
fcegda bcegda fgabcd geba bcafedg cdbge gceda bgd cebfd gb | bfgcda egba geba dgbceaf
fagdeb bgf cdefgba dbgea agfe fg dgaecb bfcda abgfd cdebgf | agfbd agcebdf fbg fbg
afdeb fb ebafgd cafbge gdfb egdba fab dbcaeg bcdegaf eacfd | bf dbacfeg fb baf
fcdae gacbed bcfgead agbd cfegbd acb edbgc cabgef ab edcba | adbg acb abfgec ba
ceda ecfab beadcf eba bcedf ae adcbegf bgafde cbfag gedcbf | gacdebf ceda cefba abe
edgc dfcbge dfbga bcgeaf bgfcd bfcde aedfcgb dcbafe gfc cg | dgec gfc cegd gacefb
dfe dcbe efdbag fcdgb fgcadb ecdgf efcga cgfbde gfdecab ed | eacfgbd dceb cfgdbe de
ecbdg daefb edgbac ac fcdeag cae cgba dfebgc dgefcab dbaec | cgadfe abdef ac aec
cdgbef debacf gafcb ea abde bedcf efacb efadgcb fea eacfdg | ecbfa ea ea fcbde
bacd eadcg gecfa edbgca edc gfcbde dgeba gfbcade cd bdafeg | aedbfg dc fabcegd cde
edfgb agebdf gc gbdafc adfebcg dfceg bcge cgebdf eafdc dgc | cdg dbefcga cg egcfd
befd gaedbf decgab badgf gcebfa bgeda dcgaebf bf bfg afgcd | bf fb edbacg febd
abcdgef egba bfcdga ecabfg dcfae ab fab fecba cfebg dfebgc | afb egab abge aebg
debgac ad afcbeg adg cfgad gdecfa daef gcaef cgbdf cebgdfa | cgfda cgfad ad da
gcbade dagcb gefba agebc bdec fagedc ce daebfgc fabgdc cea | decb bfgea ce ecbd
ecdfbga dcabfe efdagc bdcae ga dfbcg gcbda agd egab eagbcd | cdebfag ecabd fbgaecd ag
acgfebd dacb acfeg bagec ebfagd cdegb bga cgadeb bfcedg ab | cgbade agecb adcb ab
edfac fdcb fdeabc fbe cfabe acdegf gbcae efgbda bf cgdefab | bgefad dcfeabg bef bfcd
gdaceb fac gafcdb cefdab fc cgabd gfbc fcdag agecfbd daegf | dagbfec dfcagb dgacf gdabcef
bgfdac db cbd fbdg cgfaedb cgafd acbef dcbfa gfceda bgaedc | agcefd cbd cdb cdabf
adcgfeb cdafbe agecbf egdcb gbaf dgfcea egf baecf fg fbgec | fcdabeg agbf bgdec efcdab
adbfec abfge cgdfea fgbaced gecfba edgbf fceba age cagb ag | age ga age gbecaf
ebfdg ecfdb fcd bfaegcd cdbg fdcgeb afbged ceabf edafgc dc | egbadf cbfgde febac gbedaf
gfacd cefagb bcegdfa ad adbf gfdcba adc bgecad gcdfe bgfac | da gebcfa fdgac gcfab
feacbg ca bgcfd gdacb dabge bgeadcf gca fdca fcdbag cdfgbe | badfgec adcf dbfegc cfdabge
acb gcbf acbefg fcade bfaeg gebadfc bc aecbdg bgeadf bface | gfcb gbcf dfegabc cab
gcfadb cdf eabdfgc df cdefga gbacd gdbfc eagcdb dabf fgecb | bcgfe cgfbe bdagce gbdfcae
gecab bgecfad defcg dae caegd ebfagc da cbaged fdecba bdag | gadb agbd eadgc ad
fabce bafdegc abfegc dabecf dbgeca gbdaf gbe eg bfega cfge | bdafg eafbg febcgad gcedab
eafdgc dabgcfe dafceb efa fbde acebf acbde ef bcgaf cgeadb | aef ef edfb cgbaf
gfdacb gfedac cefbd cgea egfdc gcd bgedfa eagfd cg fcagebd | cg fgdcae cg cg
ebfga dgbcef gbc cdbef agdfbec cfdg cabegd dcfeba gcfeb cg | afbeg cg bgcdfe gefba
cfgdb efbgdca gde gafcbe gacbed acbeg ed dcgeb fagdec adbe | ecgdb deg gbeafc gaecbdf
gdefbc egfb fcbeda gabecd fcdbg bg bgc dfgac dfbec bfgadce | afbdce bdcefa bg acfdbeg
ed gbcedaf gbde defcab efbdga adfbg adegf ecfga dae dcgfba | cafgbd bdeg bcfdga gefbcad
gfeabd ceafd gca agdcfbe badge agbecf dcgeab gcbd gcade gc | daecf dbeag agc gc
cgbaf dbagf da gdcfbae febdg gbfeca bdca cdegaf gfdacb fad | cgafbd bcafg fbdga badc`
