package netfile

type  FileObject struct {
	Name     string `json:"name" bson:"name"`
	FullName string `json:"-" bson:"full_name"`
	FullPath string `json:"-" bson:"full_path"`
	Rights   string `json:"rights" bson:"rights"`
	Size     string `json:"size" bson:"size"`
	Ext      string `json:"-" bson:"ext"`
	Date     string `json:"date" bson:"date"`
	FileType string `json:"type" bson:"file_type"`
	Content  []byte `json:"-" bson:"content"`
	Status   int8 `json:"status" bson:"status"`
}

type NetFile interface {
	Create(filename string, data []byte)
	Read()
	Write()
	Delete()
	Rename()
	Move()
	List(path string)
}