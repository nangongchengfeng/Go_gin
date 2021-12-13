package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Debug().Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return

}

func GetTagTotal(maps interface{}) (count int) {
	db.Debug().Model(&Tag{}).Where(maps).Count(&count)
	return
}

//判断tag是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Debug().Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) bool{
	db.Debug().Create(&Tag {
		Name : name,
		State : state,
		CreatedBy : createdBy,
	})

	return true
}