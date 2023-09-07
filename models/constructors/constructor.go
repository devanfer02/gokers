package constructor

import (
	"github.com/devanfer02/gokers/models"
	"github.com/devanfer02/gokers/helpers"
	"github.com/google/uuid"
)

func NewKRS(studentId, classId uuid.UUID, sks uint) *models.KRS {
	krs := &models.KRS{}
	krs.StudentID = studentId
	krs.ClassID = classId
	krs.Semester = helpers.GetCurrentSemester()
	krs.Grade = "K"
	krs.Sks = sks 

	return krs
}

func NewKrsDetail(id uuid.UUID, studentId uuid.UUID) *models.KrsDetail {
	krsDtl := &models.KrsDetail{}
	krsDtl.ID = id
	krsDtl.StudentID = studentId
	krsDtl.MaxSks = 24
	return krsDtl
}

func NewClass(courseId, lecturerId uuid.UUID, className, classRoom string, quota, current uint) *models.Class {
	class := &models.Class{}

	class.CourseID = courseId
	class.LecturerID = lecturerId
	class.ClassName = className
	class.ClassRoom = classRoom
	class.Quota = quota 
	class.Current = current 

	return class;
}