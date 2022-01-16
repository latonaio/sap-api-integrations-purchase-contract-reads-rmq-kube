package sap_api_output_formatter

type PurchaseContract struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	APISchema        string `json:"api_schema"`
	PurchaseContract string `json:"purchase_contract"`
	Deleted          bool   `json:"deleted"`
}    
    
type Header struct {
	PurchaseContract               string      `json:"PurchaseContract"`
	PurchaseContractType           string      `json:"PurchaseContractType"`
	CompanyCode                    string      `json:"CompanyCode"`
	PurchasingDocumentDeletionCode string      `json:"PurchasingDocumentDeletionCode"`
	CreationDate                   string      `json:"CreationDate"`
	Supplier                       string      `json:"Supplier"`
	PurchasingOrganization         string      `json:"PurchasingOrganization"`
	PurchasingGroup                string      `json:"PurchasingGroup"`
	PaymentTerms                   string      `json:"PaymentTerms"`
	NetPaymentDays                 string      `json:"NetPaymentDays"`
	DocumentCurrency               string      `json:"DocumentCurrency"`
	ExchangeRate                   string      `json:"ExchangeRate"`
	ValidityStartDate              string      `json:"ValidityStartDate"`
	ValidityEndDate                string      `json:"ValidityEndDate"`
	SupplierRespSalesPersonName    string      `json:"SupplierRespSalesPersonName"`
	SupplierPhoneNumber            string      `json:"SupplierPhoneNumber"`
	IncotermsClassification        string      `json:"IncotermsClassification"`
	PurchaseContractTargetAmount   string      `json:"PurchaseContractTargetAmount"`
	InvoicingParty                 string      `json:"InvoicingParty"`
	ReleaseCode                    string      `json:"ReleaseCode"`
	LastChangeDateTime             string      `json:"LastChangeDateTime"`
	PurchasingProcessingStatus     string      `json:"PurchasingProcessingStatus"`
	PurchasingProcessingStatusName string      `json:"PurchasingProcessingStatusName"`
	PurgContractIsInPreparation    bool        `json:"PurgContractIsInPreparation"`
    ToItem                         string      `json:"to_PurchaseContractItem"`	
}

type Item struct {
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
	ToItemCondition                string      `json:"to_PurContrItemCondition"`
	ToItemAddress                  string      `json:"to_PurCtrAddress"`
}

type ToItem struct {
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
	ToItemCondition                string      `json:"to_PurContrItemCondition"`
	ToItemAddress                  string      `json:"to_PurCtrAddress"`
}

type ToItemAddress struct {
	PurchaseContract           string `json:"PurchaseContract"`
	AddressID                  string `json:"AddressID"`
	PurchaseContractItem       string `json:"PurchaseContractItem"`
	CityName                   string `json:"CityName"`
	PostalCode                 string `json:"PostalCode"`
	StreetName                 string `json:"StreetName"`
	Country                    string `json:"Country"`
	CorrespondenceLanguage     string `json:"CorrespondenceLanguage"`
	Region                     string `json:"Region"`
	PhoneNumber                string `json:"PhoneNumber"`
	FaxNumber                  string `json:"FaxNumber"`
}


type ToItemCondition struct {
	PurchaseContract             string      `json:"PurchaseContract"`
	PurchaseContractItem         string      `json:"PurchaseContractItem"`
	ConditionValidityEndDate     string      `json:"ConditionValidityEndDate"`
	ConditionType                string      `json:"ConditionType"`
	ConditionRecord              string      `json:"ConditionRecord"`
	ConditionSequentialNumber    string      `json:"ConditionSequentialNumber"`
	ConditionValidityStartDate   string      `json:"ConditionValidityStartDate"`
	PricingScaleType             string      `json:"PricingScaleType"`
	PricingScaleBasis            string      `json:"PricingScaleBasis"`
	ConditionScaleQuantity       string      `json:"ConditionScaleQuantity"`
	ConditionScaleQuantityUnit   string      `json:"ConditionScaleQuantityUnit"`
	ConditionScaleAmount         string      `json:"ConditionScaleAmount"`
	ConditionScaleAmountCurrency string      `json:"ConditionScaleAmountCurrency"`
	ConditionCalculationType     string      `json:"ConditionCalculationType"`
	ConditionRateValue           string      `json:"ConditionRateValue"`
	ConditionRateValueUnit       string      `json:"ConditionRateValueUnit"`
	ConditionQuantity            string      `json:"ConditionQuantity"`
	ConditionQuantityUnit        string      `json:"ConditionQuantityUnit"`
	BaseUnit                     string      `json:"BaseUnit"`
	ConditionIsDeleted           bool        `json:"ConditionIsDeleted"`
	PaymentTerms                 string      `json:"PaymentTerms"`
	ConditionReleaseStatus       string      `json:"ConditionReleaseStatus"`
}
