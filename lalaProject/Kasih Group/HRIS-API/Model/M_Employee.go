package Model

type M_Employee struct {
	Emp_ID                   string `json:"ID" gorm:"DEFAULT:NEWID()"`
	Emp_NIK                  string `form:"Emp_NIK" json:"Emp_NIK" gorm:"column:Emp_NIK"`
	Password                 string `form:"Password" json:"Password" gorm:"column:Password"`
	IsResetPassword          string `form:"IsResetPassword" json:"IsResetPassword" gorm:"column:IsResetPassword"`
	Gender_ID                string `form:"Gender_ID" json:"Gender_ID" gorm:"column:Gender_ID"`
	Religion_ID              string `form:"Religion_ID" json:"Religion_ID" gorm:"column:Religion_ID"`
	Company_ID               string `form:"Company_ID" json:"Company_ID" gorm:"column:Company_ID"`
	Date_Of_Hire             string `form:"Date_Of_Hire" json:"Date_Of_Hire" gorm:"column:Date_Of_Hire"`
	Emp_Fullname             string `form:"Emp_Fullname" json:"Emp_Fullname" gorm:"column:Emp_Fullname"`
	Emp_NickName             string `form:"Emp_NickName" json:"Emp_NickName" gorm:"column:Emp_NickName"`
	Emp_Address_KTP          string `form:"Emp_Address_KTP" json:"Emp_Address_KTP" gorm:"column:Emp_Address_KTP"`
	Emp_Address_Domisili     string `form:"Emp_Address_Domisili" json:"Emp_Address_Domisili" gorm:"column:Emp_Address_Domisili"`
	Emp_City                 string `form:"Emp_City" json:"Emp_City" gorm:"column:Emp_City"`
	Emp_State                string `form:"Emp_State" json:"Emp_State" gorm:"column:Emp_State"`
	Emp_Identity_Card        string `form:"Emp_Identity_Card" json:"Emp_Identity_Card" gorm:"column:Emp_Identity_Card"`
	Emp_Zip_Code             string `form:"Emp_Zip_Code" json:"Emp_Zip_Code" gorm:"column:Emp_Zip_Code"`
	Emp_BirthDate            string `form:"Emp_BirthDate" json:"Emp_BirthDate" gorm:"column:Emp_BirthDate"`
	Emp_BirthPlace           string `form:"Emp_BirthPlace" json:"Emp_BirthPlace" gorm:"column:Emp_BirthPlace"`
	Emp_Marital_Status       string `form:"Emp_Marital_Status" json:"Emp_Marital_Status" gorm:"column:Emp_Marital_Status"`
	Emp_Phone_Home           string `form:"Emp_Phone_Home" json:"Emp_Phone_Home" gorm:"column:Emp_Phone_Home"`
	Emp_HP_Num               string `form:"Emp_HP_Num" json:"Emp_HP_Num" gorm:"column:Emp_HP_Num"`
	Emp_BPJS_Kesehatan       string `form:"Emp_BPJS_Kesehatan" json:"Emp_BPJS_Kesehatan" gorm:"column:Emp_BPJS_Kesehatan"`
	Emp_BPJS_Ketenagakerjaan string `form:"Emp_BPJS_Ketenagakerjaan" json:"Emp_BPJS_Ketenagakerjaan" gorm:"column:Emp_BPJS_Ketenagakerjaan"`
	Image_Url                string `form:"Image_Url" json:"Image_Url" gorm:"column:Image_Url"`
	NPWP_Flag                string `form:"NPWP_Flag" json:"NPWP_Flag" gorm:"column:NPWP_Flag"`
	Emp_NPWP_Number          string `form:"Emp_NPWP_Number" json:"Emp_NPWP_Number" gorm:"column:Emp_NPWP_Number"`
	Emp_NPWP_Name            string `form:"Emp_NPWP_Name" json:"Emp_NPWP_Name" gorm:"column:Emp_NPWP_Name"`
	Emp_NPWP_Address         string `form:"Emp_NPWP_Address" json:"Emp_NPWP_Address" gorm:"column:Emp_NPWP_Address"`
	Emp_NPWP_File            string `form:"Emp_NPWP_File" json:"Emp_NPWP_File" gorm:"column:Emp_NPWP_File"`
	Emp_NPWP_Status          string `form:"Emp_NPWP_Status" json:"Emp_NPWP_Status" gorm:"column:Emp_NPWP_Status"`
	Status_Schedule          string `form:"Status_Schedule" json:"Status_Schedule" gorm:"column:Status_Schedule"`
	Emp_Active               string `form:"Emp_Active" json:"Emp_Active" gorm:"column:Emp_Active"`
	Emp_NonActiveDate        string `form:"Emp_NonActiveDate" json:"Emp_NonActiveDate" gorm:"column:Emp_NonActiveDate"`
	TotalCuti                string `form:"TotalCuti" json:"TotalCuti" gorm:"column:TotalCuti"`
	Insert_Usert             string `form:"Insert_Usert" json:"Insert_Usert" gorm:"column:Insert_Usert"`
	Insert_Date              string `form:"Insert_Date" json:"Insert_Date" gorm:"column:Insert_Date"`
	Update_User              string `form:"Update_User" json:"Update_User" gorm:"column:Update_User"`
	Update_Date              string `form:"Update_Date" json:"Update_Date" gorm:"column:Update_Date"`
	Status_Baching           string `form:"Status_Baching" json:"Status_Baching" gorm:"column:Status_Baching"`
	Baching_Date             string `form:"Baching_Date" json:"Baching_Date" gorm:"column:Baching_Date"`
}

func (M_Employee) TableName() string {
	return "M_Employee"
}
