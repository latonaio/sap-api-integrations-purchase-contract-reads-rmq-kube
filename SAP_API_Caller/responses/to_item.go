package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseContract               string      `json:"PurchaseContract"`
			PurchaseContractItem           string      `json:"PurchaseContractItem"`
			PurchasingContractDeletionCode string      `json:"PurchasingContractDeletionCode"`
			PurchaseContractItemText       string      `json:"PurchaseContractItemText"`
			CompanyCode                    string      `json:"CompanyCode"`
			Plant                          string      `json:"Plant"`
			StorageLocation                string      `json:"StorageLocation"`
			MaterialGroup                  string      `json:"MaterialGroup"`
			SupplierMaterialNumber         string      `json:"SupplierMaterialNumber"`
			OrderQuantityUnit              string      `json:"OrderQuantityUnit"`
			TargetQuantity                 string      `json:"TargetQuantity"`
			PurgDocReleaseOrderQuantity    string      `json:"PurgDocReleaseOrderQuantity"`
			OrderPriceUnit                 string      `json:"OrderPriceUnit"`
			ContractNetPriceAmount         string      `json:"ContractNetPriceAmount"`
			DocumentCurrency               string      `json:"DocumentCurrency"`
			NetPriceQuantity               string      `json:"NetPriceQuantity"`
			TaxCode                        string      `json:"TaxCode"`
			TaxCountry                     string      `json:"TaxCountry"`
			StockType                      string      `json:"StockType"`
			IsInfoRecordUpdated            string      `json:"IsInfoRecordUpdated"`
			PriceIsToBePrinted             bool        `json:"PriceIsToBePrinted"`
			PurgDocEstimatedPrice          bool        `json:"PurgDocEstimatedPrice"`
			PlannedDeliveryDurationInDays  string      `json:"PlannedDeliveryDurationInDays"`
			OverdelivTolrtdLmtRatioInPct   string      `json:"OverdelivTolrtdLmtRatioInPct"`
			UnlimitedOverdeliveryIsAllowed bool        `json:"UnlimitedOverdeliveryIsAllowed"`
			UnderdelivTolrtdLmtRatioInPct  string      `json:"UnderdelivTolrtdLmtRatioInPct"`
			PurchasingDocumentItemCategory string      `json:"PurchasingDocumentItemCategory"`
			AccountAssignmentCategory      string      `json:"AccountAssignmentCategory"`
			GoodsReceiptIsExpected         bool        `json:"GoodsReceiptIsExpected"`
			GoodsReceiptIsNonValuated      bool        `json:"GoodsReceiptIsNonValuated"`
			InvoiceIsExpected              bool        `json:"InvoiceIsExpected"`
			InvoiceIsGoodsReceiptBased     bool        `json:"InvoiceIsGoodsReceiptBased"`
			ManualDeliveryAddressID        string      `json:"ManualDeliveryAddressID"`
			VolumeUnit                     string      `json:"VolumeUnit"`
			Subcontractor                  string      `json:"Subcontractor"`
			EvaldRcptSettlmtIsAllowed      bool        `json:"EvaldRcptSettlmtIsAllowed"`
			Material                       string      `json:"Material"`
			ProductType                    string      `json:"ProductType"`
			MaterialType                   string      `json:"MaterialType"`
			ToItemCondition struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PurContrItemCondition"`
			ToItemAddress struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PurCtrAddress"`
		} `json:"results"`
	} `json:"d"`
}
