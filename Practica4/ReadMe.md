# Minimum Path Sum (Práctica 4)

### Fecha de entrega: jueves 16 de noviembre, 2023

## Hecho por : Emma Alicia Jimenez Sanchez (320046155) 

## Sobre la practica:

Dada una matriz `m x n` llena de números no negativos, encuentre un camino desde arriba a la izquierda hasta abajo a la derecha, que minimice la suma de todos los números a lo largo de su camino.

## Compilar y Ejecutar

Para compilar el programa se debe de ejecutar el siguiente comando en la terminal:

```
go build Solution.go
```

Para correr el programa se debe escribir el siguiente comando en la terminal:

```
go run Solution.go
```
### Explicación del algoritmo 
Dada una matriz como parametro se debe de calcular la ruta más corta desde la posición `(0,0)` hasta la posición `(n,m)` de la matrix.

Para esto debemos ver que se va a crear una matriz para poder los valores de la suma, se tiene que recorrer la matriz `grid`, para esto se utiliza 2 for's anidados, donde `i` va a corresponder al número de los renglones y `j` el número de columnas, y se debe de checar los siguientes casos:
 +   Si `i` es igual a 0 y `j` es diferente a 0:
 Como solo el problema indica que solo se puede mover de abajo y a la derecha, entonces solo debemos checar el número en la anterior posición pero en el mismo renglón y se suma con el número que esta en la misma posición de la matriz `grid`.

 Ejemplo:

 Primer renglón de la matriz `grid`: `[1,2,3,4]`

 Primer renglón de la matriz nueva : `[0,0,0,0]`

 Como no se toma en cuenta la primera posición, entonces pasamos a la posición `(0, 1)`, donde en la matriz va a ser 0(porque se encuentran en la anteior posición en el mismo renglón) más el número que se encuentra en esa misma posición en la matriz `grid`, por lo que el número que va a ir en `(0,1)` es 2.

 Se sigue checando de esta manera para los demás números del primer renglón de la nueva matriz, siendo así que su primer renglón es : `[0,2,5,9]`

 +   Si `j` es igual a 0 e `i` es diferente a 0:
Para la columna se repite el mismo proceso que en el anteior caso, solo que en lugar de checar la posición anterior del renglón, ahora es el de la columna. 

 +   Si `j` e `i` es diferente de 0:
Cunado ya son diferentes, debemos repetir el proceso que se explico antes para los renglones, solo que ahora se debe de tomar en cuenta ambos valores, la posición anterior del renglón como el de la columna, sin embargo aquí hay que checar, ya que se pide que sea la suma mínima para poder llegar al valor que se quiere, por lo que se creó una función auxiliar `min` el cuál dice que valor es el mínimo dependiendo de los valores ingresados.
Al tener ya el valor mínimo se suma al número que se encuentra en la misma posición de la matriz `grid`.

Veamos que no se toma el caso cuando ambas con 0 ya que desde ese punto iniciamos la suma y vamos recorriendo lo demás.

Al final solo se regresa el número en la posición `n-1,m-1` más el número que se encuentra en la posición `(0,0)` de la matriz `grid`.

### ¿Cómo se llegó a la solución?

Al leer el problema y por haber visto el semestre pasado un poyecto similar a este problema, lo primero que pensé es que se podía usar gráficas y con el algoritmo de Dijsktra se podía llegar a la solución, sin embargo implementar gráficas en golang esta un poco tedioso, ya que aunque te den un paquete con el código no se si uno mismo lo debía de implementar o solo era de copiar y pegar para poder usarlo, pero aun así el código se hizo muy grande y no me gustó, aparte no le entendí como utilizar las gráficas en go.

Después trate de ininciar por el valor final e ir guardando la suma mínima en una variable auxiliar, es decir, iniciar en donde queremos llegar. Por lo que con esto pensé en solo usar la función auxiliar mínimo para poder ver que valor es el mínimo y así llegar a donde se debía empezar, al checar lanzaba el error `indexerror` (ya que no tomaba en cuenta cuando `i` y `j`, y tenía que el valor mínimo como: `min(grid[i-1])(grid[j-1])`) por lo que cuando era 0 daba un valor que no estaba en la matriz, entonces se tuvo que hacer los casos cuando eran iguales a 0, en otras palabras cuando `i` es diferente d 0 y `j` es igual a 0, y su viceversa.

Después ví que no era necesario usar `min()` cuando estabamos en el primer renglón y en la primera columna, ya que como lo estaba implementando solo se podía mover a la izquierda o hacia arriba, y como en la primera columna no hay ningún número al cuál ver su valor mínimo , solo se tomaba el que estaba antes de su posición, al igual con el renglón, solo que aquí no había un número arriba por el cuál comprara, por lo que solo se tomaba el valor que estaba antes.

Sin embargo no salía este hasta que ya que la variable auxiliar no daba el resultado correcto, así que decide mejor guardar los valores en una matriz auxiliar de las misma dimensiones que la matriz `grid` (solo que aquí surgió otro problema ya que no se puede inicializar una matriz con valores de tipo `int` deben ser constantes) y solo regresar la posición `(0,0)` donde iba a estar la suma mínima acumulada desde el final hasta el inicio, pero aun no salía.

Después trate de resolver el problema a mano para ver que estaba fallando y por pura curiosidad decidi cambiarle como lo recorría, ahora lo recorría desde la posición inicial hasta la final, primero con un for recorriendo renglones y otro for anidado recorriendo columnas y solo regresar la posición `(n-1,m-1)`, y daba un valor más cercano a lo que el leetcode daba de ejemplo, sin embargo al imprimir porque no daba el valor deseado note que le faltaba sumar el primer valor de la matriz, es decir el número de la posición `(0,0)` de la matriz `grid`, lo sume al valor que regresaba y ya dió.

Solo que aún me estaba causando duda que ¿por qué así lo acepataba y por que de la otra manero no?, y me metí a checar el algoritmo en youtube y entendí mejor como es que funcionaba.
