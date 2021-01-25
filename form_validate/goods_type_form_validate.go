package form_validate

type GoodsTypeForm struct {
	Id             int64     `form:"id"`
	GoodsId        int64     `form:"goods_id"`                         // 商品ID
	TypeKey        string    `form:"type_key"`   // 属性名称：如颜色，输入商品的人可以自定义
	TypeVale       string    `form:"type_vale"` // 属性文字：入库数据格式 ["白色", "蓝色", "黄色"]
	IsDispay       int8      `form:"is_dispay"`       // 0 不显示 1 显示
	IsCreate 	   int    	 `form:"_create"`
}
