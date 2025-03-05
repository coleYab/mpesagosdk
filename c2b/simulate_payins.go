package c2b

import "github.com/coleYab/mpesasdk/types"

type SimulateCustomerInititatedPayment struct {
	CommandID types.CommandId `json:"CommandID" validate:"required"`
	Amount uint64 `json:"Amount" validate:"required,gte=1"`
	Msisdn string `json:"Msisdn" validate:"required,len=10,numeric"`
	BillRefNumber string `json:"BillRefNumber" validate:"required,min=6,max=20"`
	ShortCode string `json:"ShortCode" validate:"required,url"`
}

