package main

import "log"

func main() {
	test, err := loadTestConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	err = test.run()
	if err != nil {
		log.Fatal("Test failed:", err)
	}
}
