package cfg

import (
	"errors"
	"os"
)

// 支持的配置文件类型常量定义
const (
	TypeToml       = "toml"
	TypeJson       = "json"
	TypeIni        = "ini"
	TypeYaml       = "yaml"
	TypeYml        = "yml"
	TypeProperties = "properties"
	TypeProps      = "props"
	TypeHcl        = "hcl"
	TypeDotenv     = "dotenv"
	TypeEnv        = "env"
)

// supportType 支持的配置文件类型
var supportType = []string{
	TypeToml,
	TypeJson,
	TypeIni,
	TypeYaml,
	TypeYml,
	TypeProperties,
	TypeProps,
	TypeHcl,
	TypeDotenv,
	TypeEnv,
}

// ConfigTypeNotSupport 配置类型不支持错误
var ConfigTypeNotSupport = errors.New("config file type do not support")

// ConfigFileNotExist 配置文件不存在错误
var ConfigFileNotExist = errors.New("config file do not exist")

// ConfigFileParseFailed 配置文件解析错误
var ConfigFileParseFailed = errors.New("config file parse failed")

// IFace config读取分析载入类
type IFace interface {
	Parse(resource interface{}, cType string, target interface{}) error
}

// IsFileExist 判断文件是否存在
// 存在返回 true 不存在返回false
func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsCfgTypeSupport 检查配置类型是否支持
func IsCfgTypeSupport(cType string) bool {
	for _, val := range supportType {
		if val == cType {
			return true
		}
	}
	return false
}
