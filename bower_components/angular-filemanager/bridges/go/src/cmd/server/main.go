package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"util"
	"netfile"
	"gopkg.in/mgo.v2"
	"fmt"
)

func list(path string) interface{} {
	/*
	{ "result": [
	    {
		"name": "magento",
		"rights": "drwxr-xr-x",
		"size": "4096",
		"date": "2016-03-03 15:31:40",
		"type": "dir"
	    }, {
		"name": "index.php",
		"rights": "-rw-r--r--",
		"size": "549923",
		"date": "2016-03-03 15:31:40",
		"type": "file"
	    }
	]}
	 */
	rs := new(result)
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Println(err)
		return rs
	}
	defer session.Close()
	mf1, err := netfile.NewMongodbFile(session, "config", "node1")
	if err != nil {
		fmt.Println("%v", err)
		return rs
	}

	rs.Result, _ = mf1.List(path)
	fmt.Println(rs.Result)
	return rs
}

func main() {
	//list("/src")
	//addFile("/Users/mebar/WebstormProjects/angular-filemanager/src")
	//return
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "../../../")

	// Route => handler
	e.Any("/handler", actionHandle)

	// Start server
	e.Logger.Fatal(e.Start(":8580"))
}

//
func actionHandle(ctx echo.Context) error {
	json_param := make(map[string]interface{})
	//解析json
	if strings.Contains(ctx.Request().Header.Get("Content-Type"), "application/json") {
		buf, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return ctx.JSONBlob(http.StatusOK, []byte("error1"))
		}
		err = json.Unmarshal(buf, &json_param)
		if err != nil {
			return ctx.JSONBlob(http.StatusOK, []byte("error2"))
		}
	}
	//uploadUrl 文件上传
	//downloadMultipleUrl
	var data interface{}
	switch util.ToStr(json_param["action"])  {
	case "list":
		data = list(util.ToStr(json_param["path"]))
	case "rename":
	case "move":
	case "copy":
	case "remove":
	case "edit":
	case "getContent":
		data, _ = getContent(util.ToStr(json_param["item"]))
	case "createFolder":
		data, _ = createFolder(util.ToStr(json_param["newPath"]))
	case "changePermissions":

	case "compress":
	case "extract":

	case "downloadMultiple":

	default:

	}
	return ctx.JSON(http.StatusOK, data)
}

type result struct {
	Result interface{} `json:"result"`
}

type resultStatus struct {
	Success bool `json:"success"`
	Error   string `json:"error"`
}

type  FileObject struct {
	Name     string `json:"name"`
	Rights   string `json:"rights"`
	Size     string `json:"size"`
	Date     string `json:"date"`
	FileType string `json:"type"`
}

func createFolder(newPath string) (interface{}, error) {
	rs := new(result)
	rsts := resultStatus{}
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Println(err)
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	defer session.Close()
	mf1, err := netfile.NewMongodbFile(session, "config", "node1")
	if err != nil {
		fmt.Println("%v", err)
		rsts.Error = err.Error()
		rsts.Success = false
		rs.Result = rsts
		return rs, err
	}
	/*
	{ "result": { "success": true, "error": null } }
	 */
	err = mf1.Create(newPath, []byte{})
	if err != nil {
		fmt.Println("---", err)
		rsts.Error = err.Error()
		rsts.Success = false
	} else {
		rsts.Success = true
	}
	rs.Result = rsts

	return rs.Result, nil
}

func getContent(item string) (interface{}, error) {
	rs := new(result)
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Println(err)
		return rs, err
	}
	defer session.Close()
	mf1, err := netfile.NewMongodbFile(session, "config", "node1")
	if err != nil {
		fmt.Println("%v", err)
		return rs, err
	}
	rs.Result, err = mf1.Read(item)
	return rs, err
}

func addFile(path string) error {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer session.Close()
	mf1, err := netfile.NewMongodbFile(session, "test", "test")
	if err != nil {
		fmt.Println("%v", err)
		return err
	}
	//"/Users/mebar/WebstormProjects/angular-filemanager/src"
	files, err := util.ListDir(path)
	if err != nil {
		fmt.Println("---", err)
		return err
	}
	for _, v := range files {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			if !strings.Contains(err.Error(), "is a directory") {
				fmt.Println("---", err)
				return err
			}
		}
		v = strings.Replace(v, path, "/templates/test", -1)
		err = mf1.Create(v, data)
		if err != nil {
			fmt.Println("---", err)
			return err
		}
	}
	path = "/Users/mebar/workspaces/golang/management1.0/bin/etc/conf.yaml"
	data, err := ioutil.ReadFile("/Users/mebar/workspaces/golang/management1.0/bin/etc/conf.yaml")
	if err != nil {
		if !strings.Contains(err.Error(), "is a directory") {
			fmt.Println("---", err)
			return err
		}
	}
	path = strings.Replace(path, "/Users/mebar/workspaces/golang/management1.0/bin/etc", "/config/test", -1)
	err = mf1.Create(path, data)
	if err != nil {
		fmt.Println("---", err)
		return err
	}
	return nil
}
