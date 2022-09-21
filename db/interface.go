package db

// mz_admin
type Admin struct {
	ID        int    `gorm:"primaryKey;column:ad_id"`
	Name      string `gorm:"column:ad_name"`
	UID       string `gorm:"column:ad_uid"`
	Password  string `gorm:"column:ad_password"`
	Problem   string `gorm:"column:ad_problem"`
	Answer    string `gorm:"column:ad_answer"`
	Classname string `gorm:"column:ad_classname"`
	Phone     string `gorm:"column:ad_phone"`
	Weixin    string `gorm:"column:ad_wx"`
	Power     string `gorm:"column:ad_power"`
	Num       int    `gorm:"column:ad_reservenum"`
}

// mz_problem
type Problem struct {
	ProblemID int    `gorm:"primaryKey;column:problem_id"`
	Problem   string `gorm:"column:problem"`
	Answer    string `gorm:"column:answer"`
}

// mz_reserve
type Reserve struct {
	ID        int    `gorm:"primaryKey;column:id"`
	Name      string `gorm:"column:name"`
	UID       string `gorm:"column:uid"`
	Time      string `gorm:"column:time"`
	Phone     string `gorm:"column:phone"`
	Classname string `gorm:"column:classname"`
	Describe  string `gorm:"column:describe"`
	Weixin    string `gorm:"column:wx"`
	Statue    string `gorm:"column:statue"`
	AID       string `gorm:"column:aid"`
}
