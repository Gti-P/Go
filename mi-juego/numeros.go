package main

import (
	"fmt"        // Para imprimir y leer del usuario
	"math/rand"  // Para generar numeros aleatorios
	"time"       // Para usar el tiempo como semilla aleatoria
)

func main() {
	// Semilla aleatoria: sin esto, siempre sale el mismo número.
	rand.Seed(time.Now().UnixNano())

	// Elegir modo de juego
	fmt.Println("=== JUEGO: ADIVINA EL NÚMERO ===")
	fmt.Println("Selecciona un modo:")
	fmt.Println("  1. Normal  → rango 1-100,  7 intentos")
	fmt.Println("  2. Difícil → rango 1-1000, 10 intentos")
	fmt.Print("Tu elección (1 o 2): ")

	var modo int
	fmt.Scan(&modo)

	// Configurar parámetros según el modo elegido
	var (
		rangoMax    int
		maxIntentos int
	)

	switch modo {
	case 2:
		rangoMax = 1000
		maxIntentos = 10
		fmt.Println("\n Modo DIFÍCIL: adivina entre 1 y 1000 en 10 intentos.")
	default:
		rangoMax = 100
		maxIntentos = 7
		fmt.Println("\n Modo NORMAL: adivina entre 1 y 100 en 7 intentos.")
	}

	// Generar número secreto dentro del rango configurado
	secreto := rand.Intn(rangoMax) + 1
	intentos := 0

	fmt.Printf("¡Comencemos! Tienes %d intentos.\n\n", maxIntentos)

	// Bucle del juego
	for intentos < maxIntentos {
		restantes := maxIntentos - intentos
		fmt.Printf("Intento %d/%d — Tu número (1-%d): ", intentos+1, maxIntentos, rangoMax)

		var intento int
		fmt.Scan(&intento)
		intentos++

		if intento < secreto {
			fmt.Printf("  ↑ Más alto. Te quedan %d intento(s).\n\n", restantes-1)
		} else if intento > secreto {
			fmt.Printf("  ↓ Más bajo. Te quedan %d intento(s).\n\n", restantes-1)
		} else {
			// ¡Acertó!
			fmt.Printf("\n ¡Correcto! Lo adivinaste en %d intento(s).\n", intentos)
			return // Salimos del programa
		}
	}

	// Se acabaron los intentos sin acertar
	fmt.Printf("\n ¡Sin intentos! El número secreto era: %d\n", secreto)
}