package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-purchase-contract-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchaseContract(purchaseContract, purchaseContractItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(purchaseContract)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(purchaseContract, purchaseContractItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(purchaseContract string) {
	headerData, err := c.callPurchaseContractSrvAPIRequirementHeader("A_PurchaseContract", purchaseContract)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": headerData, "function": "PurchaseContractHeader"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemData, "function": "PurchaseContractItem"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemAddressData, err := c.callToItemAddress(itemData[0].ToItemAddress)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemAddressData, "function": "PurchaseContractAddress"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemAddressData)

	itemConditionData, err := c.callToItemCondition(itemData[0].ToItemCondition)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemConditionData, "function": "PurchaseContractItemCondition"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemConditionData)

}

func (c *SAPAPICaller) callPurchaseContractSrvAPIRequirementHeader(api, purchaseContract string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASECONTRACT_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, purchaseContract)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemAddress(url string) ([]sap_api_output_formatter.ToItemAddress, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemCondition(url string) ([]sap_api_output_formatter.ToItemCondition, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemCondition(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(purchaseContract, purchaseContractItem string) {
	itemData, err := c.callPurchaseContractSrvAPIRequirementItem("A_PurchaseContractItem", purchaseContract, purchaseContractItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemData, "function": "PurchaseContractItem"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemAddressData, err := c.callToItemAddress(itemData[0].ToItemAddress)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemAddressData, "function": "PurchaseContractAddress"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemAddressData)

	itemConditionData, err := c.callToItemCondition(itemData[0].ToItemCondition)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemConditionData, "function": "PurchaseContractItemCondition"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemConditionData)

}

func (c *SAPAPICaller) callPurchaseContractSrvAPIRequirementItem(api, purchaseContract, purchaseContractItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASECONTRACT_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, purchaseContract, purchaseContractItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, purchaseContract string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseContract eq '%s'", purchaseContract))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, purchaseContract, purchaseContractItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseContract eq '%s' and PurchaseContractItem eq '%s'", purchaseContract, purchaseContractItem))
	req.URL.RawQuery = params.Encode()
}
