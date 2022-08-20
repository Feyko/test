package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	b, err := os.ReadFile("crashes.json")
	if err != nil {
		log.Fatal(err)
	}

	var data interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err)
	}
	m := data.([]interface{})
	b, err = os.ReadFile("FG.log")

	regexes := CrashesToRegexes(m)

	for _, reg := range regexes {
		reg.Match(b)
	}

	//regen := gen(regexes...)
	//
	//outs := make([]<-chan bool, 0, runtime.NumCPU())
	//
	//for i := 0; i < runtime.NumCPU(); i++ {
	//	outs = append(outs, regexer(regen, b))
	//}
	//
	//merged := merge(outs...)
	//
	//for range merged {
	//
	//}

	fmt.Printf("Done in %v", time.Since(start))
}

func merge(cs ...<-chan bool) <-chan bool {
	var wg sync.WaitGroup
	out := make(chan bool)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan bool) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(strings ...*regexp.Regexp) <-chan *regexp.Regexp {
	out := make(chan *regexp.Regexp)
	go func() {
		for _, s := range strings {
			out <- s
		}
		close(out)
	}()
	return out
}

func regexer(in <-chan *regexp.Regexp, data []byte) <-chan bool {
	out := make(chan bool)
	go func() {
		for regex := range in {
			b := regex.Match(data)
			out <- b
		}
		close(out)
	}()
	return out
}

func CrashesToRegexes(crashes []interface{}) []*regexp.Regexp {
	r := make([]*regexp.Regexp, len(crashes))
	for i, c := range crashes {
		r[i] = CrashToRegex(c)
	}
	return r
}

func CrashToRegex(crash interface{}) *regexp.Regexp {

	reg := crash.(map[string]interface{})["crash"].(string)
	regex, err := regexp.Compile(reg)

	if err != nil {
		log.Fatal(err)
	}

	return regex
}
