// Biblioteca Digital - Sistema de gestión de libros electrónicos con programación funcional.
// Ejecutar desde esta carpeta: go run .
package main

import (
	"fmt"

	"biblioteca/lectura"
	"biblioteca/libros"
	"biblioteca/prestamos"
	"biblioteca/reportes"
)

func main() {
	fmt.Println("===== BIBLIOTECA DIGITAL =====")

	// Módulo libros
	catalogo := libros.CatalogoInicial()
	fmt.Println("Disponibles:", libros.Titulos(libros.Disponibles(catalogo)))
	fmt.Println("Novelas:", libros.Titulos(libros.PorGenero(catalogo, "novela")))

	l, encontrado := libros.Buscar(catalogo, 1)
	fmt.Println("Buscar ID 1:", l.Titulo, encontrado)

	// Módulo prestamos
	p1, err := prestamos.Prestar(catalogo, 1, 1, "Oliver", 7)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	catalogo = libros.CambiarDisponibilidad(catalogo, 1, false)
	fmt.Println("Préstamo", p1.ID, "de", l.Titulo, "->", p1.Estado)

	_, err = prestamos.Prestar(catalogo, 2, 4, "Ana", 7)
	fmt.Println("Intento de prestar 1984:", err)

	multaDiaria := prestamos.Multa(0.25)
	fmt.Printf("Multa por 3 días de retraso: $%.2f\n", multaDiaria(3))

	// Módulo lectura
	progreso := lectura.Progreso{Libro: l, PaginasLeidas: 0}
	progreso = lectura.Avanzar(progreso, 55)
	fmt.Printf("Progreso de %s: %.1f%%\n", l.Titulo, lectura.Porcentaje(progreso))

	// Módulo reportes
	p2 := prestamos.Devolver(p1)
	todos := []prestamos.Prestamo{p1, p2}
	fmt.Println("Préstamos activos:", len(reportes.Activos(todos)))
	fmt.Printf("Total multas: $%.2f\n", reportes.TotalMultas([]int{3, 0, 5}, multaDiaria))
}
