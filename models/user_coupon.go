package models

import (
	"ganji/common"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserCoupon struct {
	BaseModel
	Id          int64     `json:"id"`
	UserId      int64     `orm:"index" json:"user_id"`
	ConponName  string    `orm:"size(128);index" json:"conpon_name"`  // 优惠券名称
	IsUsed      int8      `orm:"default(0)" json:"is_used"`           // 0 未使用； 1:已经使用
	TotalAmount float64   `orm:"default(150);digits(22);decimals(8)" json:"total_amount"`
	StartTime   time.Time `orm:"auto_now_add;type(datetime);index" json:"start_time"`
	EndTime     time.Time `orm:"auto_now_add;type(datetime);index" json:"end_time"`
}

func (this *UserCoupon) TableName() string {
	return common.TableName("user_coupon")
}

func (this *UserCoupon) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *UserCoupon) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *UserCoupon) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func GetMyCoupon(user_id int64) (*UserCoupon, error) {
	var user_cp UserCoupon
	err := user_cp.Query().Filter("UserId", user_id).One(&user_cp)
	if err != nil {
		return nil, err
	}
	return &user_cp, nil
}

