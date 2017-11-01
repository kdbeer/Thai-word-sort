package main

import "fmt"

func main() {
	fmt.Println('ข')

	step := 30

	val := 1000
	text := 3585
	for i := 0; i < 46; i++ {
		fmt.Println("Alphabet[\"", string(text+i), "\"]=", val)
		val += step
	}

	fmt.Println()

	lowerCase := 97
	upperCase := 65
	for i := 0; i < 26; i++ {
		fmt.Println("Alphabet[\"", string(lowerCase+i), "\"]=", val)
		val += step
		fmt.Println("Alphabet[\"", string(upperCase+i), "\"]=", val)
		val += step
	}

	fmt.Println()

	fmt.Println("Alphabet[\"๐\"]=", val)
	val += step
	fmt.Println("Alphabet[\"๑\"]=", val)
	val += step
	fmt.Println("Alphabet[\"๒\"]=", val)
	val += step
	fmt.Println("Alphabet[\"๓\"]=", val)
	val += step
	fmt.Println("Alphabet[\"๔\"]=", val)
	val += step
	fmt.Println("Alphabet[\"๕\"]=", val)
	val += step
	fmt.Println("Alphabet[\"๖\"]=", val)
	val += step
	fmt.Println("Alphabet[\"๗\"]=", val)
	val += step
	fmt.Println("Alphabet[\"๘\"]=", val)
	val += step
	fmt.Println("Alphabet[\"๙\"]=", val)
	val += step

	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Println("Alphabet[\"", i, "\"]=", val)
		val += step
	}

	fmt.Println()
	characters := 91
	for i := 0; i < 15; i++ {
		fmt.Println("Alphabet[\"", string(characters+i), "\"]=", val)
		val += step
	}

	fmt.Println()
	characters = 58
	for i := 0; i < 7; i++ {
		fmt.Println("Alphabet[\"", string(characters+i), "\"]=", val)
		val += step
	}

	fmt.Println()
	characters = 91
	for i := 0; i < 7; i++ {
		fmt.Println("Alphabet[\"", string(characters+i), "\"]=", val)
		val += step
	}

}
