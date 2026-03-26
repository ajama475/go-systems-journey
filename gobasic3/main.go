package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	//split by word
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		//increment a word when scanned
		words[scan.Text()]++
	}
	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	//use go sort
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss[:3] {
		fmt.Println(s.key, "appears", s.val, "times")
	}

}