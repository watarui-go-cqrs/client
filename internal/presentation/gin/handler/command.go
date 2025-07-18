package handler

import (
	"client/internal/presentation/gin/helper"

	"github.com/gin-gonic/gin"
)

type CommandHandler struct {
	helper *helper.CommandHelper
}

func NewCommandHandler(helper *helper.CommandHelper) *CommandHandler {
	return &CommandHandler{helper: helper}
}

// カテゴリ登録をする
// @tags カテゴリ
// @Summary 登録する
// @Description 新しいカテゴリを登録する
// @ID add-category
// @Accept application/json
// @Produce json
// @Param newcategory body pb.CategoryUpParam true "crud:1(更新種別),name(カテゴリ名)"
// @Success 200 {object} pb.Category
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /category/add [post]
func (h *CommandHandler) CreateCategory(ctx *gin.Context) {
	h.helper.CreateCategory(ctx)
}

// 商品カテゴリを変更する
// @tags カテゴリ
// @Summary 変更する
// @Description 指定された識別子のカテゴリを変更する
// @ID update-category
// @Accept application/json
// @Produce json
// @Param updatecategory body pb.CategoryUpParam true "crud:2(更新種別),id(カテゴリID),name(カテゴリ名)"
// @Success 200 {object} pb.Category
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /category/update [put]
func (h *CommandHandler) UpdateCategory(ctx *gin.Context) {
	h.helper.UpdateCategory(ctx)
}

// カテゴリを削除する
// @tags カテゴリ
// @Summary 削除する
// @Description 指定された識別子でカテゴリを削除する
// @ID delete-category
// @Accept application/json
// @Produce json
// @Param deletecategory body pb.CategoryUpParam true "crud:3(更新種別),id(カテゴリID)"
// @Success 200 {object} pb.Category
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /category/delete [delete]
func (h *CommandHandler) DeleteCategory(ctx *gin.Context) {
	h.helper.DeleteCategory(ctx)
}

// 商品を登録する
// @tags 商品
// @Summary 登録する
// @Description 新商品を登録する
// @ID add-product
// @Accept application/json
// @Produce json
// @Param newproduct body pb.ProductUpParam true "crud:1(更新種別),name(商品名),price(単価),categoryId(商品カテゴリ番号)"
// @Success 200 {object} pb.Product
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/add [post]
func (h *CommandHandler) CreateProduct(ctx *gin.Context) {
	h.helper.CreateProduct(ctx)
}

// 商品を変更する
// @tags 商品
// @Summary 変更する
// @Description 指定された識別子の商品を変更する
// @ID update-product
// @Accept application/json
// @Produce json
// @Param updateproduct body pb.ProductUpParam true "crud:2(更新種別),id(商品ID),name(商品名),price(単価),categoryId(商品カテゴリ番号)"
// @Success 200 {object} pb.Product
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/update [put]
func (h *CommandHandler) UpdateProduct(ctx *gin.Context) {
	h.helper.UpdateProduct(ctx)
}

// 商品を削除する
// @tags 商品
// @Summary 削除する
// @Description 指定された識別子で商品を削除する
// @ID delete-product
// @Accept application/json
// @Produce json
// @Param deleteproduct body pb.ProductUpParam true "crud:3(更新種別),id(商品ID)"
// @Success 200 {object} pb.Product
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/delete [delete]
func (h *CommandHandler) DeleteProduct(ctx *gin.Context) {
	h.helper.DeleteProduct(ctx)
}
