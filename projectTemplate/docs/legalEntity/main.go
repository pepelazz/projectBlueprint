package legalEntity

import (
	t "github.com/pepelazz/projectGenerator/types"
	"github.com/pepelazz/projectGenerator/utils"
)

const (
	name            = "legal_entity"
	name_ru         = "юрЛица"
	name_ru_plural  = "юрЛица"
	menu_icon       = "https://image.flaticon.com/icons/svg/469/469323.svg"
	breadcrumb_icon = "far fa-file-alt"
)

func GetDoc() t.DocType {
	doc := t.DocType{
		Name:       name,
		NameRu:     name_ru,
		PathPrefix: "docs",
		Flds: []t.FldType{
			t.GetFldTitle(),
			t.GetFldString("inn", "ИНН", 30, [][]int{{2, 1}}, "col-4"),
			t.GetFldString("kpp", "КПП", 30, [][]int{{2, 2}}, "col-4"),
			t.GetFldString("type", "тип организации", 50, [][]int{{2, 2}}),
			t.GetFldString("address_legal", "юр адрес", 0,   [][]int{{3, 1}}, "col-8"),
		},
		Vue: t.DocVue{
			RouteName:      name,
			MenuIcon:       menu_icon,
			BreadcrumbIcon: breadcrumb_icon,
			Roles:          []string{},
		},
		//Templates:   map[string]*t.DocTemplate{"webClient_item.vue": {},},
		IsBaseTemapltes: t.DocIsBaseTemapltes{true, true},
		Sql: t.DocSql{
			IsSearchText: true,
			IsBeforeTrigger: true,
			IsAfterTrigger: true,
		},
	}
	// создаем стандартные методы sql "list", "update", "get_by_id" с возможностью ограничения по ролям
	doc.Sql.FillBaseMethods(doc.Name)

	doc.Vue.I18n = map[string]string{
		"listTitle": utils.UpperCaseFirst(name_ru_plural),
		"listDeletedTitle": "Удаленные " + name_ru_plural,
	}

	doc.Init()

	return doc
}
