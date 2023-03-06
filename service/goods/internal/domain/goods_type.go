package domain

// domain 下定义的结构体进行转换

type GoodsType struct {
	ID        int64
	Name      string
	TypeCode  string
	NameAlias string
	IsVirtual bool
	Desc      string
	Sort      int32
}
