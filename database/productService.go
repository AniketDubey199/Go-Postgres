package database

import (
	"encoding/json"
	"net/http"

	db "github.com/AniketDubey199/PostgreSQL/db/sqlc"
)

//	type Product struct {
//		ProductName string
//		UnitPrice   int32
//		StockQty    int32
//		ProductID   pgtype.UUID
//	}

type ProductHandler struct {
	Name     *string `json:"prod_name"`
	Price    *int32  `json:"price"`
	StockQty *int32  `json:"quantity"`
}

func (h *UserHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ProductHandler
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.queries.CreateProduct(r.Context(), db.CreateProductParams{
		ProductName: *req.Name,
		UnitPrice:   *req.Price,
		StockQty:    *req.StockQty,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
