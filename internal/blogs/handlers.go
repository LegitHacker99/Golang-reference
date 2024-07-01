package blogs

import (
	db "chap1/internal/database"
	users "chap1/internal/users"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate
var DB *gorm.DB

func init() {
	validate = validator.New()
	DB = db.GormDB
}

func Get_blog_data(w http.ResponseWriter, r *http.Request) {
	var blogs []Blog
	if err := DB.Find(&blogs).Error; err != nil {
		log.Println("Error fetching blogs:", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(blogs)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		http.Error(w, "Error serializing data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func Post_blog(w http.ResponseWriter, r *http.Request) {

	// READ FROM REQUEST BODY, BUT ALWAYS REURNS []BYTE
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("server: could not read request body: \n", err)
	}

	// CONVERT []BYTE TO STRING AND BECOMES READABLE
	log.Print("\n", string(reqBody))

	// READ FROM REQUEST BODY TO A VAR
	// MAP[STRING]INTERFACE{} FOR UNKNOWN DATA STRUCTURE OF JSON
	var dumdum map[string]interface{}
	err = json.Unmarshal(reqBody, &dumdum)
	if err != nil {
		log.Print("Error here in unmarshalling to map[string]interface{}")
	}

	// UNMARSHAL() DOES NOT FORMAT JSON
	log.Println("\n\nchampoon::::::::::\n", dumdum)
	// FORMAT USING MARSHALINDENT()
	formattedJson, _ := json.MarshalIndent(dumdum, "", "\t")
	log.Println(string(formattedJson))

	// PARSE INTO KNOWN STRUCT
	var blog Blog
	err = json.Unmarshal(reqBody, &blog)
	if err != nil {
		log.Println("error unmarshalling to struct", err)
	}

	// UNFORMATTED JSON
	log.Println("\n\nblogger::::::::\n", blog)
	formattedJson2, _ := json.MarshalIndent(blog, "", "\t")
	log.Println(string(formattedJson2))

	// YOU CAN USE OS.STDOUT TO WRITE TO THE CONSOLE
	os.Stdout.Write(formattedJson)

	// POST OPERATION
	newEntry := DB.Create(&blog)
	if newEntry.Error != nil {
		log.Println("newest:::::", newEntry.Error)
		http.Error(w, "POST FAILED", http.StatusInternalServerError)
	}

}

func Put_blog(w http.ResponseWriter, r *http.Request) {
	var db = db.GormDB
	var userData users.User

	payload, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	err := json.Unmarshal(payload, &userData)
	if err != nil {
		log.Println("Unmarshal error")
	}

	res := db.Model(&users.User{}).Where("user_uuid = ?", userData.User_Uuid).Updates(userData)
	if res.Error != nil {
		log.Println("problem here:::::::", res.Error)
		http.Error(w, "Query Fail", http.StatusInternalServerError)
	}
}

func Patch_blog(w http.ResponseWriter, r *http.Request) {

}

func Del_blog(w http.ResponseWriter, r *http.Request) {

}
