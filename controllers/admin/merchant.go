package controllers

import (
	"ganji/form_validate"
	"ganji/global"
	"ganji/global/response"
	"ganji/services"
	"github.com/gookit/validate"
	"strconv"
	"strings"
)

type MerchantController struct {
	baseController
}

func (Self *MerchantController) Index() {
	var merchantService services.MerchantService
	data, pagination := merchantService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "merchant/index.html"
}

func (Self *MerchantController) Add() {
	Self.Layout = "public/base.html"
	Self.TplName = "merchant/add.html"
}


func (Self *MerchantController) Create() {
	var MerchantForm form_validate.MerchantForm
	var merchantService services.MerchantService
	if err := Self.ParseForm(&MerchantForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	v := validate.Struct(MerchantForm)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}
	//默认头像
	MerchantForm.Logo = "/static/admin/images/avatar.png"

	//添加管理员
	var adminUserService services.AdminUserService
	if adminUserService.IsExistName(strings.TrimSpace(MerchantForm.UserName), 0) {
		response.ErrorWithMessage("登陆账号已经存在", Self.Ctx)
	}

	insertId := merchantService.Create(&MerchantForm)
	url := global.URL_BACK

	userInsertId := adminUserService.Create(&form_validate.AdminUserForm{Username: MerchantForm.UserName,Nickname: MerchantForm.MerchantName,Password: "1qaz369*",MerchantId:insertId})

	if MerchantForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}
	if insertId > 0 && userInsertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *MerchantController) Edit() {
	id, _ := Self.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", Self.Ctx)
	}
	var merchantService services.MerchantService

	merchant := merchantService.GetMerchantById(id)
	if merchant == nil {
		response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
	}

	Self.Data["data"] = merchant
	Self.Layout = "public/base.html"
	Self.TplName = "merchant/edit.html"
}

func (Self *MerchantController) Update(){
	var merchantForm form_validate.MerchantForm
	if err := Self.ParseForm(&merchantForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}

	if merchantForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", Self.Ctx)
	}

	v := validate.Struct(merchantForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	//商家验重
	var merchantService services.MerchantService
	if merchantService.IsExistName(strings.TrimSpace(merchantForm.MerchantName), merchantForm.Id) {
		response.ErrorWithMessage("账号已经存在", Self.Ctx)
	}
	if merchantService.Update(&merchantForm) > 0 {
		response.Success(Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *MerchantController) Del() {
	idStr := Self.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		Self.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}
	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", Self.Ctx)
	}
	var merchantService services.MerchantService
	if merchantService.Del(idArr) > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

