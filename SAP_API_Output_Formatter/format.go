package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-quotation-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
	SalesQuotation:                 data.SalesQuotation,
	SalesQuotationType:             data.SalesQuotationType,
	SalesOrganization:              data.SalesOrganization,
	DistributionChannel:            data.DistributionChannel,
	OrganizationDivision:           data.OrganizationDivision,
	SalesGroup:                     data.SalesGroup,
	SalesOffice:                    data.SalesOffice,
	SalesDistrict:                  data.SalesDistrict,
	SoldToParty:                    data.SoldToParty,
	CreationDate:                   data.CreationDate,
	LastChangeDate:                 data.LastChangeDate,
	PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
	CustomerPurchaseOrderDate:      data.CustomerPurchaseOrderDate,
	SalesQuotationDate:             data.SalesQuotationDate,
	TotalNetAmount:                 data.TotalNetAmount,
	TransactionCurrency:            data.TransactionCurrency,
	SDDocumentReason:               data.SDDocumentReason,
	PricingDate:                    data.PricingDate,
	RequestedDeliveryDate:          data.RequestedDeliveryDate,
	ShippingCondition:              data.ShippingCondition,
	CompleteDeliveryIsDefined:      data.CompleteDeliveryIsDefined,
	ShippingType:                   data.ShippingType,
	HeaderBillingBlockReason:       data.HeaderBillingBlockReason,
	DeliveryBlockReason:            data.DeliveryBlockReason,
	BindingPeriodValidityStartDate: data.BindingPeriodValidityStartDate,
	BindingPeriodValidityEndDate:   data.BindingPeriodValidityEndDate,
	HdrOrderProbabilityInPercent:   data.HdrOrderProbabilityInPercent,
	ExpectedOrderNetAmount:         data.ExpectedOrderNetAmount,
	IncotermsClassification:        data.IncotermsClassification,
	CustomerPaymentTerms:           data.CustomerPaymentTerms,
	CustomerPriceGroup:             data.CustomerPriceGroup,
	PriceListType:                  data.PriceListType,
	PaymentMethod:                  data.PaymentMethod,
	CustomerTaxClassification1:     data.CustomerTaxClassification1,
	ReferenceSDDocument:            data.ReferenceSDDocument,
	ReferenceSDDocumentCategory:    data.ReferenceSDDocumentCategory,
	SalesQuotationApprovalReason:   data.SalesQuotationApprovalReason,
	SalesDocApprovalStatus:         data.SalesDocApprovalStatus,
	OverallSDProcessStatus:         data.OverallSDProcessStatus,
	TotalCreditCheckStatus:         data.TotalCreditCheckStatus,
	OverallSDDocumentRejectionSts:  data.OverallSDDocumentRejectionSts,
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
	SalesQuotation:                data.SalesQuotation,
	SalesQuotationItem:            data.SalesQuotationItem,
	SalesQuotationItemCategory:    data.SalesQuotationItemCategory,
	SalesQuotationItemText:        data.SalesQuotationItemText,
	PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
	Material:                      data.Material,
	MaterialByCustomer:            data.MaterialByCustomer,
	RequestedQuantity:             data.RequestedQuantity,
	RequestedQuantityUnit:         data.RequestedQuantityUnit,
	ItemOrderProbabilityInPercent: data.ItemOrderProbabilityInPercent,
	ItemGrossWeight:               data.ItemGrossWeight,
	ItemNetWeight:                 data.ItemNetWeight,
	ItemWeightUnit:                data.ItemWeightUnit,
	ItemVolume:                    data.ItemVolume,
	ItemVolumeUnit:                data.ItemVolumeUnit,
	TransactionCurrency:           data.TransactionCurrency,
	NetAmount:                     data.NetAmount,
	MaterialGroup:                 data.MaterialGroup,
	MaterialPricingGroup:          data.MaterialPricingGroup,
	Batch:                         data.Batch,
	Plant:                         data.Plant,
	IncotermsClassification:       data.IncotermsClassification,
	CustomerPaymentTerms:          data.CustomerPaymentTerms,
	ProductTaxClassification1:     data.ProductTaxClassification1,
	SalesDocumentRjcnReason:       data.SalesDocumentRjcnReason,
	WBSElement:                    data.WBSElement,
	ProfitCenter:                  data.ProfitCenter,
	ReferenceSDDocument:           data.ReferenceSDDocument,
	ReferenceSDDocumentItem:       data.ReferenceSDDocumentItem,
	SDProcessStatus:               data.SDProcessStatus,
	}

	return item, nil
}
