package main

import (
	"log"

	termui "github.com/gizak/termui/v3"

	"github.com/drakeu/k8s-dev-tools/config"
	"github.com/drakeu/k8s-dev-tools/ui"
)

func main() {
	cfg := config.NewConfig()
	cfg.LoadConfiguration()

	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	contextList := ui.NewContextsList(cfg)
	contextList.Render()

	// use the current context in kubeconfig
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(config)

}
