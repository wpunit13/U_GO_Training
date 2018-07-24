package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}
	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	buckets := make([]int,200)

	for scanner.Scan{
		n != HashBucket
	}

}
