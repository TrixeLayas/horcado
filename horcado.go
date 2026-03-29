//eesta cosa de Rand.seed lo cheque y es para generar un numero alatorio para elegir una palabra alotoria
//de aqui sale https://www.geeksforgeeks.org/go-language/generating-random-numbers-in-golang/
// el time.now sirve para checar la hora actual  que son las  11:59pm
// y el otro el unixnano ps es para convertir eso a un numero gigante
// y ps ya por ultimo el rand.intt sirve para generar un numero alatorio

// el len sirve me sirvio para saber cantidad de letras que tiene la palabra
// el Scanln sirve para leer una letra del usuario y la guarda en la variable letra
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

	for {
		fmt.Println("Adivina cual de todas las palabras te toco")
		fmt.Println()
		fmt.Println("pa que no te pierdas la palabra tiene", len(palabra), "letras")
		fmt.Println()
		fmt.Println("escribe una letra para que la adivines")
		fmt.Println()
		fmt.Scanln(&letra)
		fmt.Println()
		if letra == string(palabra[0]) {
			fmt.Println("porfin que le atinas a una ya era mucho")
		} else {
			intentos++
			fmt.Println("ya mero te quedan sin oportunidad, intentos:", intentos)
			fmt.Println()
		}
		if intentos == 3 {
			fmt.Println("ya ni modo lo dejaste tieso")
			fmt.Println()

			fmt.Println("la palabra es por qeu estas medio wey:", palabra)
			fmt.Println()
			return
		}
	}
}
