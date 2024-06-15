package payment

type (
	ItemServiceReq struct {
	}

	ItemServiceReqDatum struct {
		ItemId string  `json:"_id" validate:"required,max=64"`
		Price  float64 `json:"price"`
	}
)
