package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RequestBody(w http.ResponseWriter, r *http.Request) string {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Something Went Wrong Reading Request Body", http.StatusInternalServerError)
		panic("error :- " + err.Error())
	}

	defer r.Body.Close()
	requestData := string(body)
	return requestData
}

func GetMapData[Reference any](s string, structs Reference) Reference {
	res_data := []byte(s)
	unMarsh := json.Unmarshal(res_data, &structs)
	if unMarsh != nil {
		panic("err" + unMarsh.Error())
	}
	return structs
}

type Response struct {
	Message string `json:"message"`
	Success int    `json:"success"`
}
type ResponseFail struct {
	Message string `json:"message"`
	Success int    `json:"success"`
}
type FinalResponse struct {
	APIResponse interface{} `json:"API Response"`
	Data        interface{} `json:"Data"`
}

func ApiSuccess(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(statusCode)

	if data == "" {
		data = []int{}
	}
	response := Response{Message: message, Success: 1}
	response_result := FinalResponse{APIResponse: response, Data: data}
	json.NewEncoder(w).Encode(response_result)
}

func ApiFailure(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(statusCode)
	response := ResponseFail{Message: message, Success: 0}
	response_result := FinalResponse{APIResponse: response, Data: []int{}}
	json.NewEncoder(w).Encode(response_result)
}

func StringToMap(stringData string) map[string]string {
	var out map[string]string

	_ = json.Unmarshal([]byte(stringData), &out)

	return out
}

func FormateDate(inputDate string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, inputDate)
	if err != nil {
		return "", err
	}
	formattedDate := parsedTime.Format("2006-01-02")
	return formattedDate, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetUserFromToken(r *gin.Context) int {
	userClaims, exists := r.Get("login_user")
	if !exists {
		return 0
	}
	user_id := userClaims.(jwt.MapClaims)["ID"].(int)
	return user_id
}
