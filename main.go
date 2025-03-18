package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	// Pré-alocando o mapa com um tamanho estimado
	data := make(map[string]Measurement, 500)

	// Aumentando o buffer do scanner
	scanner := bufio.NewScanner(measurements)
	const bufferSize = 256 * 1024 // 256KB
	buf := make([]byte, bufferSize)
	scanner.Buffer(buf, bufferSize)

	for scanner.Scan() {
		rawData := scanner.Text()
		parts := strings.Split(rawData, ";")
		location := parts[0]
		temp, _ := strconv.ParseFloat(parts[1], 64)

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

	// Cria uma slice com os nomes das localizações
	locations := make([]string, 0, len(data))
	for location := range data {
		locations = append(locations, location)
	}

	// Ordena a slice alfabeticamente
	sort.Strings(locations)

	// Exibe os dados no formato {Location=min,avg,max, ...}
	fmt.Print("{")
	for i, location := range locations {
		measurement := data[location]
		avg := measurement.Sum / float64(measurement.Count)

		// Adiciona vírgula entre as localizações, exceto para a última
		if i > 0 {
			fmt.Print(", ")
		}

		fmt.Printf("%s=%.1f,%.1f,%.1f",
			location,
			measurement.Min,
			avg,
			measurement.Max,
		)
	}
	fmt.Println("}")
}
