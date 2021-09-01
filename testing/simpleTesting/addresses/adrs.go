package addresses

import "strings"

//AddresType verifies if an address is valid and returns it
func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia", "praça"}
	addressLowerCase := strings.ToLower(address)
	firtsAddressWord := strings.Split(addressLowerCase, " ")[0]

	isAddressValid := false

	for _, addressType := range validTypes {
		if addressType == firtsAddressWord {
			isAddressValid = true
		}
	}

	if isAddressValid {
		return strings.Title(firtsAddressWord)
	}

	return "Invalid address type"
}
