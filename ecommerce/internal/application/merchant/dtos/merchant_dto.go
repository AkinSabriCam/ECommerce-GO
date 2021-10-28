package dtos

import (
	"ecommerce/internal/domain/merchant"
	"github.com/google/uuid"
)

type MerchantDto struct {
	Id uuid.UUID
	BusinessName string
	Contact merchant.MerchantContact
}

func MapToMerchantDto (merchant merchant.Merchant) MerchantDto {
	return MerchantDto{
		Id: merchant.GetId(),
		BusinessName: merchant.GetBusinessName(),
		Contact: merchant.GetMerchantContact(),
	}
}

func MapToMerchantSliceDto (merchants []merchant.Merchant) []MerchantDto {
	var merchantList []MerchantDto
	for _,merchant := range merchants {
		merchantList = append(merchantList, MapToMerchantDto(merchant))
	}

	return merchantList
}