package main


import(
	"fmt"
	"treino/storage"
	"bufio"
	"os"
)



func AdicionarExercicioBanco() error{
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
	fmt.Println("8 - Posterior Coxa")
	fmt.Println("9 - Panturrilha")
	fmt.Println("10 - Abdômen")
	fmt.Println("0 - Voltar")
	fmt.Printf("\nSelecione um grupamento muscular: ")
fmt.Scan(&grupoEscolhido)

if grupoEscolhido == 0{
	return nil
}

// LIMPAR O BUFFER após o Scan
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan() // ← Esta linha consome o \n residual


// Agora sim ler o nome do exercício
fmt.Printf("\nEscreva o exercício que deseja adicionar: ")
scanner.Scan()
nomeExercicio = scanner.Text()
	
	// Criar o struct Exercicio com o nome digitado
	novoExercicio := storage.Exercicio{Nome: nomeExercicio}
	
	switch grupoEscolhido {
	case 1:
		storage.BancoExercicios.Peito = append(storage.BancoExercicios.Peito, novoExercicio)
	case 2:
		storage.BancoExercicios.Costas = append(storage.BancoExercicios.Costas, novoExercicio)
	case 3:
		storage.BancoExercicios.Biceps = append(storage.BancoExercicios.Biceps, novoExercicio)
	case 4:
		storage.BancoExercicios.Triceps = append(storage.BancoExercicios.Triceps, novoExercicio)
	case 5:
		storage.BancoExercicios.Ombro = append(storage.BancoExercicios.Ombro, novoExercicio)
	case 6:
		storage.BancoExercicios.Antebraco = append(storage.BancoExercicios.Antebraco, novoExercicio)
	case 7:
		storage.BancoExercicios.Quadriceps = append(storage.BancoExercicios.Quadriceps, novoExercicio)
	case 8:
		storage.BancoExercicios.Posterior = append(storage.BancoExercicios.Posterior, novoExercicio)
	case 9:
		storage.BancoExercicios.Panturrilha = append(storage.BancoExercicios.Panturrilha, novoExercicio)
	case 10:
		storage.BancoExercicios.Abdomen = append(storage.BancoExercicios.Abdomen, novoExercicio)
	default:
		return fmt.Errorf("grupo inválido: %d", grupoEscolhido)
	}

	fmt.Printf("\nExercício adicionado com sucesso!\n")
	
	return storage.SalvarExercicios()
}

	
