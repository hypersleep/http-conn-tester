package main

import(
	"time"
)

type(
	Test struct {
		DNS      string
		URL      string
		Interval time.Duration
		Timeout  time.Duration
		File     string
	}

	TestResult struct {
		DNSResolve time.Duration
		Connection time.Duration
		Download   time.Duration
		Status     string
	}
)

func loadTestConfig() (*Test, error) {
	// defaults
	test := &Test{
		Interval: time.Duration(time.Second),
		Timeout: time.Duration(time.Second),
		File: "result",
	}

	return test, nil
}

func (test *Test) run() error {
	return nil
}
