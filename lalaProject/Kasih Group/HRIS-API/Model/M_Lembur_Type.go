package Model

type M_Lembur_Type struct {
	ID     string `json:"ID" gorm:"column:ID"`
	Reason string `form:"Reason" json:"Reason" gorm:"column:Reason"`
	Active string `form:"Active" json:"Active" gorm:"column:Active"`
}

func (M_Lembur_Type) TableName() string {
	return "M_Lembur_Type"
}
