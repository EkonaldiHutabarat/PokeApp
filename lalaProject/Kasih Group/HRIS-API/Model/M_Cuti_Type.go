package Model

type M_Cuti_Type struct {
	ID     string `json:"ID" gorm:"column:ID"`
	Reason string `form:"Reason" json:"Reason" gorm:"column:Reason"`
	Days   string `form:"Days" json:"Days" gorm:"column:Days"`
	Active string `form:"Active" json:"Active" gorm:"column:Active"`
}

func (M_Cuti_Type) TableName() string {
	return "M_Cuti_Type"
}
