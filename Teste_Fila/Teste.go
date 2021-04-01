package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Node struct {
	Nome      string
	Sobrenome string
	next      *Node
}

type Row struct {
	head *Node
	Size int
}

var MyRow Row

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Abrir_Arquivo() []Node {
	var VetorTeste []Node

	FileTeste, err := os.Open("Teste.json")
	Check(err)

	defer FileTeste.Close()

	byteValueTeste, err := ioutil.ReadAll(FileTeste)
	Check(err)

	json.Unmarshal(byteValueTeste, &VetorTeste)

	return VetorTeste
}

func Lancar_Dado_Codigo(VetTeste []Node) {

	for i := 0; i < len(VetTeste); i++ {
		node := &Node{}

		node.Nome = VetTeste[i].Nome
		node.Sobrenome = VetTeste[i].Sobrenome

		Lancar_Dado_Row(node)
	}
}

func Lancar_Dado_Arquivo() {
	var VetorTeste []Node
	pointer := MyRow.head

	for Size := 0; Size < MyRow.Size; Size++ {
		VetorTeste = append(VetorTeste, *pointer)
		pointer = pointer.next
	}

	dataTeste, err := json.MarshalIndent(VetorTeste, "", "  ")
	Check(err)

	err = ioutil.WriteFile("Teste.json", dataTeste, 0600)
}

func Lancar_Dado_Row(Data *Node) {

	if MyRow.Size == 0 {
		MyRow.head = Data
	} else {
		pointer := MyRow.head

		for pointer.next != nil {
			pointer = pointer.next
		}

		pointer.next = Data

	}
	MyRow.Size++
}

func Deletar_Dado() {
	MyRow.head = MyRow.head.next
	MyRow.Size--
}

func Lancar_Dados_Registro() {
	Lancar := 0

	for Lancar != 2 {
		node := &Node{}
		fmt.Println("Nome: ")
		fmt.Scanln(&node.Nome)
		fmt.Println("Sobrenome: ")
		fmt.Scanln(&node.Sobrenome)

		Lancar_Dado_Row(node)

		fmt.Println("Deseja lançar mais registros? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)
	}
}

func PrintRow() {
	Print := MyRow.head

	for Size := 0; Size < MyRow.Size; Size++ {
		fmt.Println(*Print)
		Print = Print.next
	}
}

func main() {
	VetorTeste := Abrir_Arquivo()
	Lancar_Dado_Codigo(VetorTeste)

	Operar := 0
	Change := 0

	for Operar != 2 {
		fmt.Println("O que deseja fazer? ", "\n", "1-Lançar registros", "\n", "2-Deletar", "\n", "3-Ver fila")
		fmt.Scanln(&Change)

		switch {
		case Change == 1:
			Lancar_Dados_Registro()
		case Change == 2:
			Deletar_Dado()
		case Change == 3:
			PrintRow()
		}

		fmt.Println("Deseja realizar mais alguma operação?", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Operar)
	}
	Lancar_Dado_Arquivo()
}
