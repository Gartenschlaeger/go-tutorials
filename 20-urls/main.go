package main

import (
	"fmt"
	"net/url"
)

func main() {
	parseUrl()
	buildUrl()
}

func parseUrl() {
	fmt.Println("parseUrl()")

	url, err := url.Parse("https://www.kaisnet.de:3000/ip.php?format=json&iv=6#tt1")
	must(err)

	fmt.Println("Scheme", url.Scheme)
	fmt.Println("Host", url.Host)
	fmt.Println("Port", url.Port())
	fmt.Println("Path", url.Path)
	fmt.Println("RawPath", url.RawPath)
	fmt.Println("ForceQuery", url.ForceQuery)
	fmt.Println("Query", url.Query())
	fmt.Println("RawQuery", url.RawQuery)
	fmt.Println("Fragment", url.Fragment)
	fmt.Println("Opaque", url.Opaque)
	fmt.Println("User", url.User)
	fmt.Println()
}

func buildUrl() {
	fmt.Println("buildUrl()")

	url := url.URL{
		Scheme: "https",
		Host:   "kaisnet.de",
		Path:   "/ip.php",
	}

	query := url.Query()
	query.Add("redirect", "true")
	query.Add("url", "https://google.de")
	query.Add("q", "golang dictionary")
	url.RawQuery = query.Encode()

	fmt.Println(url.String())
	fmt.Println()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
