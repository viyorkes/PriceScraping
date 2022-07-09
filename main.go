package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	//book link for this course
	resp, err := http.Get("https://www.packtpub.com/application-development/hands-go-programming")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	stringBody := string(data)

	re := regexp.MustCompile(`.*price*`)

	priceMatches := re.FindStringSubmatch(stringBody)

	fmt.Printf(
		"Book Price: %s\n",
		priceMatches,
	)

}
