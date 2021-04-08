package models

type Login  struct {
	Username string `json:"Account", db:"Account"`
	Password string `json:"Password", db:"Passwd"`
}

type User struct {
	Id int `json:"Id"`
	Username string `json:"Username"`
	Role string `json:"Role"`
}

type LoginSuccess struct {
	Token string `json:"Token"`
	Success bool `json:"Success"`
}

type LoginFail struct {
	Msg string `json:Message`
	Success bool `json:"Success"`
}

type Authorized struct{
	Msg string `json:"Message"`
	Author bool `json:"Authorized"`
}

type CreateDataSuccess struct{
	Msg string `json:"Message"`
	Success bool `json:"Success"`
}

type ApiDataTenDinhDanh struct{
	Authen bool `json:"Authentication"`
	Data DataTenDinhDanh `json:"AllData"`
	Success bool `json:"Success"`
}

type ApiDataDashboardSms struct {
	Authen bool `json:"Authentication"`
	Data DataDashboardSms `json:"DataSms"`
	Success bool `json:"Success"`
}

type ApiDataDashboardCall struct {
	Authen bool `json:"Authentication"`
	Data DataDashboardCall `json:"DataCall"`
	Success bool `json:"Success"`
}

type DataTenDinhDanh struct{
	Data []TenDinhDanh `json:"Data"`
}

type DataDashboardSms struct {
	Data []DashboardSms `json:"DataSms"`
}

type DataDashboardCall struct {
	Data []DashboardCall `json:"DataCall"`
}

type Users []User

type TenDinhDanh struct{
	Id int `json:"id"`
	Date string `json:"date"`
	SmsBlockedInten int `json:"smsBlockedInten"`
	SmsBlockedExten int `json:"smsBlockedExten"`
	PhoneBlockedIntenSms int `json:"phoneBlockedIntenSms"`
	PhoneBlockedExtenSms int `json:"phoneBlockedExtenSms"`
	Total5656Sms int `json:"total5656Sms"`
	Handled5656Sms int `json:"handled5656Sms"`
	Verify5656Sms int `json:"verify5656Sms"`
	Wrong5656Sms int `json:"wrong5656Sms"`
	PhoneDoubtCall int `json:"phoneDoubtCall"`
	PhoneBlockedCall int `json:"phoneBlockedCall"`
	PhoneFromDoubtCall int `json:"phoneFromDoubtCall"`
	PhoneFromBlockedCall int `json:"phoneFromBlockedCall"`
	Total5656Call int `json:"total5656Call"`
	Handled5656Call int `json:"handled5656Call"`
	Verify5656Call int `json:"verify5656Call"`
	Wrong5656Call int `json:"wrong5656Call"`
	TelecomBrand string `json:"telecomBrand"`
}

type DashboardSms struct{
	Id int `json:"id"`
	Date string `json:"date"`
	SmsBlockedInten int `json:"smsBlockedInten"`
	SmsBlockedExten int `json:"smsBlockedExten"`
	PhoneBlockedIntenSms int `json:"phoneBlockedIntenSms"`
	PhoneBlockedExtenSms int `json:"phoneBlockedExtenSms"`
	Total5656Sms int `json:"total5656Sms"`
	Handled5656Sms int `json:"handled5656Sms"`
	Verify5656Sms int `json:"verify5656Sms"`
	Wrong5656Sms int `json:"wrong5656Sms"`
	TelecomBrand string `json:"telecomBrand"`
}

type DashboardCall struct{
	Id int `json:"id"`
	Date string `json:"date"`
	PhoneDoubtCall int `json:"phoneDoubtCall"`
	PhoneBlockedCall int `json:"phoneBlockedCall"`
	PhoneFromDoubtCall int `json:"phoneFromDoubtCall"`
	PhoneFromBlockedCall int `json:"phoneFromBlockedCall"`
	Total5656Call int `json:"total5656Call"`
	Handled5656Call int `json:"handled5656Call"`
	Verify5656Call int `json:"verify5656Call"`
	Wrong5656Call int `json:"wrong5656Call"`
	TelecomBrand string `json:"telecomBrand"`
}

func (d *DataTenDinhDanh) AddData(data TenDinhDanh){
	d.Data = append(d.Data,data)
}

func (d *DataDashboardSms) AddData(data DashboardSms){
	d.Data = append(d.Data,data)
}

func (d *DataDashboardCall) AddData(data DashboardCall){
	d.Data = append(d.Data,data)
}