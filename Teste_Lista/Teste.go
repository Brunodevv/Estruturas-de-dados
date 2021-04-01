package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var MyList LinkedList

//Struct que irá conter os dados do nó
type Node struct {
	Name      string
	Sobrename string
	Idade     int
	Next      *Node `json:"next,omitempty"`
}

//Struct que irá conter o nó cabeça (automáticamente ela irá conter os outros nó)
type LinkedList struct {
	Head *Node
	Size int
}

//Função para tratar erros
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Função para abrir o arquivo e atribuír os valores dele em um vetor do tipo "Node"
func AbrirArquivo() []Node {
	//Declaração do vetor que irá conter os dados do arquivo
	var VetorTeste []Node

	//Abri o arquivo "Teste.json" para leitura
	FileTeste, err := os.Open("Teste.json")
	//Tratamento de erro
	Check(err)

	//Fechará o arquivo após o mesmo ser utilizado
	defer FileTeste.Close()

	//Atribuí os bytes do arquivo a variável "byteValueTeste"
	byteValueTeste, err := ioutil.ReadAll(FileTeste)
	//Tratamento de erro
	Check(err)

	//Converte os bytes em dados e atribuí ao vetor "VetorTeste"
	json.Unmarshal(byteValueTeste, &VetorTeste)

	//Retorna a variável "VetorTeste"
	return VetorTeste
}

//Função para lançar os dados do arquivo para serem manipulados no código
func Lancar_Dado_Codigo(VetorTeste []Node) {
	//Enquanto i for menor que o comprimento do vetor ele irá fazer as seguintes operações...
	for i := 0; i < len(VetorTeste); i++ {
		//node irá conter os campos do struct "Node"
		node := &Node{}

		//Atribuíndo valores aos campos de node
		node.Name = VetorTeste[i].Name
		node.Sobrename = VetorTeste[i].Sobrename
		node.Idade = VetorTeste[i].Idade

		//Passa variável node como parâmetro para a função "Lancar_List" que irá encadear essa lista
		//tornando ela manipulavél de forma encadeada no código
		Lancar_List(node)
	}
}

//Função para lançar os dados do código no arquivo
func Lancar_Dado_Arquivo() {

	//Declaração de uma variável vetor do tipo "Node"
	var VetorTeste []Node
	//pointer irá receber o local da memória da variável "Head" que contém o primeiro item da lista encadeada
	pointer := MyList.Head

	Size := 0

	//Enquanto "Size" for menor que "MyList.Size" que indica o tamanho da lista as seguintes operações serão feitas...
	for Size < MyList.Size {

		//Se Size for igual a zero, o que significa que é o primeiro item da lista
		if Size == 0 {
			//pointer recebe o endereço de memória do valor apontado por "myList.Head"
			pointer = MyList.Head.Next
			//MyList.Head.Next recebe um valor nil

			/* O arquivo já é uma lista encadeada então não precisamos dizer a ele qual é o próximo elemento
			pois ele já sabe, mas quando precisamos manipular esses dados do arquivo temos que passar eles para um vetor
			que não tem tamanho definido, o que faz com que cada dado seja armazenado em locais aleatórios da memória
			então para mantermos o controle dos dados fazemos com que o um registro aponte para o local da memória do próximo registro */

			//Por isso definimos que "MyList.Head.Next" seja um valor nil, para que ele não seja arquivado
			MyList.Head.Next = nil

			//VetorTeste irá acrescentar os dados de "MyList.Head"
			VetorTeste = append(VetorTeste, *MyList.Head)

			//Se Size não for igual a zero então as seguintes operações serão realizadas...
		} else {
			//auxiliar receberá o local da memória apontado pela variável pointer
			auxiliar := pointer
			//pointer irá apontar para o local de memória indicada pelo antigo valor de pointer
			pointer = pointer.Next

			//auxiliar.Next receberá um valor nil para não ser arquivado de forma ecadeada
			auxiliar.Next = nil

			//Irá acrescentar ao VetorTeste os dados da variável auxiliar
			VetorTeste = append(VetorTeste, *auxiliar)
		}

		Size++
	}

	DataTeste, err := json.MarshalIndent(VetorTeste, "", "  ")
	Check(err)

	err = ioutil.WriteFile("Teste.json", DataTeste, 0600)
}

//Função para atribuír dados ao registro
func Lancar_Dados() {
	Lancar := 0

	for Lancar != 2 {
		node := &Node{}
		fmt.Println("Nome: ")
		fmt.Scanln(&node.Name)
		fmt.Println("Sobrenome")
		fmt.Scanln(&node.Sobrename)
		fmt.Println("Idade:")
		fmt.Scanln(&node.Idade)

		//Irá passar como parâmetro para a função "Lancar_List" a variável node
		//essa função irá encadear os registros
		Lancar_List(node)

		fmt.Println("Deseja lançar mais um registro?", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)
	}
}

//Função para lançar na lista encadeada os dados inseridos na função acima
func Lancar_List(Valor *Node) {

	//Se "MyList.Size" for igual a zero, o que significa que este vai ser o primeiro elemento da lista
	if MyList.Size == 0 {
		/*Então a variável "Head" em português "cabeça" receberá o local da memória da variável "Valor"
		que contém os dados do registro */
		MyList.Head = Valor

		//Se não
	} else {
		//Então pointer irá apontar pa o local da memória em que "Head" está
		pointer := MyList.Head

		/*Enquanto pointer next for diferente de nil ele vai avançar no encadeamento
		quando o pointer.Next for nil o mesmo irá receber o local da memória  da variável "Valor"
		que contém os dados do registro  	 */
		for pointer.Next != nil {
			pointer = pointer.Next
		}

		pointer.Next = Valor
	}

	//Aumenta o tamanho da lista
	MyList.Size++
}

//Função para imprimir os dados da lista
func PrintList() {
	/*Como nossa struct "LinkedList" não é um vetor e sim é constituída de um ponteiro que aponta para um local da memória
	então não podemos dar apenas um fmt.Println(LinkedList)*/

	//Se torna necessário percorrer todo o encadeamento e imprimir os valores de cada local da memória percorrido

	//Print recebe "MyList.Head" que é o primeiro elemento da lista
	Print := MyList.Head
	//Size é uma variável auxiliar que irá ajudar a percorrer o encadeamento
	Size := 0

	//Enquanto Size for menor que o tamanho da lista encadeada faremos o seguinte...
	for Size < MyList.Size {

		//Imprimiremos o valor do ponteiro cuja variável "Print" está apontando
		fmt.Println(*Print)

		//Cada vez que a variável imprimir o valor que nela contém ela receberá um novo valor
		//Que será o próximo elemento da lista
		Print = Print.Next

		Size++
	}
}

//Função para deletar um deteminado registro da lista
func Deletar() {
	/*O metódo de consulta a lista escolhido foi o nome, mas vale lembrar que é possível usar qualquer campo como referência de busca
	podemos até mesmo usar mais do que um campo como referência para consulta*/

	Name := ""
	//Variável pointer irá receber o local de memória do primeiro elemento da lista
	pointer := MyList.Head

	fmt.Println("Informe o nome que deseja excluir: ")
	fmt.Scanln(&Name)

	//Se o nome procurado se encontrar na primeira posição então o seguinte será feito...
	if MyList.Head.Name == Name {
		//A cabeça da lista receberá um novo local de memória, que é na verdade o local que para o qual a mesma está apontando
		MyList.Head = MyList.Head.Next
		//É necessário reduzir o tamanho da lista ao final de uma operação de exclusão
		MyList.Size--

		fmt.Println("Registro deletado com sucesso!!")

		//Se o nome não estiver no primeiro local as seguintes operações serão feitas
	} else {

		/*Enquanto pointer.Next não for igual ao nome e nem o valor para o qual ele está apontando for igual a nil
		a variável pointer receberá um novo endereço de memória*/
		for pointer.Next.Name != Name && pointer.Next.Next != nil {
			pointer = pointer.Next
		}

		//Caso pointe.Next.Name for igual ao nome iremos fazer o seguinte ...
		if pointer.Next.Name == Name {
			//auxiliar irá receber o local da memória a qual a pointer está apontando
			auxiliar := pointer.Next
			//pointer então passará a apontar para o local da memória que o seu descendente apontava
			pointer.Next = auxiliar.Next

			//Com isso o local da memória deletado não será mais referênciado, logo não fará mais parte da lista

			MyList.Size--

			fmt.Println("Registro deletado com sucesso!!")

			//Caso ponteiro.Next.Name não seja igual a Name, significa que chegamos no final da lista e não achamos o nome
			//ou seja o nome não existe nesta lista
		} else {
			fmt.Println("O registro não existe nessa lista!!")
		}
	}
}

//Função main, que chamará as respectivas funções acima
func main() {

	VetorTeste := AbrirArquivo()
	Lancar_Dado_Codigo(VetorTeste)

	Operar := 0
	Change := 0

	for Operar != 2 {
		fmt.Println("O que deseja fazer? ", "\n", "1-Lançar dados", "\n", "2-Ler dados inseridos", "\n", "3-Excluir")
		fmt.Scanln(&Change)

		switch {
		case Change == 1:
			Lancar_Dados()
		case Change == 2:
			PrintList()
		case Change == 3:
			Deletar()
		}

		fmt.Println("Deseja realizar mais alguma operação? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Operar)

	}

	Lancar_Dado_Arquivo()
}
