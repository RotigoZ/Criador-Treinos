package main


import(
	"fmt"
	"treino/storage"
	"bufio"
	"slices"
	"strconv"
)



func RemoverExercicioBanco(scanner *bufio.Scanner) error{

	grupoEscolhido := 0
	var nomeExercicio string

	fmt.Println("\nGRUPAMENTOS")
	fmt.Println("1 - Peito")
	fmt.Println("2 - Costas")
	fmt.Println("3 - Bíceps")
	fmt.Println("4 - Tríceps")
	fmt.Println("5 - Ombro")
	fmt.Println("6 - Antebraço")
	fmt.Println("7 - Quadríceps")
	fmt.Println("8 - Posterior de Coxa")
	fmt.Println("9 - Panturrilha")
	fmt.Println("10 - Abdômen")
	fmt.Println("0 - Voltar")
	fmt.Printf("\nSelecione um grupamento muscular: ")
	scanner.Scan()
	input := scanner.Text()
		grupoEscolhido, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print(err)
		}

		if grupoEscolhido == 0{
		return nil
		}


		fmt.Printf("\nEscreva o exercício que deseja remover: ")
		scanner.Scan()
		nomeExercicio = scanner.Text()
		exercicioRemovido := false
			
	switch grupoEscolhido {
	case 1:
		tamanhoAntes := len(storage.BancoExercicios.Peito)
		storage.BancoExercicios.Peito = slices.DeleteFunc(storage.BancoExercicios.Peito, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
	case 2:
		tamanhoAntes := len(storage.BancoExercicios.Costas)
		storage.BancoExercicios.Costas = slices.DeleteFunc(storage.BancoExercicios.Costas, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
	case 3:
		tamanhoAntes := len(storage.BancoExercicios.Biceps)
		storage.BancoExercicios.Biceps = slices.DeleteFunc(storage.BancoExercicios.Biceps, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
	case 4:
		tamanhoAntes := len(storage.BancoExercicios.Triceps)
		storage.BancoExercicios.Triceps = slices.DeleteFunc(storage.BancoExercicios.Triceps, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
	case 5:
		tamanhoAntes := len(storage.BancoExercicios.Ombro)
		storage.BancoExercicios.Ombro = slices.DeleteFunc(storage.BancoExercicios.Ombro, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
	case 6:
		tamanhoAntes := len(storage.BancoExercicios.Antebraco)
		storage.BancoExercicios.Antebraco = slices.DeleteFunc(storage.BancoExercicios.Antebraco, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
	case 7:
		tamanhoAntes := len(storage.BancoExercicios.Quadriceps)
		storage.BancoExercicios.Quadriceps = slices.DeleteFunc(storage.BancoExercicios.Quadriceps, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
	case 8:
		tamanhoAntes := len(storage.BancoExercicios.Posterior)
		storage.BancoExercicios.Posterior = slices.DeleteFunc(storage.BancoExercicios.Posterior, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
	case 9:
		tamanhoAntes := len(storage.BancoExercicios.Panturrilha)
		storage.BancoExercicios.Panturrilha = slices.DeleteFunc(storage.BancoExercicios.Panturrilha, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
	case 10:
		tamanhoAntes := len(storage.BancoExercicios.Abdomen)
		storage.BancoExercicios.Abdomen = slices.DeleteFunc(storage.BancoExercicios.Abdomen, func(ex storage.Exercicio) bool {
			return ex.Nome == nomeExercicio
		})
		exercicioRemovido = len(storage.BancoExercicios.Peito) < tamanhoAntes
		case 0:
		return nil
	default:
		return fmt.Errorf("grupo inválido: %d", grupoEscolhido)
	}
	if !exercicioRemovido {
    fmt.Printf("\nExercício '%s' não encontrado neste grupo!\n", nomeExercicio)
    return nil
}

	fmt.Printf("\nExercício %s removido com sucesso!\n", nomeExercicio)
	
	return storage.SalvarExercicios()
}

	
