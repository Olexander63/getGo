package main

import "fmt"

func main() {
	soup := mirepoix(nil)
	fmt.Println(soup)
}

func mirepoix(ingredient []string) []string {
	return append(ingredient, "onion", "carrot", "celery")
}
