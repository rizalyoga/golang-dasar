package main

import "fmt"

type RepositoryManager interface {
	AddToCart(userID string, productID string) (err error)
}

type ShoppingCart struct {
	repo RepositoryManager // Ketergantungan (dependency) pada interface
}

type FakeRepo struct {
	name string
}

func (f *FakeRepo) AddToCart(userID string, productID string) (err error) {
	fmt.Printf("✅ [FakeRepo] Sukses menambahkan produk '%s' ke keranjang user '%s'\n", productID, userID)
	return nil
}

func NewShoppingCart(repo RepositoryManager) ShoppingCart {
	return ShoppingCart{
		repo: repo, // Mengisi field 'repo' di ShoppingCart dengan implementasi yang diberikan
	}
}

func main() {
	dataStore := &FakeRepo{}
	var cart = NewShoppingCart(dataStore)
	err := cart.repo.AddToCart("12345678", "87654321")

	if err != nil {
		fmt.Printf("Error bro")
	}
}
