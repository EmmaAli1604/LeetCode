# Projection Area of 3D Shapes (Práctica 2)

### Fecha de entrega: jueves 31 de agosto, 2023

## Hecho por : Emma Alicia Jimenez Sanchez (320046155) 

## Sobre la practica:

Se le da un grid de n x n donde colocamos unos cubos de 1 x 1 x 1 que están alineados con los ejes x, y y z.

Cada valor v = grid[i][j] representa una torre de v cubos colocados encima de la celda (i, j).

Vemos la proyección de estos cubos sobre los planos xy, yz y zx.

Una proyección es como una sombra, que mapea nuestra figura tridimensional a un plano bidimensional. Estamos viendo la "sombra" cuando miramos los cubos desde arriba, de frente y de lado.

Devuelve el área total de las tres proyecciones.

Para compilar el programa se debe de ejecutar el siguiente comando en la terminal:

```
go build Solution.go
```

Para correr el programa se debe escribir el siguiente comando en la terminal:

```
go run Solution.go
```
### Explicación del algoritmo 
Primero lo que hacemos es declarar las variables que vamos a utilizar que sería ` result `.\
Ahora recorremos los valores de nuestra matriz primero la fila, antes de pasar al siguiente `for` debemos inicializar las variables `max_X` y  `max_Y` iguales a 0.\
Después la columna, siendo así nuestro algoritmo va a ser `O(n²)`, donde debemos checar si en la casilla el valor es diferente de cero, si lo es entonces sumamos un 1 a la variable `result`, después verificamos si el valor es mayor a al maxímo que fijamos de la fila, si lo es reasigamos la casilla como nuevo valor máximo de la fila.\
Finalmente checamos el valor maximo de la columna, si el valor de la casilla es mayor que el valor maximo que tenemos lo reasignamos.\
Cuando se finalize el `for` de la columna debemos hacer la siguiente suma `result = result + max_X + max_Y`, para ir guardando los valores que ya tenemos. Checamos esto con todas las filas y columnas de la matriz `grid` para obtener el area de la sombra 3D. Al final regresamos la variable `result`.

### ¿Cómo se llegó a la solución?
Primero pense en la solucion donde cuantos valores de las casillas de `grid` son diferentes de cero y el valor se iba acumulando en la variable `result`, después sumar los valores de la última fila y de la última columana, al final se regresaba el resulatdo de sumar `result + max_X +max_Y `, esto funcionaba para los ejemplos que daba LeetCode, sin embargo al momento de dar `Submit`, no pasaba todos los test, ya que no conte con la siguiente matriz `grid = [[1,4],[1,1]]`. Con este algoritmo el resultado era `11`, cuando debía ser `14`.\
Realmente no entendi como fue que se saco ese resulatdo, así que vi una solución y tampoco le entendi, hasta que hice un dibujo y la matriz para poder ver el patron, donde vi que en el cuadrante donde tenemos a los dos bloques de 1 , la sombra de uno de ellos se ve opcada por el de tamaño 4, así teniendo así la siguiente suma `4 + + (4+1) + (4+1) = 14`, así vi que debemos ver el valor maximo tanto de la fila como de la columna, y se debían de sumar con el resultado.\
Ahora un error que cometí fue no reiniciar los valores de las variables `max_X` y `max_Y`, ya que me deba un valor mucho más elevado de lo que debía, siendo así que se tomaba los máximos valores de otras columnas y filas, los sumaba y daba valores erroneos al momento de sumar como `17` (lo más curioso es que solo fue en este caso).
Así fue como llegue a la solución, después de unas horas frente a la computadora.

## Authors

- [@EmmaAli1604](https://github.com/EmmaAli1604)

