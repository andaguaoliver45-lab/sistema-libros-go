// Paquete libros: catálogo de libros electrónicos.
package libros

import "biblioteca/funcional"

type Libro struct {
	ID         int
	Titulo     string
	Autor      string
	Genero     string
	Paginas    int
	Disponible bool
}

// CatalogoInicial devuelve los libros de ejemplo de la biblioteca.
func CatalogoInicial() []Libro {
	return []Libro{
		{1, "Huasipungo", "Jorge Icaza", "novela", 220, true},
		{2, "Cien años de soledad", "Gabriel García Márquez", "novela", 470, true},
		{3, "El principito", "Antoine de Saint-Exupéry", "infantil", 96, true},
		{4, "1984", "George Orwell", "ciencia ficcion", 328, false},
	}
}

// Buscar devuelve dos valores: el libro y si fue encontrado.
func Buscar(lista []Libro, id int) (Libro, bool) {
	for _, l := range lista {
		if l.ID == id {
			return l, true
		}
	}
	return Libro{}, false
}

// Disponibles devuelve solo los libros que se pueden prestar.
func Disponibles(lista []Libro) []Libro {
	return funcional.Filter(lista, func(l Libro) bool { return l.Disponible })
}

// PorGenero devuelve los libros de un género específico.
func PorGenero(lista []Libro, genero string) []Libro {
	return funcional.Filter(lista, func(l Libro) bool { return l.Genero == genero })
}

// CambiarDisponibilidad devuelve un NUEVO catálogo con el libro actualizado.
func CambiarDisponibilidad(lista []Libro, id int, disponible bool) []Libro {
	return funcional.Map(lista, func(l Libro) Libro {
		if l.ID == id {
			l.Disponible = disponible
		}
		return l
	})
}

// Titulos transforma la lista de libros en una lista de títulos.
func Titulos(lista []Libro) []string {
	return funcional.Map(lista, func(l Libro) string { return l.Titulo })
}
