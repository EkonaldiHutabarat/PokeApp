package Model

type T_Cuti struct {
	ID                 string `json:"ID" gorm:"DEFAULT:NEWID()"`
	Emp_NIK            string `form:"Emp_NIK" json:"Emp_NIK" gorm:"column:Emp_NIK"`
	Start_Date         string `form:"Start_Date" json:"Start_Date" gorm:"type:time"`
	End_Date           string `form:"End_Date" json:"End_Date" gorm:"type:time"`
	Length_Days        string `form:"Length_Days" json:"Length_Days" gorm:"column:Length_Days"`
	Description        string `form:"Description" json:"Description" gorm:"column:Description"`
	Note               string `form:"Note" json:"Note" gorm:"column:Note"`
	AppByHead          string `form:"AppByHead" json:"AppByHead" gorm:"column:AppByHead"`
	Type_Approve_Head  string `form:"Type_Approve_Head" json:"Type_Approve_Head" gorm:"column:Type_Approve_Head"`
	Done_Head          string `form:"Done_Head" json:"Done_Head" gorm:"column:Done_Head"`
	Approve_Head_Date  string `form:"Approve_Head_Date" json:"Approve_Head_Date" gorm:"type:time"`
	AppByHead2         string `form:"AppByHead2" json:"AppByHead2" gorm:"column:AppByHead2"`
	Type_Approve_Head2 string `form:"Type_Approve_Head2" json:"Type_Approve_Head2" gorm:"column:Type_Approve_Head2"`
	Done_Head2         string `form:"Done_Head2" json:"Done_Head2" gorm:"column:Done_Head2"`
	Approve_Head_Date2 string `form:"Approve_Head_Date2" json:"Approve_Head_Date2" gorm:"type:time"`
	AppByHRD           string `form:"AppByHRD" json:"AppByHRD" gorm:"column:AppByHRD"`
	Type_Approve_HRD   string `form:"Type_Approve_HRD" json:"Type_Approve_HRD" gorm:"column:Type_Approve_HRD"`
	Done_HRD           string `form:"Done_HRD" json:"Done_HRD" gorm:"column:Done_HRD"`
	Approve_HRD_Date   string `form:"Approve_HRD_Date" json:"Approve_HRD_Date" gorm:"type:time"`
	Insert_User        string `form:"Insert_User" json:"Insert_User" gorm:"column:Insert_User"`
	Insert_Date        string `form:"Insert_Date" json:"Insert_Date" gorm:"type:time"`
	Update_User        string `form:"Update_User" json:"Update_User" gorm:"column:Update_User"`
	Update_Date        string `form:"Update_Date" json:"Update_Date" gorm:"type:time"`
	Cuti_File          string `form:"Cuti_File" json:"Cuti_File" gorm:"column:Cuti_File"`
	SisaCuti           string `form:"SisaCuti" json:"SisaCuti" gorm:"column:SisaCuti"`
}

func (T_Cuti) TableName() string {
	return "T_Cuti"
}
