package Model

type T_Ijin struct {
	ID                 string `json:"ID" gorm:"DEFAULT:NEWID()"`
	Emp_NIK            string `form:"Emp_NIK" json:"Emp_NIK" gorm:"column:Emp_NIK"`
	Date               string `form:"Date" json:"Date" gorm:"column:Date"`
	Start_Time         string `form:"Start_Time" json:"Start_Time" gorm:"column:Start_Time"`
	End_Time           string `form:"End_Time" json:"End_Time" gorm:"column:End_Time"`
	Length             string `form:"Length" json:"Length" gorm:"column:Length"`
	Description        string `form:"Description" json:"Description" gorm:"column:Description"`
	Reason             string `form:"Reason" json:"Reason" gorm:"column:Reason"`
	AppByHead          string `form:"AppByHead" json:"AppByHead" gorm:"column:AppByHead"`
	Type_Approve_Head  string `form:"Type_Approve_Head" json:"Type_Approve_Head" gorm:"column:Type_Approve_Head"`
	Done_Head          string `form:"Done_Head" json:"Done_Head" gorm:"column:Done_Head"`
	Approve_Head_Date  string `form:"Approve_Head_Date" json:"Approve_Head_Date" gorm:"column:Approve_Head_Date"`
	AppByHead2         string `form:"AppByHead2" json:"AppByHead2" gorm:"column:AppByHead2"`
	Type_Approve_Head2 string `form:"Type_Approve_Head2" json:"Type_Approve_Head2" gorm:"column:Type_Approve_Head2"`
	Done_Head2         string `form:"Done_Head2" json:"Done_Head2" gorm:"column:Done_Head2"`
	Approve_Head_Date2 string `form:"Approve_Head_Date2" json:"Approve_Head_Date2" gorm:"column:Approve_Head_Date2"`
	AppByHRD           string `form:"AppByHRD" json:"AppByHRD" gorm:"column:AppByHRD"`
	Type_Approve_HRD   string `form:"Type_Approve_HRD" json:"Type_Approve_HRD" gorm:"column:Type_Approve_HRD"`
	Done_HRD           string `form:"Done_HRD" json:"Done_HRD" gorm:"column:Done_HRD"`
	Approve_HRD_Date   string `form:"Approve_HRD_Date" json:"Approve_HRD_Date" gorm:"column:Approve_HRD_Date"`
	Ijin_File          string `form:"Ijin_File" json:"Ijin_File" gorm:"column:Ijin_File"`
	Insert_User        string `form:"Insert_User" json:"Insert_User" gorm:"column:Insert_User"`
	Insert_Date        string `form:"Insert_Date" json:"Insert_Date" gorm:"column:Insert_Date"`
	Update_User        string `form:"Update_User" json:"Update_User" gorm:"column:Update_User"`
	Update_Date        string `form:"Update_Date" json:"Update_Date" gorm:"column:Update_Date"`
	End_Date           string `form:"End_Date" json:"End_Date" gorm:"column:End_Date"`
	Length_Days        string `form:"Length_Days" json:"Length_Days" gorm:"column:Length_Days"`
}

func (T_Ijin) TableName() string {
	return "T_Ijin"
}

// type uniqueidentifier interface {
// 	GormDataType() string
// }
