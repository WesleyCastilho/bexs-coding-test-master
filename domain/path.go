package domain

import (
	"errors"
	"sort"
)

//Atenção os dados abaixo serão utilizados para montar o grafo de Dijkstra.
// Airport como vértice
type Airport struct {
	code string
}

// Connection como aresta
type Connection struct {
	from  *Airport
	to    *Airport
	price uint64
}

type PathConnectionPrice struct {
	airportCodePath    []string
	totalPrice         uint64
	hasLastDestination bool
}

// Rotas como caminho do grafo  (s a t)
type Routes struct {
	connections []*Connection
	airport     []*Airport
}

const Infinity = uint64(^uint64(0) >> 1)

func New() *Routes {
	return &Routes{
		airport:     []*Airport{},
		connections: []*Connection{},
	}
}

func NewAirport(code string) *Airport {
	return &Airport{code}
}

func (r *Routes) AddConnection(from *Airport, to *Airport, price uint64) {
	connection := &Connection{
		from:  from,
		to:    to,
		price: price,
	}

	r.connections = append(r.connections, connection)
	r.AddAirport(from)
	r.AddAirport(to)
}

func (r *Routes) GetAllConnections() []*Connection {
	return r.connections
}

func (r *Routes) GetAllAirport() []*Airport {
	return r.airport
}

func (r *Routes) HasConnection(from *Airport, to *Airport) bool {
	for _, c := range r.connections {
		if c.from.code == from.code && c.to.code == to.code {
			return true
		}
	}
	return false
}

func (r *Routes) AddAirport(airport *Airport) {
	var isAirportPresent bool
	for _, a := range r.airport {
		if a.code == airport.code {
			isAirportPresent = true
		}
	}
	if !isAirportPresent {
		r.airport = append(r.airport, airport)
	}
}

func (r *Routes) FindAirportByCode(code string) *Airport {
	var airport *Airport
	airport = nil
	for _, a := range r.airport {
		if a.code == code {
			airport = a
		}
	}
	return airport
}

func (r *Routes) GetConnectionsFromAirport(airport *Airport) (connections []*Connection) {
	for _, connection := range r.connections {
		if connection.from.code == airport.code {
			connections = append(connections, connection)
		}
	}
	return connections
}

func (r *Routes) BestPriceRoute(from *Airport, to *Airport, origin *Airport, accumulatedPath []string, accumulatedPrice uint64) (error, bool, []string, uint64) {
	if from.code == to.code && from.code == origin.code && to.code == origin.code {
		return errors.New("(Atenção) Rotas circulares não são permitidas."), false, nil, 0
	}

	pathConnectionPrices := []*PathConnectionPrice{}

	connections := r.GetConnectionsFromAirport(from)

	if from.code == to.code && from.code != origin.code {
		return nil, true, append(accumulatedPath, from.code), accumulatedPrice
	} else if len(connections) == 0 {
		return nil, false, append(accumulatedPath, from.code), accumulatedPrice
	} else {
		for _, conn := range connections {
			if origin.code != to.code && conn.to.code == origin.code {
				continue
			}
			_, hasLastDestination, recursionPath, recursionPrice := r.BestPriceRoute(conn.to, to, origin, append(accumulatedPath, from.code), (accumulatedPrice + conn.price))
			if hasLastDestination {
				pathConnectionPrices = append(pathConnectionPrices, &PathConnectionPrice{recursionPath, recursionPrice, hasLastDestination})
			}
		}

		sort.Slice(pathConnectionPrices, func(i, j int) bool {
			return pathConnectionPrices[i].totalPrice < pathConnectionPrices[j].totalPrice
		})

		if len(pathConnectionPrices) == 0 {
			return nil, false, append(accumulatedPath, from.code), accumulatedPrice
		} else {
			return nil, pathConnectionPrices[0].hasLastDestination, pathConnectionPrices[0].airportCodePath, pathConnectionPrices[0].totalPrice
		}
	}
}
