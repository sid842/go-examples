package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func top3Words() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv // to store reverse mapping

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val // to sort in descending order
	})

	for _, s := range ss[:3] { // to consider only top 3 elements
		fmt.Println(s.key, "appears", s.val, "times")
	}
}
