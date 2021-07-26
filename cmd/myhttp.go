package main

import (
	"fmt"

	"github.com/faroukelkholy/myhttp"
)

func main() {
	limit, urls := myhttp.ParseCLI()

	if err := myhttp.Start(limit,urls); err != nil {
		fmt.Println(err.Error())
	}
}


