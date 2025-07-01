package helper

import (
	"client/internal/infrastructure/provider"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type QueryHelper struct {
	provider *provider.QueryClientProvider
}

func NewQueryHelper(provider *provider.QueryClientProvider) *QueryHelper {
	return &QueryHelper{provider: provider}
}

func (h *QueryHelper) CategoryList(ctx *gin.Context) {
	param := &emptypb.Empty{}
	if result, err := h.provider.Category.List(ctx, param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.GetError())
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategories())
	}
}

func (h *QueryHelper) ProductList(ctx *gin.Context) {
	param := &emptypb.Empty{}
	if result, err := h.provider.Product.List(ctx, param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.GetError())
			return
		}
		ctx.JSON(http.StatusOK, result.GetProducts())
	}
}

func (h *QueryHelper) ProductListStream(ctx *gin.Context) {
	param := &emptypb.Empty{}
	stream, err := h.provider.Product.ListStream(ctx, param)
	if err != nil {
		ErrResp(ctx, err)
		return
	}

	products := make([]*pb.Product, 0)
	for {
		product, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			ErrResp(ctx, err)
			return
		}
		products = append(products, product)
	}
	ctx.JSON(http.StatusOK, products)
}

func (h *QueryHelper) CategoryById(ctx *gin.Context) {
	param := &pb.CategoryParam{Id: ctx.Param("id")}
	if result, err := h.provider.Category.ById(ctx, param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.GetError())
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategory())
	}
}

func (h *QueryHelper) ProductById(ctx *gin.Context) {
	param := &pb.ProductParam{Id: ctx.Param("id")}
	if result, err := h.provider.Product.ById(ctx, param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.GetError())
			return
		}
		ctx.JSON(http.StatusOK, result.GetProduct())
	}
}

func (h *QueryHelper) ProductByKeyword(ctx *gin.Context) {
	param := &pb.ProductParam{Keyword: ctx.Param("keyword")}
	if result, err := h.provider.Product.ByKeyword(ctx, param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.GetError())
			return
		}
		ctx.JSON(http.StatusOK, result.GetProducts())
	}
}
