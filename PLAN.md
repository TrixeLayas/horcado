el plan que tuve fue simple 

primero desarrolle un sistema de lista para mis palabras como dijo por que voya estar inventado  y con base a eso como mi inicial 
desarrolle esta lista 

palabras_alo_wey := []string{"hola", "wey", "comoestas", "pepito", "boca"}
y tambien para declarar 
var letra string
	intentos := 0
	progreso := ""
  para ir iniciando teniendo de donde inciar y que hacer 
  
con base a esto como mi incial 
dije vamos a desarrollar el primer mensaje que esta es la base de 
fmt.Println("Adivina cual de todas las palabras te toco")

luego te esto dije okey como doy espacios  ycomo no me queria complicar buscando como dar espacio 
duplique  esto debajo de todos 
fmt.Println(" ")

eso lo que hace es dejar un mensaje pero vacio dejando espacio  lo hice en todos 

luego dije ya tengo la lista y el texto al iniciar 

solo falta desarrollar el como vamos a detectar que el usuario ya ingreso una palabra correcta y pase ala siguiente 

entonces usamos esto

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

este bloque lo que hace es  si letra  es tipo texto  vamos a agarrar palabra  y vamos a obtener la longitu de progreso 
luego declaramos que progreso es igual a progreso mas letra  y lo imprimimos en terminal 
esto lo que hace que yo al atinar una palabra  hace esto w entonces pregunta otra palabra y si le atino hace esto we luego otro intento wey y dice has acertado la palabra si te equivocas por eso esta el else para decirte que fallaste y sumarlo a intentos 



dije okey ya tengo las palabras entonces como vamos a a aleatorizarlos 
busque y me encontre con estas funciones 
rand.Seed(time.Now().UnixNano())
	palabra := palabras_alo_wey[rand.Intn(len(palabras_alo_wey))]


RAND.SEED
TIME.NOW
UNIXNANO

ESTAS ME SERVIERON MUCHO PARA ESTO  YA QUE LA PRIMERA ES PARA ALEATORIZAR TODO 
LA SEGUNDA PARA LA HORA EXACTA 
LA TERCERA  unixnano ps es para convertir eso a un numero  DIGAMOS  TENGO 5 PALABRAS  NTONCES LO QUE HACE AGARRA LAS 5 LAS ALEATORIADO EN NUMEROS DE 0 4 SI NO ME EQUIVOCO PARA HACER EL 1 AL 5  YASI PODER DECIR ESTA EJECUCION SERA COMOESTAS LA SIGUIENTE PUEDE TOCAR CODO Y ASI
ESAS ME LAS ENCONTRE AQUI https://www.geeksforgeeks.org/go-language/generating-random-numbers-in-golang/

Y YA POR FINALIZAR ES SIMPLEMENTE IMPRIMIR CADA UNO DE ESTOS COMO ACERTASTE O FALLASTE O GANASTE 
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



    ESTO YA ES PARA FINZALIZAR CON EL PROYECTO SIMPLENTE ES UN IF ANIDADOS PARA ASI PODER USAR PROGRESO SERA IGUAL IGUAL A PALABRA  Y SI ACERTASTE UNA LETRA  TE DIRA ACERTASTE UNA  POR QUE YA LO VALIDAMOS CON EL IF 
    EL OTRO IF ES PARA VALIDAR INTENTOS QUE SOLO DAMOS 3 CADA FALLO SE RESTARA UN INTENTO HASTA LLEGAR A 0 Y DECIR EL SIGUIENTE MENSAJE LA PALABRA ESTA COMO VEO QUE NO LE ATINASTE ES " , PALABRA ASI MOSTRANDO LA PALABRA COMPLETA 
