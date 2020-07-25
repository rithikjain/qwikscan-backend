package cart

import (
	"github.com/jinzhu/gorm"
	"github.com/rithikjain/quickscan-backend/pkg"
	"github.com/rithikjain/quickscan-backend/pkg/entities"
)

type Repository interface {
	CreateCart(cart *entities.Cart) (*entities.Cart, error)

	ChangeCartName(cartID string, name string) (*entities.Cart, error)

	GetCarts(userID string) (*[]entities.Cart, error)

	CreateCartItem(cartItem *entities.CartItem) (*entities.CartItem, error)

	UpdateCartItemCount(cartItemID string, newCount int) (*entities.CartItem, error)

	DeleteCartItem(cartItemID string) error

	GetCartItems(cartID string) (*[]entities.CartItem, error)
}

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateCart(cart *entities.Cart) (*entities.Cart, error) {
	result := r.DB.Create(cart)
	if result.Error != nil {
		return nil, pkg.ErrDatabase
	}
	return cart, nil
}

func (r *repo) ChangeCartName(cartID string, name string) (*entities.Cart, error) {
	cart := &entities.Cart{}
	r.DB.Where("uuid = ?", cartID).First(cart)
	if cart.CartName == "" {
		return nil, pkg.ErrNotFound
	}
	cart.CartName = name
	err := r.DB.Save(cart).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return cart, nil
}

func (r *repo) GetCarts(userID string) (*[]entities.Cart, error) {
	var carts []entities.Cart
	err := r.DB.Where("user_id = ?", userID).Find(&carts).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return &carts, nil
}

func (r *repo) CreateCartItem(cartItem *entities.CartItem) (*entities.CartItem, error) {
	result := r.DB.Create(cartItem)
	if result.Error != nil {
		return nil, pkg.ErrDatabase
	}
	return cartItem, nil
}

func (r *repo) UpdateCartItemCount(cartItemID string, newCount int) (*entities.CartItem, error) {
	cartItem := &entities.CartItem{}
	r.DB.Where("uuid = ?", cartItemID).First(cartItem)
	if cartItem.ItemName == "" {
		return nil, pkg.ErrNotFound
	}
	cartItem.ItemQuantity = newCount
	err := r.DB.Save(cartItem).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return cartItem, nil
}

func (r *repo) DeleteCartItem(cartItemID string) error {
	cartItem := &entities.CartItem{}
	r.DB.Where("uuid = ?", cartItemID).First(cartItem)
	if cartItem.ItemName == "" {
		return pkg.ErrNotFound
	}
	err := r.DB.Delete(cartItem).Error
	if err != nil {
		return pkg.ErrDatabase
	}
	return nil
}

func (r *repo) GetCartItems(cartID string) (*[]entities.CartItem, error) {
	var cartItems []entities.CartItem
	err := r.DB.Where("cart_id = ?", cartID).Find(&cartItems).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return &cartItems, nil
}
