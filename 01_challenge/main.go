package main

import (
	"fmt"
)

// Product 📦 representa un producto en el inventario
type Product struct {
	Stock          int // 🔢 Cantidad de unidades disponibles
	AlertThreshold int // ⚠️ Stock mínimo para generar una alerta
}

// UpdateInventory 🔄 actualiza el inventario basado en las ventas y genera alertas
/*
🎯 Objetivo:
    1. Actualizar el inventario basado en las ventas
    2. Generar alertas para productos con bajo stock
	3. Generar alertas para ventas que exceden el stock disponible

📥 Parámetros:
    - inventory: mapa de productos en el inventario
    - sales: mapa de ventas realizadas
	Nota: Revisa la función main para ver los datos de ejemplo.

📤 Retorno:
    - []string: slice de alertas generadas

📝 Formatos de alerta:
    1. Stock bajo:
       "Alerta: Stock bajo para [nombre del producto]. Quedan [cantidad] unidades."
    2. Venta excede stock:
       "Error: Venta de [cantidad] unidades de [nombre del producto] excede el stock disponible de [stock actual]."
*/
// TODO: Implementa esta función
func UpdateInventory(inventory map[string]Product, sales map[string]int) []string {
	// Tu código aquí
	return nil
}

func main() {
	// 🏪 Ejemplo de inventario inicial
	inventory := map[string]Product{
		"Laptop":   {Stock: 50, AlertThreshold: 10},
		"Mouse":    {Stock: 20, AlertThreshold: 5},
		"Keyboard": {Stock: 15, AlertThreshold: 5},
	}

	// 💰 Ejemplo de ventas realizadas
	sales := map[string]int{
		"Laptop":   45,
		"Mouse":    16,
		"Keyboard": 12,
	}

	// 🚀 Ejecutar la actualización del inventario
	alerts := UpdateInventory(inventory, sales)

	// 📢 Imprimir alertas generadas
	fmt.Println("🚨 Alertas generadas:")
	for _, alert := range alerts {
		fmt.Println(alert)
	}

	// 📊 Imprimir inventario actualizado
	fmt.Println("\n📦 Inventario actualizado:")
	for product, details := range inventory {
		fmt.Printf("%s: %d en stock\n", product, details.Stock)
	}
}
