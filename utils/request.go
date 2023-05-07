package utils

// import (
// 	"fmt"
// 	"net/http"
// 	"net/url"
// 	"strings"

// 	"github.com/labibramadhan/golang-underscore"

// 	"encoding/json"
// 	"github.com/parnurzeal/gorequest"
// 	"github.com/spf13/cast"
// 	"time"
// )

// type MRequest struct {
// 	HttpClient *gorequest.SuperAgent

// 	BaseURL       string
// 	LastQueryData url.Values
// }

// func Request(baseUrl string) *MRequest {
// 	req := &MRequest{
// 		HttpClient: gorequest.New().Timeout(90 * time.Second),
// 		BaseURL:    baseUrl,
// 	}

// 	req.HttpClient.Debug = IsDev()

// 	return req
// }

// func (mRequest *MRequest) Get(urls interface{}, queries interface{}) *gorequest.SuperAgent {
// 	mRequest.ApplyQueries(queries)

// 	q := mRequest.HttpClient.Get(mRequest.GetUrl(urls)).
// 		Retry(2, 500*time.Millisecond, http.StatusRequestTimeout, http.StatusBadGateway, http.StatusTooManyRequests)

// 	mRequest.SyncClient(q)

// 	return q
// }

// func (mRequest *MRequest) Post(urls interface{}, queries interface{}, payload interface{}) *gorequest.SuperAgent {
// 	mRequest.ApplyQueries(queries)

// 	q := mRequest.HttpClient.Post(mRequest.GetUrl(urls)).
// 		Retry(2, 500*time.Millisecond, http.StatusRequestTimeout, http.StatusBadGateway, http.StatusTooManyRequests)

// 	if payload != nil {
// 		q = q.Send(payload)
// 	}

// 	mRequest.SyncClient(q)

// 	return q
// }

// func (mRequest *MRequest) Put(urls interface{}, queries interface{}, payload interface{}) *gorequest.SuperAgent {
// 	mRequest.ApplyQueries(queries)

// 	q := mRequest.HttpClient.Put(mRequest.GetUrl(urls)).
// 		Retry(2, 500*time.Millisecond, http.StatusRequestTimeout, http.StatusBadGateway, http.StatusTooManyRequests)

// 	if payload != nil {
// 		q = q.Send(payload)
// 	}

// 	mRequest.SyncClient(q)

// 	return q
// }

// func (mRequest *MRequest) Delete(urls interface{}, queries interface{}) *gorequest.SuperAgent {
// 	mRequest.ApplyQueries(queries)

// 	q := mRequest.HttpClient.Delete(mRequest.GetUrl(urls)).
// 		Retry(2, 500*time.Millisecond, http.StatusRequestTimeout, http.StatusBadGateway, http.StatusTooManyRequests)
// 	mRequest.SyncClient(q)

// 	return q
// }

// func (mRequest *MRequest) GetUrl(urls interface{}) string {
// 	urlsCast := cast.ToStringSlice(urls)
// 	urlsFiltered := underscore.Reject(urlsCast, func(val string, _ int) bool {
// 		return val == ""
// 	})

// 	urlsFilteredVal, urlsFilteredOK := urlsFiltered.([]string)

// 	targetUrls := []string{mRequest.BaseURL}

// 	if urlsFilteredOK {
// 		targetUrls = append(targetUrls, urlsFilteredVal...)
// 	}

// 	return strings.Join(targetUrls, "/")
// }

// func (mRequest *MRequest) ApplyQueries(queries interface{}) {
// 	if queries != nil {
// 		if IsMap(queries) {
// 			queries := cast.ToStringMapString(queries)
// 			for queryKey, queryVal := range queries {
// 				mRequest.HttpClient.Query(fmt.Sprintf("%s=%s", queryKey, queryVal))
// 			}

// 			mRequest.LastQueryData = mRequest.HttpClient.QueryData
// 		} else {
// 			queriesJSON, errJSON := json.Marshal(queries)

// 			if errJSON != nil {
// 				panic(errJSON)
// 				return
// 			}

// 			queriesMap := make(map[string]interface{})

// 			errMap := json.Unmarshal(queriesJSON, &queriesMap)
// 			if errMap != nil {
// 				panic(errMap)
// 				return
// 			}

// 			mRequest.ApplyQueries(queriesMap)
// 		}
// 	}
// }

// func (mRequest *MRequest) SyncClient(q *gorequest.SuperAgent) {
// 	q.QueryData = mRequest.LastQueryData
// }

// func (mRequest *MRequest) BasicAuth(username string, password string) *MRequest {
// 	mRequest.HttpClient.SetBasicAuth(username, password)
// 	return mRequest
// }

// func (mRequest *MRequest) BuildUrl() string {
// 	return mRequest.GetUrl([]string{mRequest.HttpClient.Url})
// }

// func (mRequest *MRequest) MakeRequest() (*http.Request, error) {
// 	return mRequest.HttpClient.MakeRequest()
// }
