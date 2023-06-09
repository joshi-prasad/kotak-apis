/*
 * Reports-KSEC
 *
 * Order Reports API's
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package reports

// Order Report for a client.
type Orders struct {
	OrderId           float64 `json:"orderId,omitempty"`
	Variety           string  `json:"variety,omitempty"`
	InstrumentName    string  `json:"instrumentName,omitempty"`
	InstrumentToken   string  `json:"instrumentToken,omitempty"`
	Exchange          string  `json:"exchange,omitempty"`
	OrderQuantity     int32   `json:"orderQuantity,omitempty"`
	PendingQuantity   int32   `json:"pendingQuantity,omitempty"`
	CancelledQuantity int32   `json:"cancelledQuantity,omitempty"`
	FilledQuantity    int32   `json:"filledQuantity,omitempty"`
	DisclosedQuantity int32   `json:"disclosedQuantity,omitempty"`
	TriggerPrice      int32   `json:"triggerPrice,omitempty"`
	Price             float64 `json:"price,omitempty"`
	Product           string  `json:"product,omitempty"`
	TransactionType   string  `json:"transactionType,omitempty"`
	OrderTimestamp    string  `json:"orderTimestamp,omitempty"`
	Validity          string  `json:"validity,omitempty"`
	StatusMessage     string  `json:"statusMessage,omitempty"`
	Tag               string  `json:"tag,omitempty"`
	// Order Status
	Status          string  `json:"status,omitempty"`
	StatusInfo      string  `json:"statusInfo,omitempty"`
	IsFNO           string  `json:"isFNO,omitempty"`
	Leg             float64 `json:"leg,omitempty"`
	Multiplier      float64 `json:"multiplier,omitempty"`
	MarketLot       float64 `json:"marketLot,omitempty"`
	StrikePrice     float64 `json:"strikePrice,omitempty"`
	MarketExchange  string  `json:"marketExchange,omitempty"`
	OptionType      string  `json:"optionType,omitempty"`
	ExchangeOrderId string  `json:"exchangeOrderId,omitempty"`
	ExpiryDate      string  `json:"expiryDate,omitempty"`
	InstrumentType  string  `json:"instrumentType,omitempty"`
}
