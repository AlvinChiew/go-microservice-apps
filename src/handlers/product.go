package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/AlvinChiew/go-microservice-apps/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// curl localhost:9090 | jq
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// curl -d '{}' localhost:9090
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return

	}

	// curl localhost:9090/1 -XPUT  -d '{"name": "tea", "description":"A nice cup of tea!"}'
	if r.Method == http.MethodPut {
		p.updateProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	// d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	// rw.Write(d)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) getIdFromUri(rw http.ResponseWriter, r *http.Request) int {
	p.l.Println("PUT", r.URL.Path)
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)

	if len(g) != 1 {
		http.Error(rw, "Invalid URI with more than an ID", http.StatusBadRequest)
		return -1
	}

	if len(g[0]) != 2 {
		http.Error(rw, "Invalid URI with more than a capture group", http.StatusBadRequest)
		return -1
	}

	idString := g[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(rw, "Invalid URI. Unable to convert ID to number", http.StatusBadRequest)
	}

	return id
}

func (p *Products) updateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	id := p.getIdFromUri(rw, r)

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
