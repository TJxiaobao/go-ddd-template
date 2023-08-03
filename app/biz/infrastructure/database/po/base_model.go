package po

import "time"

type BaseModel struct {
	Id        int64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:gmt_create" json:"-"`
	UpdatedAt time.Time `gorm:"column:gmt_modify" json:"-"`
	IsDeleted int64     `gorm:"column:is_deleted" json:"-"`
}
