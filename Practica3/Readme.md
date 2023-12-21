# Sort Colors (Práctica 3)

### Fecha de entrega: jueves 19 de octubre, 2023

## Hecho por : Emma Alicia Jimenez Sanchez (320046155) 

## Sobre la practica:

Dada una matriz `nums` con n objetos de color rojo, blanco o azul, ordénelos en su lugar de modo que los objetos del mismo color sean adyacentes, con los colores en el orden rojo, blanco y azul.

Usaremos los enteros 0, 1 y 2 para representar el color rojo, blanco y azul, respectivamente.

Debes resolver este problema sin utilizar la función de ordenación de la biblioteca.

Para compilar el programa se debe de ejecutar el siguiente comando en la terminal:

```
go build Solution.go
```

Para correr el programa se debe escribir el siguiente comando en la terminal:

```
go run Solution.go
```
### Explicación del algoritmo 
Iniciamos a 3 variables donde van a ser contadores para los números que nos mandan en el parametro, teniendo así `nums_0`,`nums_1` y `nums_2`, además de `index`, la cuál nos va a ayudar a guardar la última posición de nuestro arreglo para poder seguir agregando los número que faltan.

Iniciamos un for que recorre el arreglo `nums` y cuenta cuantos 0,1's y 2's hay en este y guarda estos datos en las variables ya mencionadas.

Revisamos los siguientes casos:
* Si el arreglo ya esta ordenado: \
    Si nos pasan un arreglo que solo tiene un elemento entonces por default el arreglo ya esta ordenado, por ejemplo: `[0]`.
* Si el arreglo solo contiene 0,1 o 2: \
    Con esto checamos que los contadores, ya sea del 0,1 o 2, por lo menos uno es diferente de 0 mientras que los demas deben ser zero, un ejemplo claro es: `[1,1,1,1]`
* Si el arreglo no contiene 0's:\
    Con el algoritmo default no funciona para este caso ya que toma en cuenta las posiciones que van para el 0, entonces para evitar que se genere problemas, entonces nada mas hacemos 2 for's lineales para poder posicionar los numeros de 1 y 2 en orden en el arreglo que nos dan.
* Default: Ordenamiento de 0,1 y 2: \
    Primero hacemos un for que va de `0 <= i < nums_0; i++` para poder poner primero la cantidad de 0 que haya en el arreglo, ademas de que el index es igualado a la última posición que `i` va a ocupar.\
    Después hacemos otro for que va de `0 <= i+1 < index+nums_1+1`, agregamos el 1 para poder indicar que va a iniciar una posición después de donde termino , para poder así ordenar ahora ordenar los 1 que hay en eñ arreglo.
    Para el for que va a ordenar la cantidad de 2 del arreglo hará el mismo proceso que el for que posiciona los 1 en el arreglo.

Con esto ordenamos los colores("números") que el parametro pasa.
### ¿Cómo se llegó a la solución?
LLegue a etsa solución ya que pense en solo tener un for para poder recorrer el arreglo, saber cuantos 0,1 y 2 hay en el arreglo `nums`, para que el espacio fuera constante, modifique el mismo arreglo del parametro. Sin embargo al momento de pasarlo por el leetcode, no pasaba los casos que se enuncian antes de hacer el for con los tres números, así que hice unos casos especificos para que el programa no tronara al momento de tener estos. Estos casos era como: `[0]`.`[1,2]`.`[2]`. Con esto el leetcode ya acepto el programa.

*Dato Curioso: En algunos casos el programa tarda 2ms, donde indica que se vencio al 25% de usuarios que usuaron golang como en este desafio, sin embargo en otras ocasiones aparace que tarda 0ms y que vencio al 100% de usuarios que implementaron golang*



## Authors

- [@EmmaAli1604](https://github.com/EmmaAli1604)
