package main

import (
	"github.com/otiai10/copy"
	"github.com/pepelazz/projectBlueprint/projectTemplate/docs/legalEntity"
	"github.com/pepelazz/projectBlueprint/projectTemplate/utils"
	"github.com/pepelazz/projectGenerator"
	t "github.com/pepelazz/projectGenerator/types"
)

func main() {
	projectGenerator.Start(getProject(), nil)
}

func getProject() t.ProjectType {
	p := &t.ProjectType{
		Name: "CompanyName",
		Docs: []t.DocType {
			legalEntity.GetDoc(),
		},
		Config: t.ProjectConfig{LocalProjectPath: "github.com/pepelazz/projectBlueprint/src"},
	}
	// названием базы маленькими буквами, без пробелов
	p.Config.Postgres = t.PostrgesConfig{"db_name", 5646, "ktulhu77", "Asia/Novosibirsk"}
	p.Config.WebServer = t.WebServerConfig{ 3091, "https://example.ru", "/home/deploy/projectName", "85.210.890.567", "root"}
	// TODO: надо прописать настройки почтового сервера для отправки email
	//Например, p.Config.Email = types.EmailConfig{"info@mail.ru", "password", "smtp.mail.ru", 465, "CompanyName"}
	p.Config.Email = t.EmailConfig{ "CompanyName"}
	p.Config.Logo = "https://cdn.pixabay.com/photo/2017/05/05/00/15/kokopelli-2285538_960_720.png"
	// формируем routes для Vue
	p.FillVueBaseRoutes()
	p.Vue.UiAppName = "CompanyName"

	// боковое меню для Vue
	p.Vue.Menu = []t.VueMenu{
		//{DocName: "client_order"},
		{Url: "users", Text: "Пользователи", Icon: "https://image.flaticon.com/icons/svg/423/423063.svg", Roles: []string{utils.RoleAdmin}},
		{DocName: "legal_entity"},
		{Text: "Справочники", Icon: "https://image.flaticon.com/icons/svg/1643/1643260.svg", IsFolder:true, LinkList:[]t.VueMenu{{DocName: "legal_entity"}}},
		{Text: "Задачи", Icon: "https://image.flaticon.com/icons/svg/1642/1642808.svg", IsFolder:true, LinkList:[]t.VueMenu{{Text: "Список задач", Url: "task"}, {Text: "Типы задач", Url: "taskType"}}},
	}
	p.FillSideMenu()

	// копируем файлы проекта (которые не шаблоны)
	err := copy.Copy("./sourceFiles", "../")
	utils.CheckErr(err, "Copy sourceFiles")

	return *p
}