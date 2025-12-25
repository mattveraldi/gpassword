package file

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Path      string
	LineQueue datatypes.JSON
}
