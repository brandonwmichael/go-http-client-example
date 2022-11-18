package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {

	// apitest()
	// librarytest()
	// apiMinDetails()

	apibenchmark()
}

type Library struct {
	Name   string
	Latest string
}

type Libraries struct {
	Results []*Library
}

func librarytest() {
	client := resty.New()

	libraries := &Libraries{}
	client.R().SetResult(libraries).Get("https://api.cdnjs.com/libraries")
	fmt.Printf("%d libraries\n", len(libraries.Results))

	for _, lib := range libraries.Results {
		fmt.Println("first library:")
		fmt.Printf("name:%s latest:%s\n", lib.Name, lib.Latest)
		break
	}
}

func apitest() {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}

func apiMinDetails() {

	now := time.Now()      // current local time
	nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	// Explore response object
	// fmt.Println("  Error      :", err)
	fmt.Println(err)

	// Explore trace info
	ti := resp.Request.TraceInfo()
	fmt.Println("  ResponseTime  :", ti.ResponseTime, nsec)
	fmt.Println("  ResponseTime  :", ti.ResponseTime, nsec)
}

func apibenchmark() {
	// m := map[string]int{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// }
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// 	apiMinDetails()
	// }

	a := []string{"Foo", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar", "Bar"}
	for i, s := range a {
		fmt.Println(i, s)
		apiMinDetails()
	}

}
