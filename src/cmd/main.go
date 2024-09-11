package main

import "teacherScheduler/src/server"

//go:generate oapi-codegen --config ../../api/configs/models.yaml ../../api/api_spec.yaml
//go:generate oapi-codegen --config ../../api/configs/client.yaml ../../api/api_spec.yaml
//go:generate oapi-codegen --config ../../api/configs/server.yaml ../../api/api_spec.yaml

func main() {
	server.Start()
}
