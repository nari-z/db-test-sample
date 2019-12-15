package main

import (
	"testing"
	
	"github.com/jinzhu/gorm"
	_ "github.com/proullon/ramsql/driver"
)

func createMockRamSQL() (*gorm.DB, error) {
	db, err := gorm.Open("ramsql", "MockRamSql")
	if err != nil {
		return nil, err
	}

	// register mock data.
	db = db.AutoMigrate(&SampleModel{})
	mockData := []*SampleModel {
		&SampleModel{ ID: 1, Name: "Model1" },
		&SampleModel{ ID: 2, Name: "Model2" },
		&SampleModel{ ID: 3, Name: "Model3" },
		&SampleModel{ ID: 4, Name: "Model4" },
		&SampleModel{ ID: 5, Name: "Model5" },
	}
	for _, d := range mockData {
		err := db.Create(d).Error
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func TestCreateSampleModel(t *testing.T) {
	db, err := createMockRamSQL()
	if err != nil {
		t.Fatalf("createMockRamSQL : %s\n", err)
	}
	defer db.Close()
	
	r := SampleRepository{ DB: db }

	// create new model.
	newModleName := "Model6"
	err = r.CreateModel(newModleName)
	if err != nil {
		t.Fatalf("SampleRepository.CreateModel : %s\n", err)
	}

	// insert check
	var lastModel SampleModel
	// TODO: WHERE, FIRST, LASTがうまく使えない。要調査。
	err = db.Debug().Order("id asc").Find(&lastModel).Error
	if err != nil {
		t.Fatalf("SampleRepository.CreateModel : %s\n", err)
	}
	if lastModel.Name != newModleName{
		t.Fatalf("SampleRepository.CreateModel : SampleModel.Name is mismatch.")
	}
}

// func TestGetSampleModel(t *testing.T) {
// 	// TODO: error "Syntax error near ~"
// 	Model1, err := r.GetModel(1)
// 	if err != nil {
// 		t.Fatalf("SampleRepository.GetModel : %s\n", err)
// 	}
// 	t.Logf(Model1.Name)
// }