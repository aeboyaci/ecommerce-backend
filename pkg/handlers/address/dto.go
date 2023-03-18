package address

type AddressResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	PhoneNumber string `json:"phoneNumber"`
}
