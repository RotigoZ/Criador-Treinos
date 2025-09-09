package main

import (
	"fmt"
	"treino/storage"
	"bufio"
	"strconv"
)




func MontadorDeTreino(TreinoAtual *[]ExercicioTreino, scanner *bufio.Scanner) {

	fmt.Println("\nGRUPAMENTOS")
	fmt.Println("1 - Peito")
	fmt.Println("2 - Costas")
	fmt.Println("3 - Bíceps")
	fmt.Println("4 - Tríceps")
	fmt.Println("5 - Ombro")
	fmt.Println("6 - Antebraço")
	fmt.Println("7 - Quadríceps")
	fmt.Println("8 - Posterior Coxa")
	fmt.Println("9 - Panturrilha")
	fmt.Println("10 - Abdômen")
	fmt.Println("0 - Voltar")
	fmt.Printf("\nSelecione um grupamento muscular: ")
	scanner.Scan()
	input := scanner.Text()

	grupoEscolhido, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Opção inválida.")
		return
	}

	switch grupoEscolhido {
	case 1:
		adicionarExercicio("Peito", storage.BancoExercicios.Peito, TreinoAtual, scanner)
	case 2:
		adicionarExercicio("Costas", storage.BancoExercicios.Costas, TreinoAtual, scanner)
	case 3:
		adicionarExercicio("Bíceps", storage.BancoExercicios.Biceps, TreinoAtual, scanner)
	case 4:
		adicionarExercicio("Tríceps", storage.BancoExercicios.Triceps, TreinoAtual, scanner)
	case 5:
		adicionarExercicio("Ombro", storage.BancoExercicios.Ombro, TreinoAtual, scanner)
	case 6:
		adicionarExercicio("Antebraço", storage.BancoExercicios.Antebraco, TreinoAtual, scanner)
	case 7:
		adicionarExercicio("Quadríceps", storage.BancoExercicios.Quadriceps, TreinoAtual, scanner)
	case 8:
		adicionarExercicio("Posterior", storage.BancoExercicios.Posterior, TreinoAtual, scanner)
	case 9:
		adicionarExercicio("Panturrilha", storage.BancoExercicios.Panturrilha, TreinoAtual, scanner)
	case 10:
		adicionarExercicio("Abdômen", storage.BancoExercicios.Abdomen, TreinoAtual, scanner)
	case 0:
		fmt.Println("Voltando.")
	default:
		fmt.Println("Opção inválida.")
	}
}


func adicionarExercicio(nomeGrupo string, grupo []storage.Exercicio, TreinoAtual *[]ExercicioTreino, scanner *bufio.Scanner) {
	if len(grupo) == 0 {
		fmt.Printf("\nO grupamento %s está vazio!\n", nomeGrupo)
		return
	}

	fmt.Printf("\nGrupo %s:\n", nomeGrupo)
	for i, ex := range grupo {
		fmt.Printf("%d - %s\n", i+1, ex.Nome)
	}

	fmt.Printf("Selecione um dos exercícios: ")
	scanner.Scan()
	input := scanner.Text()

	exEscolhido, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Insira uma opção válida (Digite o número)")
			return
		}
	exEscolhido--

	if exEscolhido < 0 || exEscolhido >= len(grupo) {
		fmt.Println("Opção inválida de exercício!")
		return
	}

	fmt.Print("Insira o número de séries do exercício: ")
	scanner.Scan()
	input = scanner.Text()

	nSeries, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Insira uma opção válida (Digite o número)")
			return
		}

	if nSeries < 1 {
		fmt.Println("Número inválido de séries (precisa ser >= 1)")
		return
	}

	novoExercicio := ExercicioTreino{
        Nome:   grupo[exEscolhido].Nome,
        Series: nSeries,
    }
    
    *TreinoAtual = append(*TreinoAtual, novoExercicio) 
    fmt.Printf("\nExercício '%s' com '%d' séries adicionado!\n", grupo[exEscolhido].Nome, nSeries)
}

