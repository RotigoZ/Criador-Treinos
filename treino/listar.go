package main


import(
	"fmt"
)

func ListarTreino(TreinoAtual []ExercicioTreino){
	i := 1
	var nomes []ExercicioTreino
	
	for _, nome := range TreinoAtual{
        nomes = append(nomes, nome)
        i++
    }

	if len(nomes) == 0 {
        fmt.Println("\nO treino atual está vazio!")
        return
    }

	indice := 1
	fmt.Println("\nTREINO ATUAL")
	for _ , v := range(TreinoAtual){
		fmt.Printf("\nEx %d - %s, %d séries.\n", indice, v.Nome, v.Series)
		indice++
	}
}