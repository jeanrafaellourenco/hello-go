package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var data = time.Now()

func banner() {
	fmt.Println("************************************")
	fmt.Printf("\tGO MENU %d/%d/%d\n", data.Day(), data.Month(), data.Year())
	fmt.Println("************************************")
}

var menuOptions = map[int]string{
	1: "Opção 1",
	2: "Opção 2",
	3: "Opção 3",
	0: "Sair",
}

func printMenu() {
	clearScreen()
	banner()
	for key, value := range menuOptions {
		fmt.Printf("\t%d - %s\n", key, value)
	}
	fmt.Println("************************************")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func opcao1() {
	clearScreen()
	fmt.Println("Selecionado a opção 1")
	fmt.Println("\n\t[ Pressione Enter para continuar ]")
	fmt.Scanln()
}

func opcao2() {
	clearScreen()
	fmt.Println("Selecionado a opção 2")
	fmt.Println("\n\t[ Pressione Enter para continuar ]")
	fmt.Scanln()
}

func opcao3() {
	clearScreen()
	fmt.Println("Selecionado a opção 3")
	fmt.Println("\n\t[ Pressione Enter para continuar ]")
	fmt.Scanln()
}

func main() {
	for {
		printMenu()
		var option int
		fmt.Print("Escolha uma opção: ")
		_, err := fmt.Scan(&option)
		if err != nil {
			clearScreen()
			continue
		}

		switch option {
		case 1:
			opcao1()
		case 2:
			opcao2()
		case 3:
			opcao3()
		case 0:
			clearScreen()
			fmt.Println("Até breve!")
			fmt.Println("\n\t[ Pressione Enter para sair ]")
			fmt.Scanln()
			clearScreen()
			os.Exit(0)
		default:
			clearScreen()
			fmt.Println("Opção inválida!")
			fmt.Println("\n\t[ Pressione Enter para continuar ]")
			fmt.Scanln()
		}
	}
}
