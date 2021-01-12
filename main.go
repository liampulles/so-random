package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	pwd, err := os.Getwd()
	failIfErr(err)
	items, err := ioutil.ReadDir(pwd)
	failIfErr(err)

	rSource := rand.NewSource(time.Now().UnixNano())
	rGen := rand.New(rSource)
	idx := rGen.Int() % len(items)

	picked := items[idx]
	fmt.Println(picked.Name())
}

func failIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
