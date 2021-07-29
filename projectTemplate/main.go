package main

import (
	"github.com/otiai10/copy"
	"github.com/pepelazz/projectBlueprint/projectTemplate/docs/legalEntity"
	"github.com/pepelazz/projectBlueprint/projectTemplate/utils"
	"github.com/pepelazz/nla_framework"
	t "github.com/pepelazz/nla_framework/types"
	"os"
)

func main() {
	projectGenerator.Start(getProject(), nil)
}

func getProject() t.ProjectType {
	p := &t.ProjectType{
		Name: "CompanyName",
	}
	p.Config.Vue.QuasarVersion = 2

	p.Docs = []t.DocType{
		legalEntity.GetDoc(p),
	}
	// названием базы маленькими буквами, без пробелов
	p.Config.Postgres = t.PostrgesConfig{"db_name", 5646, "xvzDV4curLidx8IWZJ6czDHQ1qa7wjfL", "Asia/Novosibirsk"}
	p.Config.WebServer = t.WebServerConfig{3091, "https://example.ru", "/home/deploy/projectName", "85.210.890.567", "root"}
	// TODO: надо прописать настройки почтового сервера для отправки email
	p.Config.Email = t.EmailConfig{Sender: "info@mail.ru", Password: "password", Host: "smtp.mail.ru", Port: 465, SenderName: "CompanyName"}
	p.Config.Logo = "https://cdn.pixabay.com/photo/2017/05/05/00/15/kokopelli-2285538_960_720.png"
	// формируем routes для Vue
	p.FillVueBaseRoutes()
	p.Vue.UiAppName = "CompanyName"

	// боковое меню для Vue
	p.Vue.Menu = []t.VueMenu{
		//{DocName: "client_order"},
		{Url: "users", Text: "Пользователи", Icon: "https://image.flaticon.com/icons/svg/423/423063.svg", Roles: []string{utils.RoleAdmin}},
		{DocName: "legal_entity"},
		{Text: "Справочники", Icon: "https://image.flaticon.com/icons/svg/1643/1643260.svg", IsFolder: true, LinkList: []t.VueMenu{{DocName: "legal_entity"}}},
	}
	p.FillSideMenu()

	// копируем файлы проекта (которые не шаблоны)
	if _, err := os.Stat("./sourceFiles"); !os.IsNotExist(err) {
		err := copy.Copy("./sourceFiles", "../")
		utils.CheckErr(err, "Copy sourceFiles")
	}

	return *p
}
