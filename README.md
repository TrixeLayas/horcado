horcado
Juego de Adivinar Palabra en Go

Cómo ejecutar
Tener Go instalado
Abrir la carpeta del proyecto en la terminal
Ejecutar: go run horcado.go

una ves ejecutado se vera asi en terminal 
C:\Users\ADMIN\Documents\horcado>go run horcado.go 
Adivina cual de todas las palabras te toco

progreso: de la palabra

escribe una letra para minimo iniciar





el juego que hice tine uan lista de 5 palabras:
hola
wey
comoestas
pepito
boca
El sistema selecciona una palabra de forma aleatoria usando rand.
El jugador debe escribir letras para ir formando la palabra poco a poco.

Ejemplo:
w
we
wey

Reglas solo avanza si aciertas si fallas devuelve intento 2 , y asi hasta llegar a 0
Si te equivocas, pierdes un intento
Tienes máximo 3 intentos
Si completas la palabra  ganas
Si fallas 3 veces  pierdes


Cómo está hecho
use estas 
//eesta cosa de Rand.seed lo cheque y es para generar un numero alatorio para elegir una palabra alotoria
//de aqui sale https://www.geeksforgeeks.org/go-language/generating-random-numbers-in-golang/
// el time.now sirve para checar la hora actual  que son las  11:59pm
// y el otro el unixnano ps es para convertir eso a un numero gigante
// y ps ya por ultimo el rand.intt sirve para generar un numero alatorio

// el len sirve me sirvio para saber cantidad de letras que tiene la palabra
// el Scanln sirve para leer una letra del usuario y la guarda en la variable letra
// para poder  tener un progreso lo que use fue  de clare en un if a letra  como tipo texto y le di como entrada a palabra luego meti corchetes
// para asi poder obtener la longitud de la palabra  luego declare progreso = progreso + letra para asi que se vaya autocomplentando

el significado del juego es 
Adivinar la palabra antes de quedarte sin intentos.
