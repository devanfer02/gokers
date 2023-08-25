package services

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	Db *configs.Database
}

func NewAuthService(db *configs.Database) *AuthService {
	return &AuthService{Db: db}
}

func (authSvc *AuthService) RegisterStudent(student *models.Student) res.Response {
	if _, err := govalidator.ValidateStruct(student); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}
	
	count := authSvc.Db.Count("email = ?", models.Student{}, student.Email)

	if count > 0 {
		return res.CreateResponseErr(status.Conflict, "email already in use", fmt.Errorf("email is already used"))
	}

	hashed, err := helpers.HashPassword(student.Password)

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	student.Password = hashed
	student.Nim, err = helpers.CreateNIM(student.Major, student.Faculty, student.Entrance, authSvc.Db)
	student.ID 		 = helpers.GenerateUUID()

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	krsDtl := models.NewKrsDetail(helpers.GenerateUUID(), student.ID)

	if err = authSvc.Db.Create(student); err != nil {
		return res.CreateResponseErr(status.Conflict, "internal server error", err)
	}

	if err = authSvc.Db.Create(&krsDtl); err != nil {
		return res.CreateResponseErr(status.Conflict, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "student registered to system", gin.H {
		"nim": student.Nim,
	})
}

func (authSvc *AuthService) LoginStudent(stdAuth *models.StudentAuth) (res.Response, string) {
	if _, err := govalidator.ValidateStruct(stdAuth); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err), ""
	}

	var student models.Student

	if err := authSvc.Db.FindFirst("nim = ?", &student, stdAuth.Nim); err != nil {
		return res.CreateResponseErr(status.Forbidden, "invalid nim or password", nil), ""
	}

	if err := helpers.CompPassword(&student.Password, &stdAuth.Password); err != nil {
		return res.CreateResponseErr(status.Forbidden, "invalid nim or password", nil), ""
	}

	token, err := helpers.GetTokenStr(student.ID, "sub")

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err), ""
	}

	return res.CreateResponse(status.Accepted, "student successfully login", nil), token
}

func (authSvc *AuthService) RegisterLecturer(lecturer *models.Lecturer) res.Response {
	if _, err := govalidator.ValidateStruct(lecturer); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	count := authSvc.Db.Count("email = ?", models.Lecturer{}, lecturer.Email)

	if count > 0 {
		return res.CreateResponseErr(status.Conflict, "email already in use", fmt.Errorf("email is already used"))
	}

	hashed, err := helpers.HashPassword(lecturer.Password)

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	lecturer.Password  = hashed
	lecturer.Ndn, err  = helpers.CreateNDN(lecturer.Major, lecturer.Faculty, authSvc.Db)
	lecturer.ID 	   = helpers.GenerateUUID()

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	if err = authSvc.Db.Create(lecturer); err != nil {
		return res.CreateResponseErr(status.Conflict, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "student registered to system", gin.H {
		"ndn": lecturer.Ndn,
	})
}

func (authSvc *AuthService) RegisterAdmin(admin *models.Admin) res.Response {
	if _, err := govalidator.ValidateStruct(admin); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	count := authSvc.Db.Count("username = ?", models.Admin{}, admin.Username)

	if count > 0 {
		return res.CreateResponseErr(status.Conflict, "username already in use", fmt.Errorf("username is already used"))		
	}

	hashed, err := helpers.HashPassword(admin.Password)

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)		
	}

	admin.Password 	= hashed
	admin.ID 		= helpers.GenerateUUID()

	if err = authSvc.Db.Create(&admin); err != nil {
		return res.CreateResponseErr(status.Conflict, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "admin registered to system", nil)
}

func (authSvc *AuthService) LoginAdmin(adminAuth *models.Admin) (res.Response, string) {
	if _, err := govalidator.ValidateStruct(adminAuth); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err), ""
	}

	var admin models.Admin

	if err := authSvc.Db.FindFirst("username = ?", &admin, adminAuth.Username); err != nil {
		return res.CreateResponseErr(status.Forbidden, "invalid username or email or password", nil), ""
	}

	if err := helpers.CompPassword(&admin.Password, &adminAuth.Password); err != nil {
		return res.CreateResponseErr(status.Forbidden, "invalid username or email or password", nil), ""
	}

	if admin.Email != adminAuth.Email {
		return res.CreateResponseErr(status.Forbidden, "invalid username or email or password", nil), ""
	}

	token, err := helpers.GetTokenStr(admin.ID, "adm")

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err), ""
	}

	return res.CreateResponse(status.Accepted, "admin successfully login", nil), token
}	