package domain

import (
	"desafio-banco-bexs/services/utils"
	"strconv"
	"strings"
	"testing"
)

type TestDataItem struct {
	from     string
	to       string
	path     string
	price    uint64
	data     [][]string
	hasError bool
}

func MockData3AirportCircularData() (Routes, []TestDataItem) {
	data := [][]string{
		{"BRC", "SCL", "5"},
		{"BRC", "ORL", "6"},
		{"SCL", "ORL", "20"},
		{"SCL", "BRC", "5"},
	}

	dataTestItems := []TestDataItem{
		{"BRC", "BRC", "Não são permitidas (consideradas) rotas circulares", 0, nil, true},
		{"SCL", "SCL", "Não são permitidas (consideradas) rotas circulares", 0, nil, true},
		{"ORL", "ORL", "Não são permitidas (consideradas) rotas circulares", 0, nil, true},
	}

	routes := Routes{}
	for _, connection := range data {
		if !routes.HasConnection(&Airport{connection[0]}, &Airport{connection[1]}) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(&Airport{connection[0]}, &Airport{connection[1]}, price)
		}
	}

	return routes, dataTestItems
}

func MockData3Airport() (Routes, []TestDataItem) {
	data := [][]string{
		{"BRC", "SCL", "5"},
		{"BRC", "ORL", "6"},
		{"SCL", "ORL", "20"},
		{"SCL", "BRC", "5"},
	}

	dataTestItems := []TestDataItem{
		{"BRC", "SCL", "BRC - SCL", 5, data, false},
		{"BRC", "ORL", "BRC - ORL", 6, data, false},
		{"SCL", "ORL", "SCL - BRC - ORL", 11, data, false},
	}

	routes := Routes{}
	for _, connection := range data {
		if !routes.HasConnection(&Airport{connection[0]}, &Airport{connection[1]}) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(&Airport{connection[0]}, &Airport{connection[1]}, price)
		}
	}

	return routes, dataTestItems
}

func MockData5Airport() (Routes, []TestDataItem) {
	data := [][]string{
		{"GRU", "BRC", "10"},
		{"BRC", "SCL", "5"},
		{"BRC", "ORL", "6"},
		{"GRU", "CDG", "75"},
		{"GRU", "SCL", "20"},
		{"GRU", "ORL", "56"},
		{"ORL", "CDG", "5"},
		{"SCL", "ORL", "20"},
	}

	dataTestItems := []TestDataItem{
		{"GRU", "BRC", "GRU - BRC", 10, data, false},
		{"GRU", "SCL", "GRU - BRC - SCL", 15, data, false},
		{"GRU", "ORL", "GRU - BRC - ORL", 16, data, false},
		{"GRU", "CDG", "GRU - BRC - ORL - CDG", 21, data, false},
		{"BRC", "SCL", "BRC - SCL", 5, data, false},
		{"BRC", "ORL", "BRC - ORL", 6, data, false},
		{"BRC", "CDG", "BRC - ORL - CDG", 11, data, false},
		{"SCL", "ORL", "SCL - ORL", 20, data, false},
		{"SCL", "CDG", "SCL - ORL - CDG", 25, data, false},
	}

	routes := Routes{}
	for _, connection := range data {
		if !routes.HasConnection(&Airport{connection[0]}, &Airport{connection[1]}) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(&Airport{connection[0]}, &Airport{connection[1]}, price)
		}
	}

	return routes, dataTestItems
}

func MockData10Airport() (Routes, []TestDataItem) {
	data := [][]string{
		{"GRU", "BRC", "10"},
		{"BRC", "SCL", "5"},
		{"GRU", "CDG", "75"},
		{"GRU", "SCL", "20"},
		{"GRU", "ORL", "56"},
		{"ORL", "CDG", "5"},
		{"SCL", "ORL", "20"},
		{"ORL", "CPH", "200"},
		{"CPH", "BLL", "79"},
		{"CPH", "FRA", "11"},
		{"CPH", "SXF", "5"},
		{"BLL", "TXL", "75"},
		{"GRU", "BLL", "955"},
	}

	dataTestItems := []TestDataItem{
		{"GRU", "BRC", "GRU - BRC", 10, data, false},
		{"GRU", "SCL", "GRU - BRC - SCL", 15, data, false},
		{"GRU", "ORL", "GRU - BRC - SCL - ORL", 35, data, false},
		{"GRU", "CDG", "GRU - BRC - SCL - ORL - CDG", 40, data, false},
		{"GRU", "CDG", "GRU - BRC - SCL - ORL - CDG", 40, data, false},
		{"GRU", "CPH", "GRU - BRC - SCL - ORL - CPH", 235, data, false},
		{"GRU", "SXF", "GRU - BRC - SCL - ORL - CPH - SXF", 240, data, false},
		{"GRU", "FRA", "GRU - BRC - SCL - ORL - CPH - FRA", 246, data, false},
		{"GRU", "BLL", "GRU - BRC - SCL - ORL - CPH - BLL", 314, data, false},
		{"GRU", "TXL", "GRU - BRC - SCL - ORL - CPH - BLL - TXL", 389, data, false},
		{"BRC", "SCL", "BRC - SCL", 5, data, false},
		{"BRC", "ORL", "BRC - SCL - ORL", 25, data, false},
		{"BRC", "CDG", "BRC - SCL - ORL - CDG", 30, data, false},
		{"BRC", "CPH", "BRC - SCL - ORL - CPH", 225, data, false},
		{"BRC", "SXF", "BRC - SCL - ORL - CPH - SXF", 230, data, false},
		{"BRC", "FRA", "BRC - SCL - ORL - CPH - FRA", 236, data, false},
		{"BRC", "BLL", "BRC - SCL - ORL - CPH - BLL", 304, data, false},
		{"BRC", "TXL", "BRC - SCL - ORL - CPH - BLL - TXL", 379, data, false},
		{"SCL", "ORL", "SCL - ORL", 20, data, false},
		{"SCL", "CDG", "SCL - ORL - CDG", 25, data, false},
		{"SCL", "CPH", "SCL - ORL - CPH", 220, data, false},
		{"SCL", "SXF", "SCL - ORL - CPH - SXF", 225, data, false},
		{"SCL", "FRA", "SCL - ORL - CPH - FRA", 231, data, false},
		{"SCL", "BLL", "SCL - ORL - CPH - BLL", 299, data, false},
		{"SCL", "TXL", "SCL - ORL - CPH - BLL - TXL", 374, data, false},
		{"ORL", "CDG", "ORL - CDG", 5, data, false},
		{"ORL", "CPH", "ORL - CPH", 200, data, false},
		{"ORL", "SXF", "ORL - CPH - SXF", 205, data, false},
		{"ORL", "FRA", "ORL - CPH - FRA", 211, data, false},
		{"ORL", "BLL", "ORL - CPH - BLL", 279, data, false},
		{"ORL", "TXL", "ORL - CPH - BLL - TXL", 354, data, false},
		{"CPH", "SXF", "CPH - SXF", 5, data, false},
		{"CPH", "FRA", "CPH - FRA", 11, data, false},
		{"CPH", "BLL", "CPH - BLL", 79, data, false},
		{"CPH", "TXL", "CPH - BLL - TXL", 154, data, false},
		{"BLL", "TXL", "BLL - TXL", 75, data, false},
	}

	routes := Routes{}
	for _, connection := range data {
		if !routes.HasConnection(&Airport{connection[0]}, &Airport{connection[1]}) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(&Airport{connection[0]}, &Airport{connection[1]}, price)
		}
	}

	return routes, dataTestItems
}

func TestNew(t *testing.T) {
	routes := Routes{}

	if len(routes.airport) == 0 && len(routes.connections) == 0 {
		t.Logf("New() PASSOU, era esperado 0 conexões e 0 aeroportos e obtivemos %v conexões e %v aeroportos.",
			routes.connections, routes.airport)
	} else {
		t.Errorf("New() FALHOU, era esperado 0 conexões e 0 aeroportos e obtivemos %v conexões e %v aeroportos.",
			routes.connections, routes.airport)
	}
}

func TestNewAirport(t *testing.T) {
	airportCode := "CGE"
	newAirport := NewAirport(airportCode)
	if newAirport.code != airportCode {
		t.Errorf("NewAirport() FALHOU, o código esperado era  '%v' mas obtivemos '%v'", airportCode, newAirport.code)
	} else {
		t.Logf("NewAirport() PASSOU, o código esperado era  '%v' e obtivemos '%v'", airportCode, newAirport.code)
	}
}

func TestGetAllConnections(t *testing.T) {
	routes5, _ := MockData5Airport()

	connections := routes5.GetAllConnections()

	if len(connections) != 8 {
		t.Errorf("GetAllConnections() FALHOU, esperado 8 connections, obtido(s) %d", len(connections))
	} else {
		t.Logf("GetAllConnections() PASSOU, esperado 8 connections, obtido(s) %d", len(connections))
	}

	routes10, _ := MockData10Airport()

	connections = routes10.GetAllConnections()

	if len(connections) != 13 {
		t.Errorf("GetAllConnections() FALHOU, esperado 13 connections, obtido(s) %d", len(connections))
	} else {
		t.Logf("GetAllConnections() PASSOU, esperado 13 connections, obtido(s) %d", len(connections))
	}
}

func TestGetAllAirport(t *testing.T) {
	routes5, _ := MockData5Airport()

	airport := routes5.GetAllAirport()

	if len(airport) != 5 {
		t.Errorf("GetAllConnections() FALHOU, esperado 5 airport, obtido(s) %d", len(airport))
	} else {
		t.Logf("GetAllConnections() PASSOU, esperado 5 aeroporto(s) obtido(s) %d", len(airport))
	}

	routes10, _ := MockData10Airport()

	airport = routes10.GetAllAirport()

	if len(airport) != 10 {
		t.Errorf("GetAllConnections() FALHOU, esperado 10 aeroporto(s) obtido(s) %d", len(airport))
	} else {
		t.Logf("GetAllConnections() PASSOU, esperado 10 aeroporto(s) obtido(s) %d", len(airport))
	}
}

func TestHasConnection(t *testing.T) {
	routes5, _ := MockData5Airport()

	hasConn5 := routes5.HasConnection(&Airport{"GRU"}, &Airport{"BRC"})

	if !hasConn5 {
		t.Errorf("HasConnection() FALHOU, esperado conexão between GRU->BRC, obtido(s) %v", hasConn5)
	} else {
		t.Logf("HasConnection() PASSOU, esperado conexão between GRU->BRC, obtido(s) %v", hasConn5)
	}

	routes10, _ := MockData10Airport()

	hasConn10 := routes10.HasConnection(&Airport{"GRU"}, &Airport{"BLL"})

	if !hasConn10 {
		t.Errorf("HasConnection() FALHOU, esperado conexão entre GRU->BLL, obtido(s) %v", hasConn10)
	} else {
		t.Logf("HasConnection() PASSOU, esperado conexão entre GRU->BLL, obtido(s) %v", hasConn10)
	}
}

func TestAirportAdded(t *testing.T) {
	routes5, _ := MockData5Airport()

	if len(routes5.airport) != 5 {
		t.Errorf("AddAirport() FALHOU, esperado 5 aeroporto(s) obtido(s) %d", len(routes5.airport))
	} else {
		t.Logf("AddAirport() PASSOU, esperado 5 aeroporto(s) obtido(s) %d", len(routes5.airport))
	}

	routes10, _ := MockData10Airport()

	if len(routes10.airport) != 10 {
		t.Errorf("AddAirport() FALHOU, esperado 10 aeroporto(s) obtido(s) %d", len(routes10.airport))
	} else {
		t.Logf("AddAirport() PASSOU, esperado 10 aeroporto(s) obtido(s) %d", len(routes10.airport))
	}
}

func TestFindAirportByCode(t *testing.T) {
	routes5, _ := MockData5Airport()

	airport := routes5.FindAirportByCode("GRU")

	if airport == nil {
		t.Errorf("FindAirportByCode() FALHOU, esperado o objeto do aeroporto GRU, obtido(s) %v", airport.code)
	} else {
		t.Logf("FindAirportByCode() PASSOU, esperado o objeto do aeroporto GRU, obtido(s) %v", airport.code)
	}

	routes10, _ := MockData10Airport()

	airport = routes10.FindAirportByCode("CPH")

	if airport == nil {
		t.Errorf("FindAirportByCode() FALHOU, esperado o objeto do aeroporto CPH, obtido(s) %v", airport.code)
	} else {
		t.Logf("FindAirportByCode() PASSOU, esperado o objeto do aeroporto CPH, obtido(s) %v", airport.code)
	}
}

func TestConnectionsAdded(t *testing.T) {
	routes5, _ := MockData5Airport()

	if len(routes5.connections) != 8 {
		t.Errorf("AddConnection() FALHOU, esperado 8 conexões obtido(s) %d", len(routes5.connections))
	} else {
		t.Logf("AddConnection() PASSOU, esperado 8 conexões obtido(s) %d", len(routes5.connections))
	}

	routes13, _ := MockData10Airport()

	if len(routes13.connections) != 13 {
		t.Errorf("AddConnection() FALHOU, esperado 13 conexões obtido(s) %d", len(routes13.connections))
	} else {
		t.Logf("AddConnection() PASSOU, esperado 13 conexões obtido(s) %d", len(routes13.connections))
	}
}

func TestConnectionsFromAirport(t *testing.T) {
	routes, _ := MockData5Airport()
	airport := &Airport{"GRU"}
	esperadoConnections := []string{"BRC", "SCL", "ORL", "CDG"}
	connections := routes.GetConnectionsFromAirport(airport)
	var convertedConnections []string

	for _, conn := range connections {
		convertedConnections = append(convertedConnections, conn.to.code)
	}

	if len(esperadoConnections) != len(convertedConnections) {
		t.Errorf("GetConnectionsFromAirport() FALHOU, esperado %d conexões obtido(s) %d", len(esperadoConnections), len(convertedConnections))
	} else if !utils.CompareStringArrays(convertedConnections, esperadoConnections) {
		t.Errorf("GetConnectionsFromAirport() FALHOU, esperado %v, obtido(s) %v", esperadoConnections, convertedConnections)
	} else {
		t.Logf("GetConnectionsFromAirport() PASSOU, esperado %v, obtido(s) %v", esperadoConnections, convertedConnections)
	}
}

func TestBestPriceRoute(t *testing.T) {
	routes, dataTestItems := MockData3Airport()

	for _, item := range dataTestItems {

		_, hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, &Airport{item.from}, []string{}, 0)

		if len(routes.airport) != 3 {
			t.Errorf("BestPriceRoute() [%v->%v] FALHOU, esperado 3 aeroporto(s) obtido(s) %d", item.from, item.to, len(routes.airport))
		} else if item.price != price || strings.Join(path, " - ") != item.path || hasLastDestination != true {
			t.Errorf("BestPriceRoute() [%v->%v] FALHOU, o caminho esperado era  '%v' mas obtivemos '%v', o preço esperado era %d mas obtivemos %d, esperado último caminho precisa ser True mas obtivemos %v", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price, hasLastDestination)
		} else {
			t.Logf("BestPriceRoute() [%v->%v] PASSOU, o caminho esperado era  '%v' e foi(foram) obtido(s) '%v', o preço esperado era %d e foi(foram) obtido(s) %d", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price)
		}
	}

	routes, dataTestItems = MockData3AirportCircularData()

	for _, item := range dataTestItems {

		err, hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, &Airport{item.from}, []string{}, 0)

		if err != nil {
			t.Logf("BestPriceRoute() [%v->%v] PASSOU, esperado erro '%v', obtido(s) error '%v', caminho '%v', possui último caminho '%v', price '%d'", item.from, item.to, item.path, err, path, hasLastDestination, price)
		} else {
			t.Errorf("BestPriceRoute() [%v->%v] FALHOU, esperado erro '%v', obtido(s) no error, caminho '%v', possui último caminho '%v', price '%d'", item.from, item.to, item.path, path, hasLastDestination, price)
		}
	}

	routes, dataTestItems = MockData5Airport()

	for _, item := range dataTestItems {

		_, hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, &Airport{item.from}, []string{}, 0)

		if len(routes.airport) != 5 {
			t.Errorf("BestPriceRoute() [%v->%v] FALHOU, esperado 5 aeroporto(s) obtido(s) %d", item.from, item.to, len(routes.airport))
		} else if item.price != price || strings.Join(path, " - ") != item.path || hasLastDestination != true {
			t.Errorf("BestPriceRoute() [%v->%v] FALHOU, o caminho esperado era  '%v' mas obtivemos '%v', o preço esperado era %d mas obtivemos %d, esperado último caminho precisa ser True mas obtivemos %v", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price, hasLastDestination)
		} else {
			t.Logf("BestPriceRoute() [%v->%v] PASSOU, o caminho esperado era  '%v' e foi(foram) obtido(s) '%v', o preço esperado era %d e foi(foram) obtido(s) %d", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price)
		}
	}

	routes, dataTestItems = MockData10Airport()

	for _, item := range dataTestItems {

		_, hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, &Airport{item.from}, []string{}, 0)

		if len(routes.airport) != 10 {
			t.Errorf("BestPriceRoute() [%v->%v] FALHOU, esperado 10 aeroporto(s) obtido(s) %d", item.from, item.to, len(routes.airport))
		} else if item.price != price || strings.Join(path, " - ") != item.path || hasLastDestination != true {
			t.Errorf("BestPriceRoute() [%v->%v] FALHOU, o caminho esperado era  '%v' mas obtivemos '%v', o preço esperado era %d mas obtivemos %d, esperado último caminho precisa ser True mas obtivemos %v", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price, hasLastDestination)
		} else {
			t.Logf("BestPriceRoute() [%v->%v] PASSOU, o caminho esperado era  '%v' e foi(foram) obtido(s) '%v', o preço esperado era %d e foi(foram) obtido(s) %d", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price)
		}
	}
}
