package customers

import (
	"context"

	"github.com/slaff-bg/stockroom/ports/graph/model"
)

type CustomerRepository interface {
	ListCustomer(ctx context.Context) *model.Customer
}
