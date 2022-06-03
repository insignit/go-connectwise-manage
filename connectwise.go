package go_connectwise_manage

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// SystemInfo returned from ConnectWise.
type SystemInfo struct {
	Version        string `json:"version"`
	IsCloud        bool   `json:"isCloud"`
	ServerTimeZone string `json:"serverTimeZone"`
	CloudRegion    string `json:"cloudRegion"`
}

// APIVersion is info from ConnectWise needed to construct the correct base url.
type APIVersion struct {
	CompanyName string `json:"CompanyName"`
	Codebase    string `json:"Codebase"`
	VersionCode string `json:"VersionCode"`
	CompanyID   string `json:"CompanyID"`
	IsCloud     bool   `json:"IsCloud"`
	SiteUrl     string `json:"SiteUrl"`
}

// CWClient is a 'holder' struct for everything needed to authenticate to the ConnectWise API.
type CWClient struct {
	APIVersion APIVersion
	ClientID   string
	CompanyID  string
	PublicKey  string
	PrivateKey string
}

// CWRequestOption makes up one (of multiple) options that we can pass to functions.
type CWRequestOption struct {
	Key   string
	Value string
}

// Post is an API primitive to post data to the ConnectWise API.
func (c CWClient) Post(path string, payload []byte, options ...CWRequestOption) (string, error) {
	baseURL := fmt.Sprintf("https://%s/%sapis/3.0", c.APIVersion.SiteUrl, c.APIVersion.Codebase)
	url := fmt.Sprintf("%s/%s", baseURL, path)
	client := &http.Client{}

	// Setup the post request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	// Authorization.
	q := req.URL.Query()
	q.Add("clientId", c.ClientID)

	auth := fmt.Sprintf("%s+%s:%s", c.CompanyID, c.PublicKey, c.PrivateKey)
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded))
	req.Header.Set("Content-Type", "application/json")

	// Query parameters.
	if len(options) > 0 {
		for _, opt := range options {
			q.Add(opt.Key, opt.Value)
		}
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 201 {
		return "", fmt.Errorf("Non-201 status: Code: %d Status: %s Message: %s", resp.StatusCode, resp.Status, body)
	}
	return string(body), nil
}

// GetSystemInfo will retrieve the needed system info from ConnectWise.
func (c CWClient) GetSystemInfo(options ...CWRequestOption) (info SystemInfo, err error) {
	j, err := c.Get("/system/info", options...)
	if err != nil {
		return info, fmt.Errorf("Can't get system info %v", err)
	}
	err = json.Unmarshal(j, &info)
	if err != nil {
		return info, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return info, nil
}

func (c CWClient) GetAll(path string, options ...CWRequestOption) (jsonPages []string, err error) {
	var currentPage int = 1
	var results []string

	for {
		page := CWRequestOption{Key: "page", Value: strconv.Itoa(currentPage)}
		newOpts := append(options, page)
		resp, err := c.Get(path, newOpts...)
		if err != nil {
			return jsonPages, err
		}

		// Some requests will do normal pagination.
		// So when they get to the end they will return an empty set.
		if string(resp) == "[]" {
			break
		}

		// Some endpoints (like `/system/info`) do not follow normal
		// pagination, they just return the same data twice,
		// so if this occurs then return the lastResponse.
		if len(results) > 0 && string(resp) == string(results[0]) {
			return []string{string(resp)}, nil
		}

		results = append(results, string(resp))
		currentPage++
	}

	return results, nil
}

// Get is an API primitive to retrieve data from the ConnectWise API.
func (c CWClient) Get(path string, options ...CWRequestOption) (jsonData []byte, err error) {
	baseURL := fmt.Sprintf("https://%s/%sapis/3.0", c.APIVersion.SiteUrl, c.APIVersion.Codebase)
	url := fmt.Sprintf("%s/%s", baseURL, path)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return jsonData, err
	}

	// Authentication.
	q := req.URL.Query()
	q.Add("clientId", c.ClientID)

	auth := fmt.Sprintf("%s+%s:%s", c.CompanyID, c.PublicKey, c.PrivateKey)
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded))

	// Query parameters.
	if len(options) > 0 {
		for _, opt := range options {
			q.Add(opt.Key, opt.Value)
		}
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return jsonData, err
	}
	defer resp.Body.Close()

	jsonData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return jsonData, err
	}

	if resp.StatusCode != 200 {
		return jsonData, fmt.Errorf("Non-200 status: Code: %d Status: %s Message: %s", resp.StatusCode, resp.Status, jsonData)
	}
	return jsonData, nil
}

// NewCWClient create a new ConnectWise client.
func NewCWClient(site string, clientId string, company string, publicKey string, privateKey string) (client CWClient, err error) {
	apiVersion, err := GetApiVersion(site, company)
	if err != nil {
		return client, fmt.Errorf("Cannot get api version for %s at %s: %w", company, site, err)
	}
	client = CWClient{
		APIVersion: apiVersion,
		ClientID:   clientId,
		CompanyID:  company,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
	return
}

// GetAPIVersion will dynamically get the API version for this client, all that.
// is required is the site and company, no authentication is needed at this point.
func GetApiVersion(site string, company string) (version APIVersion, err error) {
	url := fmt.Sprintf("https://%s/login/companyinfo/%s", site, company)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &version)
	if err != nil {
		return version, fmt.Errorf("Can't get unmarshal data %v", err)
	}

	return
}
