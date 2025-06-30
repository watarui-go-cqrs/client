package helper

import (
	"client/internal/infrastructure/provider"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watarui-go-cqrs/pb/pb"
)

type CommandHelper struct {
	provider *provider.CommandClientProvider
}

func NewCommandHelper(provider *provider.CommandClientProvider) *CommandHelper {
	return &CommandHelper{provider: provider}
}

func (h *CommandHelper) CreateCategory(ctx *gin.Context) {
	var param pb.CategoryUpParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	if result, err := h.provider.Category.Create(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.Error)
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategory())
	}
}

func (h *CommandHelper) UpdateCategory(ctx *gin.Context) {
	var param pb.CategoryUpParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	if result, err := h.provider.Category.Update(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.Error)
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategory())
	}
}

func (h *CommandHelper) DeleteCategory(ctx *gin.Context) {
	var param pb.CategoryUpParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	if result, err := h.provider.Category.Delete(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.Error)
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategory())
	}
}

func (h *CommandHelper) CreateProduct(ctx *gin.Context) {
	var param pb.ProductUpParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	if result, err := h.provider.Product.Create(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.Error)
			return
		}
		ctx.JSON(http.StatusOK, result.GetProduct())
	}
}

func (h *CommandHelper) UpdateProduct(ctx *gin.Context) {
	var param pb.ProductUpParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	if result, err := h.provider.Product.Update(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.Error)
			return
		}
		ctx.JSON(http.StatusOK, result.GetProduct())
	}
}

func (h *CommandHelper) DeleteProduct(ctx *gin.Context) {
	var param pb.ProductUpParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	if result, err := h.provider.Product.Delete(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.Error)
			return
		}
		ctx.JSON(http.StatusOK, result.GetProduct())
	}
}
