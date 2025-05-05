package models

import "github.com/beego/beego/v2/client/orm"

const TableName = "otp_code"

type OTPCode struct {
	Id     uint `orm:"auto"`
	Label  string
	User   string
	Secret string
	Algo   string
	Digits int
	Period int
}

func NewOTPCode(label, user, secret, algo string, digits, period int) *OTPCode {
	otp := &OTPCode{
		Label:  label,
		User:   user,
		Secret: secret,
		Algo:   algo,
		Digits: digits,
		Period: period,
	}

	if otp.Algo == "" {
		otp.Algo = "SHA1"
	}

	if otp.Digits == 0 {
		otp.Digits = 6
	}

	if otp.Period == 0 {
		otp.Period = 30
	}

	return otp
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
