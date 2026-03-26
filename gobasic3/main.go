/*
This program reads in words, count how many times the
word shows up and stores them using key value pair. 
It also sorts them in descending order using the number of
times the word showed up
*/

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
		//increment a word when read
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