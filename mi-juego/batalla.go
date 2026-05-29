package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ─── PARTE 1: Structs ────────────────────────────────────────────────────────

type Personaje struct {
	Nombre  string
	Vida    int
	Ataque  int
	Defensa int
}

type Habilidad struct {
	Nombre string
	Poder  int
}

// ─── PARTE 3: Métodos del Personaje ─────────────────────────────────────────

// EstaVivo comprueba si el personaje sigue en combate.
func (p *Personaje) EstaVivo() bool {
	return p.Vida > 0
}

// Atacar aplica daño al objetivo usando puntero para modificar el original.
func (p *Personaje) Atacar(objetivo *Personaje) int {
	variacion := rand.Intn(11) - 5 // valor entre -5 y +5
	dano := p.Ataque - objetivo.Defensa + variacion
	if dano < 1 {
		dano = 1 // daño mínimo siempre es 1
	}
	objetivo.Vida -= dano
	return dano
}

// ─── PARTE 4: Elegir habilidad ───────────────────────────────────────────────

func elegirHabilidad(habilidades []Habilidad) Habilidad {
	fmt.Println("\n⚔️  Elige tu habilidad:")
	for i, h := range habilidades {
		fmt.Printf("  %d. %s (Poder: +%d)\n", i+1, h.Nombre, h.Poder)
	}
	fmt.Print("Tu elección: ")

	var opcion int
	fmt.Scan(&opcion)

	// Validación: evitar pánico por índice fuera de rango
	if opcion < 1 || opcion > len(habilidades) {
		fmt.Println("  ⚠️  Opción inválida. Se usará la primera habilidad.")
		return habilidades[0]
	}
	return habilidades[opcion-1] // ajuste de índice base-0
}

// ─── PARTE 5: Función main ───────────────────────────────────────────────────

func main() {
	rand.Seed(time.Now().UnixNano())

	// Definición de personajes
	heroe := Personaje{
		Nombre:  "Ingeniero",
		Vida:    100,
		Ataque:  22,
		Defensa: 8,
	}
	dragon := Personaje{
		Nombre:  "Dragón de Datos",
		Vida:    120,
		Ataque:  20,
		Defensa: 5,
	}

	// ─── PARTE 2: Slices de habilidades ─────────────────────────────────────

	habilidadesHeroe := []Habilidad{
		{Nombre: "Golpe Básico", Poder: 0},
		{Nombre: "Golpe Fuerte", Poder: 8},
		{Nombre: "Magia Arcana", Poder: 15},
	}

	habilidadesDragon := []Habilidad{
		{Nombre: "Zarpazo", Poder: 0},
		{Nombre: "Llamarada", Poder: 10},
		{Nombre: "Rugido Devastador", Poder: 18},
	}

	fmt.Println("══════════════════════════════════════════")
	fmt.Println("   ⚔️   LA BATALLA POR LA CUMBRE DE GO   ⚔️")
	fmt.Println("══════════════════════════════════════════")
	fmt.Printf("🧙 %s  VS  🐉 %s\n", heroe.Nombre, dragon.Nombre)
	fmt.Println("──────────────────────────────────────────")

	// Bucle principal: continúa mientras ambos estén vivos
	turno := 1
	for heroe.EstaVivo() && dragon.EstaVivo() {
		fmt.Printf("\n═══ Turno %d ═══\n", turno)
		fmt.Printf("  💚 %s: %d HP  |  🔴 %s: %d HP\n",
			heroe.Nombre, heroe.Vida, dragon.Nombre, dragon.Vida)

		// — Turno del héroe —
		habilidadHeroe := elegirHabilidad(habilidadesHeroe)
		heroe.Ataque += habilidadHeroe.Poder            // bono temporal
		danoHeroe := heroe.Atacar(&dragon)
		heroe.Ataque -= habilidadHeroe.Poder            // revertir bono
		fmt.Printf("  🗡️  %s usa %s → %d de daño al %s\n",
			heroe.Nombre, habilidadHeroe.Nombre, danoHeroe, dragon.Nombre)

		if !dragon.EstaVivo() {
			break
		}

		// — Turno del dragón (IA aleatoria) —
		habilidadDragon := habilidadesDragon[rand.Intn(len(habilidadesDragon))]
		dragon.Ataque += habilidadDragon.Poder          // bono temporal
		danoDragon := dragon.Atacar(&heroe)
		dragon.Ataque -= habilidadDragon.Poder          // revertir bono
		fmt.Printf("  🔥 %s usa %s → %d de daño al %s\n",
			dragon.Nombre, habilidadDragon.Nombre, danoDragon, heroe.Nombre)

		turno++
	}

	// ─── Resultado final ─────────────────────────────────────────────────────
	fmt.Println("\n══════════════════════════════════════════")
	if heroe.EstaVivo() {
		fmt.Println("🏆 ¡VICTORIA! Compilaste al Dragón de Datos.")
	} else {
		fmt.Println("💀 DERROTA. Tu ejecución fue terminada.")
	}
	fmt.Println("══════════════════════════════════════════")
}