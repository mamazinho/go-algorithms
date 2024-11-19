package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"slices"
	"time"
)

type Input struct {
	Name string
}

func main() {
	input := []*Input{
		{Name: "matheus"},
		{Name: "josue"},
		{Name: "marcel"},
		{Name: "andre"},
		{Name: "allan"},
		{Name: "ronconi"},
		{Name: "jhon"},
		{Name: "silvao"},
	}
	toCompare := []*Input{
		{Name: "chris"},
		{Name: "stain"},
		{Name: "gabriel"},
		{Name: "marcello"},
		{Name: "lucas"},
		{Name: "dario"},
		{Name: "ronconi"},
		{Name: "diogo"},
	}

	working(input, toCompare)
	josueSuggestion(input, toCompare)
	matheusSuggestion(input, toCompare)
}

func working(input []*Input, toCompare []*Input) {
	var categoryInfo []*Input
	start := time.Now()
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&memStatsBefore)

	set := make(map[string]bool, len(input))
	for _, attr := range input {
		set[attr.Name] = true
	}

	for _, meliAttr := range toCompare {
		if !set[meliAttr.Name] {
			categoryInfo = append(categoryInfo, &Input{Name: meliAttr.Name})
		}
	}

	runtime.GC()
	runtime.ReadMemStats(&memStatsAfter)
	fmt.Printf("Working MEM: %.2fKB\n", float64(memStatsAfter.Alloc - memStatsBefore.Alloc) / 1024)
	fmt.Println("Working TIME: ", time.Since(start))
	b, _ := json.Marshal(categoryInfo)
	fmt.Println("Result: ", string(b))
}

func josueSuggestion(input []*Input, toCompare []*Input) {
	var categoryInfo []*Input
	start := time.Now()
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&memStatsBefore)

	set := make(map[string]struct{}, len(input))
	for _, attr := range input {
		set[attr.Name] = struct{}{}
	}

	for _, meliAttr := range toCompare {
		if _, exists := set[meliAttr.Name]; !exists {
			categoryInfo = append(categoryInfo, &Input{Name: meliAttr.Name})
		}
	}

	runtime.GC()
	runtime.ReadMemStats(&memStatsAfter)
	fmt.Printf("Josue MEM: %.2fKB\n", float64(memStatsAfter.Alloc - memStatsBefore.Alloc) / 1024)
	fmt.Println("Josue TIME: ", time.Since(start))
	b, _ := json.Marshal(categoryInfo)
	fmt.Println("Result: ", string(b))
}

func matheusSuggestion(input []*Input, toCompare []*Input) {
	var categoryInfo []*Input
	start := time.Now()
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&memStatsBefore)

	my_array := make([]string, 0, len(input))
	for _, attr := range input {
		my_array = append(my_array, attr.Name)
	}
	set := slices.Compact(my_array)

	for _, meliAttr := range toCompare {
		if !slices.Contains(set, meliAttr.Name) {
			categoryInfo = append(categoryInfo, &Input{Name: meliAttr.Name})
		}
	}

	runtime.GC()
	runtime.ReadMemStats(&memStatsAfter)
	fmt.Printf("Matheus MEM: %.2fKB\n", float64(memStatsAfter.Alloc - memStatsBefore.Alloc) / 1024)
	fmt.Println("Matheus TIME: ", time.Since(start))
	b, _ := json.Marshal(categoryInfo)
	fmt.Println("Result: ", string(b))
}
