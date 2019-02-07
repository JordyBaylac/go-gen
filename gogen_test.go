package main

import (
	"encoding/json"
	"fmt"
	"testing"

	bytefmt "code.cloudfoundry.org/bytefmt"
	auto "github.com/JordyBaylac/go-gen/samples"
)

func BenchmarkUnmarshal(b *testing.B) {

	solutionLen := len(getSolution())

	testCases := []struct {
		numberOfSolutions int
		totalBytes        int
	}{
		// {100, solutionLen * 100},
		// {1000, solutionLen * 1000},
		// {10000, solutionLen * 10000},
		// {100000, solutionLen * 100000},
		// {200000, solutionLen * 200000},
		{300000, solutionLen * 300000},
		// {500000, solutionLen * 500000},
	}

	for _, test := range testCases {
		size := bytefmt.ByteSize(uint64(test.totalBytes))
		title := fmt.Sprintf("Solutions %d. Size of %s bytes", test.numberOfSolutions, size)
		b.Run(title, func(b *testing.B) {
			jsonData := getJSON(test.numberOfSolutions)
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				var results auto.ResultList
				b.StartTimer()
				if err := json.Unmarshal(jsonData, &results); err != nil {
					b.Errorf("Unexpected Error: %v", err)
				}
				b.StopTimer()
				for range results {
				}
			}
		})
	}

}
