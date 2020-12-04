package Structs
import (
	"time"
	"net/http"
)

func (client *ImanageClient) UpdateToken(token string) {
	client.Access_token = token
}
type ImanageClient struct {
	BaseUrl string
	Api int
	UserAgent string
	HttpClient *http.Client
	Client_id string
	Client_secret string
	Grant_type string
	Scope string
	Access_token string
	Customer_ID int
}

type ApiResponse struct {
	Data struct {
		AuthStatus string `json:"auth_status"`
		User       struct {
			CustomerID int    `json:"customer_id"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Email      string `json:"email"`
		} `json:"user"`
		DmsVersion string `json:"dms_version"`
		Work       struct {
			PreferredLibrary string `json:"preferred_library"`
			Libraries        []struct {
				Alias  string `json:"alias"`
				Type   string `json:"type"`
				Region string `json:"region"`
			} `json:"libraries"`
		} `json:"work"`
		Region       string        `json:"region"`
		Capabilities []interface{} `json:"capabilities"`
		Versions     []struct {
			Name    string `json:"name"`
			URL     string `json:"url"`
			Version string `json:"version"`
		} `json:"versions"`
	} `json:"data"`
}

type DocumentProfile struct {
	Data struct {
		Access             string    `json:"access"`
		Author             string    `json:"author"`
		Class              string    `json:"class"`
		ClassDescription   string    `json:"class_description"`
		ContentType        string    `json:"content_type"`
		CreateDate         time.Time `json:"create_date"`
		Database           string    `json:"database"`
		Declared           bool      `json:"declared"`
		DefaultSecurity    string    `json:"default_security"`
		DocumentNumber     int       `json:"document_number"`
		EditDate           time.Time `json:"edit_date"`
		EditProfileDate    time.Time `json:"edit_profile_date"`
		Extension          string    `json:"extension"`
		FileCreateDate     time.Time `json:"file_create_date"`
		FileEditDate       time.Time `json:"file_edit_date"`
		ID                 string    `json:"id"`
		InUse              bool      `json:"in_use"`
		InUseBy            string    `json:"in_use_by"`
		Indexable          bool      `json:"indexable"`
		IsCheckedOut       bool      `json:"is_checked_out"`
		IsDeclared         bool      `json:"is_declared"`
		IsExternal         bool      `json:"is_external"`
		IsExternalAsNormal bool      `json:"is_external_as_normal"`
		IsHipaa            bool      `json:"is_hipaa"`
		IsInUse            bool      `json:"is_in_use"`
		IsRelated          bool      `json:"is_related"`
		IsRestorable       bool      `json:"is_restorable"`
		Iwl                string    `json:"iwl"`
		LastUser           string    `json:"last_user"`
		Name               string    `json:"name"`
		Operator           string    `json:"operator"`
		RetainDays         int       `json:"retain_days"`
		Size               int       `json:"size"`
		Type               string    `json:"type"`
		TypeDescription    string    `json:"type_description"`
		Version            int       `json:"version"`
		WorkspaceID        string    `json:"workspace_id"`
		WorkspaceName      string    `json:"workspace_name"`
		Wstype             string    `json:"wstype"`
	} `json:"data"`
}

type WorkspaceProfile struct {
	Data struct {
		Author                 string    `json:"author"`
		Class                  string    `json:"class"`
		ClassDescription       string    `json:"class_description"`
		ContentType            string    `json:"content_type"`
		CreateDate             time.Time `json:"create_date"`
		Custom1                string    `json:"custom1"`
		Custom2                string    `json:"custom2"`
		Database               string    `json:"database"`
		Declared               bool      `json:"declared"`
		DefaultSecurity        string    `json:"default_security"`
		DocumentNumber         int       `json:"document_number"`
		EditDate               time.Time `json:"edit_date"`
		EditProfileDate        time.Time `json:"edit_profile_date"`
		EffectiveSecurity      string    `json:"effective_security"`
		Extension              string    `json:"extension"`
		FileCreateDate         time.Time `json:"file_create_date"`
		FileEditDate           time.Time `json:"file_edit_date"`
		HasSubfolders          bool      `json:"has_subfolders"`
		ID                     string    `json:"id"`
		InUse                  bool      `json:"in_use"`
		Indexable              bool      `json:"indexable"`
		IsCheckedOut           bool      `json:"is_checked_out"`
		IsContainerSavedSearch bool      `json:"is_container_saved_search"`
		IsContentSavedSearch   bool      `json:"is_content_saved_search"`
		IsDeclared             bool      `json:"is_declared"`
		IsExternal             bool      `json:"is_external"`
		IsExternalAsNormal     bool      `json:"is_external_as_normal"`
		IsHidden               bool      `json:"is_hidden"`
		IsHipaa                bool      `json:"is_hipaa"`
		IsInUse                bool      `json:"is_in_use"`
		IsRelated              bool      `json:"is_related"`
		IsRestorable           bool      `json:"is_restorable"`
		Iwl                    string    `json:"iwl"`
		LastUser               string    `json:"last_user"`
		MyMattersShortcutID    string    `json:"my_matters_shortcut_id"`
		Name                   string    `json:"name"`
		Operator               string    `json:"operator"`
		Owner                  string    `json:"owner"`
		RetainDays             int       `json:"retain_days"`
		Size                   int       `json:"size"`
		Subtype                string    `json:"subtype"`
		Type                   string    `json:"type"`
		TypeDescription        string    `json:"type_description"`
		Version                int       `json:"version"`
		WorkspaceID            string    `json:"workspace_id"`
		Wstype                 string    `json:"wstype"`
	} `json:"data"`
}

type DocumentDownloadHeaders struct {

}