/*
 * Portfolio - API
 *
 * This is a sample server for Kotak Trade API - Portfolio
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package portfolio

// Scripwise Transaction Details
type TransactionsScripwise struct {
	// Name of the instrument
	InstrumentName string `json:"instrumentName,omitempty"`
	// Buy or Sell Type of Transaction.
	TransactionType string `json:"transactionType,omitempty"`
	// Buy Quantity of instrument.
	BoughtQty float32 `json:"boughtQty,omitempty"`
	// Buy Rate of intrument.
	BoughtRate float32 `json:"boughtRate,omitempty"`
	// Buy Value of intrument.
	BoughtValue float32 `json:"boughtValue,omitempty"`
	// Sell Quantity of instrument.
	SoldQty float32 `json:"soldQty,omitempty"`
	// Sell Rate of intrument.
	SoldRate float32 `json:"soldRate,omitempty"`
	// Sell Value of intrument.
	SoldValue float32 `json:"soldValue,omitempty"`
	// Date of Transaction for instrument
	BillDate string `json:"billDate,omitempty"`
	// Bill Number of a Transaction for instrument
	BillNo string `json:"billNo,omitempty"`
	// Days to Long Term Capital Gain.
	DaysToLtcg float64 `json:"daysToLtcg,omitempty"`
	// Acquisition Type of Transaction
	AcquisitionType string `json:"acquisitionType,omitempty"`
}
