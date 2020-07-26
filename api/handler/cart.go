package handler

import (
	"encoding/json"
	"github.com/rithikjain/quickscan-backend/api/middleware"
	"github.com/rithikjain/quickscan-backend/api/view"
	"github.com/rithikjain/quickscan-backend/pkg/cart"
	"github.com/rithikjain/quickscan-backend/pkg/entities"
	"net/http"
)

func createCart(svc cart.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}

		var cart entities.Cart
		if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
			view.Wrap(err, w)
			return
		}

		claims, err := middleware.ValidateAndGetClaims(r.Context(), "user")
		if err != nil {
			view.Wrap(err, w)
			return
		}
		cart.UserID = claims["id"].(string)

		c, err := svc.CreateCart(&cart)
		if err != nil {
			view.Wrap(err, w)
			return
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Cart Created",
			"cart":    c,
		})
	})
}

func changeCartName(svc cart.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}

		type Req struct {
			CartID  string `json:"cart_id"`
			NewName string `json:"new_name"`
		}
		var req Req
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			view.Wrap(err, w)
			return
		}

		c, err := svc.ChangeCartName(req.CartID, req.NewName)
		if err != nil {
			view.Wrap(err, w)
			return
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Cart Name Updated",
			"cart":    c,
		})
	})
}

func showMyCarts(svc cart.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}

		claims, err := middleware.ValidateAndGetClaims(r.Context(), "user")
		if err != nil {
			view.Wrap(err, w)
			return
		}

		carts, err := svc.GetCarts(claims["id"].(string))
		if err != nil {
			view.Wrap(err, w)
			return
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Carts Fetched",
			"carts":   carts,
		})
	})
}

func createItem(svc cart.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}

		var cartItem entities.CartItem
		if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
			view.Wrap(err, w)
			return
		}

		ci, err := svc.CreateCartItem(&cartItem)
		if err != nil {
			view.Wrap(err, w)
			return
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Item Created",
			"item":    ci,
		})
	})
}

func updateItemCount(svc cart.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}

		type Req struct {
			ItemID   string `json:"item_id"`
			NewCount int    `json:"new_count"`
		}
		var req Req
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			view.Wrap(err, w)
			return
		}

		c, err := svc.UpdateCartItemCount(req.ItemID, req.NewCount)
		if err != nil {
			view.Wrap(err, w)
			return
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Cart Item Count Updated",
			"item":    c,
		})
	})
}

func deleteItem(svc cart.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}

		type Req struct {
			ItemID string `json:"item_id"`
		}
		var req Req
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			view.Wrap(err, w)
			return
		}

		err := svc.DeleteCartItem(req.ItemID)
		if err != nil {
			view.Wrap(err, w)
			return
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Item Deleted",
		})
	})
}

func showItems(svc cart.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}

		cartID := r.URL.Query().Get("cart_id")

		items, err := svc.GetCartItems(cartID)
		if err != nil {
			view.Wrap(err, w)
			return
		}

		var total = 0
		for _, item := range *items {
			total += item.ItemPrice * item.ItemQuantity
		}

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message":     "Items Fetched",
			"items":       items,
			"total_price": total,
		})
	})
}

// Handler
func MakeCartHandler(r *http.ServeMux, svc cart.Service) {
	r.Handle("/api/cart/create", middleware.Validate(createCart(svc)))
	r.Handle("/api/cart/changename", middleware.Validate(changeCartName(svc)))
	r.Handle("/api/cart/showmycarts", middleware.Validate(showMyCarts(svc)))
	r.Handle("/api/cart/additem", middleware.Validate(createItem(svc)))
	r.Handle("/api/cart/updateitemcount", middleware.Validate(updateItemCount(svc)))
	r.Handle("/api/cart/deleteitem", middleware.Validate(deleteItem(svc)))
	r.Handle("/api/cart/showitems", showItems(svc))
}
