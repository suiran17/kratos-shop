package domain

type Specification struct {
	ID                 int64
	TypeID             int64
	Name               string
	Sort               int32
	Status             bool
	IsSKU              bool
	IsSelect           bool
	SpecificationValue []*SpecificationValue
}

type SpecificationValue struct {
	ID     int64
	AttrId int64
	Value  string
	Sort   int32
}

// IsTypeIDEmpty 新增判断类型不能为空的方法
func (b *Specification) IsTypeIDEmpty() bool {
	return b.TypeID == 0
}

// IsValueEmpty 规格下面的参数不能为空的方法
func (b *Specification) IsValueEmpty() bool {
	return b.SpecificationValue == nil
}

type SpecificationList []*Specification

func (p SpecificationList) FindById(id int64) *Specification {
	for _, item := range p {
		if item.ID == id {
			return item
		}
	}
	return nil
}
