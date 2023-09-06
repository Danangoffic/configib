package Tools

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Danangoffic/configib/app/models"
)

func UserCredentialsToOBM(requestOrPasswordFormOrUserMap interface{}) map[string]string {
	return userCredentialsToOBM(requestOrPasswordFormOrUserMap, "")
}

func userCredentialsToOBM(requestOrPasswordFormOrUserMap interface{}, key string) map[string]string {
	// var credential string
	var passwordForm *models.PasswordForm
	var request *http.Request
	var userMap map[string]interface{}
	var OBMParameterMap map[string]string = make(map[string]string)

	if key == "" || len(key) < 1 {
		if requestOrPasswordFormOrUserMap == nil {
			return OBMParameterMap
		}

		switch requestOrPasswordFormOrUserMap.(type) {
		case *http.Request:
			request = requestOrPasswordFormOrUserMap.(*http.Request)
			if request.FormValue("OBMParameterWeb") != "" {
				jsonMap := make(map[string]interface{})
				err := json.Unmarshal([]byte(request.FormValue("OBMParameterWeb")), &jsonMap)
				if err == nil {
					OBMParameterMap = jsonMap["OBMParameter"].(map[string]string)
				}
			}
		case map[string]interface{}:
			userMap = requestOrPasswordFormOrUserMap.(map[string]interface{})
			if credential, ok := userMap["easyPin"]; ok {
				credential = credential.(string)
				if credential != "" {
					OBMParameterMap = parseCredentials(credential.(string))
				}
			} else if credential, ok := userMap["loginPassword"]; ok {
				credential = credential.(string)
				if credential != "" {
					OBMParameterMap = parseCredentials(credential.(string))
				}
			} else if credential, ok := userMap["userPassword"]; ok {
				credential = credential.(string)
				if credential != "" {
					OBMParameterMap = parseCredentials(credential.(string))
				}
			}
		case models.PasswordForm:
			passwordForm = requestOrPasswordFormOrUserMap.(*models.PasswordForm)
			OBMParameterMap = passwordForm.OBMParameter
		default:
			return OBMParameterMap
		}
	} else {
		if passwordForm, ok := requestOrPasswordFormOrUserMap.(*models.PasswordForm); ok {
			OBMParameterMap = passwordForm.OBMParameter
		} else if userMap, ok := requestOrPasswordFormOrUserMap.(map[string]interface{}); ok {
			if key == "OLD" {
				if credential, ok := userMap["oldEasyPin"]; ok {
					credential = credential.(string)
					if credential != "" {
						OBMParameterMap = parseCredentials(credential.(string))
					}
				}
			} else if key == "NEW" {
				if credential, ok := userMap["easyPin"]; ok {
					credential = credential.(string)
					if credential != "" {
						OBMParameterMap = parseCredentials(credential.(string))
					}
				}
			} else if key == "CONFIRMNEWPASSWORD" {
				if credential, ok := userMap["confirmNewPassword"]; ok {
					credential = credential.(string)
					if credential != "" {
						OBMParameterMap = parseCredentials(credential.(string))
					}
				}
			}
		}
	}

	return OBMParameterMap
}

func parseCredentials(credential string) map[string]string {
	var OBMParameterMap map[string]string = make(map[string]string)
	credentialsArray := strings.Split(credential, "|")
	for _, credential := range credentialsArray {
		configCodeOBMKeyValue := strings.Split(credential, ":")
		if len(configCodeOBMKeyValue) == 2 {
			OBMParameterMap[configCodeOBMKeyValue[0]] = configCodeOBMKeyValue[1]
		}
	}
	return OBMParameterMap
}
