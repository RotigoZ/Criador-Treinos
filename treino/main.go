package main


import(
	"fmt"
	"treino/storage"
	"log"
	"bufio"
	"os"
	"strconv"
)

type ExercicioTreino struct {
    Nome   string `json:"nome"`
    Series int    `json:"series"`
}


func main(){
	var TreinoAtual []ExercicioTreino 


	err := storage.CarregarExercicios()
    if err != nil {
        log.Fatalf("Erro ao carregar banco de exercícios: %v", err)
    }

	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Printf("\nBem vindo ao montador de treinos do Rodrigo!")
	
	for{
	fmt.Printf(`
1-Adicionar exercício no treino.
2-Remover exercício do treino.
3-Adicionar exercício (Banco).
4-Remover exercício (Banco)
5-Listar treino atual.
6-Salvar treino atual (TXT).
0-Finalizar programa.
`)
		fmt.Printf("\nInsira uma das opções (Digite o número): ")
		scanner.Scan()
		input := scanner.Text()


		if input == "" {
            continue
        }
		
		r, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Insira uma opção válida (Digite o número)")
			continue
		}

		switch r {
		case 1:
			MontadorDeTreino(&TreinoAtual, scanner)
		case 2:
			ListarTreino(TreinoAtual)
			RemovedorDeExercicio(&TreinoAtual, scanner)
		case 3:
			err := AdicionarExercicioBanco()
    		if err != nil {
        	fmt.Printf("Erro: %v\n", err)
			}
		case 4:
			 err := RemoverExercicioBanco(scanner)
    		if err != nil {
        	fmt.Printf("Erro: %v\n", err)
			}
		case 5:
			ListarTreino(TreinoAtual)
		case 6:
			Salvar(TreinoAtual)
		case 0:
			fmt.Printf("Programa encerrado.")
			return
		default:
			fmt.Printf("Insira uma opção válida (Digite o número)")
		}
		
	} 

}