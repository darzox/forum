package app

import (
	controller "forum/internal/controller/http"
)

func Run() {
	controller.RunServer()
}
