package dictionaryapi

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/httputils"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
)

const (
	dictionaryApiUrl = "https://api.dictionaryapi.dev"
)

type (
	// Client represents the dictionary api client.
	Client struct {
		httpClient *resty.Client
		Word       WordsManager
	}

	// ClientOption are additional settings that can be passed to the Client.
	ClientOption func(*Client)

	// Error represents the error returned by the dictionary api.
	Error struct {
		Title      string `json:"title"`
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
		Resolution string `json:"resolution"`
	}
)

// NewClient returns a instance of Client with the default settings.
// These settings can be optionally modified using ClientOptions that can be passed to NewClient.
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		httpClient: resty.New(),
	}

	c.Word = &wordManager{Client: c}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithHTTPClient can be used as a ClientOption to pass a custom httpClient.
func WithHTTPClient(httpClient *resty.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithWordsManager can be used as a ClientIOption to pass a custom implementation of the WordsManage interface.
func WithWordsManager(wm WordsManager) ClientOption {
	return func(c *Client) {
		c.Word = wm
	}
}

// request formats and returns a base http request that can be extended later.
// as part of this the baseUrl and the default headers are set in the http client.
func (c *Client) request() *resty.Request {
	c.httpClient.SetBaseURL(dictionaryApiUrl)
	c.httpClient.SetDisableWarn(true)

	return c.httpClient.R().SetHeader(httputils.ContentTypeHeaderKey, httputils.ApplicationJsonMIMEType).
		SetHeader(httputils.AcceptHeaderKey, httputils.ApplicationJsonMIMEType)
}

// get makes a GET API call to the dictionaryapi based on the URI passed as input.
// The result is unmarshal-ed into the result interface.
// The method returns an errors.RestErr if
//	- there is an error while making the http request
// 	- the Status code returned by the dictionary API is not 200 - Status Ok
func (c *Client) get(uri string, result interface{}) *errors.RestErr {
	apiError := new(Error)
	resp, err := c.request().SetResult(result).SetError(apiError).Get(uri)
	logger.RestyDebugLogs(resp)
	if err != nil {
		return &errors.RestErr{
			Message: http.StatusText(http.StatusInternalServerError),
			StatusCode: http.StatusInternalServerError,
			Error: err.Error(),
		}
	}

	if resp.StatusCode() != http.StatusOK {
		return &errors.RestErr{
			Message: apiError.Message,
			StatusCode: resp.StatusCode(),
			Error: apiError.Title,
		}
	}

	return nil
}
