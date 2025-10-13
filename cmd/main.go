package main

import (
	"flag"
	"fmt"
	"github.com/yourusername/ssp_grpc/internal/config"
)

func main() {
	var flagConfig = flag.String("c", "./config/config.yaml", "config path file")
	flag.Parse() // 解析命令行参数
	load, err := config.Load(*flagConfig)
	if err != nil {
		fmt.Printf("load config file failed, err: %v\n", err)
		return
	}
	fmt.Printf("load config file success: %s \n", load)

	i := len("cff7b4d9d92cace1ff7ffaaa125480ba")
	fmt.Println(i)
}
