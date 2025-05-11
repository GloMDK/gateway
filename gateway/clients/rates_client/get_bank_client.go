package rates_client

import (
	"fmt"
	"gateway/service"
)

func (r *RatesClient) GetBankClientByName(bankName service.BankName) (service.BankClient, error) {
	bankClient, found := r.banksClients[bankName]
	if !found {
		return nil, fmt.Errorf("there is no bankClient with name: %s", bankName)
	}

	return bankClient, nil
}
