package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
// Usamos un if con la funcion len(os.args) para poder introducir el dato sin necesidad de entrar y usar fmt.Scanln
// Definimos que la funcion debe ser menor que 2 para que el programa cuando no lea ningnun parametro imprima "ingresa un nombre"
	if len(os.Args) < 2 {
		fmt.Println("Ingresa un nombre ")

		return
	}
	name := os.Args[1]
	//TTS : strings package 
	// Usamos strings.ToLower del paquete strings para convertir todos los datos ingresados a miniscula todos y no tener problemas cuando,
	//el usuario ingrese alguna letra mayuscula
	name = strings.ToLower(name)
	//TTS : map literal 
	// Creamos un maps para poder darle valor numerico a las letras ingresadas 
	valornum := make(map[string]int)
	valornum["a"] = 1
	valornum["b"] = 2
	valornum["c"] = 3
	valornum["d"] = 4
	valornum["e"] = 5
	valornum["f"] = 6
	valornum["g"] = 7
	valornum["h"] = 8
	valornum["i"] = 9
	valornum["j"] = 10
	valornum["k"] = 11
	valornum["l"] = 12
	valornum["m"] = 13
	valornum["n"] = 14
	valornum["o"] = 15
	valornum["p"] = 16
	valornum["q"] = 17
	valornum["r"] = 18
	valornum["s"] = 19
	valornum["t"] = 20
	valornum["u"] = 21 
	valornum["v"] = 22
	valornum["w"] = 23
	valornum["x"] = 24
	valornum["y"] = 25
	valornum["z"] = 26

	// Inicializamos la variable sum en 0 
    sum := 0
	//usamos un ciclo for range para que haga la iteracion y sume los valores de cada letra de nuestra variable valor num
	for _ , l := range name{
		sum += valornum[string(l)]
	}
	//la formula para hallar nuestro primer numero
	div := sum / 3  
	//la formula para hallar nuestro segundo numero que es el primer digito del resultado de div
	seg := fmt.Sprintf("a%d\n", div)
	fmt.Println(string(seg[1]))
	val, err := strconv.Atoi(string(seg[1]))
	if err != nil {
		panic(err)
	}
	//tercera formula para hallar nuestro tercer bloque 
	ter := div % 10*3

	fmt.Printf("%X%X%X \n",div,val*3,ter)
}
// things to study : for range, String package functions, map literals, division entre enteros