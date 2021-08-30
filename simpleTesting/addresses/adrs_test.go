//TESTE DE UNIDADE
package addresses

import "testing"

type testCase struct {
	addressInput   string
	expectedReturn string
}

func TestAddressType(t *testing.T) {
	testCases := []testCase{
		{"Rua Apinajes", "Rua"},
		{"Rodovia BR21", "Rodovia"},
		{"Avenida Paulista", "Avenida"},
		{"Praça da Sé", "Praça"},
		{"Estrada da noite", "Estrada"},
		{"RUA SANTA VIRGINIA", "Rua"},
		{"Planalto ABC", "Invalid address type"},
	}

	for _, cases := range testCases {
		receivedAddressType := AddressType(cases.addressInput)
		if receivedAddressType != cases.expectedReturn {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				receivedAddressType,
				cases.expectedReturn,
			)
		}
	}
}
