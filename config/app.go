/*
Package config
载入环境变量和默认值
优先级为：环境变量 > 配置文件 > 默认值
example: 以app.name为例
1. 环境变量的键需要大写，中间使用_分隔，且加上前缀，默认的前缀为DOCKERLAB_
export DOCKERLAB_APP_NAME=dockerlab
2. 配置文件
---
app:
  name: "dockerlab"
---
3. 默认值
config.Load("app.name", "dockerlab")
默认值为：dockerlab
---
*/
package config

import "YangCodeCamp/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Load("app.name", "yangcodecamp"),
			// app的key，用于关键内容的签字
			"key": config.Load("app.key", "yangcodecamp"),
			// 时区 Local = Asia/Shanghai
			"time_zone": config.Load("app.time_zone", "Asia/Shanghai"),
			// 是否开启debug模式，在debug模式下会自动插入数据，跳过鉴权
			"debug": config.Load("app.debug", false),
		}
	})
}

// Initialize 用于触发包内的init函数
func Initialize() {}
