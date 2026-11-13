# 📚 Biblioteca Digital

Sistema de gestión de libros electrónicos desarrollado en **Go** aplicando **programación funcional**.

## Objetivo
Administrar el catálogo de libros electrónicos, los préstamos, el progreso de lectura y los reportes de una biblioteca digital, usando funciones puras, funciones de orden superior, closures e inmutabilidad de datos.

## Módulos (paquetes)
| Paquete | Descripción |
|---|---|
| `funcional` | Funciones genéricas `Map`, `Filter` y `Reduce` |
| `libros` | Catálogo: listar, buscar, filtrar por género y cambiar disponibilidad |
| `prestamos` | Prestar y devolver libros, cálculo de multas |
| `lectura` | Progreso de lectura y porcentaje leído |
| `reportes` | Préstamos activos y total de multas |

## Paquetes utilizados
Solo librería estándar de Go (`fmt`, `errors`). No se usan paquetes de terceros.

## Cómo ejecutar
```bash
go run .
```
