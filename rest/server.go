package main

import (
	domain "desafio-banco-bexs/domain"
	csvparser "desafio-banco-bexs/services/csv"
	"desafio-banco-bexs/services/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type SearchResponse struct {
	From  string
	To    string
	Path  string
	Price uint64
}

type Connection struct {
	From  string
	To    string
	Price uint64
}

type RoutesWrapper struct {
	routes domain.Routes
	data   [][]string
}

func main() {
	routesWrapper := RoutesWrapper{domain.Routes{}, nil}
	loadDataFromCSV(&routesWrapper)

	fmt.Println("[INFO] Iniciando servidor na porta 8080...")
	http.HandleFunc("/", ApiStatus)
	http.Handle("/route", &routesWrapper)
	http.ListenAndServe(":8080", nil) //Definimos a porta do servidor por padrao utilizo a 8080
}

func (routesWrapper *RoutesWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fromParams := r.URL.Query()["from"]
		toParams := r.URL.Query()["to"]

		if len(fromParams) == 1 && len(toParams) == 1 {
			fromParams[0] = utils.TrimAndUpper(fromParams[0])
			toParams[0] = utils.TrimAndUpper(toParams[0])
			fromAirport := routesWrapper.routes.FindAirportByCode(fromParams[0])
			toAirport := routesWrapper.routes.FindAirportByCode(toParams[0])
			if fromAirport == nil {
				errorMsg := "Origem '" + fromParams[0] + "' não encontrada. Tente informar uma origem existente, ou criar uma nova com esta origem."
				http.Error(w, errorMsg, http.StatusBadRequest)
				fmt.Printf("\n(Erro!)%v | %v", getTimeNowFormatted(), errorMsg)
				return
			} else if toAirport == nil {
				errorMsg := "Destino '" + toParams[0] + "' não encontrado. Tente informar um destino existente, ou criar uma nova com este destino."
				http.Error(w, errorMsg, http.StatusBadRequest)
				fmt.Printf("\n(Erro!)%v | %v",
					getTimeNowFormatted(), errorMsg)
				return
			} else {
				_, _, path, price := routesWrapper.routes.BestPriceRoute(fromAirport, toAirport, fromAirport, []string{}, 0)
				pathString := strings.Join(path, " - ")

				if len(path) == 1 && price == 0 {
					errorMsg := "Caminho não encontrado para '" + fromParams[0] + "->" + toParams[0] + "'. Tente novamente com um caminho adequado."
					http.Error(w, errorMsg, http.StatusBadRequest)
					fmt.Printf("\n(Erro!)'%v' | %v",
						getTimeNowFormatted(), errorMsg)
					return
				} else {
					js, err := json.Marshal(SearchResponse{fromParams[0], toParams[0], pathString, price})
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						fmt.Printf("\n(Erro!)%v | %v",
							getTimeNowFormatted(), err.Error())
						return
					}

					w.Header().Set("Content-Type", "application/json")
					w.Write(js)
				}
			}
		} else {
			errorMsg := "Faltaram argumentos, ou foram informados argumentos inválidos"
			http.Error(w, errorMsg, http.StatusBadRequest)
			fmt.Printf("\n(Erro!)%v | %v",
				getTimeNowFormatted(), errorMsg)
			return
		}

	case http.MethodPost:
		var conn Connection
		err := json.NewDecoder(r.Body).Decode(&conn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Printf("\n(Erro!)%v | %v",
				getTimeNowFormatted(), err.Error())
			return
		}

		from := domain.NewAirport(utils.TrimAndUpper(conn.From))
		to := domain.NewAirport(utils.TrimAndUpper(conn.To))
		if routesWrapper.routes.HasConnection(from, to) {
			errorMsg := "Conexão '" + conn.From + "->" + conn.To + "' já está cadastrada."
			fmt.Printf("\n(Erro!)%v | %v",
				getTimeNowFormatted(), errorMsg)
			http.Error(w, errorMsg, http.StatusBadRequest)
			return
		}

		routesWrapper.routes.AddConnection(from, to, conn.Price)
		routesWrapper.data = append(routesWrapper.data, []string{conn.From, conn.To, strconv.FormatUint(conn.Price, 10)})

		defaultCsvFilePath := "./data/input-routes.csv"
		err, newData := csvparser.CreateWrite(defaultCsvFilePath, routesWrapper.data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Printf("(Erro!)%v | Não foi possível gravar o arquivo csv devido ao erro: '%v'\n", getTimeNowFormatted(), err)
			return
		}

		routesWrapper.data = newData

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	default:
		errorMsg := "Método REST não autorizado."
		http.Error(w, errorMsg, http.StatusMethodNotAllowed)
		fmt.Printf("\n(Erro!)%v | %v",
			getTimeNowFormatted(), errorMsg)
		return
	}
}

func ApiStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Ola-Bem-vindo", "Por aqui tudo funcionando!")
	w.Header().Set("Banco-Bexs", "Servidor de busca da menor rota possivel")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func loadDataFromCSV(routesWrapper *RoutesWrapper) {
	defaultCsvFilePath := "./data/input-routes.csv" //Caso nao informe um csv ao iniciar o servidor, o padrao será apontado.

	err, data := csvparser.Read(defaultCsvFilePath)
	if err != nil {
		fmt.Printf("(Erro!)%v | Não foi possível ler o arquivo csv devido ao erro: '%v'\n", getTimeNowFormatted(), err)
		os.Exit(1)
	}

	routesWrapper.data = data

	for _, connection := range data {
		if !routesWrapper.routes.HasConnection(domain.NewAirport(connection[0]), domain.NewAirport(connection[1])) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routesWrapper.routes.AddConnection(domain.NewAirport(connection[0]), domain.NewAirport(connection[1]), price)
		}
	}
}

func getTimeNowFormatted() string {
	datetime := time.Now()
	return datetime.Format("31-07-1989 15:04:05")
}
