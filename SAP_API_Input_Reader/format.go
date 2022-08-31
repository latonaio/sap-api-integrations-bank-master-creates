package sap_api_input_reader

import (
	"sap-api-integrations-bank-master-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToBank() *requests.Bank {
	data := sdc.Bank
	return &requests.Bank{
		BankCountry:         data.BankCountry,
		BankInternalID:      data.BankInternalID,
		BankName:            data.BankName,
		Region:              data.Region,
		ShortStreetName:     data.ShortStreetName,
		ShortCityName:       data.ShortCityName,
		SWIFTCode:           data.SWIFTCode,
		BankNetworkGrouping: data.BankNetworkGrouping,
		// IsMarkedForDeletion: data.IsMarkedForDeletion,
		Bank:         data.Bank,
		BankBranch:   data.BankBranch,
		BankCategory: data.BankCategory,
	}
}
