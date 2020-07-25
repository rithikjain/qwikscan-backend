package cart

import (
	uuid2 "github.com/nu7hatch/gouuid"
	"github.com/rithikjain/quickscan-backend/pkg/entities"
)

type Service interface {
	CreateCart(cart *entities.Cart) (*entities.Cart, error)

	ChangeCartName(cartID string, name string) (*entities.Cart, error)

	GetCarts(userID string) (*[]entities.Cart, error)

	CreateCartItem(cartItem *entities.CartItem) (*entities.CartItem, error)

	UpdateCartItemCount(cartItemID string, newCount int) (*entities.CartItem, error)

	DeleteCartItem(cartItemID string) error

	GetCartItems(cartID string) (*[]entities.CartItem, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateCart(cart *entities.Cart) (*entities.Cart, error) {
	uuid, err := uuid2.NewV4()
	if err != nil {
		return nil, err
	}
	cart.UUID = uuid.String()

	return s.repo.CreateCart(cart)
}

func (s *service) ChangeCartName(cartID string, name string) (*entities.Cart, error) {
	return s.repo.ChangeCartName(cartID, name)
}

func (s *service) GetCarts(userID string) (*[]entities.Cart, error) {
	return s.repo.GetCarts(userID)
}

func (s *service) CreateCartItem(cartItem *entities.CartItem) (*entities.CartItem, error) {
	uuid, err := uuid2.NewV4()
	if err != nil {
		return nil, err
	}
	cartItem.UUID = uuid.String()

	return s.repo.CreateCartItem(cartItem)
}

func (s *service) UpdateCartItemCount(cartItemID string, newCount int) (*entities.CartItem, error) {
	return s.repo.UpdateCartItemCount(cartItemID, newCount)
}

func (s *service) DeleteCartItem(cartItemID string) error {
	return s.repo.DeleteCartItem(cartItemID)
}

func (s *service) GetCartItems(cartID string) (*[]entities.CartItem, error) {
	return s.repo.GetCartItems(cartID)
}
