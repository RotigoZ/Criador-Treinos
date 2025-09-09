package main

import (
	"fmt"
	"os"
	"encoding/json"
)




func Salvar(TreinoAtual []ExercicioTreino){
	nTreino := 1
	i := 1

	var nomes []ExercicioTreino
	
	for _, nome := range TreinoAtual{
        nomes = append(nomes, nome)
        i++
    }

	if len(nomes) == 0 {
        fmt.Println("\nO treino atual está vazio!.")
        return
    }

	fmt.Println("\nSalvando o treino atual!")

	jsonBytes, err := json.MarshalIndent(TreinoAtual, "", "  ")
	if err != nil{
		fmt.Println(err)
    	return
	}
	jsonString := string(jsonBytes)
	
	arquivo, err := os.Create("C:/Users/rod/Desktop/projetos/criadorDeTreino/treinosProntos/treino.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer arquivo.Close()

	nTreino++

	_, err = arquivo.Write([]byte(jsonString))
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("\nArquivo salvo com sucesso!")
}