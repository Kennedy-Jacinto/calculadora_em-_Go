package main

import (
	"fmt"
)

var num1 float32
var num2 float32

func inputDate() {
	fmt.Println("insira o Primeiro Número: ")
	fmt.Scanf("%f", &num1)
	fmt.Println("insira o Segundo Número: ")
	fmt.Scanf("%f", &num2)
}

func menu() {
	fmt.Println("********************CALCULADORA********************")
	fmt.Println("1 - CALCULADORA SIMPLES ")
	fmt.Println("2 - CALCULADORA CIENTIFICA ")
	fmt.Println("3 - SAIR ")
	fmt.Println("Escolha a operação: ")
}

func menu2() {
	fmt.Println("******************** CALCULADORA CIENTIFICA ********************")
	fmt.Println("1 - potência ")
	fmt.Println("2 - raiz quadrada ")
	fmt.Println("3 - sair ")
	fmt.Println("Escolha a operação: ")
}

func menu1() {
	fmt.Println("******************** CALCULADORA SIMPLES ********************")
	fmt.Println("1 - somar ")
	fmt.Println("2 - subtrair ")
	fmt.Println("3 - dividir ")
	fmt.Println("4 - sair ")
	fmt.Println("Escolha a operação: ")
}

func somar(num1 float32, num2 float32) float32 {
	return num1 + num2
}

func subtrair(num1 float32, num2 float32) float32 {
	return num1 - num2
}

func dividir(num1 float32, num2 float32) float32 {
	return num1 / num2
}

func calcPotencia() float32 {
	var i float32 = 1
	var res float32 = num1
	for i < num2 {
		res *= res
		i++
	}
	return res
}

func calcRaizQuadrada() float32 {
	res := (num1 / 1.0 / 2.0)
	return res
}

func optPotencia() {
	fmt.Print("deseja calcular a potência de :")
	fmt.Scanf("%f", &num1)
	fmt.Print("\nelevado á :")
	fmt.Scanf("%f", &num2)
}

func optRaizQuadrada() {
	fmt.Print("Calcular a raiz Quadrada de :")
	fmt.Scanf("%f", &num1)
}

func main() {

	var option int = 0
	for option != 3 {
		var operation int = 0
		menu()
		fmt.Scanf("%d", &option)
		if option == 1 {
			menu1()
			fmt.Scanf("%d", &operation)
			switch operation {
			case 1:
				inputDate()
				fmt.Printf("o Resultado da soma de %f e %f é : %2.f \n ", num1, num2, somar(num1, num2))
				break
			case 2:
				inputDate()
				fmt.Printf("o Resultado da subtração de %f e %f é : %2.f \n", num1, num2, subtrair(num1, num2))
				break
			case 3:
				inputDate()
				fmt.Printf("o Resultado da divisão de %f e %f é : %2.f \n", num1, num2, dividir(num1, num2))
				break
			case 4:
				fmt.Println("Escolheu sair bye!!!")
				break
			default:
				fmt.Println("opção invalida !!! Tente novamente")
				break
			}

		} else if option == 2 {
			menu2()
			fmt.Scanf("%d", &operation)
			switch operation {
			case 1:
				optPotencia()
				fmt.Printf("\nA potência de %f elevado a %f é : %2.f \n", num1, num2, calcPotencia())
				break
			case 2:
				optRaizQuadrada()
				fmt.Printf("\nA raiz quadrada de %f é : %2.f\n", num1, calcRaizQuadrada())
				break
			case 3:
				fmt.Println("Escolheu sair bye!!!")
				break
			default:
				fmt.Println("opção invalida !!! Tente novamente")
			}
		} else if option == 3 {
			fmt.Println("Escolheu sair ... Tchau !!!")
			break
		} else {
			fmt.Println("Option Invalida!")
		}

	}

}
