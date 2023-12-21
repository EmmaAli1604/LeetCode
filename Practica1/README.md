# Product of Array Except Self (Práctica 1)

### Fecha de entrega: jueves 31 de agosto, 2023

## Hecho por : Emma Alicia Jimenez Sanchez (320046155) 

## Sobre la practica:

Dada un arreglo de enteros nums, devuelve un arreglo answer tal que answer[i] es igual al producto de todos los elementos de nums excepto nums[i].

Se garantiza que el producto de cualquier prefijo o sufijo de nums cabe en un entero de 32 bits.

Debes escribir un algoritmo que se ejecute en tiempo O(n) y sin utilizar la operación de división.

### Para compilar y correr el programa

Para compilar  y correr el programa se debe de ejecutar el siguiente comando en la terminal:
```
go run Solution.go
```
### Explicación del algoritmo 
Con el ejemplo que da leetcode:
```
[1 2 3 4]
```
Primero lo que hacemos es multiplicar de izquierda a derecha y en dos en dos el arreglo dado. 

El primer recorrido del arreglo los resultados los vamos guardando en una lista:
```
[1 2 6 24]
```
Después lo recorremos de derecha a izquierda, igual multiplicando en dos en dos y los resultados los vamos guardando en una segunda lista:

```
[24 24 12 4]
```

Ya con esto, recorremos el arreglo por tecera vez, pero ahora checamos en el índice que se encuentra cada número; si se encuentra en el primer lugar lo que hacemos es checar la segunda lista y ver que número tiene en la derecha, como no tiene ningún número a su izquierda, ese sería el primer resultado que se guarda en el arreglo `answer`, si el índice del número se encuentra en `i > 0 && i < n-1` checamos su índice en la primera lista y ver que número tiene en la izquierda, este se va a multiplicar con el número que esta en la derecha del índice en la segunda lista, y si es el último número del arreglo entonces es la vicerversa del primer caso.
Teniendo así que el arreglo `answer` es:

```
[24 12 8 6]
```

### ¿Cómo se llegó a la solución?
Después de ver en el laboratorio que tenemos esta dos arrgelos de número multiplicados pues vi que en la primera lista, si retrocedemos un número de donde se encuentra el índice obtenemos la multiplicación de los números que van antes del número que esta en el índice y en la segunda lista ocurre los mismo pero con los que van después del índice en el que estamos.
El único pronblema era cuando eran el primero y el último, ya que no tienen un número con que multiplicar, entonces nada más se agarra el número que tiene a la derecha o izquierda de la primera o segunda lista.
Con esto se consigue multiplicar el arreglo mismo excepto por el índice en el que estamos.
### ¿Por qué se ocupó las listas?
Pense en guardar los valores numéricos en listas ya que como son dinámicas podemos guardar elementos sin establecer un tamaño predeterminado, porque como no sabemos de que tamaño va a ser arreglo que nos van a mandar, no podemos definir el tamaño.

