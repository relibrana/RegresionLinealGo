package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Estructura de datos para gestionar los cálculos parciales
type partialCalc struct {
	sumX  float64
	sumY  float64
	sumXY float64
	sumX2 float64
}

// Función concurrente para calcular sumatorias parciales
func calculatePartialSums(x []float64, y []float64, startIndex int, endIndex int, wg *sync.WaitGroup, results chan partialCalc, Mutex *sync.Mutex) {
	defer wg.Done()

	partial := partialCalc{}
	for i := startIndex; i < endIndex; i++ {
		partial.sumX += x[i]
		partial.sumY += y[i]
		partial.sumXY += x[i] * y[i]
		partial.sumX2 += x[i] * x[i]
	}

	Mutex.Lock()
	results <- partial
	Mutex.Unlock()
}

// Función para calcular el modelo de regresión lineal de forma concurrente
func concurrentLinearRegression(x []float64, y []float64) (float64, float64) {
	numDataPoints := len(x)
	numRoutines := 4 // Número de goroutines a utilizar
	results := make(chan partialCalc, numRoutines)
	var wg sync.WaitGroup
	var Mutex sync.Mutex // Mutex para sincronizar el acceso a 'total'
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		start := i * (numDataPoints / numRoutines)
		end := (i + 1) * (numDataPoints / numRoutines)
		if i == numRoutines-1 {
			end = numDataPoints
		}
		go calculatePartialSums(x, y, start, end, &wg, results, &Mutex)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	total := partialCalc{}
	for partial := range results {
		Mutex.Lock()
		total.sumX += partial.sumX
		total.sumY += partial.sumY
		total.sumXY += partial.sumXY
		total.sumX2 += partial.sumX2
		Mutex.Unlock()
	}

	n := float64(numDataPoints)
	coefB := (n*total.sumXY - total.sumX*total.sumY) / (n*total.sumX2 - total.sumX*total.sumX)
	coefA := (total.sumY / n) - coefB*(total.sumX/n)

	return coefA, coefB
}
func main() {
	file, err := os.Create("elapsed_times.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for i := 0; i < 1000; i++ {
		var x []float64
		var y []float64
		for i := 1; i <= 1000000; i++ {
			x = append(x, float64(i))
			y = append(y, float64(i*2))
		}

		startTime := time.Now()
		concurrentLinearRegression(x, y)

		elapsedTimeMicroseconds := time.Since(startTime).Seconds() * 1000000

		_, err := fmt.Fprintf(file, "%.6f\n", elapsedTimeMicroseconds)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
