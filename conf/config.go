package conf

import "fmt"

type DockerContainerConfig struct {
	ImageRepository string // 容器名字
	ImageTag        string // 容器版本
}

type VolumePathConfig struct {
	CodeHostPath   string // 存放代码宿主机路径
	CodeTargetPath string // 挂载到容器内的代码路径
	CaseHostPath   string // 存放测试样例文件（输入文件）宿主机路径
	CaseTargetPath string // 挂载到容器内样例文件路径

	AnswerHostPath string // 答案文件主机路径
}

type RabbitMQConfig struct {
	Host string
	Username string
	Password string
}

func (volume *VolumePathConfig) GetCodePath() string {
	return volume.CodeHostPath + ":" + volume.CodeTargetPath
}

func (volume *VolumePathConfig) GetCasePath() string {
	return volume.CaseHostPath + ":" + volume.CaseTargetPath
}

func (conf DockerContainerConfig) GetImageName() string {
	return conf.ImageRepository + ":" + conf.ImageTag
}

func (conf RabbitMQConfig) GetURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:5672/", conf.Username, conf.Password, conf.Host)
}