package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	response, err := http.Get("https://www.kaisnet.de/ip.php")
	must(err)

	fmt.Println("Status", response.Status)
	fmt.Println("StatusCode", response.StatusCode)
	fmt.Println("ContentLength", response.ContentLength)
	fmt.Println("Header", response.Header)
	fmt.Println("TLS", response.TLS)
	fmt.Println("TransferEncoding", response.TransferEncoding)
	fmt.Println("Proto", response.Proto)
	fmt.Println("ProtoMajor", response.ProtoMajor)
	fmt.Println("ProtoMinor", response.ProtoMinor)
	fmt.Println("Trailer", response.Trailer)
	fmt.Println("Uncompressed", response.Uncompressed)

	data, err := ioutil.ReadAll(response.Body)
	must(err)

	fmt.Println("Body", string(data))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
