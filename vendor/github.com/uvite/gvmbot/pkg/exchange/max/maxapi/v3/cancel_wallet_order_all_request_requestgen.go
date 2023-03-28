// Code generated by "requestgen -method DELETE -url /api/v3/wallet/:walletType/orders -type CancelWalletOrderAllRequest -responseType []Order"; DO NOT EDIT.

package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/uvite/gvmbot/pkg/exchange/max/maxapi"
	"net/url"
	"reflect"
	"regexp"
)

func (c *CancelWalletOrderAllRequest) Side(side string) *CancelWalletOrderAllRequest {
	c.side = &side
	return c
}

func (c *CancelWalletOrderAllRequest) Market(market string) *CancelWalletOrderAllRequest {
	c.market = &market
	return c
}

func (c *CancelWalletOrderAllRequest) GroupID(groupID uint32) *CancelWalletOrderAllRequest {
	c.groupID = &groupID
	return c
}

func (c *CancelWalletOrderAllRequest) WalletType(walletType max.WalletType) *CancelWalletOrderAllRequest {
	c.walletType = walletType
	return c
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (c *CancelWalletOrderAllRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (c *CancelWalletOrderAllRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check side field -> json key side
	if c.side != nil {
		side := *c.side

		// assign parameter of side
		params["side"] = side
	} else {
	}
	// check market field -> json key market
	if c.market != nil {
		market := *c.market

		// assign parameter of market
		params["market"] = market
	} else {
	}
	// check groupID field -> json key groupID
	if c.groupID != nil {
		groupID := *c.groupID

		// assign parameter of groupID
		params["groupID"] = groupID
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (c *CancelWalletOrderAllRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := c.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if c.isVarSlice(_v) {
			c.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (c *CancelWalletOrderAllRequest) GetParametersJSON() ([]byte, error) {
	params, err := c.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (c *CancelWalletOrderAllRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check walletType field -> json key walletType
	walletType := c.walletType

	// TEMPLATE check-required
	if len(walletType) == 0 {
		return nil, fmt.Errorf("walletType is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of walletType
	params["walletType"] = walletType

	return params, nil
}

func (c *CancelWalletOrderAllRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (c *CancelWalletOrderAllRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (c *CancelWalletOrderAllRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (c *CancelWalletOrderAllRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := c.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (c *CancelWalletOrderAllRequest) Do(ctx context.Context) ([]max.Order, error) {

	params, err := c.GetParameters()
	if err != nil {
		return nil, err
	}
	query := url.Values{}

	apiURL := "/api/v3/wallet/:walletType/orders"
	slugs, err := c.GetSlugsMap()
	if err != nil {
		return nil, err
	}

	apiURL = c.applySlugsToUrl(apiURL, slugs)

	req, err := c.client.NewAuthenticatedRequest(ctx, "DELETE", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := c.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []max.Order
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}