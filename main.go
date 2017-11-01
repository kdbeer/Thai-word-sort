package main

import (
	"fmt"
	"math"
	"sort"
)

var Alphabet map[string]float64

type Word struct {
	Text   string
	Weight string
}

type ByWeight []Word

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByWeight) Less(i, j int) bool { return a[i].Weight < a[j].Weight }

func main() {
	Alphabet = InitData()
	testByWords()
	//testCompare()
	//CalculateSortWeight("คุณเอ๋")

}

func testCompare() {
	testData := []string{
		"บ บ บ  raw develop",
		"คุณแอน",
	}
	words := []Word{}
	for _, k := range testData {
		cmpValue := CalculateSortWeight(k)
		words = append(words, Word{k, cmpValue})
	}
	sort.Sort(ByWeight(words))
	for _, k := range words {
		fmt.Println(k.Text, "\t", k.Weight)
	}
}

func testByWords() {
	testData := []string{
		"กกกกกกกกกงงงงง2",
		"กกกกกกกกกงงงงง1",
		"แต้วเอง",
		"ทดสอบร้าน3",
		"ทดสอบร้าน0",
		"ร้าน After brown 3",
		"ร้าน After brown 1",
		"ร้าน After brown 2",
		"ร้านบราวนี่1",
	}
	words := []Word{}
	for _, k := range testData {
		cmpValue := CalculateSortWeight(k)
		words = append(words, Word{k, cmpValue})
	}
	sort.Sort(ByWeight(words))
	for _, k := range words {
		fmt.Println(k.Text, "\t", k.Weight)
	}
}

func CalculateSortWeight(text string) string {

	textTest := ""
	var vowelValue float64
	characterCounter := 0.0
	vowelCounter := 0.0
	for _, v := range text {
		if Alphabet[string(v)] == 0 {
			continue
		}

		if Alphabet[string(v)] >= 100 {
			alphabetValue := Alphabet[string(v)]
			vowelCounter = 0 //If found character reset vowelCounter
			if characterCounter == 0.0 {
				textTest = fmt.Sprint(alphabetValue+vowelValue, ".")
			} else {
				textTest += fmt.Sprint(alphabetValue + vowelValue)
			}
			vowelValue = 0
			characterCounter++
		} else {
			temp := int(vowelValue + (Alphabet[string(v)] / (math.Pow(10, vowelCounter))))
			vowelValue = float64(temp)
			vowelCounter++
		}
	}
	return textTest
}

func InitData() map[string]float64 {
	alphabet := make(map[string]float64)
	alphabet["เ"] = 5
	alphabet["แ"] = 6
	alphabet["โ"] = 7
	alphabet["ใ"] = 8
	alphabet["ไ"] = 9
	alphabet["ะ"] = 10
	alphabet["ั"] = 11
	alphabet["า"] = 12
	alphabet["ำ"] = 13
	alphabet["ิ"] = 14
	alphabet["ี"] = 15
	alphabet["ึ"] = 16
	alphabet["ื"] = 17
	alphabet["ุ"] = 18
	alphabet["ู"] = 19

	alphabet["็"] = 20
	alphabet["์"] = 21
	alphabet["่"] = 22
	alphabet["้"] = 23
	alphabet["๊"] = 24
	alphabet["๋"] = 25

	alphabet["ก"] = 1000
	alphabet["ข"] = 1030
	alphabet["ฃ"] = 1060
	alphabet["ค"] = 1090
	alphabet["ฅ"] = 1120
	alphabet["ฆ"] = 1150
	alphabet["ง"] = 1180
	alphabet["จ"] = 1210
	alphabet["ฉ"] = 1240
	alphabet["ช"] = 1270
	alphabet["ซ"] = 1300
	alphabet["ฌ"] = 1330
	alphabet["ญ"] = 1360
	alphabet["ฎ"] = 1390
	alphabet["ฏ"] = 1420
	alphabet["ฐ"] = 1450
	alphabet["ฑ"] = 1480
	alphabet["ฒ"] = 1510
	alphabet["ณ"] = 1540
	alphabet["ด"] = 1570
	alphabet["ต"] = 1600
	alphabet["ถ"] = 1630
	alphabet["ท"] = 1660
	alphabet["ธ"] = 1690
	alphabet["น"] = 1720
	alphabet["บ"] = 1750
	alphabet["ป"] = 1780
	alphabet["ผ"] = 1810
	alphabet["ฝ"] = 1840
	alphabet["พ"] = 1870
	alphabet["ฟ"] = 1900
	alphabet["ภ"] = 1930
	alphabet["ม"] = 1960
	alphabet["ย"] = 1990
	alphabet["ร"] = 2020
	alphabet["ฤ"] = 2050
	alphabet["ล"] = 2080
	alphabet["ฦ"] = 2110
	alphabet["ว"] = 2140
	alphabet["ศ"] = 2170
	alphabet["ษ"] = 2200
	alphabet["ส"] = 2230
	alphabet["ห"] = 2260
	alphabet["ฬ"] = 2290
	alphabet["อ"] = 2320
	alphabet["ฮ"] = 2350

	alphabet["a"] = 2380
	alphabet["A"] = 2410
	alphabet["b"] = 2440
	alphabet["B"] = 2470
	alphabet["c"] = 2500
	alphabet["C"] = 2530
	alphabet["d"] = 2560
	alphabet["D"] = 2590
	alphabet["e"] = 2620
	alphabet["E"] = 2650
	alphabet["f"] = 2680
	alphabet["F"] = 2710
	alphabet["g"] = 2740
	alphabet["G"] = 2770
	alphabet["h"] = 2800
	alphabet["H"] = 2830
	alphabet["i"] = 2860
	alphabet["I"] = 2890
	alphabet["j"] = 2920
	alphabet["J"] = 2950
	alphabet["k"] = 2980
	alphabet["K"] = 3010
	alphabet["l"] = 3040
	alphabet["L"] = 3070
	alphabet["m"] = 3100
	alphabet["M"] = 3130
	alphabet["n"] = 3160
	alphabet["N"] = 3190
	alphabet["o"] = 3220
	alphabet["O"] = 3250
	alphabet["p"] = 3280
	alphabet["P"] = 3310
	alphabet["q"] = 3340
	alphabet["Q"] = 3370
	alphabet["r"] = 3400
	alphabet["R"] = 3430
	alphabet["s"] = 3460
	alphabet["S"] = 3490
	alphabet["t"] = 3520
	alphabet["T"] = 3550
	alphabet["u"] = 3580
	alphabet["U"] = 3610
	alphabet["v"] = 3640
	alphabet["V"] = 3670
	alphabet["w"] = 3700
	alphabet["W"] = 3730
	alphabet["x"] = 3760
	alphabet["X"] = 3790
	alphabet["y"] = 3820
	alphabet["Y"] = 3850
	alphabet["z"] = 3880
	alphabet["Z"] = 3910

	alphabet["๐"] = 3940
	alphabet["๑"] = 3970
	alphabet["๒"] = 4000
	alphabet["๓"] = 4030
	alphabet["๔"] = 4060
	alphabet["๕"] = 4090
	alphabet["๖"] = 4120
	alphabet["๗"] = 4150
	alphabet["๘"] = 4180
	alphabet["๙"] = 4210

	alphabet["0"] = 4240
	alphabet["1"] = 4270
	alphabet["2"] = 4300
	alphabet["3"] = 4330
	alphabet["4"] = 4360
	alphabet["5"] = 4390
	alphabet["6"] = 4420
	alphabet["7"] = 4450
	alphabet["8"] = 4480
	alphabet["9"] = 4510

	alphabet["["] = 4540
	alphabet["\\"] = 4570
	alphabet["]"] = 4600
	alphabet["^"] = 4630
	alphabet["_"] = 4660
	alphabet["`"] = 4690
	alphabet["a"] = 4720
	alphabet["b"] = 4750
	alphabet["c"] = 4780
	alphabet["d"] = 4810
	alphabet["e"] = 4840
	alphabet["f"] = 4870
	alphabet["g"] = 4900
	alphabet["h"] = 4930
	alphabet["i"] = 4960

	alphabet[":"] = 4990
	alphabet[";"] = 5020
	alphabet["<"] = 5050
	alphabet["="] = 5080
	alphabet[">"] = 5110
	alphabet["?"] = 5140
	alphabet["@"] = 5170

	return alphabet
}
