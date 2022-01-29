package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-contract-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			PurchaseContract:               data.PurchaseContract,
			PurchaseContractType:           data.PurchaseContractType,
			CompanyCode:                    data.CompanyCode,
			PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
			CreationDate:                   data.CreationDate,
			Supplier:                       data.Supplier,
			PurchasingOrganization:         data.PurchasingOrganization,
			PurchasingGroup:                data.PurchasingGroup,
			PaymentTerms:                   data.PaymentTerms,
			NetPaymentDays:                 data.NetPaymentDays,
			DocumentCurrency:               data.DocumentCurrency,
			ExchangeRate:                   data.ExchangeRate,
			ValidityStartDate:              data.ValidityStartDate,
			ValidityEndDate:                data.ValidityEndDate,
			SupplierRespSalesPersonName:    data.SupplierRespSalesPersonName,
			SupplierPhoneNumber:            data.SupplierPhoneNumber,
			IncotermsClassification:        data.IncotermsClassification,
			PurchaseContractTargetAmount:   data.PurchaseContractTargetAmount,
			InvoicingParty:                 data.InvoicingParty,
			ReleaseCode:                    data.ReleaseCode,
			LastChangeDateTime:             data.LastChangeDateTime,
			PurchasingProcessingStatus:     data.PurchasingProcessingStatus,
			PurchasingProcessingStatusName: data.PurchasingProcessingStatusName,
			PurgContractIsInPreparation:    data.PurgContractIsInPreparation,
			ToItem:                         data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			PurchaseContract:               data.PurchaseContract,
			PurchaseContractItem:           data.PurchaseContractItem,
			PurchasingContractDeletionCode: data.PurchasingContractDeletionCode,
			PurchaseContractItemText:       data.PurchaseContractItemText,
			CompanyCode:                    data.CompanyCode,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			MaterialGroup:                  data.MaterialGroup,
			SupplierMaterialNumber:         data.SupplierMaterialNumber,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			TargetQuantity:                 data.TargetQuantity,
			PurgDocReleaseOrderQuantity:    data.PurgDocReleaseOrderQuantity,
			OrderPriceUnit:                 data.OrderPriceUnit,
			ContractNetPriceAmount:         data.ContractNetPriceAmount,
			DocumentCurrency:               data.DocumentCurrency,
			NetPriceQuantity:               data.NetPriceQuantity,
			TaxCode:                        data.TaxCode,
			TaxCountry:                     data.TaxCountry,
			StockType:                      data.StockType,
			IsInfoRecordUpdated:            data.IsInfoRecordUpdated,
			PriceIsToBePrinted:             data.PriceIsToBePrinted,
			PurgDocEstimatedPrice:          data.PurgDocEstimatedPrice,
			PlannedDeliveryDurationInDays:  data.PlannedDeliveryDurationInDays,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
			GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
			InvoiceIsExpected:              data.InvoiceIsExpected,
			InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
			ManualDeliveryAddressID:        data.ManualDeliveryAddressID,
			VolumeUnit:                     data.VolumeUnit,
			Subcontractor:                  data.Subcontractor,
			EvaldRcptSettlmtIsAllowed:      data.EvaldRcptSettlmtIsAllowed,
			Material:                       data.Material,
			ProductType:                    data.ProductType,
			MaterialType:                   data.MaterialType,
			ToItemCondition:                data.ToItemCondition.Deferred.URI,
			ToItemAddress:                  data.ToItemAddress.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			PurchaseContract:               data.PurchaseContract,
			PurchaseContractItem:           data.PurchaseContractItem,
			PurchasingContractDeletionCode: data.PurchasingContractDeletionCode,
			PurchaseContractItemText:       data.PurchaseContractItemText,
			CompanyCode:                    data.CompanyCode,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			MaterialGroup:                  data.MaterialGroup,
			SupplierMaterialNumber:         data.SupplierMaterialNumber,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			TargetQuantity:                 data.TargetQuantity,
			PurgDocReleaseOrderQuantity:    data.PurgDocReleaseOrderQuantity,
			OrderPriceUnit:                 data.OrderPriceUnit,
			ContractNetPriceAmount:         data.ContractNetPriceAmount,
			DocumentCurrency:               data.DocumentCurrency,
			NetPriceQuantity:               data.NetPriceQuantity,
			TaxCode:                        data.TaxCode,
			TaxCountry:                     data.TaxCountry,
			StockType:                      data.StockType,
			IsInfoRecordUpdated:            data.IsInfoRecordUpdated,
			PriceIsToBePrinted:             data.PriceIsToBePrinted,
			PurgDocEstimatedPrice:          data.PurgDocEstimatedPrice,
			PlannedDeliveryDurationInDays:  data.PlannedDeliveryDurationInDays,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
			GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
			InvoiceIsExpected:              data.InvoiceIsExpected,
			InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
			ManualDeliveryAddressID:        data.ManualDeliveryAddressID,
			VolumeUnit:                     data.VolumeUnit,
			Subcontractor:                  data.Subcontractor,
			EvaldRcptSettlmtIsAllowed:      data.EvaldRcptSettlmtIsAllowed,
			Material:                       data.Material,
			ProductType:                    data.ProductType,
			MaterialType:                   data.MaterialType,
			ToItemCondition:                data.ToItemCondition.Deferred.URI,
			ToItemAddress:                  data.ToItemAddress.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemAddress(raw []byte, l *logger.Logger) ([]ToItemAddress, error) {
	pm := &responses.ToItemAddress{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemAddress. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemAddress := make([]ToItemAddress, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemAddress = append(toItemAddress, ToItemAddress{
			PurchaseContract:       data.PurchaseContract,
			AddressID:              data.AddressID,
			PurchaseContractItem:   data.PurchaseContractItem,
			CityName:               data.CityName,
			PostalCode:             data.PostalCode,
			StreetName:             data.StreetName,
			Country:                data.Country,
			CorrespondenceLanguage: data.CorrespondenceLanguage,
			Region:                 data.Region,
			PhoneNumber:            data.PhoneNumber,
			FaxNumber:              data.FaxNumber,
		})
	}

	return toItemAddress, nil
}

func ConvertToToItemCondition(raw []byte, l *logger.Logger) ([]ToItemCondition, error) {
	pm := &responses.ToItemCondition{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemCondition. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemCondition := make([]ToItemCondition, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemCondition = append(toItemCondition, ToItemCondition{
			PurchaseContract:             data.PurchaseContract,
			PurchaseContractItem:         data.PurchaseContractItem,
			ConditionValidityEndDate:     data.ConditionValidityEndDate,
			ConditionType:                data.ConditionType,
			ConditionRecord:              data.ConditionRecord,
			ConditionSequentialNumber:    data.ConditionSequentialNumber,
			ConditionValidityStartDate:   data.ConditionValidityStartDate,
			PricingScaleType:             data.PricingScaleType,
			PricingScaleBasis:            data.PricingScaleBasis,
			ConditionScaleQuantity:       data.ConditionScaleQuantity,
			ConditionScaleQuantityUnit:   data.ConditionScaleQuantityUnit,
			ConditionScaleAmount:         data.ConditionScaleAmount,
			ConditionScaleAmountCurrency: data.ConditionScaleAmountCurrency,
			ConditionCalculationType:     data.ConditionCalculationType,
			ConditionRateValue:           data.ConditionRateValue,
			ConditionRateValueUnit:       data.ConditionRateValueUnit,
			ConditionQuantity:            data.ConditionQuantity,
			ConditionQuantityUnit:        data.ConditionQuantityUnit,
			BaseUnit:                     data.BaseUnit,
			ConditionIsDeleted:           data.ConditionIsDeleted,
			PaymentTerms:                 data.PaymentTerms,
			ConditionReleaseStatus:       data.ConditionReleaseStatus,
		})
	}

	return toItemCondition, nil
}
