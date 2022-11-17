package demo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"time"
)

type Template struct {
	ProjectName string
	ProjectPath string // working directory + project name
	// TODO 补充信息
	Author     string
	Email      string
	Copyright  string
	CreateTime string
}

func getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error in get working directory.")
		errWrapper(err)
	}
	return dir
}

func errWrapper(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewTemplate(projectName string) *Template {
	return &Template{
		ProjectName: projectName,
		ProjectPath: path.Join(getwd(), projectName),
	}
}

func (t *Template) projectIsExist() bool {
	_, err := os.Stat(t.ProjectPath)
	return err == nil || os.IsExist(err)
}

func (t *Template) createDir(dirName string) error {
	p := path.Join(t.ProjectPath, dirName)
	return os.MkdirAll(p, 0755)
}

func (t *Template) createFile(dirPath, fileName, fileContent string) error {
	err := os.Chdir(path.Join(t.ProjectPath, dirPath))
	if err != nil {
		return err
	}
	f, _ := os.Create(fileName)
	defer f.Close()

	if err != nil {
		return err
	}
	_, err = f.Write([]byte(fileContent))
	return err
}

func (t *Template) modInit() string {
	v := runtime.Version()
	r, _ := regexp.Compile("go(\\d\\.\\d+)(.*)")
	res := r.FindStringSubmatch(v)
	version := "1.13"
	if len(res) >= 1 {
		version = res[1]
	}
	return fmt.Sprintf(GoModuleStr, t.ProjectName, version)
}

func (t *Template) Run() {
	isExist := t.projectIsExist()

	if isExist {
		errWrapper(errors.New(fmt.Sprintf("project %s is existed", t.ProjectName)))
	}

	var (
		services_api_config         = "services/api/config"
		services_api_models         = "services/api/models"
		services_api_routers        = "services/api/routers"
		services_api_controllers    = "services/api/controllers"
		services_api_utils          = "services/api/utils"
		services_api_web            = "services/api/web"
		services_api_cmd            = "services/api/cmd" // 多个api.go入口服务
		services_auth_cmd           = "services/auth/cmd"
		services_gateway_cmd        = "services/gateway/cmd"
		services_mq_cmd             = "services/mq/cmd"
		services_mq_etc             = "services/mq/etc"
		services_mq_internal_config = "services/mq/internal/config"
		services_mq_internal_logic  = "services/mq/internal/logic"
		services_mq_internal_server = "services/mq/internal/server"
		services_mq_internal_svc    = "services/mq/internal/svc"
		services_mq_internal_types  = "services/mq/internal/types"
		services_mq_proto           = "services/mq/proto"
	)
	dirs := []string{
		"data", "deploy", "docs", "test", "services",
		services_api_config,
		services_api_models,
		services_api_routers,
		services_api_controllers,
		services_api_utils,
		services_api_web,
		services_api_cmd,
		services_auth_cmd,
		services_gateway_cmd,
		services_mq_cmd,
		services_mq_etc,
		services_mq_internal_config,
		services_mq_internal_logic,
		services_mq_internal_server,
		services_mq_internal_svc,
		services_mq_internal_types,
		services_mq_proto,
	}
	for _, dir := range dirs {
		errWrapper(t.createDir(dir))
	}

	// 根目录
	errWrapper(t.createFile("data", "README.md", ReadmeStr))
	errWrapper(t.createFile("deploy", "README.md", ReadmeStr))
	errWrapper(t.createFile("docs", "README.md", ReadmeStr))
	errWrapper(t.createFile("test", "README.md", ReadmeStr))
	errWrapper(t.createFile("services", "README.md", ReadmeStr))
	errWrapper(t.createFile("", "README.md", fmt.Sprintf(ProjectReadmeStr, t.ProjectName, t.ProjectName, t.ProjectName, t.ProjectName, t.ProjectName, time.Now().Format("2006-01-02 15:04:05"))))
	errWrapper(t.createFile("", "操作使用说明.md", GuideStr))
	errWrapper(t.createFile("", "go.mod", t.modInit()))

	// serives 目录

	// 1. api 服务
	errWrapper(t.createFile(services_api_cmd, "api.go", fmt.Sprintf(ApiStr, t.ProjectName)))
	errWrapper(t.createFile(services_api_controllers, "controller.go", ControllerStr))
	errWrapper(t.createFile(services_api_routers, "router.go", RouterStr))
	errWrapper(t.createFile(services_api_models, "model.go", ModelStr))
	errWrapper(t.createFile(services_api_config, "app.ini", ""))

	// 2. auth 服务
	errWrapper(t.createFile(services_auth_cmd, "gateway.go", CmdStr))

	// 3. gateway 服务
	errWrapper(t.createFile(services_gateway_cmd, "gateway.go", CmdStr))

	// 4. mq 服务
	errWrapper(t.createFile(services_mq_cmd, "mq.go", CmdStr))
	errWrapper(t.createFile(services_mq_etc, "mq.yaml", ""))

}
