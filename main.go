package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/alexedwards/scs/v2"
	//"github.com/alexedwards/scs/v2/memstore"
	"github.com/asaskevich/govalidator"
	"github.com/casbin/casbin"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golangSql/m/auth"
	"golangSql/m/models"
	"golangSql/m/routes"
	"golangSql/m/utils"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const(
	layoutDate = "2006-01-02" //YYYY-MM-DD
	layoutDate2 = "02-01-2006"
)

//api
var getAllDataByDay = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	startDate := r.FormValue("startDate")

	tStart, _ := time.Parse(layoutDate,startDate)

	result, err := models.Connection().Prepare("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms" +
		",phoneBlockedIntenSms,total5656Sms,handled5656Sms,verify5656Sms," +
		"wrong5656Sms,phoneDoubtCall," +
		"phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall," +
		"total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand " +
		"from dataColletion where date = ? order by date desc")

	conn , err := result.Query(tStart.Format("2006-04-02"))

	datasTenDinhDanh := models.DataTenDinhDanh{}

	for conn.Next() {
		var id int
		var date time.Time
		var res string
		var smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms int
		var total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms int
		var phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall int
		var total5656Call,handled5656Call,verify5656Call,wrong5656Call int

		conn.Scan(&id,&date,&smsBlockedInten,&smsBlockedExten,&phoneBlockedIntenSms,&phoneBlockedExtenSms,&total5656Sms,&handled5656Sms,&verify5656Sms,&wrong5656Sms,
			&phoneDoubtCall,&phoneBlockedCall,&phoneFromDoubtCall,&phoneFromBlockedCall,&total5656Call,&handled5656Call,&verify5656Call,&wrong5656Call,&res)

		datasTenDinhDanh.AddData(models.TenDinhDanh{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			SmsBlockedInten:      smsBlockedInten,
			SmsBlockedExten:      smsBlockedExten,
			PhoneBlockedIntenSms: phoneBlockedIntenSms,
			PhoneBlockedExtenSms: phoneBlockedExtenSms,
			Total5656Sms:         total5656Sms,
			Handled5656Sms:       handled5656Sms,
			Verify5656Sms:        verify5656Sms,
			Wrong5656Sms:         wrong5656Sms,
			PhoneDoubtCall:       phoneDoubtCall,
			PhoneBlockedCall:     phoneBlockedCall,
			PhoneFromDoubtCall:   phoneFromDoubtCall,
			PhoneFromBlockedCall: phoneFromBlockedCall,
			Total5656Call:        total5656Call,
			Handled5656Call:      handled5656Call,
			Verify5656Call:       verify5656Call,
			Wrong5656Call: wrong5656Call,
			TelecomBrand: res,
		})
	}
	
	dataLast := &models.ApiDataTenDinhDanh{
		Authen:  true,
		Data:    datasTenDinhDanh,
		Success: false,
	}
	
	userBytes, err := json.Marshal(dataLast)

	w.Write(userBytes)

	if err != nil{
		panic(err)
	}

	models.Connection().Close()
})

var getAllDataBrandName = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	telecomBrand, _ := auth.ExtractUsernameFromToken(r)

	p := bluemonday.UGCPolicy()
	res := p.Sanitize(telecomBrand)

	result, err := models.Connection().Query("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms" +
		",phoneBlockedIntenSms,total5656Sms,handled5656Sms,verify5656Sms," +
		"wrong5656Sms,phoneDoubtCall," +
		"phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall," +
		"total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand " +
		"from dataColletion where telecomBrand = ? order by date desc", res)

	datasTenDinhDanh := models.DataTenDinhDanh{}

	for result.Next() {
		var id int
		var date time.Time
		var res string
		var smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms int
		var total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms int
		var phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall int
		var total5656Call,handled5656Call,verify5656Call,wrong5656Call int

		result.Scan(&id,&date,&smsBlockedInten,&smsBlockedExten,&phoneBlockedIntenSms,&phoneBlockedExtenSms,&total5656Sms,&handled5656Sms,&verify5656Sms,&wrong5656Sms,
			&phoneDoubtCall,&phoneBlockedCall,&phoneFromDoubtCall,&phoneFromBlockedCall,&total5656Call,&handled5656Call,&verify5656Call,&wrong5656Call,&res)

		datasTenDinhDanh.AddData(models.TenDinhDanh{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			SmsBlockedInten:      smsBlockedInten,
			SmsBlockedExten:      smsBlockedExten,
			PhoneBlockedIntenSms: phoneBlockedIntenSms,
			PhoneBlockedExtenSms: phoneBlockedExtenSms,
			Total5656Sms:         total5656Sms,
			Handled5656Sms:       handled5656Sms,
			Verify5656Sms:        verify5656Sms,
			Wrong5656Sms:         wrong5656Sms,
			PhoneDoubtCall:       phoneDoubtCall,
			PhoneBlockedCall:     phoneBlockedCall,
			PhoneFromDoubtCall:   phoneFromDoubtCall,
			PhoneFromBlockedCall: phoneFromBlockedCall,
			Total5656Call:        total5656Call,
			Handled5656Call:      handled5656Call,
			Verify5656Call:       verify5656Call,
			Wrong5656Call: wrong5656Call,
			TelecomBrand: res,
		})
	}

	dataLast := &models.ApiDataTenDinhDanh{
		Authen:  true,
		Data : datasTenDinhDanh,
		Success: true,
	}

	userBytes, err := json.Marshal(dataLast)

	w.Write(userBytes)

	if err != nil{
		panic(err)
	}

	models.Connection().Close()
})

var uploadDataFromBrandName = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	checkN := true
	checkD := true

	stmt, err := models.Connection().Prepare("INSERT INTO dataColletion(date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms," +
		"total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms,phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall," +
		"phoneFromBlockedCall,total5656Call,handled5656Call,verify5656Call,telecomBrand,wrong5656Call,phoneBlockedExtenSms) " +
		"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)

	json.Unmarshal(body, &keyVal)

	date := keyVal["date"]
	if auth.CheckDate(date) == false{
		checkD = false
	}
	smsBlockedInten := keyVal["smsBlockedInten"]
	if auth.CheckNum(smsBlockedInten) == false{
		checkN = false
	}
	smsBlockedExten := keyVal["smsBlockedExten"]
	if auth.CheckNum(smsBlockedExten) == false{
		checkN = false
	}
	phoneBlockedIntenSms := keyVal["phoneBlockedIntenSms"]
	if auth.CheckNum(phoneBlockedIntenSms) == false{
		checkN = false
	}
	phoneBlockedExtenSms := keyVal["phoneBlockedExtenSms"]
	if auth.CheckNum(phoneBlockedExtenSms) == false{
		checkN = false
	}
	total5656Sms := keyVal["total5656Sms"]
	if auth.CheckNum(total5656Sms) == false{
		checkN = false
	}
	handled5656Sms := keyVal["handled5656Sms"]
	if auth.CheckNum(handled5656Sms) == false{
		checkN = false
	}
	verify5656Sms := keyVal["verify5656Sms"]
	if auth.CheckNum(verify5656Sms) == false{
		checkN = false
	}
	wrong5656Sms := keyVal["wrong5656Sms"]
	if auth.CheckNum(wrong5656Sms) == false{
		checkN = false
	}
	phoneDoubtCall := keyVal["phoneDoubtCall"]
	if auth.CheckNum(phoneDoubtCall) == false{
		checkN = false
	}
	phoneBlockedCall := keyVal["phoneBlockedCall"]
	if auth.CheckNum(phoneBlockedCall) == false{
		checkN = false
	}
	phoneFromDoubtCall := keyVal["phoneFromDoubtCall"]
	if auth.CheckNum(phoneFromDoubtCall) == false{
		checkN = false
	}
	phoneFromBlockedCall := keyVal["phoneFromBlockedCall"]
	if auth.CheckNum(phoneFromBlockedCall) == false{
		checkN = false
	}
	total5656Call := keyVal["total5656Call"]
	if auth.CheckNum(total5656Call) == false{
		checkN = false
	}
	handled5656Call := keyVal["handled5656Call"]
	if auth.CheckNum(handled5656Call) == false{
		checkN = false
	}
	verify5656Call := keyVal["verify5656Call"]
	if auth.CheckNum(verify5656Call) == false{
		checkN = false
	}
	wrong5656Call := keyVal["wrong5656Call"]
	if auth.CheckNum(wrong5656Call) == false{
		checkN = false
	}
	telecomBrand, _ := auth.ExtractUsernameFromToken(r)

	if checkN == true && checkD == true {
		convertTime, _ := time.Parse(layoutDate2, date)
		dateInsert := convertTime.Format("2006-01-02")

		_, err = stmt.Exec(dateInsert, smsBlockedInten, smsBlockedExten, phoneBlockedIntenSms,total5656Sms, handled5656Sms,
			verify5656Sms, wrong5656Sms, phoneDoubtCall, phoneBlockedCall, phoneFromDoubtCall,
			phoneFromBlockedCall, total5656Call, handled5656Call, verify5656Call,telecomBrand,wrong5656Call,phoneBlockedExtenSms)

		if err != nil {
			panic(err.Error())
		}

		suc := &models.CreateDataSuccess{
			Msg:     "Đã thêm dữ liệu thành công",
			Success: true,
		}

		userBytes, err := json.Marshal(suc)

		w.Write(userBytes)

		if err != nil{
			panic(err)
		}
	}else if checkN == false && checkD == true{
		suc := &models.CreateDataSuccess{
			Msg:     "Số nhập đang không đúng hoặc vượt quá 999.999",
			Success: false,
		}

		userBytes, err := json.Marshal(suc)

		w.Write(userBytes)

		if err != nil{
			panic(err)
		}
	}else if checkD == false && checkN == true{
		suc := &models.CreateDataSuccess{
			Msg:     "Định dạng ngày không đúng hoặc đang ở thì tương lai",
			Success: false,
		}

		userBytes, err := json.Marshal(suc)

		w.Write(userBytes)

		if err != nil{
			panic(err)
		}
	}else if checkD == false && checkN == false{
		suc := &models.CreateDataSuccess{
			Msg:     "Nhập sai toàn bộ định dạng ngày và số",
			Success: false,
		}

		userBytes, err := json.Marshal(suc)

		w.Write(userBytes)

		if err != nil{
			panic(err)
		}
	}

	models.Connection().Close()
})

var dashboardSms = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	success := true

	telecomBrand, _ := auth.ExtractUsernameFromToken(r)

	p := bluemonday.UGCPolicy()
	res := p.Sanitize(telecomBrand)

	result, err := models.Connection().Query("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms" +
		",phoneBlockedIntenSms,total5656Sms,handled5656Sms,verify5656Sms," +
		"wrong5656Sms, telecomBrand " +
		"from dataColletion where telecomBrand = ? order by date desc", res)

	if err != nil{
		success = false
		panic(err)
	}


	dataDashboardSms := models.DataDashboardSms{}

	for result.Next() {
		var id int
		var date time.Time
		var smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms int
		var total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms int

		result.Scan(&id,&date,&smsBlockedInten,&smsBlockedExten,&phoneBlockedIntenSms,&phoneBlockedExtenSms,&total5656Sms,&handled5656Sms,&verify5656Sms,&wrong5656Sms,&res)

		dataDashboardSms.AddData(models.DashboardSms{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			SmsBlockedInten:      smsBlockedInten,
			SmsBlockedExten:      smsBlockedExten,
			PhoneBlockedIntenSms: phoneBlockedIntenSms,
			PhoneBlockedExtenSms: phoneBlockedExtenSms,
			Total5656Sms:         total5656Sms,
			Handled5656Sms:       handled5656Sms,
			Verify5656Sms:        verify5656Sms,
			Wrong5656Sms:         wrong5656Sms,
			TelecomBrand:         res,
		})
	}

	dataLast := &models.ApiDataDashboardSms{
		Authen:  true,
		Data : dataDashboardSms,
		Success: success,
	}

	userBytes, err := json.Marshal(dataLast)

	w.Write(userBytes)

	if err != nil{
		panic(err)
	}

	models.Connection().Close()
})

var dashboardCall = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	success := true

	telecomBrand, _ := auth.ExtractUsernameFromToken(r)

	p := bluemonday.UGCPolicy()
	res := p.Sanitize(telecomBrand)

	result, err := models.Connection().Query("SELECT id,date, phoneDoubtCall," +
		"phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall," +
		"total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand " +
		"from dataColletion where telecomBrand = ? order by date desc", res)

	if  err != nil {
		success = false
	}

	dataDashboardCall := models.DataDashboardCall{}

	for result.Next() {
		var id int
		var date time.Time
		var phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall int
		var total5656Call,handled5656Call,verify5656Call,wrong5656Call int

		result.Scan(&id,&date,&phoneDoubtCall,&phoneBlockedCall,&phoneFromDoubtCall,&phoneFromBlockedCall,&total5656Call,&handled5656Call,&verify5656Call,&wrong5656Call,&res)

		dataDashboardCall.AddData(models.DashboardCall{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			PhoneDoubtCall:       phoneDoubtCall,
			PhoneBlockedCall:     phoneBlockedCall,
			PhoneFromDoubtCall:   phoneFromDoubtCall,
			PhoneFromBlockedCall: phoneFromBlockedCall,
			Total5656Call:        total5656Call,
			Handled5656Call:      handled5656Call,
			Verify5656Call:       verify5656Call,
			Wrong5656Call: wrong5656Call,
			TelecomBrand: res,
		})
	}

	dataLast := &models.ApiDataDashboardCall{
		Authen:  true,
		Data : dataDashboardCall,
		Success: success,
	}

	userBytes, err := json.Marshal(dataLast)

	w.Write(userBytes)

	if err != nil{
		panic(err)
	}

	models.Connection().Close()
})

//login - logout

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	creds := &models.Login{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(creds)

	//Santize parameters username and password
	username := routes.Santize(creds.Username)
	password, err := routes.Hash(creds.Password)

	//check isNull
	if govalidator.IsNull(username) || govalidator.IsNull(password) {
		mes := "Tên đăng nhập hoặc mật khẩu đang trống"
		success := false

		data := &models.LoginFail{
			Msg:     mes,
			Success: success,
		}

		userBytes, err := json.Marshal(data)

		if err != nil{
			log.Println("Error")
		}

		w.Write(userBytes)
		return
	}

	//get password from db
	result := models.Connection().QueryRow("select passwd from users where account= ?", creds.Username)
	if  err != nil {
		// If there is an issue with the database, return a 500 error
		utils.JSON(w, 500, "Internal Server Error")
		return
	}

	storedCreds := &models.Login{}

	err = result.Scan(&storedCreds.Password)

	if  err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) success
		if err == sql.ErrNoRows {
			mes := "Mật khẩu không đúng hoặc nhập sai kiểu dữ liệu"
			success := false

			data := &models.LoginFail{
				Msg:     mes,
				Success: success,
			}

			userBytes, err := json.Marshal(data)

			if err != nil{
				log.Println("Error")
			}

			w.Write(userBytes)
			return
		}
		// If the error is of any other type, send a 500 success
		mes := "Internal Server Error"
		success := false

		data := &models.LoginFail{
			Msg:     mes,
			Success: success,
		}

		userBytes, err := json.Marshal(data)

		if err != nil{
			log.Println("Error")
		}

		w.Write(userBytes)
		return
	}

	//compare pass
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password));  err != nil {
		mes := "Mật khẩu không đúng hoặc nhập sai kiểu dữ liệu"
		success := false

		data := &models.LoginFail{
			Msg:     mes,
			Success: success,
		}

		userBytes, err := json.Marshal(data)

		if err != nil{
			log.Println("Error")
		}

		w.Write(userBytes)
	}else {
		//create jwt
		token, errCreate := auth.Create(username)
		success := true

		data := &models.LoginSuccess{
			Token: token,
			Success: success,
		}

		if errCreate != nil {
			mes := "Internal Server Error"
			success := false

			data := &models.LoginFail{
				Msg:     mes,
				Success: success,
			}

			userBytes, err := json.Marshal(data)

			if err != nil{
				log.Println("Error")
			}

			w.Write(userBytes)
			return
		}

		userBytes, err := json.Marshal(data)

		if err != nil{
			log.Println("Error")
		}

		w.Write(userBytes)
	}

	if models.CheckRole(username,password) == 1{
		http.Redirect(w,r,"/",http.StatusSeeOther)
	}else if models.CheckRole(username,password) == 2{
		http.Redirect(w,r,"/Admin/GetData",http.StatusSeeOther)
	}else if models.CheckRole(username,password) != 2 && models.CheckRole(username,password) != 1{
		http.Redirect(w,r,"/Warning",http.StatusSeeOther)
	}
})

var ChangePasswd = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var account, _ = auth.ExtractUsernameFromToken(r)

	result, err := models.Connection().Prepare("UPDATE users SET passwd = ? where account = ?")

	if err != nil {
		// If there is an issue with the database, return a 500 error
		utils.JSON(w, 500, "Internal Server Error")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	newTitle := keyVal["password"]

	passHash, _ := routes.Hash(newTitle)

	_, err = result.Exec(passHash, account)

	if err != nil {
		panic(err.Error())
	} else {
		utils.JSON(w, 200 , `{"success" :`+ `"true"}`)
		models.Connection().Close()
	}
})
//func not use in final version
var Register = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := models.Connection().Prepare("INSERT INTO Users (account, passwd,[role],brandName) VALUES (?,?,?,?)")

	if  err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if  err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	account := keyVal["account"]
	password := keyVal["password"]
	role := keyVal["role"]
	brandName := keyVal["brandName"]

	passHash, _ := routes.Hash(password)

	_, err = stmt.Exec(account,passHash,role,brandName)

	if  err != nil {
		panic(err.Error())
	}

	mes := "Tạo tài khoản thành công"
	success := true

	data := &models.LoginFail{
		Msg:     mes,
		Success: success,
	}

	userBytes, err := json.Marshal(data)
	w.Write(userBytes)
})

//func logoutHandler() http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if err := session.Renew(r); err != nil {
//			routes.WriteError(http.StatusInternalServerError, "ERROR", w, err)
//			return
//		}
//		writeSuccess("SUCCESS", w)
//	})
//}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/login.html")
}

//admin
//tong 2 cai
var getAllDataTelecom = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

	tmpl := template.Must(template.ParseFiles("./static/admin/admin.html"))

	result, err := models.Connection().Prepare("SELECT id,date,smsBlockedInten" +
		",smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms,total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms," +
		"phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall,total5656Call,handled5656Call,verify5656Call," +
		"wrong5656Call,telecomBrand from dataColletion order by date asc")

	conn , err := result.Query()

	if err != nil {
		fmt.Println(":POISERASREAS")
	}

	datasTenDinhDanh := models.DataTenDinhDanh{}

	for conn.Next() {
		var id int
		var date time.Time
		var res string
		var smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms int
		var total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms int
		var phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall int
		var total5656Call,handled5656Call,verify5656Call, wrong5656Call int

		conn.Scan(&id,&date,&smsBlockedInten,&smsBlockedExten,&phoneBlockedIntenSms,&phoneBlockedExtenSms,&total5656Sms,&handled5656Sms,&verify5656Sms,&wrong5656Sms,
			&phoneDoubtCall,&phoneBlockedCall,&phoneFromDoubtCall,&phoneFromBlockedCall,&total5656Call,&handled5656Call,&verify5656Call,&wrong5656Call,&res)

		datasTenDinhDanh.AddData(models.TenDinhDanh{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			SmsBlockedInten:      smsBlockedInten,
			SmsBlockedExten:      smsBlockedExten,
			PhoneBlockedIntenSms: phoneBlockedIntenSms,
			PhoneBlockedExtenSms: phoneBlockedExtenSms,
			Total5656Sms:         total5656Sms,
			Handled5656Sms:       handled5656Sms,
			Verify5656Sms:        verify5656Sms,
			Wrong5656Sms:         wrong5656Sms,
			PhoneDoubtCall:       phoneDoubtCall,
			PhoneBlockedCall:     phoneBlockedCall,
			PhoneFromDoubtCall:   phoneFromDoubtCall,
			PhoneFromBlockedCall: phoneFromBlockedCall,
			Total5656Call:        total5656Call,
			Handled5656Call:      handled5656Call,
			Verify5656Call:       verify5656Call,
			Wrong5656Call:        wrong5656Call,
			TelecomBrand:         res,
		})
	}
	tmpl.Execute(w, datasTenDinhDanh)
})

var dashboardAllSms = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("./static/admin/adminSms.html"))

	result, _ := models.Connection().Query("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms" +
		",total5656Sms,handled5656Sms,verify5656Sms," +
		"wrong5656Sms, telecomBrand" +
		"from dataColletion order by date desc")

	dataDashboardSms := models.DataDashboardSms{}

	for result.Next() {
		var id int
		var res string
		var date time.Time
		var smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms int
		var total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms int

		result.Scan(&id,&date,&smsBlockedInten,&smsBlockedExten,&phoneBlockedIntenSms,&phoneBlockedExtenSms,&total5656Sms,&handled5656Sms,&verify5656Sms,&wrong5656Sms,&res)

		dataDashboardSms.AddData(models.DashboardSms{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			SmsBlockedInten:      smsBlockedInten,
			SmsBlockedExten:      smsBlockedExten,
			PhoneBlockedIntenSms: phoneBlockedIntenSms,
			PhoneBlockedExtenSms: phoneBlockedExtenSms,
			Total5656Sms:         total5656Sms,
			Handled5656Sms:       handled5656Sms,
			Verify5656Sms:        verify5656Sms,
			Wrong5656Sms:         wrong5656Sms,
			TelecomBrand:         res,
		})
	}
	tmpl.Execute(w, dataDashboardSms)
	models.Connection().Close()
})

var dashboardAllCall = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("./static/admin/adminCall.html"))

	result, _ := models.Connection().Query("SELECT id,date, phoneDoubtCall," +
		"phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall," +
		"total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand" +
		"from dataColletion order by date desc")

	dataDashboardCall := models.DataDashboardCall{}

	for result.Next() {
		var id int
		var res string
		var date time.Time
		var phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall int
		var total5656Call,handled5656Call,verify5656Call,wrong5656Call int

		result.Scan(&id,&date,&phoneDoubtCall,&phoneBlockedCall,&phoneFromDoubtCall,&phoneFromBlockedCall,&total5656Call,&handled5656Call,&verify5656Call,&wrong5656Call,&res)

		dataDashboardCall.AddData(models.DashboardCall{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			PhoneDoubtCall:       phoneDoubtCall,
			PhoneBlockedCall:     phoneBlockedCall,
			PhoneFromDoubtCall:   phoneFromDoubtCall,
			PhoneFromBlockedCall: phoneFromBlockedCall,
			Total5656Call:        total5656Call,
			Handled5656Call:      handled5656Call,
			Verify5656Call:       verify5656Call,
			Wrong5656Call: wrong5656Call,
			TelecomBrand: res,
		})
	}

	tmpl.Execute(w, dataDashboardCall)
	models.Connection().Close()
})

//theo day
var getAllDataOneTelecomByDay = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	tmpl := template.Must(template.ParseFiles("./static/admin/admin.html"))

	telecomBrand := r.FormValue("telecomBrand")

	startDate := r.FormValue("startDate")

	tStart, _ := time.Parse(layoutDate,startDate)

	var result *sql.Stmt
	var conn *sql.Rows

	if(telecomBrand == "null" && startDate != ""){
		result ,_ = models.Connection().Prepare("SELECT id,date,smsBlockedInten," +
			"smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms,total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms," +
			"phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall,total5656Call,handled5656Call,verify5656Call," +
			"wrong5656Call,telecomBrand from dataColletion where date = ?")
		conn ,_= result.Query(tStart.Format("2006-01-02"))
	}else if(startDate == "" && telecomBrand != "null"){
		result ,_ = models.Connection().Prepare("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms," +
			"phoneBlockedExtenSms,total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms,phoneDoubtCall," +
			"phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall,total5656Call,handled5656Call,verify5656Call,wrong5656Call," +
			"telecomBrand from dataColletion where telecomBrand = ?")
		conn ,_= result.Query(telecomBrand)
	}else if(telecomBrand == "null" && startDate == ""){
		result, _ = models.Connection().Prepare("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms," +
			"phoneBlockedExtenSms,total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms,phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall," +
			"phoneFromBlockedCall,total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand from dataColletion order by date asc")
		conn , _ = result.Query()
	}else if(telecomBrand != "null" && startDate != "") {
		result ,_ = models.Connection().Prepare("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms," +
			"total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms,phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall," +
			"phoneFromBlockedCall,total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand from dataColletionfrom where telecomBrand = ? and date = ?")
		conn , _ = result.Query(telecomBrand, tStart.Format("2006-01-02"))
	}

	datasTenDinhDanh := models.DataTenDinhDanh{}

	for conn.Next() {
		var id int
		var date time.Time
		var res string
		var smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms int
		var total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms int
		var phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall int
		var total5656Call,handled5656Call,verify5656Call, wrong5656Call int

		conn.Scan(&id,&date,&smsBlockedInten,&smsBlockedExten,&phoneBlockedIntenSms,&phoneBlockedExtenSms,&total5656Sms,&handled5656Sms,&verify5656Sms,&wrong5656Sms,
			&phoneDoubtCall,&phoneBlockedCall,&phoneFromDoubtCall,&phoneFromBlockedCall,&total5656Call,&handled5656Call,&verify5656Call,&wrong5656Call,&res)

		datasTenDinhDanh.AddData(models.TenDinhDanh{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			SmsBlockedInten:      smsBlockedInten,
			SmsBlockedExten:      smsBlockedExten,
			PhoneBlockedIntenSms: phoneBlockedIntenSms,
			PhoneBlockedExtenSms: phoneBlockedExtenSms,
			Total5656Sms:         total5656Sms,
			Handled5656Sms:       handled5656Sms,
			Verify5656Sms:        verify5656Sms,
			Wrong5656Sms:         wrong5656Sms,
			PhoneDoubtCall:       phoneDoubtCall,
			PhoneBlockedCall:     phoneBlockedCall,
			PhoneFromDoubtCall:   phoneFromDoubtCall,
			PhoneFromBlockedCall: phoneFromBlockedCall,
			Total5656Call:        total5656Call,
			Handled5656Call:      handled5656Call,
			Verify5656Call:       verify5656Call,
			Wrong5656Call:        wrong5656Call,
			TelecomBrand:         res,
		})
	}
	tmpl.Execute(w, datasTenDinhDanh)
})

var getAllDataSmsBrandNameByDay = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	tmpl := template.Must(template.ParseFiles("./static/admin/adminSms.html"))

	telecomBrand := r.FormValue("telecomBrand")

	startDate := r.FormValue("startDate")

	tStart, _ := time.Parse(layoutDate,startDate)

	var conn *sql.Stmt
	var result *sql.Rows

	if(telecomBrand == "null" && startDate != ""){
		conn, _ = models.Connection().Prepare("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms" +
			",total5656Sms,handled5656Sms,verify5656Sms," +
			"wrong5656Sms, telecomBrand " +
			"from dataColletion where date = ? order by date desc")

		result, _ = conn.Query(tStart.Format("2006-04-02"))
	}else if(startDate == "" && telecomBrand != "null"){
		conn, _ = models.Connection().Prepare("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms" +
			",total5656Sms,handled5656Sms,verify5656Sms," +
			"wrong5656Sms, telecomBrand " +
			"from dataColletion where telecomBrand = ? order by date desc")

		result, _ = conn.Query(telecomBrand)
	}else if (telecomBrand == "null" && startDate == "") {
		conn, _ = models.Connection().Prepare("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms" +
			",total5656Sms,handled5656Sms,verify5656Sms," +
			"wrong5656Sms, telecomBrand " +
			"from dataColletion order by date desc")

		result, _ = conn.Query()
	}else if(telecomBrand != "null" && startDate != "") {
		conn, _ = models.Connection().Prepare("SELECT id,date,smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms" +
			",total5656Sms,handled5656Sms,verify5656Sms," +
			"wrong5656Sms, telecomBrand " +
			"from dataColletion where telecomBrand = ? and date = ? order by date desc")

		result, _ = conn.Query(telecomBrand,tStart.Format("2006-04-02"))
	}

	dataDashboardSms := models.DataDashboardSms{}

	for result.Next() {
		var id int
		var res string
		var date time.Time
		var smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms int
		var total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms int

		result.Scan(&id,&date,&smsBlockedInten,&smsBlockedExten,&phoneBlockedIntenSms,&phoneBlockedExtenSms,&total5656Sms,&handled5656Sms,&verify5656Sms,&wrong5656Sms,&res)

		dataDashboardSms.AddData(models.DashboardSms{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			SmsBlockedInten:      smsBlockedInten,
			SmsBlockedExten:      smsBlockedExten,
			PhoneBlockedIntenSms: phoneBlockedIntenSms,
			PhoneBlockedExtenSms: phoneBlockedExtenSms,
			Total5656Sms:         total5656Sms,
			Handled5656Sms:       handled5656Sms,
			Verify5656Sms:        verify5656Sms,
			Wrong5656Sms:         wrong5656Sms,
			TelecomBrand:         res,
		})
	}
	tmpl.Execute(w, dataDashboardSms)
	models.Connection().Close()
})

var getAllDataCallBrandNameByDay = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	tmpl := template.Must(template.ParseFiles("./static/admin/adminCall.html"))

	telecomBrand := r.FormValue("telecomBrand")

	startDate := r.FormValue("startDate")

	tStart, _ := time.Parse(layoutDate,startDate)

	var conn *sql.Stmt
	var result *sql.Rows

	if(telecomBrand == "null" && startDate != ""){
		conn, _ = models.Connection().Prepare("SELECT id,date, phoneDoubtCall," +
			"phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall," +
			"total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand " +
			"from dataColletion where date = ? order by date desc")

		result, _ = conn.Query(tStart.Format("2006-04-02"))
	}else if(startDate == "" && telecomBrand != "null"){
		conn, _ = models.Connection().Prepare("SELECT id,date, phoneDoubtCall," +
			"phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall," +
			"total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand " +
			"from dataColletion where telecomBrand = ? order by date desc")

		result, _ = conn.Query(telecomBrand)
	}else if(telecomBrand == "null" && startDate == ""){
		conn, _ = models.Connection().Prepare("SELECT id,date, phoneDoubtCall," +
			"phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall," +
			"total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand " +
			"from dataColletion order by date desc")

		result, _ = conn.Query()
	}else if(telecomBrand != "null" && startDate != "") {
		conn, _ = models.Connection().Prepare("SELECT id,date, phoneDoubtCall," +
			"phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall," +
			"total5656Call,handled5656Call,verify5656Call,wrong5656Call,telecomBrand " +
			"from dataColletion where telecomBrand = ? and date = ? order by date desc")

		result, _ = conn.Query(telecomBrand,tStart.Format("2006-04-02"))
	}

	dataDashboardCall := models.DataDashboardCall{}

	for result.Next() {
		var id int
		var res string
		var date time.Time
		var phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall int
		var total5656Call,handled5656Call,verify5656Call,wrong5656Call int

		result.Scan(&id,&date,&phoneDoubtCall,&phoneBlockedCall,&phoneFromDoubtCall,&phoneFromBlockedCall,&total5656Call,&handled5656Call,&verify5656Call,&wrong5656Call,&res)

		dataDashboardCall.AddData(models.DashboardCall{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			PhoneDoubtCall:       phoneDoubtCall,
			PhoneBlockedCall:     phoneBlockedCall,
			PhoneFromDoubtCall:   phoneFromDoubtCall,
			PhoneFromBlockedCall: phoneFromBlockedCall,
			Total5656Call:        total5656Call,
			Handled5656Call:      handled5656Call,
			Verify5656Call:       verify5656Call,
			Wrong5656Call: wrong5656Call,
			TelecomBrand: res,
		})
	}

	tmpl.Execute(w, dataDashboardCall)
	models.Connection().Close()
})

//dowload Excel
func PrepareAndReturnExcel(brandName string) *excelize.File{
	result, err := models.Connection().Query("SELECT id,date,smsBlockedInten" +
		",smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms,total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms," +
		"phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall,total5656Call,handled5656Call,verify5656Call," +
		"wrong5656Call,telecomBrand from dataColletion where telecomBrand = ? order by date asc",brandName)

	var data []models.TenDinhDanh

	for result.Next() {
		var id int
		var date time.Time
		var res string
		var smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms int
		var total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms int
		var phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall int
		var total5656Call,handled5656Call,verify5656Call,wrong5656Call int

		result.Scan(&id,&date,&smsBlockedInten,&smsBlockedExten,&phoneBlockedIntenSms,&phoneBlockedExtenSms,&total5656Sms,&handled5656Sms,&verify5656Sms,&wrong5656Sms,
			&phoneDoubtCall,&phoneBlockedCall,&phoneFromDoubtCall,&phoneFromBlockedCall,&total5656Call,&handled5656Call,&verify5656Call,&wrong5656Call,&res)

		data = append(data,models.TenDinhDanh{
			Id:                   id,
			Date:                 date.Format("02-01-2006"),
			SmsBlockedInten:      smsBlockedInten,
			SmsBlockedExten:      smsBlockedExten,
			PhoneBlockedIntenSms: phoneBlockedIntenSms,
			PhoneBlockedExtenSms: phoneBlockedExtenSms,
			Total5656Sms:         total5656Sms,
			Handled5656Sms:       handled5656Sms,
			Verify5656Sms:        verify5656Sms,
			Wrong5656Sms:         wrong5656Sms,
			PhoneDoubtCall:       phoneDoubtCall,
			PhoneBlockedCall:     phoneBlockedCall,
			PhoneFromDoubtCall:   phoneFromDoubtCall,
			PhoneFromBlockedCall: phoneFromBlockedCall,
			Total5656Call:        total5656Call,
			Handled5656Call:      handled5656Call,
			Verify5656Call:       verify5656Call,
			Wrong5656Call: wrong5656Call,
			TelecomBrand: res,
		})
	}

	if err != nil {
		fmt.Println(err)
	}

	f := excelize.NewFile()

	sheet1Name := "Sheet1"

	style, err := f.NewStyle(`
                {
                        "alignment":{
								"horizontal":"center",
								"vertical":"center",
								"ident":1,
								"reading_order": 0,
								"relative_indent":1,
								"shrink_to_fit":true,
								"wrap_text":true
                        },
				"border":[
                                {
                                        "type":"bottom",
                                        "color":"000000",
                                        "style":1
                                },
								{
										"type":"left",
										"color":"000000",
										"style":1
								},
								{
										"type":"right",
										"color":"000000",
										"style":1
								},
								{
										"type":"top",
										"color":"000000",
										"style":1
								}
                        ],
                        "fill":{
                                        "type":"pattern",
                                        "color":[
                                                "#c6e0b4"
                                        ],
                                        "pattern":1
                                },
						"font":
    							{
      									"bold": true,
        								"family": "Times New Roman",
        								"size": 11
    							}
                }
	`)

	categories := map[string]string{
		"A1":"BÁO CÁO SỐ LIỆU TIN NHẮN RÁC, CUỘC GỌI RÁC CỦA " + strings.ToUpper(brandName),
		"A3": "Ngày báo cáo", "B3": "Tin nhắn rác","J3": "Cuộc gọi rác",
		"B4":"Số lượng thuê bao đã thực hiện ngăn chặn","D4":"Số lượng tin nhắn rác đã chặn","F4":"Phản ánh 5656","J4":"Số lượng thuê bao nghi ngờ",
		"K4":"Số lượng thuê bao đã thực hiện ngăn chặn",
		"L4": "Số cuộc gọi phát sinh từ thuê bao nghi ngờ", "M4":"Số cuộc gọi phát sinh từ các thuê bao đã chặn","N4":"Phản ánh 5656",
		"B5":"Nội mạng","C5":"Ngoại mạng","D5":"Nội mạng","E5":"Ngoại mạng","F5":"Tổng phản ánh","G5":"Đã xử lý","H5":"Đang xác minh","I5":"Phản ánh không hợp lệ",
		"N5":"Tổng phản ánh","O5":"Đã xử lý","P5":"Đang xác minh","Q5":"Phản ánh không hợp lệ",
	}

	for k, v := range categories {
		f.SetCellValue(sheet1Name, k, v)
		f.MergeCell(sheet1Name,"a1","q1")
		f.MergeCell(sheet1Name,"a3","a5")
		f.MergeCell(sheet1Name,"B4","C4")
		f.MergeCell(sheet1Name,"D4","E4")
		f.MergeCell(sheet1Name,"F4","I4")
		f.MergeCell(sheet1Name,"b3","i3")
		f.MergeCell(sheet1Name,"j3","q3")
		f.MergeCell(sheet1Name,"n4","q4")
		f.MergeCell(sheet1Name,"j4","j5")
		f.MergeCell(sheet1Name,"k4","k5")
		f.MergeCell(sheet1Name,"l4","l5")
		f.MergeCell(sheet1Name,"m4","m5")
		f.SetCellStyle(sheet1Name,"a1","q1",style)
		f.SetCellStyle(sheet1Name,"a3","q3",style)
		f.SetCellStyle(sheet1Name,"a4","q4",style)
		f.SetCellStyle(sheet1Name,"a5","q5",style)
		f.SetColWidth(sheet1Name,"A","A",15)
		f.SetColWidth(sheet1Name,"B","B",15)
		f.SetColWidth(sheet1Name,"C","C",15)
		f.SetColWidth(sheet1Name,"D","D",15)
		f.SetColWidth(sheet1Name,"E","E",15)
		f.SetColWidth(sheet1Name,"F","F",15)
		f.SetColWidth(sheet1Name,"G","G",15)
		f.SetColWidth(sheet1Name,"H","H",15)
		f.SetColWidth(sheet1Name,"I","I",15)
		f.SetColWidth(sheet1Name,"J","J",15)
		f.SetColWidth(sheet1Name,"K","K",15)
		f.SetColWidth(sheet1Name,"L","L",15)
		f.SetColWidth(sheet1Name,"M","M",15)
		f.SetColWidth(sheet1Name,"N","N",15)
		f.SetColWidth(sheet1Name,"O","O",15)
		f.SetColWidth(sheet1Name,"P","P",15)
		f.SetColWidth(sheet1Name,"Q","Q",15)
		f.SetRowHeight(sheet1Name,1,25)
		f.SetRowHeight(sheet1Name,2,8)
		f.SetRowHeight(sheet1Name,3,35)
		f.SetRowHeight(sheet1Name,4,60)
		f.SetRowHeight(sheet1Name,5,40)
	}

	for i, each := range data {
		f.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+6), each.Date)
		f.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+6), each.SmsBlockedInten)
		f.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+6), each.SmsBlockedExten)
		f.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+6), each.PhoneBlockedIntenSms)
		f.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+6), each.PhoneBlockedExtenSms)
		f.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+6), each.Total5656Sms)
		f.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+6), each.Handled5656Sms)
		f.SetCellValue(sheet1Name, fmt.Sprintf("H%d", i+6), each.Verify5656Sms)
		f.SetCellValue(sheet1Name, fmt.Sprintf("I%d", i+6), each.Wrong5656Sms)
		f.SetCellValue(sheet1Name, fmt.Sprintf("J%d", i+6), each.PhoneDoubtCall)
		f.SetCellValue(sheet1Name, fmt.Sprintf("K%d", i+6), each.PhoneBlockedCall)
		f.SetCellValue(sheet1Name, fmt.Sprintf("L%d", i+6), each.PhoneFromDoubtCall)
		f.SetCellValue(sheet1Name, fmt.Sprintf("M%d", i+6), each.PhoneFromBlockedCall)
		f.SetCellValue(sheet1Name, fmt.Sprintf("N%d", i+6), each.Total5656Call)
		f.SetCellValue(sheet1Name, fmt.Sprintf("O%d", i+6), each.Handled5656Call)
		f.SetCellValue(sheet1Name, fmt.Sprintf("P%d", i+6), each.Verify5656Call)
		f.SetCellValue(sheet1Name, fmt.Sprintf("Q%d", i+6), each.Wrong5656Call)
	}

	return f
}

func PrepareAndReturnExcelAll() *excelize.File{
	//result, err := models.Connection().Query("SELECT id,date,smsBlockedInten" +
	//	",smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms,total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms," +
	//	"phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall,total5656Call,handled5656Call,verify5656Call," +
	//	"wrong5656Call,telecomBrand from dataColletion order by date asc")
	//
	//var data []models.TenDinhDanh
	//
	//for result.Next() {
	//	var id int
	//	var date time.Time
	//	var res string
	//	var smsBlockedInten,smsBlockedExten,phoneBlockedIntenSms,phoneBlockedExtenSms int
	//	var total5656Sms,handled5656Sms,verify5656Sms,wrong5656Sms int
	//	var phoneDoubtCall,phoneBlockedCall,phoneFromDoubtCall,phoneFromBlockedCall int
	//	var total5656Call,handled5656Call,verify5656Call,wrong5656Call int
	//
	//	result.Scan(&id,&date,&smsBlockedInten,&smsBlockedExten,&phoneBlockedIntenSms,&phoneBlockedExtenSms,&total5656Sms,&handled5656Sms,&verify5656Sms,&wrong5656Sms,
	//		&phoneDoubtCall,&phoneBlockedCall,&phoneFromDoubtCall,&phoneFromBlockedCall,&total5656Call,&handled5656Call,&verify5656Call,&wrong5656Call,&res)
	//
	//	data = append(data,models.TenDinhDanh{
	//		Id:                   id,
	//		Date:                 date.Format("02-01-2006"),
	//		SmsBlockedInten:      smsBlockedInten,
	//		SmsBlockedExten:      smsBlockedExten,
	//		PhoneBlockedIntenSms: phoneBlockedIntenSms,
	//		PhoneBlockedExtenSms: phoneBlockedExtenSms,
	//		Total5656Sms:         total5656Sms,
	//		Handled5656Sms:       handled5656Sms,
	//		Verify5656Sms:        verify5656Sms,
	//		Wrong5656Sms:         wrong5656Sms,
	//		PhoneDoubtCall:       phoneDoubtCall,
	//		PhoneBlockedCall:     phoneBlockedCall,
	//		PhoneFromDoubtCall:   phoneFromDoubtCall,
	//		PhoneFromBlockedCall: phoneFromBlockedCall,
	//		Total5656Call:        total5656Call,
	//		Handled5656Call:      handled5656Call,
	//		Verify5656Call:       verify5656Call,
	//		Wrong5656Call: wrong5656Call,
	//		TelecomBrand: res,
	//	})
	//}
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	f := excelize.NewFile()

	sheet1Name := "Sheet1"

	style, _ := f.NewStyle(`
                {
                        "alignment":{
								"horizontal":"center",
								"vertical":"center",
								"ident":1,
								"reading_order": 0,
								"relative_indent":1,
								"shrink_to_fit":true,
								"wrap_text":true
                        },
				"border":[
                                {
                                        "type":"bottom",
                                        "color":"000000",
                                        "style":1
                                },
								{
										"type":"left",
										"color":"000000",
										"style":1
								},
								{
										"type":"right",
										"color":"000000",
										"style":1
								},
								{
										"type":"top",
										"color":"000000",
										"style":1
								}
                        ],
                        "fill":{
                                        "type":"pattern",
                                        "color":[
                                                "#c6e0b4"
                                        ],
                                        "pattern":1
                                },
						"font":
    							{
      									"bold": true,
        								"family": "Times New Roman",
        								"size": 11
    							}
                }
	`)

	stype2, _ := f.NewStyle(`{
				"border":[
                                {
                                        "type":"top",
                                        "color":"000000",
                                        "style":1
                                },
                                {
                                        "type":"bottom",
                                        "color":"000000",
                                        "style":1
                                }
                        ],
                        "fill":{
                                        "type":"pattern",
                                        "color":[
                                                "#c6e0b4"
                                        ],
                                        "pattern":1
                                },
						"font":
    							{
      									"bold": true,
        								"family": "Times New Roman",
        								"size": 14
    							}
                }`)

	report5656 := map[string]string{
		"B1": "Phản ánh 5656",
		"d2":"Mobifone","f2":"Vinaphone",
		"h2":"Viettel","j2":"Vietnamobile","l2":"Khác",
		"d3":"SMS","e3":"V","f3":"SMS","g3":"V","h3":"SMS","i3":"V",
		"j3":"SMS","k3":"V","l3":"SMS","m3":"V",
	}

	dataOwn := map[string]string{
		"N1":"Chặn tin nhắn rác",
		"N2":"Số lượng thuê bao đã thực hiện ngăn chặn","ac2":"Số lượng sms đã thực hiện ngăn chặn","ar2":"Số lượng phản ánh đã xử lý(trên đầu số 5656)",
		"n3":"Mobifone","q3":"Vinaphone","t3":"Viettel","w3":"Vietnamobile","z3":"Others",
		"ac3":"Mobifone","af3":"Vinaphone","ai3":"Viettel","al3":"Vietnamobile","ao3":"Others",
		"ar3":"Mobifone","av3":"Vinaphone","az3":"Viettel","bd3":"Vietnamobile","bh3":"Others",
		"Bl1":"Chặn cuộc gọi rác",
		"bl2":"Số thuê bao nghi ngờ",
		"bl3":"Mobifone","bm3":"Vinaphone","bn3":"Viettel","bo3":"Vietnamobile","bp3":"Others",
		"bq2":"Số lượng thuê bao đã thực hiện ngăn chặn",
		"bq3":"Mobifone","br":"Vinaphone","bs3":"Viettel","bt3":"Vietnamobile","bu3":"Others",
		"bv2":"Số lượng cuộc gọi phát sinh từ thuê bao nghi ngờ",
		"bv3":"Mobifone","bw3":"Vinaphone","bx3":"Viettel","by3":"Vietnamobile","bz3":"Others",
		"ca2":"Số lượng cuộc gọi phát sinh từ thuê bao đã chặn",
		"ca3":"Mobifone","cb3":"Vinaphone","cc3":"Viettel","cd3":"Vietnamobile","ce3":"Others",
		"cf2":"Số lượng phản ánh đã xử lý(trên đầu số 5656)",
		"cf3":"Mobifone","cj3":"Vinaphone","cn3":"Viettel","cr3":"Vietnamobile","cv3":"Others",
	}

	registerDNC := map[string]string{
		"CZ1":"Đăng ký DNC",
		"CZ2":"Đăng ký DNC","DE2":"Hủy DNC",
	}

	unknownBrandName := map[string]string{
		"DJ1":"Cấp tên định danh",
		"DJ2":"Hợp lệ(đã cấp)","DK2":"Không hợp lệ",
	}

	for k, v := range report5656 {
		f.SetCellValue(sheet1Name, k, v)
		f.MergeCell(sheet1Name,"B1","M1")
		f.MergeCell(sheet1Name,"D2","E2")
		f.MergeCell(sheet1Name,"F2","G2")
		f.MergeCell(sheet1Name,"H2","I2")
		f.MergeCell(sheet1Name,"J2","K2")
		f.MergeCell(sheet1Name,"L2","M2")
		f.SetCellStyle(sheet1Name,k,k,style)
	}

	for k, v := range dataOwn {
		f.SetCellValue(sheet1Name, k, v)
		f.MergeCell(sheet1Name,"N1","CV1")

		//MSG
		f.MergeCell(sheet1Name,"N3","p3")
		f.MergeCell(sheet1Name,"q3","s3")
		f.MergeCell(sheet1Name,"t3","v3")
		f.MergeCell(sheet1Name,"w3","y3")
		f.MergeCell(sheet1Name,"z3","ab3")
		f.MergeCell(sheet1Name,"N2","ab2")

		f.MergeCell(sheet1Name,"ac3","ae3")
		f.MergeCell(sheet1Name,"af3","ah3")
		f.MergeCell(sheet1Name,"ai3","ak3")
		f.MergeCell(sheet1Name,"al3","an3")
		f.MergeCell(sheet1Name,"ao3","aq3")
		f.MergeCell(sheet1Name,"ac2","aq2")

		f.MergeCell(sheet1Name,"ar3","au3")
		f.MergeCell(sheet1Name,"av3","ay3")
		f.MergeCell(sheet1Name,"az3","bc3")
		f.MergeCell(sheet1Name,"bd3","bg3")
		f.MergeCell(sheet1Name,"bh3","bk3")
		f.MergeCell(sheet1Name,"ar2","bk2")

		//CALL
		f.MergeCell(sheet1Name,"bl1","cy1")

		f.MergeCell(sheet1Name,"bl2","bp2")

		f.MergeCell(sheet1Name,"bq2","bu2")

		f.MergeCell(sheet1Name,"bv2","bz2")

		f.MergeCell(sheet1Name,"ca2","ce2")

		f.MergeCell(sheet1Name,"cf2","cy2")
		f.MergeCell(sheet1Name,"cf3","ci3")
		f.MergeCell(sheet1Name,"cj3","cm3")
		f.MergeCell(sheet1Name,"cn3","cq3")
		f.MergeCell(sheet1Name,"cr3","cu3")
		f.MergeCell(sheet1Name,"cv3","cy3")

		f.SetCellStyle(sheet1Name,k,k,stype2)
	}

	for k, v := range registerDNC{
		f.SetCellValue(sheet1Name, k, v)
		f.MergeCell(sheet1Name,"CZ1","DI1")
		f.MergeCell(sheet1Name,"CZ2","DD2")
		f.MergeCell(sheet1Name,"DE2","DI1")
		f.SetCellStyle(sheet1Name,k,k,stype2)
	}

	for k, v := range unknownBrandName{
		f.SetCellValue(sheet1Name, k, v)
		f.MergeCell(sheet1Name,"DJ1","DK1")
		f.SetCellStyle(sheet1Name,k,k,stype2)
	}

	//for i, each := range data {
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+6), each.Date)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+6), each.SmsBlockedInten)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+6), each.SmsBlockedExten)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+6), each.PhoneBlockedIntenSms)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+6), each.PhoneBlockedExtenSms)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+6), each.Total5656Sms)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+6), each.Handled5656Sms)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("H%d", i+6), each.Verify5656Sms)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("I%d", i+6), each.Wrong5656Sms)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("J%d", i+6), each.PhoneDoubtCall)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("K%d", i+6), each.PhoneBlockedCall)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("L%d", i+6), each.PhoneFromDoubtCall)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("M%d", i+6), each.PhoneFromBlockedCall)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("N%d", i+6), each.Total5656Call)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("O%d", i+6), each.Handled5656Call)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("P%d", i+6), each.Verify5656Call)
	//	f.SetCellValue(sheet1Name, fmt.Sprintf("Q%d", i+6), each.Wrong5656Call)
	//}

	return f
}

var downloadExcel = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	telecomBrand := r.FormValue("brandName")
	// Get the Excel file with the user input data
	file := PrepareAndReturnExcel(telecomBrand)

	// Set the headers necessary to get browsers to interpret the downloadable file
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename= " +time.Now().UTC().Format("data-20060102.xlsx"))
	w.Header().Set("File-Name", time.Now().UTC().Format("data-20060102.xlsx"))
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	_ = file.Write(w)
})

var downloadExcelAll = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
	// Get the Excel file with the user input data
	file := PrepareAndReturnExcelAll()

	// Set the headers necessary to get browsers to interpret the downloadable file
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename= " +time.Now().UTC().Format("data-20060102.xlsx"))
	w.Header().Set("File-Name", time.Now().UTC().Format("AllData-20060102.xlsx"))
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	_ = file.Write(w)
})

func main() {
	m := mux.NewRouter().StrictSlash(true)
	authEnforcer, err := casbin.NewEnforcerSafe("./access/auth_model.conf", "./access/policy.csv")

	if err != nil{
		log.Fatal(err)
	}

	//setup session store
	session := scs.New()
	session.IdleTimeout = 30*time.Minute
	session.Cookie.Persist = true
	session.Cookie.Secure = true

	users := createUsers()

	//api
	//create data from brandname
	m.Handle("/TelecomBrandName/Create", routes.CheckJwt(uploadDataFromBrandName)).Methods("POST")
	//get data from brandname
	m.Handle("/TelecomBrandName/GetData", routes.CheckJwt(getAllDataBrandName)).Methods("GET")
	//get data from xx/xx/xxxx to xx/xx/xxxx
	m.Handle("/TelecomBrandName/GetDataByDay",routes.CheckJwt(getAllDataByDay)).Methods("GET")
	//get data sms from xx/xx/xxxx to xx/xx/xxxx
	m.Handle("/TelecomBrandName/GetDataSmsByDay",routes.CheckJwt(getAllDataSmsBrandNameByDay)).Methods("GET")
	//get data call from xx/xx/xxxx to xx/xx/xxxx
	m.Handle("/TelecomBrandName/GetDataCallByDay",routes.CheckJwt(getAllDataCallBrandNameByDay)).Methods("GET")
	//dashboard sms from brandname
	m.Handle("/Dashboard/Sms", routes.CheckJwt(dashboardSms)).Methods("GET")
	//dashboard call from brandname
	m.Handle("/Dashboard/Call", routes.CheckJwt(dashboardCall)).Methods("GET")

	//mvt
	//get all data from adminu7
	m.HandleFunc("/Admin/GetData", getAllDataTelecom).Methods("GET")
	//dashboard admin sms
	m.Handle("/Admin/Dashboard/Sms", dashboardAllSms).Methods("GET")
	//dashboard admin call
	m.Handle("/Admin/Dashboard/Call", dashboardAllCall).Methods("GET")

	//byDay
	m.HandleFunc("/Admin/GetDataBrandByDay", getAllDataOneTelecomByDay).Methods("GET")

	m.HandleFunc("/Admin/GetDataSmsBrandByDay", getAllDataSmsBrandNameByDay).Methods("GET")

	m.HandleFunc("/Admin/GetAllDataCallBrandByDay", getAllDataCallBrandNameByDay).Methods("GET")

	//download
	m.HandleFunc("/Admin/Download/Excel", downloadExcel).Methods("GET")
	m.HandleFunc("/Admin/Download/All", downloadExcelAll).Methods("GET")

	//client register (not use)
	m.HandleFunc("/Auth/Register", Register).Methods("POST")
	//client login
	m.HandleFunc("/Auth/Login", Login).Methods("POST")
	// get static file
	m.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	m.HandleFunc("/", rootHandler)
	m.HandleFunc("/login",loginHandler)

	//authorization

	log.Fatal(http.ListenAndServe(":8080",session.LoadAndSave(routes.Authorizer(authEnforcer,users)(m))))
}

func createUsers() models.Users {
	users := models.Users{}
	users = append(users, models.User{Id: 1, Username: "Admin", Role: "admin"})
	users = append(users, models.User{Id: 2, Username: "Sabine", Role: "member"})
	users = append(users, models.User{Id: 3, Username: "Sepp", Role: "member"})
	return users
}