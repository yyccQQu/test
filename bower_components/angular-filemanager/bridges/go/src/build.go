// +build ignore
/*
go run build.go setup
go run build.go govendor fetch +out

go run build.go build server

govendor

*/
package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	goarch      string
	goos        string
	gocc        string
	gocxx       string
	cgo         string
	version     string = "v1.0.1"
	race        bool
	environment string
)

//TODO 自动打包到docker
//或者 bin
func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	flag.StringVar(&goarch, "goarch", runtime.GOARCH, "GOARCH")
	flag.StringVar(&goos, "goos", runtime.GOOS, "GOOS")
	flag.StringVar(&gocc, "cc", "", "CC")
	flag.StringVar(&gocxx, "cxx", "", "CXX")
	flag.StringVar(&cgo, "cgo-enabled", "", "CGO_ENABLED")
	flag.BoolVar(&race, "race", race, "Use race detector")
	flag.StringVar(&environment, "environment", "development", "Use environment testing production development")

	flag.Parse()
	//设置gopath
	ensureGoPath()

	if flag.NArg() == 0 {
		log.Println("Usage: go run build.go build")
		return
	}
	cmds := flag.Args()
	//workingDir, _ = os.Getwd()
	//对应的命令
	args := []string{}
	for k, cmd := range cmds {
		if k == 0 {
			continue
		}
		args = append(args, cmd)
	}
	fmt.Println(cmds[0], args)
	switch cmds[0] {
	case "govendor":
		fmt.Println(cmds[0], args)
		runPrint("govendor", args...)
	//runPrint("govendor", "fetch","+out")
	case "build-run":
		clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin name]")
			return
		}
		build(args[0], "./cmd/"+args[0], []string{})
		runPrint("../bin/" + args[0])
	case "build":
		clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin name]")
			return
		}
		build(args[0], "./cmd/"+args[0], []string{})
	case "builds":
		clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin names]")
			return
		}
		for _, binary := range args {
			build(binary, "./src/cmd/"+binary, []string{})
		}
	case "test":
		test("./...")

	case "sha-dist":
		shaFilesInDist()

	case "clean":
		clean()
	default:
		log.Fatalf("Unknown command %q", cmds[0])
	}

}

//配置 gopath
func ensureGoPath() {
	//govendor add +external

	g_gopath := os.Getenv("GOPATH")

	if g_gopath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		gopath := filepath.Clean(filepath.Join(cwd, "/../")) //+ ":" + filepath.Clean(cwd)
		log.Println("GOPATH is", gopath)
		os.Setenv("GOPATH", gopath)
	} else {
		//设置gopath
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		gopath := filepath.Clean(filepath.Join(cwd, "/../")) + ":" + g_gopath
		log.Println("GOPATH is", gopath)
		os.Setenv("GOPATH", gopath)
	}
}

func ChangeWorkingDir(dir string) {
	os.Chdir(dir)
}

//依赖管理
func setup() {
	vendor := ""
	//配置依赖
	str_arr := strings.Split(vendor, "\n")
	for _, v := range str_arr {
		if len(v) > 0 {
			runPrint("go", "get", "-v", v)
		}
	}
	//runPrint("go", "get", "-v", "github.com/kardianos/govendor")
	//runPrint("go", "install", "-v", "./pkg/cmd/grafana-server")
}

func test(pkg string) {
	setBuildEnv()
	runPrint("go", "test", "-short", "-timeout", "60s", pkg)
}

//编译
func build(binaryName, pkg string, tags []string) {
	binary := "../bin/" + binaryName
	if goos == "windows" {
		binary += ".exe"
	}

	rmr(binary, binary+".md5")
	args := []string{"build", "-ldflags", ldflags()}
	if len(tags) > 0 {
		args = append(args, "-tags", strings.Join(tags, ","))
	}
	if race {
		args = append(args, "-race")
	}
	args = append(args, "-o", binary)
	args = append(args, pkg)
	setBuildEnv()

	runPrint("go", "version")
	runPrint("go", args...)

	// Create an md5 checksum of the binary, to be included in the archive for
	// automatic upgrades.
	err := md5File(binary)
	if err != nil {
		log.Fatal(err)
	}
}

func ldflags() string {
	var b bytes.Buffer
	b.WriteString("-w")
	b.WriteString(fmt.Sprintf(" -X config.VERSION=%s", version))
	b.WriteString(fmt.Sprintf(" -X main.ENVIRONMENT=%s", environment))
	//b.WriteString(fmt.Sprintf(" -X main.commit=%s", getGitSha()))
	//b.WriteString(fmt.Sprintf(" -X main.buildstamp=%d", buildStamp()))
	return b.String()
}

func rmr(paths ...string) {
	for _, path := range paths {
		log.Println("rm -r", path)
		os.RemoveAll(path)
	}
}

//配置文件
func clean() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	rmr(filepath.Join(cwd, "../bin"))
	//runPrint("mkdir", "-p", filepath.Join(cwd, "../bin/log"))
	//runPrint("mkdir", "-p", filepath.Join(cwd, "../bin/etc"))
	//runPrint("cp", "-r", filepath.Join(cwd, "etc/conf.yaml"), filepath.Join(cwd, "../bin/etc/conf.yaml"))

	//rmr("tmp")
	//os.Getenv("GOPATH")  批量问题
	//rmr(filepath.Join(os.Getenv("GOPATH"), fmt.Sprintf("pkg/%s_%s/github.com/grafana", goos, goarch)))
}

func setBuildEnv() {
	os.Setenv("GOOS", goos)
	if strings.HasPrefix(goarch, "armv") {
		os.Setenv("GOARCH", "arm")
		os.Setenv("GOARM", goarch[4:])
	} else {
		os.Setenv("GOARCH", goarch)
	}
	if goarch == "386" {
		os.Setenv("GO386", "387")
	}
	if cgo != "" {
		os.Setenv("CGO_ENABLED", cgo)
	}
	if gocc != "" {
		os.Setenv("CC", gocc)
	}
	if gocxx != "" {
		os.Setenv("CXX", gocxx)
	}
}

func runPrint(cmd string, args ...string) {
	log.Println(cmd, strings.Join(args, " "))
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	err := ecmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func md5File(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	h := md5.New()
	_, err = io.Copy(h, fd)
	if err != nil {
		return err
	}

	out, err := os.Create(file + ".md5")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "%x\n", h.Sum(nil))
	if err != nil {
		return err
	}

	return out.Close()
}

func shaFilesInDist() {
	filepath.Walk("./dist", func(path string, f os.FileInfo, err error) error {
		if path == "./dist" {
			return nil
		}

		if strings.Contains(path, ".sha256") == false {
			err := shaFile(path)
			if err != nil {
				log.Printf("Failed to create sha file. error: %v\n", err)
			}
		}
		return nil
	})
}

func shaFile(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	h := sha256.New()
	_, err = io.Copy(h, fd)
	if err != nil {
		return err
	}

	out, err := os.Create(file + ".sha256")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "%x\n", h.Sum(nil))
	if err != nil {
		return err
	}

	return out.Close()
}
