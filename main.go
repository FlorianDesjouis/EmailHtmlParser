package main

import "os"
import "fmt"
import "net/http"
import "regexp"
import "io/ioutil"

func main()  {
	args := os.Args[1:]

	index := 0

	for _, arg := range args {
		resp, err := http.Get(arg)

		if err != nil {
			panic(err)
		}

		var bytesArr []byte
		bytesArr, _ = ioutil.ReadAll(resp.Body)

		body := string(bytesArr[:])

		re, _ := regexp.Compile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
		results := re.FindAllString(body, -1)

		for _, res := range results {
			email := string(res[:])
			fmt.Println(email)
		}
		index++
	}
}
