package main

import (
	"bufio"
	domain "desafio-banco-bexs/domain"
	csvparser "desafio-banco-bexs/services/csv"
	utils "desafio-banco-bexs/services/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Conferido argumentos
	if len(os.Args) != 2 {
		fmt.Println("(Erro!) Por favor informe o caminho para o arquivo .csv com as rotas a serem verificadas!")
		os.Exit(1)
	} else {
		matched, err := regexp.MatchString(`^.*\.csv$`, os.Args[1])
		if err != nil || !matched {
			fmt.Printf("(Erro!) Um caminho de arquivo csv é esperado com seguinte padrao 'input-routes.csv', diferente de:  %v\n", os.Args[1])
			os.Exit(1)
		}
	}

	err, data := csvparser.Read(os.Args[1])
	if err != nil {
		fmt.Printf("(Erro!)Não foi possível ler o arquivo devido ao erro: '%v'\n", err)
		os.Exit(1)
	}

	routes := domain.Routes{}
	for _, connection := range data {
		if !routes.HasConnection(domain.NewAirport(connection[0]), domain.NewAirport(connection[1])) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(domain.NewAirport(connection[0]), domain.NewAirport(connection[1]), price)
		}
	}

	fmt.Println("=========================================================================================")
	fmt.Println("Iniciando.... Banco Bexs Melhor Rota, para sair digite 'Sair' ou, pressione ctrl+C.")
	fmt.Println("=========================================================================================")

	for {
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("\nPor favor digite a rota, no padrão ORI-DES: ")
		sentence, err := buf.ReadBytes('\n')
		if utils.TrimAndUpper(string(sentence)) == "SAIR" {
			break
		} else if err != nil {
			fmt.Print(err)
		} else {
			splitRoutes := strings.Split(string(sentence), "-")
			if len(splitRoutes) == 2 {
				splitRoutes[0] = utils.TrimAndUpper(splitRoutes[0])
				splitRoutes[1] = utils.TrimAndUpper(splitRoutes[1])
				from := domain.NewAirport(splitRoutes[0])
				to := domain.NewAirport(splitRoutes[1])

				if routes.FindAirportByCode(splitRoutes[0]) == nil {
					fmt.Printf("Rota de Origem %v não encontrada. Insira uma rota existente, ou adicione uma rota para esta origem.", splitRoutes[0])
				} else if routes.FindAirportByCode(splitRoutes[1]) == nil {
					fmt.Printf("Rota de Destino %v não encontrada. Insira uma rota existente, ou adicione uma rota para este destino.", splitRoutes[1])
				} else {
					_, _, path, price := routes.BestPriceRoute(from, to, from, []string{}, 0)
					pathString := strings.Join(path, " - ")

					if len(path) == 1 && price == 0 {
						fmt.Printf("Não encontramos alternaticas para %v->%v. Por favor, tente novamente", splitRoutes[0], splitRoutes[1])
					} else {
						fmt.Printf("Uhull a melhor rota para esta viagem é: %v > $%v", pathString, price)
					}
				}

			} else {
				fmt.Printf(" Este Formato de rota  %v está errado. Por favor, tente com o formato ORI-DES (GRU-GIG) por exemplo", strings.Trim(string(sentence), "\t \n"))
			}
		}
	}
}
