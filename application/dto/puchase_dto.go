package dto

// PurchaseDTO represents the data needed to purchase a premium package
type PurchaseDTO struct {
	Package string `json:"package" binding:"required"`
}
