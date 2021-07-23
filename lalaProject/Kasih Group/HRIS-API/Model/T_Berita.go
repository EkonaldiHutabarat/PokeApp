package Model

type T_Berita struct {
	ID            string `json:"ID" gorm:"DEFAULT:NEWID()"`
	Title         string `form:"Title" json:"Title" gorm:"column:Title"`
	News_Content  string `form:"News_Content" json:"News_Content" gorm:"column:News_Content"`
	Expired_Date  string `form:"Expired_Date" json:"Expired_Date" gorm:"column:Expired_Date"`
	Company_ID    string `form:"Company_ID" json:"Company_ID" gorm:"column:Company_ID"`
	Department_ID string `form:"Department_ID" json:"Department_ID" gorm:"column:Department_ID"`
	Division_ID   string `form:"Division_ID" json:"Division_ID" gorm:"column:Division_ID"`
	Unit_ID       string `form:"Unit_ID" json:"Unit_ID" gorm:"column:Unit_ID"`
	Insert_User   string `form:"Insert_User" json:"Insert_User" gorm:"column:Insert_User"`
	Insert_Date   string `form:"Insert_Date" json:"Insert_Date" gorm:"column:Insert_Date"`
	Update_User   string `form:"Update_User" json:"Update_User" gorm:"column:Update_User"`
	Update_Date   string `form:"Update_Date" json:"Update_Date" gorm:"column:Update_Date"`
}

func (T_Berita) TableName() string {
	return "T_Berita"
}
