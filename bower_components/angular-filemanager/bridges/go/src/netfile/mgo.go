package netfile

import (
	"gopkg.in/mgo.v2"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"path/filepath"
	"time"
	"util"
	"strings"
	"os"
	"fmt"
)

/*
配置文件
config  [db]
    name  [table]
    	{文件数据} [row]
模版
templates
    name
    	{文件数据}
缓存
cache
    siteid_cache
   	{文件数据}
*/
var (
	SESSION_IS_NIL = errors.New("mgo.Session is nil")
	COLLATION_IS_NIL = errors.New("mgo.Collation is nil")
	TABLE_NAME_IS_NIL = errors.New("mgo.Collation table is nil")
)

const (
	DIR = "dir"
	FILE = "file"
)

type MongodbFile struct {
	sess      *mgo.Session `json:"-"`
	collation *mgo.Collection `json:"-"`
	table     string

	FileObject
}

func NewMongodbFile(sess *mgo.Session, db, table string) (*MongodbFile, error) {
	mf := new(MongodbFile)
	mf.sess = sess
	if mf.sess == nil {
		return nil, SESSION_IS_NIL
	}
	mf.collation = mf.sess.DB(db).C(table)
	if mf.collation == nil {
		return nil, COLLATION_IS_NIL
	}

	return mf, nil

}

//创建 文件
func (mf *MongodbFile)Create(path string, data []byte) error {
	fo := FileObject{}
	ext := filepath.Ext(path)
	fo.FullPath = filepath.Dir(path)
	fo.FullName = filepath.Clean(path)
	fo.Name = strings.Replace(filepath.Clean(path), fo.FullPath + string(os.PathSeparator), "", -1)
	fo.Content = data
	size := len(data)
	if size == 0 {
		fo.FileType = DIR
	} else {
		fo.FileType = FILE
		fo.Ext = ext
	}
	fo.Date = time.Now().Format("2006-01-02 15:04:05")
	fo.Size = util.ToStr(len(data))
	fo.Status = 1

	if strings.HasPrefix(path, "/config") {
		tableName := util.GetSecondPathName(path)
		if tableName == "" {
			return TABLE_NAME_IS_NIL
		}
		tables, err := mf.sess.DB("config").CollectionNames()
		if err != nil {
			fmt.Println(err)
			return err
		}
		hasTable := false
		for _, v := range tables {
			if v == tableName {
				hasTable = true
				break
			}
		}
		//第二级的新建
		if tableName == fo.Name && fo.FileType == DIR {
			mf.collation = mf.sess.DB("config").C(tableName) //.Create(new(mgo.CollectionInfo))
			err = mf.collation.Create(new(mgo.CollectionInfo))
			if err != nil {
				fmt.Println(err)
				return err
			}
			idx := mgo.Index{}
			idx.Key = []string{"full_name"}
			idx.Unique = true
			err = mf.collation.EnsureIndex(idx)
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else {
			mf.collation = mf.sess.DB("config").C(tableName)
			_, err = mf.collation.Upsert(bson.M{"full_name": fo.FullName}, &fo)
			//添加index
			if !hasTable && err == nil {
				idx := mgo.Index{}
				idx.Key = []string{"full_name"}
				idx.Unique = true
				mf.collation.EnsureIndex(idx)
			}
		}
	} else if strings.HasPrefix(path, "/templates") {
		tableName := util.GetSecondPathName(path)
		if tableName == "" {
			return TABLE_NAME_IS_NIL
		}
		tables, err := mf.sess.DB("templates").CollectionNames()
		if err != nil {
			fmt.Println(err)
			return err
		}
		hasTable := false
		for _, v := range tables {
			if v == tableName {
				hasTable = true
				break
			}
		}
		if tableName == fo.Name && fo.FileType == DIR {
			mf.collation = mf.sess.DB("templates").C(tableName) //.Create(new(mgo.CollectionInfo))
			err = mf.collation.Create(new(mgo.CollectionInfo))
			if err != nil {
				fmt.Println(err)
				return err
			}
			idx := mgo.Index{}
			idx.Key = []string{"full_name"}
			idx.Unique = true
			err = mf.collation.EnsureIndex(idx)
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else {
			mf.collation = mf.sess.DB("templates").C(tableName)
			_, err = mf.collation.Upsert(bson.M{"full_name": fo.FullName}, &fo)
			//添加index
			if !hasTable && err == nil {
				idx := mgo.Index{}
				idx.Key = []string{"full_name"}
				idx.Unique = true
				mf.collation.EnsureIndex(idx)
			}
		}
	} else if strings.HasPrefix(path, "/cache") {
		tableName := util.GetSecondPathName(path)
		if tableName == "" {
			return TABLE_NAME_IS_NIL
		}
		tables, err := mf.sess.DB("cache").CollectionNames()
		if err != nil {
			fmt.Println(err)
			return err
		}
		hasTable := false
		for _, v := range tables {
			if v == tableName {
				hasTable = true
				break
			}
		}
		if tableName == fo.Name && fo.FileType == DIR {
			mf.collation = mf.sess.DB("cache").C(tableName) //.Create(new(mgo.CollectionInfo))
			err = mf.collation.Create(new(mgo.CollectionInfo))
			if err != nil {
				fmt.Println(err)
				return err
			}
			idx := mgo.Index{}
			idx.Key = []string{"full_name"}
			idx.Unique = true
			err = mf.collation.EnsureIndex(idx)
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else {
			mf.collation = mf.sess.DB("cache").C(tableName)
			_, err = mf.collation.Upsert(bson.M{"full_name": fo.FullName}, &fo)
			//添加index
			if !hasTable && err == nil {
				idx := mgo.Index{}
				idx.Key = []string{"full_name"}
				idx.Unique = true
				mf.collation.EnsureIndex(idx)
			}
		}
	}

	return nil
}

//3个库
func (mf *MongodbFile)List(path string) (*[]FileObject, error) {
	/*
	配置文件
	config  [db]
	name  [table]
	{
		文件数据
	} [row]
	模版
	templates
	name
	{
		文件数据
	}
	缓存
	cache
	siteid_cache
	{
		文件数据
	}*/
	fs := make([]FileObject, 0)
	if strings.HasPrefix(path, "/config") {
		tableName := util.GetSecondPathName(path)
		if tableName == "" {
			tables, err := mf.sess.DB("config").CollectionNames()
			if err != nil {
				fmt.Println(err)
			}
			for _, v := range tables {
				if v == "system.indexes" {
					continue
				}
				f1 := FileObject{}
				f1.Name = v
				f1.Rights = "drwxr-xr-x"
				f1.Size = "0"
				f1.Date = "2017-03-03 15:31:40"
				f1.FileType = "dir"
				fs = append(fs, f1)
			}
		} else {
			mf.collation = mf.sess.DB("config").C(tableName)
			mf.collation.Find(bson.M{"full_path": path}).All(&fs)
		}
	} else if strings.HasPrefix(path, "/templates") {
		tableName := util.GetSecondPathName(path)
		if tableName == "" {
			tables, err := mf.sess.DB("templates").CollectionNames()
			if err != nil {
				fmt.Println(err)
			}
			for _, v := range tables {
				if v == "system.indexes" {
					continue
				}
				f1 := FileObject{}
				f1.Name = v
				f1.Rights = "drwxr-xr-x"
				f1.Size = "0"
				f1.Date = "2017-03-03 15:31:40"
				f1.FileType = "dir"
				fs = append(fs, f1)
			}
		} else {
			mf.collation = mf.sess.DB("templates").C(tableName)
			mf.collation.Find(bson.M{"full_path": path}).All(&fs)
		}
	} else if strings.HasPrefix(path, "/cache") {
		tableName := util.GetSecondPathName(path)
		if tableName == "" {
			tables, err := mf.sess.DB("cache").CollectionNames()
			if err != nil {
				fmt.Println(err)
			}
			for _, v := range tables {
				if v == "system.indexes" {
					continue
				}
				f1 := FileObject{}
				f1.Name = v
				f1.Rights = "drwxr-xr-x"
				f1.Size = "0"
				f1.Date = "2017-03-03 15:31:40"
				f1.FileType = "dir"
				fs = append(fs, f1)
			}
		} else {
			mf.collation = mf.sess.DB("cache").C(tableName)
			mf.collation.Find(bson.M{"full_path": path}).All(&fs)
		}
	} else {
		if path == "/" {
			f1 := FileObject{}
			f1.Name = "config"
			f1.Rights = "drwxr-xr-x"
			f1.Size = "0"
			f1.Date = "2017-03-03 15:31:40"
			f1.FileType = "dir"
			fs = append(fs, f1)

			f2 := FileObject{}
			f2.Name = "templates"
			f2.Rights = "drwxr-xr-x"
			f2.Size = "0"
			f2.Date = "2017-03-03 15:31:40"
			f2.FileType = "dir"
			fs = append(fs, f2)

			f3 := FileObject{}
			f3.Name = "cache"
			f3.Rights = "drwxr-xr-x"
			f3.Size = "0"
			f3.Date = "2017-03-03 15:31:40"
			f3.FileType = "dir"
			fs = append(fs, f3)
		}
	}
	return &fs, nil
}

func (mf *MongodbFile)Read(path string) (string, error) {
	fs := FileObject{}
	if strings.HasPrefix(path, "/config") {
		tableName := util.GetSecondPathName(path)
		mf.collation = mf.sess.DB("config").C(tableName)
		err := mf.collation.Find(bson.M{"full_name": path}).One(&fs)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	} else if strings.HasPrefix(path, "/templates") {
		tableName := util.GetSecondPathName(path)
		mf.collation = mf.sess.DB("templates").C(tableName)
		err := mf.collation.Find(bson.M{"full_name": path}).One(&fs)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	} else if strings.HasPrefix(path, "/cache") {
		tableName := util.GetSecondPathName(path)
		mf.collation = mf.sess.DB("cache").C(tableName)
		err := mf.collation.Find(bson.M{"full_name": path}).One(&fs)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	return string(fs.Content), nil
}


