package unittest

import (
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestDbInit(t *testing.T) {
	dsn := "root:root123@tcp(tx.tk93.top:13306)/shop_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 传递可变参数
	if err := db.AutoMigrate(modelList()...); err != nil {
		panic("failed to migrate database")
	}
}

func modelList() []interface{} {
	return []interface{}{
		&User{},
		&Category{},
		&GoodsType{},

		&SpecificationsAttr{},
		&SpecificationsAttrValue{},
	}
}

type User struct {
	ID          int64      `gorm:"primarykey"`
	Mobile      string     `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号码，用户唯一标识';not null"`
	Password    string     `gorm:"type:varchar(100);not null "` // 用户密码的保存需要注意是否加密
	NickName    string     `gorm:"type:varchar(25) comment '用户昵称'"`
	Birthday    *time.Time `gorm:"column:birthday;default:'2006-01-02 15:04:05';type:datetime comment '出生日日期'"`
	Gender      string     `gorm:"column:gender;default:male;type:varchar(16) comment 'female:女,male:男'"`
	Role        int        `gorm:"column:role;default:1;type:int comment '1:普通用户，2:管理员'"`
	CreatedAt   time.Time  `gorm:"column:add_time"`
	UpdatedAt   time.Time  `gorm:"column:update_time"`
	DeletedAt   gorm.DeletedAt
	IsDeletedAt bool
}

type Category struct {
	ID               int32          `gorm:"primarykey;type:int" json:"id"`
	Name             string         `gorm:"type:varchar(50);not null;comment:分类名称" json:"name"`
	ParentCategoryID int32          `json:"parent_id"`
	ParentCategory   *Category      `json:"-"`
	SubCategory      []*Category    `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32          `gorm:"column:level;default:1;not null;type:int;comment:分类的级别" json:"level"`
	IsTab            bool           `gorm:"comment:是否显示;default:false" json:"is_tab"`
	Sort             int32          `gorm:"comment:分类排序;default:99;not null;type:int" json:"sort"`
	CreatedAt        time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

type GoodsType struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null;comment:商品类型名称" json:"name"`
	TypeCode  string         `gorm:"type:varchar(50);not null;comment:商品类型编码" json:"type_code"`
	NameAlias string         `gorm:"type:varchar(50);not null;comment:商品类型别名" json:"name_alias"`
	IsVirtual bool           `gorm:"comment:是否是虚拟商品显示;default:false" json:"is_virtual"`
	Desc      string         `gorm:"type:varchar(50);not null;comment:商品类型描述" json:"desc"`
	Sort      int32          `gorm:"comment:类型排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// SpecificationsAttr 规格参数信息表
type SpecificationsAttr struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	TypeID    int64          `gorm:"index:type_id;type:int;comment:商品类型ID;not null"`
	Name      string         `gorm:"type:varchar(250);not null;comment:规格参数名称" json:"name"`
	Sort      int32          `gorm:"comment:规格排序;default:99;not null;type:int" json:"sort"`
	Status    bool           `gorm:"comment:参数状态;default:false" json:"status"`
	IsSKU     bool           `gorm:"comment:是否通用的SKU持有;default:false" json:"is_sku"`
	IsSelect  bool           `gorm:"comment:是否可查询;default:false" json:"is_select"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// SpecificationsAttrValue 规格参数信息选项表
type SpecificationsAttrValue struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	AttrId    int64          `gorm:"index:attr_id;type:int;comment:规格ID;not null"`
	Value     string         `gorm:"type:varchar(250);not null;comment:规格参数信息值" json:"value"`
	Sort      int32          `gorm:"comment:规格参数值排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}