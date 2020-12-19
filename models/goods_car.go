package models

import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsCar struct {
	BaseModel
	Id           int64      `orm:"column(id);auto;size(11)" json:"id" form:"id"`
	GoodsId      int64      `orm:"default(1)" json:"goods_id"`  // 商品ID
	Logo         string     `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"logo" form:"logo"` // 商品LOGO
	GoodsTitle   string     `orm:"size(64)" json:"goods_title"`                              // 商品标题
	GoodsName    string     `orm:"size(512);index" json:"goods_name" form:"goods_name"`      // 产品名称
	UserId       int64      `orm:"size(64);index" json:"user_id"`                            // 购买用户
	BuyNums      int64      `orm:"default(0)" json:"buy_nums"`                               // 购买数量
	PayAmount    float64    `orm:"default(0);digits(22);decimals(8)" json:"pay_amount"`      // 支付金额
}

func (this *GoodsCar) TableName() string {
	return common.TableName("goods_car")
}


func (this *GoodsCar) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsCar) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCar) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCar) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsCar) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

