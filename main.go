package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	bearMountainURL = "https://www.snowforecast.com/resorts/4336-bear-mountain-resort"
	brianHeadURL    = "https://www.snowforecast.com/resorts/4789-brian-head-resort"
	separator       = "==================================================="
)

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	var wg sync.WaitGroup
	wg.Add(2)

	go getReport(bearMountainURL, "Bear Mountain Snow Report", &wg)
	go getReport(brianHeadURL, "Brianhead Snow Report", &wg)

	wg.Wait()
	fmt.Println("\nDone.")
}

func getReport(url string, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var report []string
	report = append(report, fmt.Sprintf("\n%s\n== %s\n%s", separator, name, separator))

	doc.Find("#tab1 .narrative .leaders li").Each(func(i int, li *goquery.Selection) {
		spans := li.ChildrenFiltered("span")
		report = append(report,
			fmt.Sprintf("\n%s: %s", spans.First().Text(), spans.Last().Text()))
	})
	report = append(report, "\n")

	fmt.Print(strings.Join(report, ""))
}
