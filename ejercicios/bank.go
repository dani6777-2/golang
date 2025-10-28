package main

import (
	"errors"
	"fmt"
)

// Definición de estructuras
type Cliente struct {
	ID     int
	Nombre string
	Email  string
	Activo bool
}

type Cuenta struct {
	Numero    string
	ClienteID int
	Saldo     float64
}

type Banco struct {
	Nombre   string
	Clientes []Cliente
	Cuentas  []Cuenta
}

// TODO: Completar las siguientes funciones

// 1. Función para crear un nuevo cliente
func (b *Banco) AgregarCliente(nombre, email string) error {
	// Validar que el email no esté vacío
	if email == "" {
		fmt.Println("El correo no debe estar vacio")
	}
	// Crear nuevo cliente con ID autoincremental
	nuevoID := len(b.Clientes) + 1
	cliente := Cliente{
		ID:     nuevoID,
		Nombre: nombre,
		Email:  email,
		Activo: true,
	}
	// Agregar al slice de clientes
	b.Clientes = append(b.Clientes, cliente)
	return nil
}

// 2. Función para crear una cuenta bancaria
func (b *Banco) CrearCuenta(clienteID int, saldoInicial float64) error {
	// Validar que el cliente exista
	client, err := b.buscarCliente(clienteID)
	if err != nil {
		return err
	}
	if !client.Activo {
		return errors.New("el cliente no está activo")
	}
	// Validar saldo inicial positivo
	if saldoInicial <= 0 {
		return errors.New("el saldo inicial debe ser positivo")
	}
	// Generar número de cuenta automático
	numeroCuenta := fmt.Sprintf("CTA-%03d", len(b.Cuentas)+1)
	cuenta := Cuenta{
		Numero:    numeroCuenta,
		ClienteID: clienteID,
		Saldo:     saldoInicial,
	}
	// Agregar al slice de cuentas
	b.Cuentas = append(b.Cuentas, cuenta)

	return nil
}

// 3. Función para depositar dinero
func (b *Banco) Depositar(numeroCuenta string, monto float64) error {
	// Buscar la cuenta
	acount, err := b.buscarCuenta(numeroCuenta)

	if err != nil {
		return err
	}
	// Validar monto positivo
	if monto <= 0 {
		return errors.New("el monto debe ser mayor a 0.")
	}
	// Actualizar saldo
	acount.Saldo += monto
	fmt.Printf("Depósito exitoso. Nuevo saldo: %.2f\n", acount.Saldo)
	return nil
}

// 4. Función para retirar dinero
func (b *Banco) Retirar(numeroCuenta string, monto float64) error {
	// Buscar la cuenta
	account, err := b.buscarCuenta(numeroCuenta)

	if err != nil {
		return err
	}
	// Validar monto positivo
	// Validación correcta:
	if monto <= 0 {
		return errors.New("el monto debe ser mayor a 0")
	}

	// Verificar fondos suficientes
	if account.Saldo < monto {
		return errors.New("fondos insuficientes")
	}

	// Actualizar saldo
	account.Saldo -= monto
	fmt.Printf("Retiro exitoso. Nuevo saldo: %.2f\n", account.Saldo)
	return nil
}

// 5. Función para transferir entre cuentas
func (b *Banco) Transferir(desdeCuenta, haciaCuenta string, monto float64) error {
	// Implementar lógica de transferencia
	accountSender, err := b.buscarCuenta(desdeCuenta)
	if err != nil {
		return err
	}

	if accountSender.Saldo <= 0 {
		return errors.New("No tiene fondos suficientes para realizar esta transerencia.")
	}

	accountAddressee, err := b.buscarCuenta(haciaCuenta)
	if err != nil {
		return err
	}

	if monto <= 0 {
		return errors.New("el monto de la transferencia debe ser mayor a 0.")
	}

	accountSender.Saldo -= monto
	accountAddressee.Saldo += monto

	fmt.Println("transaccion realizada con exito, desde: ", accountSender.Numero, " hacia: ", accountAddressee.Numero, "por un monto de: ", monto)

	return nil
}

// 6. Función para mostrar información del banco
func (b *Banco) MostrarInformacion() {
	// Mostrar todos los clientes y sus cuentas
	fmt.Printf("Banco: %s\n", b.Nombre)
	fmt.Println("Clientes:")
	for _, cliente := range b.Clientes {
		status := "Inactivo"
		if cliente.Activo {
			status = "Activo"
		}
		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Estado: %s\n", cliente.ID, cliente.Nombre, cliente.Email, status)
	}
	fmt.Println("Cuentas:")
	for _, cuenta := range b.Cuentas {
		fmt.Printf("Número: %s, ClienteID: %d, Saldo: %.2f\n", cuenta.Numero, cuenta.ClienteID, cuenta.Saldo)
	}
}

// Función auxiliar para buscar cliente
func (b *Banco) buscarCliente(id int) (*Cliente, error) {
	for i := range b.Clientes {
		if b.Clientes[i].ID == id && b.Clientes[i].Activo {
			return &b.Clientes[i], nil
		}
	}
	return nil, errors.New("cliente no encontrado")
}

// Función auxiliar para buscar cuenta
func (b *Banco) buscarCuenta(numero string) (*Cuenta, error) {
	for i := range b.Cuentas {
		if b.Cuentas[i].Numero == numero {
			return &b.Cuentas[i], nil
		}
	}
	return nil, errors.New("cuenta no encontrada")
}

func main() {
	// Crear banco
	banco := Banco{
		Nombre: "Mi Banco Go",
	}

	// Menú interactivo (opcional - ejercicio avanzado)
	fmt.Println("=== SISTEMA BANCARIO ===")

	// Aquí puedes agregar un menú interactivo o pruebas manuales
	// Ejemplo de uso:

	// Agregar clientes
	banco.AgregarCliente("Juan Pérez", "juan@email.com")
	banco.AgregarCliente("María García", "maria@email.com")

	// Crear cuentas
	banco.CrearCuenta(1, 1000.0)
	banco.CrearCuenta(2, 500.0)

	// Mostrar información
	banco.MostrarInformacion()

	// Realizar operaciones
	banco.Depositar("CTA-001", 200.0)
	banco.Retirar("CTA-001", 150.0)
	banco.Transferir("CTA-001", "CTA-002", 100.0)

	banco.MostrarInformacion()
}
