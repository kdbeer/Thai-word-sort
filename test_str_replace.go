package main

import (
	"fmt"
	"strings"
)

func main() {
	k := []string{
		"/images/kdkdkdk",
		"/images/kdlldosoos",
		"/images/kdllos",
		"/images/kdsoos",
		"/images/kdlldos",
		"/images/kdlldosss",
	}

	for _, img := range k {
		img = strings.Replace(img, "images", "merchant", -1)
		fmt.Println(img)
	}

	fmt.Println(k)
}
