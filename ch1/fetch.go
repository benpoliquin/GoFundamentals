package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const urlPrefix = "https://"

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, urlPrefix) {
			url = urlPrefix + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %v\n", url, err)
		}
		fmt.Printf("%s", b)
	}
}
