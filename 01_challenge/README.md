# 🚀 Challenge 01: Sistema de Gestión de Inventario para una Tienda en Línea

En este short challenge, tu objetivo es implementar una función en Go que gestione el inventario de una tienda en línea, actualizando las cantidades de productos después de realizar ventas y generando alertas para productos con bajo stock y cuando la venta excede la cantidad en stock.

## 📋 Descripción

Este es un problema común en sistemas de comercio electrónico. Un short challenge perfecto para poner a prueba tus habilidades en Go y tu capacidad para resolver problemas del mundo real.

### Reglas:
1. Cada producto tiene un nombre, una cantidad en stock y un umbral de alerta.
2. Después de cada venta, se debe actualizar la cantidad en stock.
3. Si la cantidad en stock cae por debajo del umbral de alerta, se debe generar una alerta.
4. Si la cantidad en stock es igual o mayor al umbral de alerta, no se genera ninguna alerta.
5. Si la cantidad en stock es menor que la cantidad vendida, se debe generar una alerta.
4. No se pueden vender más unidades de las que hay en stock.

Nota: En main.go encontrarás la estructura que deberán tener las alertas generadas.

## 🛠️ Instrucciones

1. **Implementa la función `UpdateInventory` en el archivo `inventory.go`:**
    - La función `UpdateInventory` debe recibir un mapa de productos (nombre del producto como clave, y un struct `Product` como valor) y un mapa de ventas (nombre del producto como clave y cantidad vendida como valor).
    - La función debe devolver un slice de strings con las alertas generadas.
    - **Importante:** No cambies el nombre de la función ni modifiques otros archivos excepto `main.go`.

2. **Ejemplo de Input y Output:**
    - **Input:**
      ```
      inventory := map[string]Product{
          "Laptop": {Stock: 50, AlertThreshold: 10},
          "Mouse": {Stock: 20, AlertThreshold: 5},
          "Keyboard": {Stock: 15, AlertThreshold: 5},
      }
      sales := map[string]int{
          "Laptop": 45,
          "Mouse": 16,
          "Keyboard": 12,
      }
      ```
    - **Output:**
      ```
      []string{
          "Alerta: Stock bajo para Laptop. Quedan 5 unidades.",
          "Alerta: Stock bajo para Mouse. Quedan 4 unidades.",
          "Alerta: Stock bajo para Keyboard. Quedan 3 unidades.",
      }
      ```

3. **Ejecuta los tests:**
    - Para validar tu solución, puedes ejecutar los tests.

    - Si estás en la raíz del proyecto, ejecuta el siguiente comando:
    ```bash
   go test -v -count=1 01_challenge
    ```
    Este comando ejecutará únicamente los tests de este challenge, mostrando la salida detallada.
    - Si estás dentro de la carpeta `01_challenge`, puedes ejecutar 
    ```bash
    go test -v -count=1
    ```

## 📺 Respuestas y Explicaciones

La respuesta completa y una [explicación detallada](https://youtu.be/Qeym4sQpIYk) están disponibles en mi canal de YouTube. ¡No olvides suscribirte para no perderte ningún detalle! [YouTube: Harley Zapata](https://www.youtube.com/@harleyzapata) 🎉

---
