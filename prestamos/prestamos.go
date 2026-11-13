// Paquete prestamos: préstamo y devolución de libros.
package prestamos

import (
	"errors"

	"biblioteca/libros"
)

type Prestamo struct {
	ID      int
	LibroID int
	Usuario string
	Dias    int
	Estado  string
}

// Prestar devuelve dos valores: el préstamo creado y un error si el libro no está disponible.
func Prestar(catalogo []libros.Libro, id int, libroID int, usuario string, dias int) (Prestamo, error) {
	l, ok := libros.Buscar(catalogo, libroID)
	if !ok {
		return Prestamo{}, errors.New("el libro no existe")
	}
	if !l.Disponible {
		return Prestamo{}, errors.New("el libro no está disponible: " + l.Titulo)
	}
	return Prestamo{ID: id, LibroID: libroID, Usuario: usuario, Dias: dias, Estado: "activo"}, nil
}

// Devolver devuelve una COPIA del préstamo marcada como devuelta.
func Devolver(p Prestamo) Prestamo {
	p.Estado = "devuelto"
	return p
}

// Multa es un closure: devuelve una función que calcula la multa con la tarifa dada.
func Multa(tarifaPorDia float64) func(diasRetraso int) float64 {
	return func(diasRetraso int) float64 {
		if diasRetraso <= 0 {
			return 0
		}
		return tarifaPorDia * float64(diasRetraso)
	}
}
