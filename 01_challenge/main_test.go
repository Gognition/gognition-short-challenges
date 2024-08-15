/*********************************
* No modifiques este archivo (🚫)
*********************************/

package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

var (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

func TestUpdateInventory(t *testing.T) {
	testCases := []struct {
		name           string
		inventory      map[string]Product
		sales          map[string]int
		expectedAlerts []string
	}{
		{
			name: "Case 1: Multiple alerts",
			inventory: map[string]Product{
				"Laptop":   {Stock: 50, AlertThreshold: 10},
				"Mouse":    {Stock: 20, AlertThreshold: 5},
				"Keyboard": {Stock: 15, AlertThreshold: 5},
			},
			sales: map[string]int{
				"Laptop":   45,
				"Mouse":    16,
				"Keyboard": 12,
			},
			expectedAlerts: []string{
				"Alerta: Stock bajo para Laptop. Quedan 5 unidades.",
				"Alerta: Stock bajo para Mouse. Quedan 4 unidades.",
				"Alerta: Stock bajo para Keyboard. Quedan 3 unidades.",
			},
		},
		{
			name: "Case 2: Single alert",
			inventory: map[string]Product{
				"Tablet":    {Stock: 30, AlertThreshold: 5},
				"Headphone": {Stock: 50, AlertThreshold: 10},
			},
			sales: map[string]int{
				"Tablet":    25,
				"Headphone": 30,
			},
			expectedAlerts: []string{
				"Alerta: Stock bajo para Tablet. Quedan 5 unidades.",
			},
		},
		{
			name: "Case 3: Alert for very low stock",
			inventory: map[string]Product{
				"Smartwatch": {Stock: 10, AlertThreshold: 3},
				"Charger":    {Stock: 100, AlertThreshold: 20},
			},
			sales: map[string]int{
				"Smartwatch": 8,
				"Charger":    50,
			},
			expectedAlerts: []string{
				"Alerta: Stock bajo para Smartwatch. Quedan 2 unidades.",
			},
		},
		{
			name: "Case 4: Error for overselling",
			inventory: map[string]Product{
				"Monitor": {Stock: 5, AlertThreshold: 2},
			},
			sales: map[string]int{
				"Monitor": 6,
			},
			expectedAlerts: []string{
				"Error: Venta de 6 unidades de Monitor excede el stock disponible de 5.",
			},
		},
		{
			name: "Case 5: Multiple alerts with exact threshold",
			inventory: map[string]Product{
				"Printer": {Stock: 15, AlertThreshold: 5},
				"Scanner": {Stock: 8, AlertThreshold: 2},
			},
			sales: map[string]int{
				"Printer": 10,
				"Scanner": 7,
			},
			expectedAlerts: []string{
				"Alerta: Stock bajo para Printer. Quedan 5 unidades.",
				"Alerta: Stock bajo para Scanner. Quedan 1 unidades.",
			},
		},
	}

	passedTests := 0
	totalTests := len(testCases)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultAlerts := UpdateInventory(tc.inventory, tc.sales)
			if !equalUnordered(resultAlerts, tc.expectedAlerts) {
				fmt.Printf("%sX Test fallido (%s): Se esperaba %v, pero se obtuvo %v. Revisa tu código por favor.%s\n",
					red, tc.name, tc.expectedAlerts, resultAlerts, reset)
				t.Fail()
			} else {
				fmt.Printf("%s✔ Test logrado satisfactoriamente (%s)%s\n", green, tc.name, reset)
				passedTests++
			}
		})
	}

	if passedTests == totalTests {
		fmt.Println(blue + "\n" + strings.Repeat("=", 40))
		fmt.Println(yellow + "🎉 ¡Felicitaciones! 🎉")
		fmt.Println(green + "Has pasado todos los tests exitosamente.")
		fmt.Println("Tu implementación del sistema de gestión de inventario es correcta.")
		fmt.Println(yellow + "Sigue así, ¡vas por buen camino!")
		fmt.Println(blue + strings.Repeat("=", 40) + reset)
	} else {
		fmt.Printf("\n%sHas pasado %d de %d tests. ¡Sigue intentando!%s\n", yellow, passedTests, totalTests, reset)
	}
}

func equalUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sortedA := make([]string, len(a))
	sortedB := make([]string, len(b))
	copy(sortedA, a)
	copy(sortedB, b)
	sort.Strings(sortedA)
	sort.Strings(sortedB)
	for i := range sortedA {
		if sortedA[i] != sortedB[i] {
			return false
		}
	}
	return true
}
