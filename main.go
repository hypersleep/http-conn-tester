package main

import(
	"log"
	"flag"
	"time"
)

func main() {
	dns      := flag.String("d", "8.8.8.8", "DNS server")
	url      := flag.String("u", "https://google.com", "URL")
	interval := flag.Duration("i", time.Duration(time.Second), "Interval between requests")
	timeout  := flag.Duration("t", time.Duration(time.Second), "Request timeout")
	output   := flag.String("o", "output", "Output file")

	flag.Parse()

	test := &Test{
		DNS:      *dns,
		URL:      *url,
		Interval: *interval,
		Timeout:  *timeout,
		Output:   *output,
	}

	err := test.Run()
	if err != nil {
		log.Fatal("Test failed:", err)
	}
}
