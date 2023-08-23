package services

import (
	//"fmt"

	//"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"github.com/devanfer02/gokers/configs"
	//"github.com/devanfer02/gokers/helpers"
	//"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type KrsService struct {
	Db *configs.Database
}

func (krsSvc *KrsService) GetKrs(krs []models.KRS, params ...interface{}) res.Response {
	if err := krsSvc.Db.FindAllCondition("student_id = ?", &krs, params); err != nil {
		return res.CreateResponseErr(
			status.ServerError, 
			"internal server error", 
			err, 
		)
	}

	return res.CreateResponse(status.Ok, "succesfully fetch krs", krs)
}

func (krsSvc *KrsService) AddClass(krs models.KRS, params ...interface{}) res.Response {
	if _, err := govalidator.ValidateStruct(krs); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	var krsTmp models.KRS 

	if err := krsSvc.Db.PreloadChainByPK([]string{"Class","Course"}, &krsTmp, krs.ClassID); err != nil {
		return res.CreateResponseErr(
			status.ServerError, 
			"internal server error", 
			err, 
		)		
	}

	// var krsDb []models.KRS 

	// if err := krsSvc.Db.FindAllCondition("student_id = ?", &krs, params); err != nil {
	// 	return res.CreateResponseErr(
	// 		status.ServerError, 
	// 		"internal server error", 
	// 		err, 
	// 	)
	// }

	//totalSks := helpers.GetTotalSks(krsDb)

	return res.Response{}
}

func (krsSvc *KrsService) RemoveClass(krs models.KRS) res.Response {
	return res.Response{}
}