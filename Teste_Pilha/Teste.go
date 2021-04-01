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
	previous  *Node
	next      *Node
}

type Stack struct {
	Top  *Node
	Base *Node
	Size int
}

var MyStackList Stack

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func OpenFile() []Node {
	var VetTeste []Node

	FileTeste, err := os.Open("Teste.json")
	Check(err)

	defer FileTeste.Close()

	byteValueTeste, err := ioutil.ReadAll(FileTeste)
	Check(err)

	json.Unmarshal(byteValueTeste, &VetTeste)

	return VetTeste
}

func Cast_in_Code(VetTeste []Node) {

	for i := 0; i < len(VetTeste); i++ {
		node := &Node{}

		node.Nome = VetTeste[i].Nome
		node.Sobrenome = VetTeste[i].Sobrenome

		StackUp(node)
	}
}

func Cast_in_File() {
	var VetTeste []Node
	pointer := MyStackList.Base

	for Size := 0; Size < MyStackList.Size; Size++ {
		VetTeste = append(VetTeste, *pointer)
		pointer = pointer.next
	}

	dataTeste, err := json.MarshalIndent(VetTeste, "", "  ")
	Check(err)

	err = ioutil.WriteFile("Teste.json", dataTeste, 0600)
}

func StackUp(Data *Node) {

	if MyStackList.Size == 0 {
		MyStackList.Base = Data
		MyStackList.Top = Data
	} else {
		pointer := MyStackList.Base

		for pointer.next != nil {
			pointer = pointer.next
		}

		pointer.next = Data
		Data.previous = MyStackList.Top
		MyStackList.Top = Data
	}

	MyStackList.Size++
}

func PrintMyStack() {
	Print := MyStackList.Top
	Size := 0

	for Size < MyStackList.Size {
		fmt.Println(*Print)

		Print = Print.previous
		Size++
	}
}

func Delete() {
	if MyStackList.Top == nil {
		fmt.Println("A pilha está vazia!!")
	} else {
		MyStackList.Top = MyStackList.Top.previous
		MyStackList.Size--
	}
}

func Insert() {
	Lancar := 0

	for Lancar != 2 {
		node := &Node{}
		fmt.Println("Nome:")
		fmt.Scanln(&node.Nome)
		fmt.Println("Sobrenome: ")
		fmt.Scanln(&node.Sobrenome)

		StackUp(node)

		fmt.Println("Deseja continuar lançando? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)
	}
}

func main() {
	VetTeste := OpenFile()
	Cast_in_Code(VetTeste)

	Operar := 0
	Change := 0

	for Operar != 2 {
		fmt.Println("O que deseja fazer? ", "\n", "1-Lançar dados", "\n", "2-Ver valores da pilha", "\n", "3-Deletar elemento")
		fmt.Scanln(&Change)

		switch {
		case Change == 1:
			Insert()
		case Change == 2:
			PrintMyStack()
		case Change == 3:
			Delete()
		}

		fmt.Println("Deseja realizar outra operação? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Operar)
	}

	Cast_in_File()
}
