package main

import "fmt"
import "strings"

var puzzleInput = []string{"mphcuiszrnjzxwkbgdzqeoyxfa", "mihcuisgrnjzxwkbgdtqeoylia", "mphauisvrnjgxwkbgdtqeiylfa", "mphcuisnrnjzxwkbgdgqeoylua", "mphcuisurnjzxwkbgdtqeoilfi", "mkhcuisvrnjzowkbgdteeoylfa", "mphcoicvrnjzxwksgdtqeoylfa", "mxhcuisvrndzxwkbgdtqeeylfa", "dphcuisijnjzxwkbgdtqeoylfa", "mihvuisvrqjzxwkbgdtqeoylfa", "mphcuisrrnvzxwkbgdtqeodlfa", "mphtuisdrnjzxskbgdtqeoylfa", "mphcutmvsnjzxwkbgdtqeoylfa", "mphcunsvrnjzswkggdtqeoylfa", "mphcuisvrwjzxwkbpdtqeoylfr", "mphcujsdrnjzxwkbgdtqeovlfa", "mpfcuisvrdjzxwkbgdtteoylfa", "mppcuisvrpjzxwkbgdtqeoywfa", "mphcuisvrnjzxwkbfptqroylfa", "mphcuisvrnjzxwkbgstoeoysfa", "mphcufsvrnjzcwkbgdeqeoylfa", "mphcuissrnjzxwkbgdkquoylfa", "sphcuxsvrnjzxwkbgdtqioylfa", "mphcuiivrhjzxwkbgdtqevylfa", "echcuisvrnjzxwkbgltqeoylfa", "mphcuisvrljexwkbvdtqeoylfa", "mpjcuisvrnjzxwkhidtqeoylfa", "mphcuisvrfjzmwkbgdtqeoylfl", "mwhcuisvrnjzxwkbgdtqeoytfm", "mphcuisvrsjzxwkbgdaqeoylfh", "mohcuisvrnjzxwkbgdtqtoymfa", "maycuisvrnjzxwkbgdtqboylfa", "pphcuisvqnjzxwkbgdtqeoylfd", "mprcuisvrnjtxwmbgdtqeoylfa", "mfhcuisgrnjzxckbgdtqeoylfa", "mphiubsvrnjzxwkbgdtqeoyufa", "dphctisvrnjzxwkbgdtqeoylfk", "mphcuisvrnjznwksgdtqeoyzfa", "mpwcuisvrnjziwkbgdtqaoylfa", "mphduzsvrnjznwkbgdtqeoylfa", "mphccisvrnjzxwebgdtqeoylqa", "xphcuisvrnjzxwkfvdtqeoylfa", "mphcupsvrnjzxwkbgdtfeoylpa", "mphcuisvrtjzjwkbgdtqeoylfe", "mpbcuisvrnjzxwkbgdmieoylfa", "mphcuisvrnjzxwkbgjtqetylaa", "mphcuisvrnjzxwpbgdtgdoylfa", "ophcufsvrqjzxwkbgdtqeoylfa", "iphcuhsvrnjzxwkbgetqeoylfa", "mphcuisvunjzxwwbgdtqeoylqa", "mphcpisvrnjzowkbgdtveoylfa", "mphcuisvrnjzxhkbgdtqeotlla", "mphcuisvrnjzxwkbodtgeoylha", "mphcuisvrjjzxwkbwdtqtoylfa", "mphcwisvrnjnxwkbgjtqeoylfa", "mplcuicqrnjzxwkbgdtqeoylfa", "mphcuisvrnjzxydbgdtqeoylfn", "ophckisvrnjzxwkbgdtqeozlfa", "mphcuisvrkjzxwkbgdtteoblfa", "yphcuisvrnjcxwkbggtqeoylfa", "mphcuisvrnazxwfbqdtqeoylfa", "mphcuisvrmjzxwkbgdtlwoylfa", "mphctksvrnjzxwibgdtqeoylfa", "mphcuisprnjzxlebgdtqeoylfa", "mphcuisnrnjzxakbgdtueoylfa", "mphcuiavrnjoxwtbgdtqeoylfa", "nphcuisvrnjzxwkbgdtqzoylfk", "mphcuisrrnjmxwkbgdtqdoylfa", "mphcuisvrujzxwkvgdtqehylfa", "mphcuisvrnfzxwkogdtqebylfa", "mphcuisvrnjwdwkbgdtqeoyxfa", "mphcuisvrntzxwkrgxtqeoylfa", "mpzcuisvrnjzxwebgdtqeoylsa", "aphcuikvrnjzxwwbgdtqeoylfa", "mphcqisvrnjzxwkpgdtqeoelfa", "mphcuusvrnjzxwkbgdtjeodlfa", "mphcuisvrnjzewkbgdtteoylza", "mphcuisvanjzxwkbgdtheoylfc", "mphcjishrnjzxwkbgltqeoylfa", "mpxcuislrnjzxwkbgdtqeoynfa", "mphcuisvrnjjxwkbgdtmeoxlfa", "mphcimsvrnjzxwkbsdtqeoylfa", "mphcxisvcnjzxwjbgdtqeoylfa", "mphcuisbrvjzxwkbgdtqeoymfa", "mplcuisvrnjzxwkbgdtaenylfa", "mphcuihvrnjzxwkygytqeoylfa", "mphcbisvrnjzxhkbgdtqezylfa", "mphcuisarnjzxwkbgatqeoylfv", "mphcumsvrnjzxwkbgdrqebylfa", "mlhcuisvrnwzxwkbgdtqeoylfx", "mpkcuisvrkjzxwkbgdtqeoylfo", "mphcuissrnjzxwkbgdtqmoylfc", "mphcuiwvrnjuxwkfgdtqeoylfa", "mphcuicvlnjzxwkbgdvqeoylfa", "mphcuisvrvvzxwkbfdtqeoylfa", "myhcuisvrnjpxwkbgntqeoylfa", "mpocuisvrnjzxwtbgitqeoylfa", "mphcuisvrnjzxwkbgdtwewyqfa", "mphcuisvtnjzxwwbgdtqeoolfa", "mphcuisvrnjzxgkbgdyqeoyyfa", "mphcuisvrdjzxwkbgpyqeoylfa", "bphcuisvrnjzxwkbgxtqefylfa", "sphcuisvrdjzxwktgdtqeoylfa", "mphcuvsvrnjmxwobgdtqeoylfa", "mphcuisvrnjzxwkbsdtqeuylfb", "mnhcmisvynjzxwkbgdtqeoylfa", "mphckisvrnjzxwkhgdkqeoylfa", "mpacuisvrnjzxwkbgdtqeoolaa", "mpgcuisvrnjzxwkbzdtqeoynfa", "mphcuisvrojzxwkbzdtqeoylga", "mphcuisvknjfxwkbydtqeoylfa", "mphcuistrnjzxwkbgdqqeuylfa", "bpvcuiszrnjzxwkbgdtqeoylfa", "mphcuxsvrnjzswkbgdtqeoelfa", "mphcuisvbnjzxwlbgdtqeoylla", "mphcuisvonczxwkbgktqeoylfa", "mphcuisvrnkzxwvbgdtquoylfa", "mphcuisvrnjzxokfgdtqeoylia", "tphcuisvrnjzxwkbjdwqeoylfa", "mihcuisvrnjzpwibgdtqeoylfa", "mphcuisvrejzxwkbgdtqjuylfa", "mprcuisvrnjixwkxgdtqeoylfa", "mpqcuiszrnjzxwkbgdtqeodlfa", "mphcuasvrnjzzakbgdtqeoylva", "mphcuisvrnjzmwkbtdtqeoycfa", "mphcuisvrnjzxwkbcdtqioylxa", "mphckisvrnjzxwkbcdtqeoylfm", "mphcuisvrnjuxwbogdtqeoylfa", "mphcuisdrnjzxwkbldtqeoylfx", "mphcuisvrnjoxwkbgdtqeyyyfa", "mphcuicvqnjzxwkbgdtqeoylna", "mpmcuisvrnjzxwkbgdtqktylfa", "mphcuisvrnqzxwkggdtqeoykfa", "mphcuisvryjzxwkbydtqejylfa", "mphcugsvrnjzxwkbghtqeeylfa", "rphcuusvrnjzxwkwgdtqeoylfa", "zphwuiyvrnjzxwkbgdtqeoylfa", "cphcuivvrnjzxwkbgdtqenylfa", "mphcuisvrnjzxwkagotqevylfa", "mprcuisvrcjzxwkbgdtqeoytfa", "mphjugsvrnezxwkbgdtqeoylfa", "mphcuisvryjzxwkbgltqeoylaa", "mphcursvrnjzxfkbgdtqeoydfa", "mphcuisvrcuzxwkbgdtqeoylfw", "mphcuisvrijzxwkbgdtqeoelfh", "xphcuisvenjzxjkbgdtqeoylfa", "mphcuisvrnazxwkbgdeqeoylaa", "mphcuisbrsjzxwkbgdtqeoygfa", "mlhvuisvrnjzxwkbgdtqeoylfh", "mphcuisvrnjzxukbgdtqeoyhfy", "mpzcuilvrnjzawkbgdtqeoylfa", "hphcuisjfnjzxwkbgdtqeoylfa", "mahcuisvrnjzxwkegdtqeoylfi", "mphcuixvrnjzcwkbgdtqetylfa", "mphcuisvrnjzxwkdgdtqeoklfj", "mlhcuisvrnjzxwkbgdteeoylka", "mphcuifvrnjbxwkrgdtqeoylfa", "mphcuasvrnjzzwkbgdtqeoylva", "mphcuisvrnjzxwkboutqeoylba", "mbhcuisvcnjzxwklgdtqeoylfa", "mpbcuisvrnjzxgkbgdtqesylfa", "mphcuisvrnjfswkbgdtqeoylfd", "mphcuisvrnjzxwkbgdoweoysfa", "uphcuisvrnjzrwkbgdtqelylfa", "mphcuisvrnjzxwkbgdtqyoylsi", "mpqcuiqvxnjzxwkbgdtqeoylfa", "mphcuisorfjzxwkbgatqeoylfa", "mphcuisvrntfxwkbzdtqeoylfa", "mphcuisvrnrzxwkbgdtueoylfl", "mphcuisvrnjzewkagdtyeoylfa", "mpocuisdrnjzxwkbgdtqeozlfa", "mphcuisvrnjjxwkbgdtoeoylfm", "mphcuisvenjzxwkbgdtqwoylza", "mpmcuisvrnjzxwkbgdtqeoxlfr", "mphcuisvgnjhxwkbgdtqeoplfa", "mphcuisvrnjzowkdgdtqeoyyfa", "mphcuisqynjzxwkbgdtqeoylda", "hphcuisvgnjzxwkbgdtbeoylfa", "iphcuipvrnuzxwkbgdtqeoylfa", "mphcuisvrnjzsikbpdtqeoylfa", "mpwcuhsvrnjzxbkbgdtqeoylfa", "mnhjuisvcnjzxwkbgdtqeoylfa", "mphcudsvrnjzxwkbgdtqloilfa", "mpncuiwvrwjzxwkbgdtqeoylfa", "mphcuisvrnjgawkbgdtqeoylya", "mphcuisvrnjzxwkbggtteoslfa", "mphcuisvrnjzxwkbgdvqeoylpe", "mphcuisvrnczxfkbgktqeoylfa", "mphcuifvrnjzxwkbgdbmeoylfa", "mphcuisvrnjytwkbgdtqeoylla", "mphcuisvrnjzxwkbgdtjeoxlfn", "mphjuisvrnjzxwkbghtqeoyffa", "mphcuisvrnjzxkrbgdtqeoylaa", "mphcbisvrnjzxwkbgttqeoylfs", "mphkuksvbnjzxwkbgdtqeoylfa", "nphcuidvrnjzxwhbgdtqeoylfa", "mphguzsvrnjzxwkbgdaqeoylfa", "mihcuisfrnjzxwkbgdtqhoylfa", "mphcuisvrnrzxwpbgdtqesylfa", "zphcuisvrnjzxwkbddtqeoylaa", "mphcuigvmnjzxwkbgdtqeoylba", "mjhcuisvrnjzxjkbgdtqeoylha", "mphnuisvrnjznwkbgdtqnoylfa", "mkhcuisvrnjcxwkbgdqqeoylfa", "mphcuisvenjzxwbbqdtqeoylfa", "qphcuisnrnjzawkbgdtqeoylfa", "mphcuisvrdjzxwkbgdtqeoywca", "mphcuzsvvnjzxwfbgdtqeoylfa", "pphcuxsvrnjzxwkbgdtmeoylfa", "mphiuvsvrnjzxlkbgdtqeoylfa", "mphlqisvrnjzxkkbgdtqeoylfa", "mmhcuisvrnjzxwkbgatqeoylea", "mphduisrrnjoxwkbgdtqeoylfa", "mphcuisvrnjnxwkvgdyqeoylfa", "mphcuvsvrnjzxgkbgdtqeoylfz", "mphcuisvryjzxwkbggtqkoylfa", "iphcuisvrdjzxwkbgotqeoylfa", "mphcuisvrnjzxwhbgdtqwoyofa", "mphcorbvrnjzxwkbgdtqeoylfa", "mghcuisvrnpzxykbgdtqeoylfa", "mphauisvrnjnxwkbzdtqeoylfa", "mphcgisvrnjzxwkwgdtqeoygfa", "mphcuisvrnjzxwkggotqeoylba", "mphcuesvrnjzxwkbgdwqebylfa", "yphcuisvrnjzxwkbgdxqeoylja", "ephyuisvrnjzywkbgdtqeoylfa", "mfhcuisqrnjzxwkbgdlqeoylfa", "mphkuisvrnjzxwkbertqeoylfa", "mphcuusgrnjzxwkbggtqeoylfa", "mphcuildrnjvxwkbgdtqeoylfa", "mphcuiuvrnjzlwkbgwtqeoylfa", "mppcuisvrljzxwkbgdtqeoylfw", "mphcwiwvrnjzxwsbgdtqeoylfa", "mphcubivrnjzxwkqgdtqeoylfa", "mphcuisvrnjpxwkngdtqeoylpa", "pchcuisvrgjzxwkbgdtqeoylfa", "mphcuisvlnjzxwkbgdtmeoylfw", "mphcuisvrnjzywkbgdvqeoylfj", "mpzcuisvrnezxwktgdtqeoylfa", "mphcuisvrnjbxwkbgzrqeoylfa", "mphcuisvrnjzxwktgdtqeodtfa", "jphcuiavrnjzxwkbgdtqeoylfv", "mphcuisvrnjzxwkbddppeoylfa", "mphcuissrkjzxwkbgxtqeoylfa", "mphcuisvrhjzxwxbgdtqeoylxa", "mphcvisvgnjjxwkbgdtqeoylfa", "mphcuisprnjwxwtbgdtqeoylfa", "mphcuissrnjzxqkbgdtqeoymfa", "mphcuiabrnjzxokbgdtqeoylfa", "mphcuisvrnczxwkbgmtpeoylfa"}

func main() {

	for i, v := range puzzleInput {
		for _, v2 := range puzzleInput[i+1:] {
			if numberOfCharsDiff(v, v2) == 1 {

				fmt.Println(numberOfCharsDiff(v, v2), v, v2)
				fmt.Println(matchingChars(v, v2))
			}
		}
	}
	fmt.Println(numberOfCharsDiff("bbasdasdwdasdwdasdwdasd", "bbasdasdadasdwdasdwdasd"))
}

func matchingChars(input1, input2 string) string {

	var s string

	for i := 0; i < len(input1); i++ {
		if input1[i] == input2[i] {
			s += string(input1[i])
		}
	}
	return s
}

func numberOfCharsDiff(input1, input2 string) int {
	//assuming all characters in low form so can compare runes.
	if len(input1) != len(input2) {
		return -1
		//fmt.Println("len 1", len(input1), "len 2", len(input2))
		//panic("aaaa")
	}

	matchingCharacters := 0

	for i := 0; i < len(input1); i++ {
		if input1[i] == input2[i] {
			matchingCharacters++
		}
	}

	return len(input1) - matchingCharacters
}

func partOne() {
	twoCount := 0
	threeCount := 0
	for _, v := range puzzleInput {
		if checkExactApperances(v, 2) {
			twoCount++
		}
		if checkExactApperances(v, 3) {
			threeCount++
		}
		//fmt.Println(v, checkExactApperances(v, 2), twoCount, checkExactApperances(v, 3), threeCount)
	}
	fmt.Println(twoCount * threeCount)
}

func checkExactApperances(input string, amount int) (exact bool) {
	input = strings.ToLower(input)
	for _, v := range input {
		if strings.Count(input, string(v)) == amount {
			return true
		}
	}
	return false
}
