package models

func GetAllGollies() ([]GolyModel, error) {
	gollies := make([]GolyModel, 0)
	tx := db.Find(&gollies)
	return gollies, tx.Error
}

func GetGoly(id uint64) (GolyModel, error) {
	goly := GolyModel{}
	tx := db.Where("id = ?", id).First(&goly)
	return goly, tx.Error
}

func CreateGoly(goly GolyModel) error {
	tx := db.Create(&goly)
	return tx.Error
}

func UpdateGoly(goly GolyModel) error {
	tx := db.Save(&goly)
	return tx.Error
}

func DeleteGoly(id uint64) error {
	tx := db.Unscoped().Delete(&GolyModel{}, id)
	return tx.Error
}

func FindByGolyUrl(url string) (GolyModel, error) {
	var goly GolyModel
	tx := db.Where("goly = ?", url).First(&goly)
	return goly, tx.Error
}
