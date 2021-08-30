//TESTE DE UNIDADE
package addresses

import "testing"

func TestAddressType(t *testing.T) {
	addresForTest := "Avenida Paulista"
	expectedAddressType := "Avenida"
	receivedAddressType := AddressType(addresForTest)

	if receivedAddressType != expectedAddressType {
		t.Error("O tipo recebido é diferente do esperado!")
	}
}
