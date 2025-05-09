package models

import (
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/creachadair/otp/otpauth"
)

const TableName = "otp_code"

type OTPCode struct {
	Id      uint `orm:"auto"`
	Issuer  string
	Account string
	Secret  string
	Algo    string
	Digits  int
	Period  int
}

func OTPCodeFromURL(url string) (*OTPCode, error) {
	url_obj, err := otpauth.ParseURL(url)
	if err != nil {
		return nil, err
	}

	otp := &OTPCode{
		Issuer:  url_obj.Issuer,
		Account: url_obj.Account,
		Secret:  strings.ToUpper(url_obj.RawSecret), // base32 requires upper case
		Algo:    url_obj.Algorithm,
		Digits:  url_obj.Digits,
		Period:  url_obj.Period,
	}
	return otp, nil
}

func (o *OTPCode) TableName() string {
	return TableName
}

func (otp *OTPCode) SaveToDB() error {
	_, err := orm.NewOrm().Insert(otp)
	return err
}

func GetOtpById(id uint) (*OTPCode, error) {
	otp := OTPCode{Id: id}
	o := orm.NewOrm()
	if err := o.Read(&otp); err != nil {
		return nil, err
	}
	return &otp, nil
}
