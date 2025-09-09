package main


import(
	"fmt"
	"slices"
	"bufio"
	"strconv"
)


func RemovedorDeExercicio(TreinoAtual *[]ExercicioTreino, scanner *bufio.Scanner) {
	
    if len(*TreinoAtual) == 0 {
        fmt.Println("\nSeu treino atual está vazio!")
        return
    }

    var exEscolhido int
    fmt.Print("\nDigite o número do exercício que deseja remover: ")
    scanner.Scan()
	input := scanner.Text()

	exEscolhido, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Insira uma opção válida (Digite o número)")
		}

    if exEscolhido < 1 || exEscolhido > len(*TreinoAtual) {
        fmt.Println("Número inválido!")
        return
    }

    indiceRemover := exEscolhido - 1
    exercicioRemovido := (*TreinoAtual)[indiceRemover].Nome
    
    *TreinoAtual = slices.Delete(*TreinoAtual, indiceRemover, indiceRemover+1)
    
    fmt.Printf("\nExercício '%s' removido com sucesso!\n", exercicioRemovido)
}


