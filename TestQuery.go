package main

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main() {
	addMainCategory()
}

type MainCategory struct {
	MainCatName string
	SubCatId    []string
}

func addMainCategory() error {

	subCat := []MainCategory{
		MainCategory{
			"ไลฟ์สไตล์",
			[]string{"LIF01", "LIF02", "LIF03", "LIF04", "LIF05"},
		},
		MainCategory{
			"อาหารและเครื่องดื่ม",
			[]string{"BEV01", "BEV02", "BEV03", "BEV04", "BEV05", "BEV06"},
		},
		MainCategory{
			"ช๊อปปิ้ง",
			[]string{"SHO01", "SHO02", "SHO03", "SHO04", "SHO05", "SHO06", "SHO07", "SHO08", "SHO09"},
		},
		MainCategory{
			"ท่องเที่ยวและโรงแรม",
			[]string{"TRA01", "TRA02", "TRA03", "TRA04", "TRA05", "TRA06", "TRA07", "TRA08", "TRA09"},
		},
		MainCategory{
			"สุขภาพและความงาม",
			[]string{"HEL01", "HEL02", "HEL03", "HEL04", "HEL05", "HEL06"},
		},
		MainCategory{
			"การศึกษา",
			[]string{"EDU01", "EDU02", "EDU03"},
		},
		MainCategory{
			"สัตว์เลี้ยง",
			[]string{"PET01", "PET02", "PET03", "PET04"},
		},
		MainCategory{
			"ร้านค้าอื่นๆ",
			[]string{"EXT01"},
		},
		MainCategory{
			"ร้านค้าประเภท Retail",
			[]string{"RETAIL"},
		},
	}

	for _, mainCat := range subCat {
		condition := bson.M{
			"category_id" : bson.M{"$in" : mainCat.SubCatId },
		}
		fmt.Println(condition)
	}
	return nil
}
