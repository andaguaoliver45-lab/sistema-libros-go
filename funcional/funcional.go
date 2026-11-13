// Paquete funcional: funciones genéricas de orden superior.
package funcional

// Filter devuelve un nuevo slice con los elementos que cumplen la condición.
func Filter[T any](xs []T, cond func(T) bool) []T {
	resultado := []T{}
	for _, x := range xs {
		if cond(x) {
			resultado = append(resultado, x)
		}
	}
	return resultado
}

// Map aplica una función a cada elemento y devuelve un nuevo slice.
func Map[T, U any](xs []T, f func(T) U) []U {
	resultado := make([]U, 0, len(xs))
	for _, x := range xs {
		resultado = append(resultado, f(x))
	}
	return resultado
}

// Reduce acumula todos los elementos en un solo valor.
func Reduce[T, A any](xs []T, inicial A, f func(A, T) A) A {
	acumulado := inicial
	for _, x := range xs {
		acumulado = f(acumulado, x)
	}
	return acumulado
}
