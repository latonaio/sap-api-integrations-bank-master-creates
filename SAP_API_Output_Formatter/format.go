package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-bank-master-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToBank(raw []byte, l *logger.Logger) (*Bank, error) {
	pm := &responses.Bank{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Bank. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D

	bank := &Bank{
		BankCountry:         data.BankCountry,
		BankInternalID:      data.BankInternalID,
		BankName:            data.BankName,
		Region:              data.Region,
		ShortStreetName:     data.ShortStreetName,
		ShortCityName:       data.ShortCityName,
		SWIFTCode:           data.SWIFTCode,
		BankNetworkGrouping: data.BankNetworkGrouping,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
		Bank:                data.Bank,
		BankBranch:          data.BankBranch,
		BankCategory:        data.BankCategory,
	}

	return bank, nil
}
