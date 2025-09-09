package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Exercicio struct {
	Nome string `json:"nome"`
}

type BancoExerciciosStruct struct {
	Peito       []Exercicio `json:"peito"`
	Costas      []Exercicio `json:"costas"`
	Biceps      []Exercicio `json:"biceps"`
	Triceps     []Exercicio `json:"triceps"`
	Ombro       []Exercicio `json:"ombro"`
	Antebraco   []Exercicio `json:"antebraco"`
	Quadriceps  []Exercicio `json:"quadriceps"`
	Posterior   []Exercicio `json:"posteriorCoxa"`
	Panturrilha []Exercicio `json:"panturrilha"`
	Abdomen     []Exercicio `json:"abdomen"`
}

var BancoExercicios BancoExerciciosStruct

func CarregarExercicios() error {
	
	caminhoArquivo := "storage/exerciciosBanco.txt"
	
	
	os.MkdirAll(filepath.Dir(caminhoArquivo), 0755)
	
	/*
	// Verificar se o arquivo existe
	if _, err := os.Stat(caminhoArquivo); os.IsNotExist(err) {
		// Se não existir, criar com dados iniciais
		return InicializarBancoPadrao(caminhoArquivo)
	}
	*/
	// Ler arquivo existente
	file, err := os.Open(caminhoArquivo)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %v", err)
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&BancoExercicios)
	if err != nil {
		return fmt.Errorf("erro ao decodificar JSON: %v", err)
	}
	
	return nil
}

// SalvarExercicios salva o banco de exercícios no arquivo JSON
func SalvarExercicios() error {
	caminhoArquivo := "storage/exerciciosBanco.txt"
	
	// Criar diretório se não existir
	os.MkdirAll(filepath.Dir(caminhoArquivo), 0755)
	
	file, err := os.Create(caminhoArquivo)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %v", err)
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(BancoExercicios)
	if err != nil {
		return fmt.Errorf("erro ao codificar JSON: %v", err)
	}
	
	return nil
}

/*
// InicializarBancoPadrao cria o arquivo com exercícios padrão
func InicializarBancoPadrao(caminhoArquivo string) error {
	BancoExercicios = BancoExerciciosStruct{
		Peito: []Exercicio{
			{"Supino Reto"},
			{"Supino Inclinado"},
			{"Supino Reto com Halteres"},
			{"Supino Inclinado com Halteres"},
			{"Crossover"},
			{"Flexão"},
		},
		Costas: []Exercicio{
			{"Puxada Alta"},
			{"Remada Curvada"},
			{"Pull-down"},
			{"Remada Unilateral"},
			{"Hiperextensão"},
		},
		Biceps: []Exercicio{
			{"Rosca Direta"},
			{"Rosca Martelo"},
			{"Rosca Concentrada"},
			{"Rosca Scott"},
		},
		Triceps: []Exercicio{
			{"Tríceps Testa"},
			{"Tríceps Corda"},
			{"Tríceps Francês"},
			{"Mergulho"},
		},
		Ombro: []Exercicio{
			{"Desenvolvimento"},
			{"Elevação Lateral"},
			{"Elevação Frontal"},
			{"Encolhimento"},
		},
		Antebraco: []Exercicio{
			{"Rosca de Punho"},
			{"Extensão de Punho"},
		},
		Quadriceps: []Exercicio{
			{"Agachamento"},
			{"Leg Press"},
			{"Cadeira Extensora"},
			{"Afundo"},
		},
		Posterior: []Exercicio{
			{"Stiff"},
			{"Mesa Flexora"},
			{"Cadeira Flexora"},
		},
		Panturrilha: []Exercicio{
			{"Panturrilha em Pé"},
			{"Panturrilha Sentado"},
			{"Panturrilha Leg Press"},
		},
		Abdomen: []Exercicio{
			{"Abdominal Crunch"},
			{"Prancha"},
			{"Elevação de Pernas"},
			{"Russian Twist"},
		},
	}
	
	return SalvarExercicios()
}
*/
