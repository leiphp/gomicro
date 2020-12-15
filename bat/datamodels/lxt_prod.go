package datamodels

//商品表
type LxtProd struct {
	ID           	int64   `gorm:"primary_key" json:"id"` //ID
	UserId       	int64   `json:"user_id"`               //上传用户id
	ProdSn       	string  `json:"prod_sn"`               //商品编号
	Platform     	string  `json:"platform"`              //客户端平台
	Name 			string  `json:"name"`      			   //商品名称
	Price           float64 `json:"price"`                 //商品价格
	CreateTime      int64   `json:"create_time"`           //创建时间
	Status          int     `json:"status"`                //0待上架，1上架，2下架, 3清仓
	Remark          string  `json:"remark"`                //备注
}

//返回表名
func (this LxtProd) TableName() string {
	return "lxt_prod"
}