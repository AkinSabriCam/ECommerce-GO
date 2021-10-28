package merchant

import "github.com/google/uuid"

type IMerchantRepository interface {
	GetById(id uuid.UUID) Merchant
	GetAll() []Merchant

}
