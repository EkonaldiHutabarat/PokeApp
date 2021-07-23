package Model

type M_Ijin_Type struct {
	ID              string `json:"ID" gorm:"column:ID"`
	Reason          string `form:"Reason" json:"Reason" gorm:"column:Reason"`
	isPersonal      string `form:"isPersonal" json:"isPersonal" gorm:"column:isPersonal"`
	Active          string `form:"Active" json:"Active" gorm:"column:Active"`
	isDifferentDays string `form:"isDifferentDays" json:"isDifferentDays" gorm:"column:isDifferentDays"`
}

func (M_Ijin_Type) TableName() string {
	return "M_Ijin_Type"
}
