package main

import (
	"fmt"
	"math"
	"sort"
)

var alphabet map[string]float64

type Word struct {
	Text   string
	Weight float64
}

type ByWeight []Word

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByWeight) Less(i, j int) bool { return a[i].Weight < a[j].Weight }

func main() {
	alphabet = InitData()
	testByWords()
	//testCompare()
	//calculateWeight("คุณเอ๋")

}

func testCompare() {
	testData := []string{
		"บ บ บ  raw develop",
		"คุณแอน",
	}
	words := []Word{}
	for _, k := range testData {
		cmpValue := calculateWeight(k)
		words = append(words, Word{k, cmpValue})
	}
	sort.Sort(ByWeight(words))
	for _, k := range words {
		fmt.Println(k.Text, "\t", k.Weight)
	}
}

func testByWords() {
	testData := []string{
		"เอ๋",
		"สุธี",
		"กบกา",
		"กบshit",
		"กสิกร",
		"คุณนก หอพัก",
		"คุณ สุธีคม",
		"คุณเอ๋",
		"คุณเอ็ม หอพัก",
		"คุณแอน ปุญณิศา",
		"เช็กแคท",
		"น้องกิ่ง",
		"นิติ",
		"บัญชีเจ้าของหอ",
		"โบท",
		"บ Raw Dev",
		"ป้ารัต เจ้าของหอ",
		"พงษ์",
		"ปีปีปีปีปีปีปีปีปีปี",
		"ปี",
		"พี่สุวิทย์",
		"เพชร โยธา",
		"แม่น้องกิ่ง",
		"แม่บ้าน",
		"อู๋เบอร์",
		"Cat 1mb",
		"Central Group",
		"Dad",
		"IT Clinic",
		"KDBEER",
		"Kimm",
		"K Puttiwat",
		"Mom",
		"P Ball Living Mobile",
		"P' Noom",
		"Rachata",
		"Raw Develop",
		"Ray",
		"Scg",
	}
	words := []Word{}
	for _, k := range testData {
		cmpValue := calculateWeight(k)
		words = append(words, Word{k, cmpValue})
	}
	sort.Sort(ByWeight(words))
	for _, k := range words {
		fmt.Println(k.Text, "\t", k.Weight)
	}

}

func calculateWeight(text string) float64 {
	//w := strings.Fields(text)
	//text = ""
	//for _, j := range w {
	//	if utf8.RuneCountInString(j) != 1 {
	//		text += j
	//		continue
	//	}
	//	if alphabet[string(j)] >= 100 && alphabet[string(j)] <= 4800 {
	//		text += j + "อ" + " "
	//	}
	//
	//}

	fmt.Println("Process for : ", text)
	characterCounter := 0.0
	vowelCounter := 0.0
	value := 0.0
	cmpValue := 0.0
	weight := 1.0
	for _, v := range text {
		if alphabet[string(v)] == 0 {
			continue
		}

		value = 0
		if alphabet[string(v)] >= 100 {
			alphabetValue := alphabet[string(v)]
			vowelCounter = 1 //If found character reset vowelCounter
			if characterCounter == 0.0 {
				cmpValue += alphabet[string(v)]
			} else {
				//weight = math.Pow(10, -(2*characterCounter + 1))
				weight = math.Pow(10, -(2*characterCounter + 1))
				value = weight * alphabetValue
			}
			cmpValue += value
			fmt.Println("C ", characterCounter, " >> ", cmpValue, weight, value)
			characterCounter++
		} else {
			vowelWeight := math.Pow(10, -(vowelCounter))
			value = weight * alphabet[string(v)] * vowelWeight
			cmpValue += value
			fmt.Println("V ", vowelCounter, "C ", characterCounter, " >> ", cmpValue, vowelWeight)
			vowelCounter++
		}
	}

	fmt.Println("---------------------------------------------------------")
	return cmpValue
}

func InitData() map[string]float64 {
	alphabet := make(map[string]float64)
	//alphabet["เ"] = -10
	//alphabet["แ"] = -8
	//alphabet["โ"] = -6
	//alphabet["ใ"] = -4
	//alphabet["ไ"] = -2
	//alphabet["ะ"] = 2
	//alphabet["ั"] = 4
	//alphabet["า"] = 6
	//alphabet["ำ"] = 8
	//alphabet["ิ"] = 10
	//alphabet["ี"] = 12
	//alphabet["ึ"] = 14
	//alphabet["ื"] = 16
	//alphabet["ุ"] = 18
	//alphabet["ู"] = 20
	//
	//alphabet["็"] = 22
	//alphabet["่"] = 24
	//alphabet["้"] = 26
	//alphabet["๊"] = 28
	//alphabet["๋"] = 30
	//alphabet["์"] = 32

	alphabet["เ"] = 20
	alphabet["แ"] = 21
	alphabet["โ"] = 22
	alphabet["ใ"] = 23
	alphabet["ไ"] = 24
	alphabet["ะ"] = 25
	alphabet["ั"] = 26
	alphabet["า"] = 27
	alphabet["ำ"] = 28
	alphabet["ิ"] = 29
	alphabet["ี"] = 30
	alphabet["ึ"] = 31
	alphabet["ื"] = 32
	alphabet["ุ"] = 33
	alphabet["ู"] = 34

	alphabet["็"] = 0
	alphabet["์"] = 0
	alphabet["่"] = 0
	alphabet["้"] = 0
	alphabet["๊"] = 0
	alphabet["๋"] = 0

	alphabet["ก"] = 100
	alphabet["ข"] = 200
	alphabet["ฃ"] = 300
	alphabet["ค"] = 400
	alphabet["ฅ"] = 500
	alphabet["ฅ"] = 600
	alphabet["ฆ"] = 700
	alphabet["ง"] = 800
	alphabet["จ"] = 900
	alphabet["ฉ"] = 1000
	alphabet["ช"] = 1100
	alphabet["ซ"] = 1200
	alphabet["ฌ"] = 1300
	alphabet["ญ"] = 1400
	alphabet["ฎ"] = 1500
	alphabet["ฎ"] = 1600
	alphabet["ฏ"] = 1700
	alphabet["ฐ"] = 1800
	alphabet["ฑ"] = 1900
	alphabet["ฒ"] = 2000
	alphabet["ณ"] = 2100
	alphabet["ณ"] = 2200
	alphabet["ด"] = 2300
	alphabet["ต"] = 2400
	alphabet["ถ"] = 2500
	alphabet["ท"] = 2600
	alphabet["ธ"] = 2700
	alphabet["น"] = 2800
	alphabet["บ"] = 2900
	alphabet["ป"] = 3000
	alphabet["ผ"] = 3100
	alphabet["ฝ"] = 3200
	alphabet["พ"] = 3300
	alphabet["ฟ"] = 3400
	alphabet["ภ"] = 3500
	alphabet["ม"] = 3600
	alphabet["ย"] = 3700
	alphabet["ร"] = 3800
	alphabet["ล"] = 3850
	alphabet["ฤ"] = 3900
	alphabet["ฦ"] = 4000
	alphabet["ว"] = 4100
	alphabet["ศ"] = 4200
	alphabet["ษ"] = 4200
	alphabet["ส"] = 4400
	alphabet["ห"] = 4500
	alphabet["ฬ"] = 4600
	alphabet["อ"] = 4700
	alphabet["ฮ"] = 4800

	alphabet["a"] = 4900
	alphabet["A"] = 5000
	alphabet["b"] = 5100
	alphabet["B"] = 5200
	alphabet["c"] = 5300
	alphabet["C"] = 5400
	alphabet["d"] = 5500
	alphabet["D"] = 5600
	alphabet["e"] = 5700
	alphabet["E"] = 5800
	alphabet["f"] = 5900
	alphabet["F"] = 6000
	alphabet["g"] = 6100
	alphabet["G"] = 6200
	alphabet["h"] = 6300
	alphabet["H"] = 6400
	alphabet["i"] = 6500
	alphabet["I"] = 6600
	alphabet["j"] = 6700
	alphabet["J"] = 6800
	alphabet["k"] = 6900
	alphabet["K"] = 7000
	alphabet["l"] = 7100
	alphabet["L"] = 7200
	alphabet["m"] = 7300
	alphabet["M"] = 7400
	alphabet["n"] = 7500
	alphabet["N"] = 7600
	alphabet["o"] = 7700
	alphabet["O"] = 7800
	alphabet["p"] = 7900
	alphabet["P"] = 8000
	alphabet["q"] = 8100
	alphabet["Q"] = 8200
	alphabet["r"] = 8300
	alphabet["R"] = 8400
	alphabet["s"] = 8500
	alphabet["S"] = 8600
	alphabet["t"] = 8700
	alphabet["T"] = 8800
	alphabet["u"] = 8900
	alphabet["U"] = 9000
	alphabet["v"] = 9100
	alphabet["V"] = 9200
	alphabet["w"] = 9300
	alphabet["W"] = 9400
	alphabet["x"] = 9500
	alphabet["X"] = 9600
	alphabet["y"] = 9700
	alphabet["Y"] = 9800
	alphabet["z"] = 9900
	alphabet["Z"] = 10000

	alphabet[" "] = 0
	alphabet["!"] = 10200
	alphabet["\""] = 10300
	alphabet["#"] = 10400
	alphabet["$"] = 10500
	alphabet["%"] = 10600
	alphabet["%"] = 10700
	alphabet["'"] = 10800
	alphabet["("] = 10900
	alphabet[")"] = 11000
	alphabet["*"] = 11100
	alphabet["+"] = 11200
	alphabet[","] = 11300
	alphabet["-"] = 11400
	alphabet["."] = 11500
	alphabet["/"] = 11600
	alphabet[":"] = 11700
	alphabet[";"] = 11800
	alphabet["<"] = 11900
	alphabet["="] = 12000
	alphabet[">"] = 12100
	alphabet["?"] = 12200
	alphabet["@"] = 12300
	alphabet["["] = 12400
	alphabet["\\"] = 12500
	alphabet["]"] = 12600
	alphabet["^"] = 12700
	alphabet["_"] = 12800
	alphabet["`"] = 12900
	alphabet["{"] = 13000
	alphabet["|"] = 13100
	alphabet["}"] = 13200
	alphabet["~"] = 13300
	return alphabet
}
