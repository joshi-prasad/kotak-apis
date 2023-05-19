/*
 * Portfolio - API
 *
 * This is a sample server for Kotak Trade API - Portfolio
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package portfolio

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type RealisedProfitLossScripwiseApiService service

/*
RealisedProfitLossScripwiseApiService Get Gain Loss Statement for Cash
Returns Gain Loss Statement for Cash
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param consumerKey Unique ID for your application
 * @param sessionToken Session ID for your application
 * @param fy Financial Year : Format - YYYY-YYYY
 * @param fromDate From Date : Format - YYYY-MM-DD
 * @param toDate To Date : Format - YYYY-MM-DD
 * @param instrumentToken Unique token of Security
 * @param optional nil or *RealisedProfitLossScripwiseApiPortfolioPnlCommodityFyFromDateToDateInstrumentTokenGetOpts - Optional Parameters:
     * @param "UserId" (optional.String) -  user Id for your application
     * @param "ClientAccount" (optional.String) -  Client account for which data requested
@return []PnlCommodity
*/

type RealisedProfitLossScripwiseApiPortfolioPnlCommodityFyFromDateToDateInstrumentTokenGetOpts struct {
	UserId        optional.String
	ClientAccount optional.String
}

func (a *RealisedProfitLossScripwiseApiService) PortfolioPnlCommodityFyFromDateToDateInstrumentTokenGet(ctx context.Context, consumerKey string, sessionToken string, fy string, fromDate string, toDate string, instrumentToken string, localVarOptionals *RealisedProfitLossScripwiseApiPortfolioPnlCommodityFyFromDateToDateInstrumentTokenGetOpts) ([]PnlCommodity, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []PnlCommodity
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/pnl/commodity/{fy}/{fromDate}/{toDate}/{instrumentToken}"
	localVarPath = strings.Replace(localVarPath, "{"+"fy"+"}", fmt.Sprintf("%v", fy), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fromDate"+"}", fmt.Sprintf("%v", fromDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toDate"+"}", fmt.Sprintf("%v", toDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instrumentToken"+"}", fmt.Sprintf("%v", instrumentToken), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["consumerKey"] = parameterToString(consumerKey, "")
	localVarHeaderParams["sessionToken"] = parameterToString(sessionToken, "")
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarHeaderParams["userId"] = parameterToString(localVarOptionals.UserId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ClientAccount.IsSet() {
		localVarHeaderParams["clientAccount"] = parameterToString(localVarOptionals.ClientAccount.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PnlCommodity
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RealisedProfitLossScripwiseApiService Get Gain Loss Statement for Cash
Returns Gain Loss Statement for Cash
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param consumerKey Unique ID for your application
 * @param sessionToken Session ID for your application
 * @param fy Financial Year : Format - YYYY-YYYY
 * @param fromDate From Date : Format - YYYY-MM-DD
 * @param toDate To Date : Format - YYYY-MM-DD
 * @param instrumentToken Unique token of Security
 * @param optional nil or *RealisedProfitLossScripwiseApiPortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGetOpts - Optional Parameters:
     * @param "UserId" (optional.String) -  user Id for your application
     * @param "ClientAccount" (optional.String) -  Client account for which data requested
@return []PnlCurrency
*/

type RealisedProfitLossScripwiseApiPortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGetOpts struct {
	UserId        optional.String
	ClientAccount optional.String
}

func (a *RealisedProfitLossScripwiseApiService) PortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGet(ctx context.Context, consumerKey string, sessionToken string, fy string, fromDate string, toDate string, instrumentToken string, localVarOptionals *RealisedProfitLossScripwiseApiPortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGetOpts) ([]PnlCurrency, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []PnlCurrency
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/pnl/currency/{fy}/{fromDate}/{toDate}/{instrumentToken}"
	localVarPath = strings.Replace(localVarPath, "{"+"fy"+"}", fmt.Sprintf("%v", fy), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fromDate"+"}", fmt.Sprintf("%v", fromDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toDate"+"}", fmt.Sprintf("%v", toDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instrumentToken"+"}", fmt.Sprintf("%v", instrumentToken), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["consumerKey"] = parameterToString(consumerKey, "")
	localVarHeaderParams["sessionToken"] = parameterToString(sessionToken, "")
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarHeaderParams["userId"] = parameterToString(localVarOptionals.UserId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ClientAccount.IsSet() {
		localVarHeaderParams["clientAccount"] = parameterToString(localVarOptionals.ClientAccount.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PnlCurrency
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RealisedProfitLossScripwiseApiService Get Gain Loss Statement for Cash
Returns Gain Loss Statement for Cash
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param consumerKey Unique ID for your application
 * @param sessionToken Session ID for your application
 * @param fy Financial Year : Format - YYYY-YYYY
 * @param fromDate From Date : Format - YYYY-MM-DD
 * @param toDate To Date : Format - YYYY-MM-DD
 * @param instrumentToken Unique token of Security
 * @param optional nil or *RealisedProfitLossScripwiseApiPortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGetOpts - Optional Parameters:
     * @param "UserId" (optional.String) -  user Id for your application
     * @param "ClientAccount" (optional.String) -  Client account for which data requested
@return []PnlDerivative
*/

type RealisedProfitLossScripwiseApiPortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGetOpts struct {
	UserId        optional.String
	ClientAccount optional.String
}

func (a *RealisedProfitLossScripwiseApiService) PortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGet(ctx context.Context, consumerKey string, sessionToken string, fy string, fromDate string, toDate string, instrumentToken string, localVarOptionals *RealisedProfitLossScripwiseApiPortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGetOpts) ([]PnlDerivative, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []PnlDerivative
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/pnl/derivative/{fy}/{fromDate}/{toDate}/{instrumentToken}"
	localVarPath = strings.Replace(localVarPath, "{"+"fy"+"}", fmt.Sprintf("%v", fy), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fromDate"+"}", fmt.Sprintf("%v", fromDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toDate"+"}", fmt.Sprintf("%v", toDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instrumentToken"+"}", fmt.Sprintf("%v", instrumentToken), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["consumerKey"] = parameterToString(consumerKey, "")
	localVarHeaderParams["sessionToken"] = parameterToString(sessionToken, "")
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarHeaderParams["userId"] = parameterToString(localVarOptionals.UserId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ClientAccount.IsSet() {
		localVarHeaderParams["clientAccount"] = parameterToString(localVarOptionals.ClientAccount.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PnlDerivative
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RealisedProfitLossScripwiseApiService Get Gain Loss Statement for Cash
Returns Gain Loss Statement for Cash
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param consumerKey Unique ID for your application
 * @param sessionToken Session ID for your application
 * @param fy Financial Year : Format - YYYY-YYYY
 * @param fromDate From Date : Format - YYYY-MM-DD
 * @param toDate To Date : Format - YYYY-MM-DD
 * @param instrumentToken Unique token of Security
 * @param optional nil or *RealisedProfitLossScripwiseApiPortfolioPnlEquityFyFromDateToDateInstrumentTokenGetOpts - Optional Parameters:
     * @param "UserId" (optional.String) -  user Id for your application
     * @param "ClientAccount" (optional.String) -  Client account for which data requested
@return []PnlEquity
*/

type RealisedProfitLossScripwiseApiPortfolioPnlEquityFyFromDateToDateInstrumentTokenGetOpts struct {
	UserId        optional.String
	ClientAccount optional.String
}

func (a *RealisedProfitLossScripwiseApiService) PortfolioPnlEquityFyFromDateToDateInstrumentTokenGet(ctx context.Context, consumerKey string, sessionToken string, fy string, fromDate string, toDate string, instrumentToken string, localVarOptionals *RealisedProfitLossScripwiseApiPortfolioPnlEquityFyFromDateToDateInstrumentTokenGetOpts) ([]PnlEquity, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []PnlEquity
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/pnl/equity/{fy}/{fromDate}/{toDate}/{instrumentToken}"
	localVarPath = strings.Replace(localVarPath, "{"+"fy"+"}", fmt.Sprintf("%v", fy), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fromDate"+"}", fmt.Sprintf("%v", fromDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toDate"+"}", fmt.Sprintf("%v", toDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instrumentToken"+"}", fmt.Sprintf("%v", instrumentToken), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["consumerKey"] = parameterToString(consumerKey, "")
	localVarHeaderParams["sessionToken"] = parameterToString(sessionToken, "")
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarHeaderParams["userId"] = parameterToString(localVarOptionals.UserId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ClientAccount.IsSet() {
		localVarHeaderParams["clientAccount"] = parameterToString(localVarOptionals.ClientAccount.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PnlEquity
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RealisedProfitLossScripwiseApiService Get Gain Loss Statement for Cash
Returns Gain Loss Statement for Cash
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param consumerKey Unique ID for your application
 * @param sessionToken Session ID for your application
 * @param fy Financial Year : Format - YYYY-YYYY
 * @param fromDate From Date : Format - YYYY-MM-DD
 * @param toDate To Date : Format - YYYY-MM-DD
 * @param instrumentToken Unique token of Security
 * @param optional nil or *RealisedProfitLossScripwiseApiPortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGetOpts - Optional Parameters:
     * @param "UserId" (optional.String) -  user Id for your application
     * @param "ClientAccount" (optional.String) -  Client account for which data requested
@return []PnlMutualfund
*/

type RealisedProfitLossScripwiseApiPortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGetOpts struct {
	UserId        optional.String
	ClientAccount optional.String
}

func (a *RealisedProfitLossScripwiseApiService) PortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGet(ctx context.Context, consumerKey string, sessionToken string, fy string, fromDate string, toDate string, instrumentToken string, localVarOptionals *RealisedProfitLossScripwiseApiPortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGetOpts) ([]PnlMutualfund, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []PnlMutualfund
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/pnl/mutualfund/{fy}/{fromDate}/{toDate}/{instrumentToken}"
	localVarPath = strings.Replace(localVarPath, "{"+"fy"+"}", fmt.Sprintf("%v", fy), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fromDate"+"}", fmt.Sprintf("%v", fromDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toDate"+"}", fmt.Sprintf("%v", toDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instrumentToken"+"}", fmt.Sprintf("%v", instrumentToken), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["consumerKey"] = parameterToString(consumerKey, "")
	localVarHeaderParams["sessionToken"] = parameterToString(sessionToken, "")
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarHeaderParams["userId"] = parameterToString(localVarOptionals.UserId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ClientAccount.IsSet() {
		localVarHeaderParams["clientAccount"] = parameterToString(localVarOptionals.ClientAccount.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PnlMutualfund
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
