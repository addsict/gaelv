package main

import "log"

func main() {
	provider := NewProvider("/tmp/gaelog.db")
	for {
		requestLog, err := provider.Next()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("%#v", requestLog)
	}
}
