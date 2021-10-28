package merchant

import "github.com/google/uuid"

type MerchantContact struct {
	ID uuid.UUID 			`gorm:"type:uuid;column:Id"`
	MerchantID uuid.UUID	`gorm:"type:uuid;column:MerchantId"`
	Country string			`gorm:"column:Country"`
	City string				`gorm:"column:City"`
	Phone string			`gorm:"column:Phone"`
	Email string			`gorm:"column:Email"`
}

func (mc *MerchantContact) GetId() uuid.UUID{
	return mc.ID
}

func (mc *MerchantContact) GetMerchantId() uuid.UUID{
	return mc.MerchantID
}

func (mc *MerchantContact) GetCountry() string{
	return mc.Country
}

func (mc *MerchantContact) GetCity() string{
	return mc.City
}

func (mc *MerchantContact) GetPhone() string{
	return mc.Phone
}

func (mc *MerchantContact) GetEmail() string{
	return mc.Email
}