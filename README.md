# 🚀 Aprendiendo Go - Día 1

Repositorio de aprendizaje de Go (Golang) donde documento mi progreso diario aprendiendo los fundamentos del lenguaje.

## 📅 Día 1 - Fundamentos Básicos

### ✅ Conceptos Aprendidos

#### 1. Variables y Tipos de Datos
- Tipos básicos: `int`, `float32`, `string`, `bool`
- Declaración de variables con `var` y con inferencia de tipo usando `:=`
- Constantes con `const`
- Variables de nivel de paquete vs variables locales

#### 2. Arrays (Arreglos)
- Declaración: `var arr [5]int`
- Tamaño fijo e inmutable
- Acceso por índice: `arr[0]`
- Modificación de elementos
- Función `len()` para obtener longitud
- Los arrays son valores (se copian al asignar)

#### 3. Slices (Arreglos Dinámicos)
- Declaración: `slice := []int{1, 2, 3}`
- Tamaño dinámico y flexible
- Creación desde arrays con slicing: `arr[0:3]`
- Función `make()`: `make([]int, length, capacity)`
- Operaciones:
  - `append()` - Agregar elementos
  - `len()` - Longitud actual
  - `cap()` - Capacidad total
  - Concatenación con `...` (spread operator)

#### 4. Paquetes y Visibilidad
- Organización de código en paquetes
- Visibilidad pública: nombres con **Mayúscula inicial**
- Visibilidad privada: nombres con **minúscula inicial**
- Alias de importación: `import alias "paquete"`
- Comandos:
  - `go mod init <nombre>` - Inicializar módulo
  - `replace` directive - Usar paquetes locales

#### 5. Funciones de Output (fmt)
- `fmt.Print()` - Imprime sin salto de línea
- `fmt.Println()` - Imprime con salto de línea automático
- `fmt.Printf()` - Imprime con formato
- Verbos de formato: `%v`, `%T`, `%d`, `%f`, `%s`, `%t`

## 📂 Estructura del Proyecto

```
golang/
├── main.go                 # Punto de entrada con demostración completa
├── go.mod                  # Definición del módulo
├── fundamentos/            # Paquete con fundamentos básicos
│   ├── variables.go        # Variables y tipos de datos
│   ├── arays.go           # Arrays (arreglos)
│   └── slices.go          # Slices (arreglos dinámicos)
├── paquetes/              # Paquete de ejemplo
│   ├── paquetes.go        # Demostración de visibilidad
│   └── go.mod
└── ejercicios/            # Ejercicios prácticos
    ├── bank.go
    └── go.mod
```

## 🎯 Ejecutar el Proyecto

```bash
# Compilar
go build -o app main.go

# Ejecutar
./app
```

## 💡 Mejores Prácticas Aplicadas

✅ Nombres de variables en camelCase  
✅ Nombres de paquetes en minúsculas  
✅ Documentación con comentarios en español  
✅ Funciones públicas documentadas con comentarios  
✅ Separación lógica de conceptos en diferentes archivos  
✅ Uso de `gofmt` para formato consistente  

## 📚 Recursos Utilizados

- [Documentación oficial de Go](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)

## 🎓 Próximos Pasos (Día 2)

- [ ] Estructuras (structs)
- [ ] Métodos y receivers
- [ ] Interfaces
- [ ] Maps (diccionarios)
- [ ] Control de flujo (if, switch, for)

---

**Fecha de inicio:** 29 de octubre de 2025  
**Lenguaje:** Go 1.x  
**Estado:** 🟢 En progreso
