package message

import (
	"fmt"
	"net/http"
)

func SendMessage(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("Get request not sent\n")
		fmt.Printf("URL: %s\n", url)
		fmt.Println(err)

		return
	} else {
		fmt.Printf("Get request sent\n")
	}

	defer res.Body.Close()
}
