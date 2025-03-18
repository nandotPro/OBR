package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurement, ok := data[location]
		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurement.Min = math.Min(measurement.Min, temp)
			measurement.Max = math.Max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		data[location] = measurement
	}

	for name, measurement := range data {
		avg := measurement.Sum / float64(measurement.Count)
		fmt.Printf("%s: min=%.1f, max=%.1f, avg=%.1f\n", name, measurement.Min, measurement.Max, avg)
	}
}
