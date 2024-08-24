package cart

import (
	"fmt"

	"github.com/teclegacy/golang-ecom/types"
)

func getCartItemsIDs(item []types.CartCheckoutItem) ([]int, error) {
	productIDs := make([]int, len(item))

	for i, item := range item {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid id for productid %d", item.ProductID)
		}
		productIDs[i] = item.ProductID
	}
	return productIDs, nil
}
