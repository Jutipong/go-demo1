package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Message string
	// Meta   interface{}
	Data   interface{}
	
}

func RespondJSON(w *gin.Context, status int,message string ,  payload interface{}) {
	fmt.Println("status ", status)
	var res ResponseData

	res.Status = status

	if res.Status  == 200 {
		res.Message="success"
		
	}else{
		res.Message=message
	}


//	res.Message= IfThenElse(status == 200, "success", message)
	// fmt.Println(aa)
	//helpers.IfThenElse(status == 200, "success", "fail")
	//res.Meta = utils.ResponseMessage(status)
	res.Data = payload

	w.JSON(200, res)
}
