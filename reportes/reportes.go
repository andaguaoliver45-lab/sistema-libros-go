// Paquete reportes: estadísticas de la biblioteca.
package reportes

import (
	"biblioteca/funcional"
	"biblioteca/prestamos"
)

// Activos devuelve los préstamos que aún no se devuelven.
func Activos(ps []prestamos.Prestamo) []prestamos.Prestamo {
	return funcional.Filter(ps, func(p prestamos.Prestamo) bool { return p.Estado == "activo" })
}

// TotalMultas aplica la función de multa a cada retraso y suma todo (Map + Reduce).
func TotalMultas(retrasos []int, calcular func(int) float64) float64 {
	multas := funcional.Map(retrasos, calcular)
	return funcional.Reduce(multas, 0.0, func(acc, m float64) float64 { return acc + m })
}
