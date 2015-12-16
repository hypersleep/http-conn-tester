package main

import(
	"log"
	"time"
	"net/url"
	"net/http"
	"io/ioutil"
	"crypto/tls"

	"github.com/bogdanovich/dns_resolver"
)

const(
	ok        string = "OK"
	timeout   string = "TIMEOUT"
	wrongURL  string = "WRONG_URL"
	dnsError  string = "DNS_ERROR"
	httpError string = "HTTP_ERROR"
)

type(
	Test struct {
		DNS      string
		URL      string
		Interval time.Duration
		Timeout  time.Duration
		Output   string
	}

	TestResult struct {
		DNSResolve time.Duration
		Connection time.Duration
		Download   time.Duration
		Total      time.Duration
		Info       string
	}
)

func (test *Test) connection() {
	var(
		testResult    *TestResult   = &TestResult{}
		endCh         chan struct{} = make(chan struct{})
		nowDNSResolve time.Time
		nowConnection time.Time
		nowDownload   time.Time
		step          int
	)

	go func() {
		// parse url
		u, err := url.Parse(test.URL)
		if err != nil {
			testResult.Info = wrongURL
			endCh <- struct{}{}
			return
		}

		nowDNSResolve = time.Now()
		step++

		// resolve A record
		resolver := dns_resolver.New([]string{test.DNS})
		ip, err := resolver.LookupHost(u.Host)
		if err != nil {
			log.Println("Failed to resolve host:", err)
			testResult.Info = dnsError
			testResult.DNSResolve = time.Since(nowDNSResolve)
			endCh <- struct{}{}
			return
		}

		testResult.DNSResolve = time.Since(nowDNSResolve)

		nowConnection = time.Now()
		step++
		// Dial host
		// get request with Host header
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client := &http.Client{Transport: tr}
		log.Println(u.Scheme  + "://" + ip[0].String() + u.Path)
		req, _ := http.NewRequest("GET", u.Scheme + "://" + ip[0].String() + u.Path, nil)
		req.Header.Add("Host", u.Host)
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Failed to send request:", err)
			testResult.Info = httpError
			testResult.Connection = time.Since(nowConnection)
			endCh <- struct{}{}
			return
		}

		testResult.Connection = time.Since(nowConnection)

		nowDownload = time.Now()
		step++
		// download response body
		qwe, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Println("Failed to read response body:", err)
			testResult.Info = httpError
			testResult.Connection = time.Since(nowDownload)
			endCh <- struct{}{}
			return
		}

		log.Println(string(qwe))

		testResult.Download = time.Since(nowDownload)

		testResult.Info = ok
		endCh <- struct{}{}
		return
	}()

	select {
		case <- endCh:
		case <- time.After(test.Timeout):
			testResult = &TestResult{Info: timeout}
			if step >= 1 {
				testResult.DNSResolve = time.Since(nowDNSResolve)
			}
			if step >= 2 {
				testResult.Connection = time.Since(nowConnection)
			}
			if step >= 3 {
				testResult.Download = time.Since(nowDownload)
			}
	}

	testResult.Total = testResult.DNSResolve + testResult.Connection + testResult.Download

	log.Println(testResult)
}

func (test *Test) Run() error {
	c := time.Tick(test.Interval)
	for range c {
		go test.connection()
	}

	return nil
}
