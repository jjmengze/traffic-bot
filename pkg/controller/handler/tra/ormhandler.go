package tra

import (
	"github.com/jinzhu/gorm"
	"k8s.io/klog/v2"
)

type ORMHandler struct {
	*gorm.DB
}

type city struct {
	CityCode string `gorm:"primary_key"`
	Name     string
}

type station struct {
	CityCode    string `gorm:"foreignkey:CityCode"`
	Name        string
	StationCode string `gorm:"primary_key"`
}

func NewORMHandler() *ORMHandler {
	return &ORMHandler{}
}

func (o *ORMHandler) PutCity(data city) error {
	klog.Info("First time query TRA City")
	if o.NewRecord(data) {
		if err := o.Create(data).Error; err != nil {
			klog.Error("put TRA city info to db error :", err)
			return err
		}
	}
	return nil
}
