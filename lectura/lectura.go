// Paquete lectura: progreso de lectura del usuario.
package lectura

import (
	"biblioteca/funcional"
	"biblioteca/libros"
)

type Progreso struct {
	Libro         libros.Libro
	PaginasLeidas int
}

// Avanzar devuelve un NUEVO progreso con las páginas sumadas (sin pasar del total).
func Avanzar(p Progreso, paginas int) Progreso {
	p.PaginasLeidas += paginas
	if p.PaginasLeidas > p.Libro.Paginas {
		p.PaginasLeidas = p.Libro.Paginas
	}
	return p
}

// Porcentaje calcula cuánto del libro se ha leído.
func Porcentaje(p Progreso) float64 {
	return float64(p.PaginasLeidas) / float64(p.Libro.Paginas) * 100
}

// TotalPaginasLeidas suma las páginas leídas de todos los libros usando Reduce.
func TotalPaginasLeidas(lista []Progreso) int {
	return funcional.Reduce(lista, 0, func(acc int, p Progreso) int { return acc + p.PaginasLeidas })
}
