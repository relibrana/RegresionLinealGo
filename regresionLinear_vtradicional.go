package main

import (
	"fmt"
	"time"
)

// Función para calcular el modelo de regresión lineal
func linearRegression(x []float64, y []float64) (float64, float64) {
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0

	// Calcular sumatorias
	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	// Calcular coeficientes regresión
	n := float64(len(x))
	coefB := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	coefA := (sumY / n) - coefB*(sumX/n)

	return coefA, coefB
}

func main() {
	// Generar un conjunto de datos de ejemplo con 1,000,000 registros
	var x []float64
	var y []float64
	for i := 1; i <= 1000000; i++ {
		x = append(x, float64(i))
		y = append(y, float64(i*2))
	}
	// Medir el tiempo de ejecución
	startTime := time.Now()

	// Calcular regresión lineal
	a, b := linearRegression(x, y)

	// Mostrar coeficientes
	fmt.Printf("Coeficiente a: %.2f\n", a)
	fmt.Printf("Coeficiente b: %.2f\n", b)

	// Predecir un nuevo valor
	newX := 1000001.0
	prediction := a + b*newX
	fmt.Printf("Predicción para x=%.1f: %.2f\n", newX, prediction)

	// Medir el tiempo transcurrido
	elapsedTime := time.Since(startTime)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsedTime)
}
