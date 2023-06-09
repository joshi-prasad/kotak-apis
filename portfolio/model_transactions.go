/*
 * Portfolio - API
 *
 * This is a sample server for Kotak Trade API - Portfolio
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package portfolio

// Transaction Details
type Transactions struct {
	// Security or Scheme Name
	Instrument string `json:"instrument,omitempty"`
	// Date of the Transaction
	Date string `json:"date,omitempty"`
	// Buy or Sell Type of Transaction.
	TransactionType string `json:"transactionType,omitempty"`
	// Net Rate of Transaction.
	NetRate float32 `json:"netRate,omitempty"`
	// Net Qunatity of Transaction.
	Quntity int32 `json:"quntity,omitempty"`
	// Net Amount of Transaction.
	Amount float32 `json:"amount,omitempty"`
	// Unrealised Gain Loss of Transaction.
	RealisedGainLoss float32 `json:"realisedGainLoss,omitempty"`
	// Acquisition Type of Transaction
	AcquisitionType string `json:"acquisitionType,omitempty"`
	// squareUp Flag of the Transaction
	SquareUp string `json:"squareUp,omitempty"`
}
