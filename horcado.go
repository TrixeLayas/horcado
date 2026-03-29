//eesta cosa de Rand.seed lo cheque y es para generar un numero alatorio para elegir una palabra alotoria
//de aqui sale https://www.geeksforgeeks.org/go-language/generating-random-numbers-in-golang/
// el time.now sirve para checar la hora actual  que son las  11:59pm
// y el otro el unixnano ps es para convertir eso a un numero gigante
// y ps ya por ultimo el rand.intt sirve para generar un numero alatorio

// el len sirve me sirvio para saber cantidad de letras que tiene la palabra
// el Scanln sirve para leer una letra del usuario y la guarda en la variable letra
// para poder  tener un progreso lo que use fue  de clare en un if a letra  como tipo texto y le di como entrada a palabra luego meti corchetes 
//para asi poder obtener la longitud de la palabra  luego declare progreso = progreso + letra para asi que se vaya autocomplentando 
package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	palabras_alo_wey := []string{"hola", "wey", "comoestas", "pepito", "boca"}
	rand.Seed(time.Now().UnixNano())
	palabra := palabras_alo_wey[rand.Intn(len(palabras_alo_wey))]
	var letra string
	intentos := 0
	progreso := ""
	for {
		fmt.Println("Adivina cual de todas las palabras te toco")
		fmt.Println()
		fmt.Println("progreso: de la palabra ", progreso)
		fmt.Println(" ")
		fmt.Println("escribe una letra para minimo iniciar")
		fmt.Println(" ")
		fmt.Scanln(&letra)
		fmt.Println(" ")
		if letra == string(palabra[len(progreso)]) {
			progreso = progreso + letra
			fmt.Println("ya mero")
		} else {
			intentos++
			fmt.Println("ya mero te quedan sin intentos:", intentos)
			fmt.Println(" ")
		}
		if progreso == palabra {
			fmt.Println("porfin das una:", palabra)
			return
		}
		if intentos == 3 {
			fmt.Println("te quedaste sin intentos sorry man")
			fmt.Println(" ")

			fmt.Println("la palabra esta como veo que no le atinaste :", palabra)
			fmt.Println(" ")
			return
		}
	}
}
