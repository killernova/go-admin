package models

func GetUserTable() (table GlobalTable) {

	// 列显示配置
	table.Info.FieldList = []FieldStruct{
		{
			Head:     "姓名",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(value string) string {
				return value
			},
		},
		{
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			ExcuFun: func(value string) string {
				if value == "1" {
					return "男"
				}
				if value == "2" {
					return "女"
				}
				return "未知"
			},
		},
	}

	table.Info.Table = "users"
	table.Info.Title = "用户表"
	table.Info.Description = "用户表"

	// 表单显示配置
	table.Form.FormList = []FormStruct{
		{
			Head:     "姓名",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "default",
		}, {
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			Default:  "",
			Editable: true,
			FormType: "text",
		},
	}

	table.Form.Table = "users"
	table.Form.Title = "用户表"
	table.Form.Description = "用户表"

	return
}
