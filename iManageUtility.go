package iManageUtility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"iManageUtility/Structs"
	"mime"
)

func InitImanageClient(server string, api int, client_id string, client_secret string, grant_type string, scope string) (Structs.ImanageClient, error) {
	client := Structs.ImanageClient{
		BaseUrl: "https://" + server,
		Api: 2,
		HttpClient: &http.Client{},
		Client_id: client_id,
		Client_secret: client_secret,
		Grant_type: grant_type,
		Scope: scope}
	
	return client, nil
}

// Get an authorization token using a password grant via Oauth
func GetLoginTokenPassword(client *Structs.ImanageClient, userid string, password string) (bool, error) {
	
	url := client.BaseUrl + "/auth/oauth2/token"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Build out our query paramters
	q := req.URL.Query()
	q.Add("username", userid)
	q.Add("password", password)
	q.Add("grant_type", client.Grant_type)
	q.Add("client_id", client.Client_id)
	q.Add("client_secret", client.Client_secret)
	q.Add("scope", client.Scope)

	req.URL.RawQuery = q.Encode()

	myClient := client.HttpClient
	resp, err := myClient.Do(req)

	if err != nil {
		return false, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	
	err = json.Unmarshal([]byte(bodyString), &data)
	if err != nil {
		return false, err
	}
	client.Access_token = data["access_token"].(string)
	return true, nil
}

// Get the customer ID for provided access_token for a logged in user. 
func GetCustomerID(client *Structs.ImanageClient) (bool, error) {
	url := client.BaseUrl + "/api"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return true, err
	}

	req.Header.Set("X-Auth-Token", client.Access_token)

	myClient := client.HttpClient
	resp, err := myClient.Do(req)

	if err != nil {
		return true, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return true, err
	}

	bodyString := string(bodyBytes)
	data := Structs.ApiResponse{}
	err = json.Unmarshal([]byte(bodyString), &data)
	if err != nil {
		return true, err
	}
	client.Customer_ID = data.Data.User.CustomerID
	return true, nil
}

// Get documents and return their profile(s)
func GetDocuments(server, access_token, limit string) (string, error) {

	return "", nil
}

func GetDocumentProfile(client Structs.ImanageClient, database string, docID string) (Structs.DocumentProfile, error) {
	url := client.BaseUrl + "/api/v2/customers/" + strconv.Itoa(client.Customer_ID) + "/libraries/" + database + "/documents/" + docID
	data := Structs.DocumentProfile{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}

	req.Header.Set("X-Auth-Token", client.Access_token)

	myClient := client.HttpClient
	resp, err := myClient.Do(req)
	if err != nil {
		return data, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		return data, err
	}

	bodyString := string(bodyBytes)
	
	err = json.Unmarshal([]byte(bodyString), &data)
	
	if err != nil {
		return data, err
	}
	
	return data, nil
}

func GetWorkspaceProfile(client Structs.ImanageClient, database string, prjID string) (Structs.WorkspaceProfile, error) {
	url := client.BaseUrl + "/api/v2/customers/" + strconv.Itoa(client.Customer_ID) + "/libraries/" + database + "/workspaces/" + prjID
	data := Structs.WorkspaceProfile{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}

	req.Header.Set("X-Auth-Token", client.Access_token)
	myClient := client.HttpClient
	resp, err := myClient.Do(req)
	if err != nil {
		return data, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		return data, err
	}

	bodyString := string(bodyBytes)
	
	err = json.Unmarshal([]byte(bodyString), &data)
	
	if err != nil {
		return data, err
	}
	
	return data, nil
}

func DownloadDocument(client Structs.ImanageClient, database string, docID string, downloadPath string) (Structs.DocumentDownloadHeaders, error) {
	url := client.BaseUrl + "/api/v2/customers/" + strconv.Itoa(client.Customer_ID) + "/libraries/" + database + "/documents/" + docID + "/download"
	data := Structs.DocumentDownloadHeaders{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}

	req.Header.Set("X-Auth-Token", client.Access_token)
	myClient := client.HttpClient
	resp, err := myClient.Do(req)
	if err != nil {
		return data, err
	}

	disposition, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	filename := params["filename"]
	fmt.Println(disposition)
	fmt.Println(filename)
	

	return data, nil
}