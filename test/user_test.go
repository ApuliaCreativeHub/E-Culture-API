package test

import (
	"E-Culture-API/models"
	"gorm.io/gorm/utils"
	"testing"
)

func TestCreate(t *testing.T) {
	u := models.User{Name: "Eugenio", Surname: "Motta", IsACurator: true, Email: "eugenio.motta@economia.it", Password: "IoMotta1!"}
	err := u.Create()
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}

func TestReadByEmail(t *testing.T) {
	u := models.User{IsACurator: true, Email: "eugenio.motta@economia.it"}
	err := u.ReadByEmail()
	utils.AssertEqual(models.User{Name: "Eugenio", Surname: "Motta", IsACurator: true, Email: "eugenio.motta@economia.it", Password: "IoMotta1!"}, u)
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}

func TestUpdate(t *testing.T) {
	u := models.User{IsACurator: true, Email: "eugenio.motta@economia.it"}
	err := u.ReadByEmail()
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
	u.Email = "eugenio.motta@economia.com"
	err = u.Update()
	utils.AssertEqual(models.User{Name: "Eugenio", Surname: "Motta", IsACurator: true, Email: "eugenio.motta@economia.com", Password: "IoMotta1!"}, u)
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}

func TestDelete(t *testing.T) {
	u := models.User{IsACurator: true, Email: "eugenio.motta@economia.com"}
	err := u.ReadByEmail()
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
	err = u.Delete()
	if err != nil {
		t.Errorf("err should be null but is %v", err)
	}
}
