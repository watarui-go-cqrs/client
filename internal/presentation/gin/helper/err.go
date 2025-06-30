package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrResp(ctx *gin.Context, err any) {
	if pbErr, ok := err.(*pb.Error); ok {
		if pbErr.Type == "Domain Error" || pbErr.Type == "CRUD Error" {
			ctx.JSON(http.StatusBadRequest, pbErr)
		} else {
			ctx.JSON(http.StatusInternalServerError, pbErr)
		}
	} else {
		other, _ := err.(error)
		if status, ok := status.FromError(other); ok {
			if status.Code() == codes.InvalidArgument {
				ctx.JSON(http.StatusBadRequest,
					pb.Error{Type: "Validate Error", Message: status.Message()})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError,
				pb.Error{Type: "InterError Error", Message: other.Error()})
		}
	}
}
