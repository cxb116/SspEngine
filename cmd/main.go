package main

import (
	"flag"
	"fmt"
	_ "github.com/cxb116/sspEngine/document/dsp" // 匿名导入,原因dsp包init() 无法执行
	"github.com/cxb116/sspEngine/implement"
	"github.com/cxb116/sspEngine/internal/engine"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"

	"github.com/cxb116/sspEngine/internal/cache"
	"github.com/cxb116/sspEngine/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var flagConfig = flag.String("c", "./config/config.yaml", "config path file")
	flag.Parse() // 解析命令行参数
	config, err := config.Load(*flagConfig)
	if err != nil {
		fmt.Printf("load config file failed, err: %v\n", err)
		return
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msgf("config: %v", config)

	RedisClient, err := cache.NewRedisClient(&config.Redis)
	if err != nil {
		log.Err(err).Msgf("redis connect failed")
		return
	}

	engine.Cache = RedisClient

	pool := goredis.NewPool(RedisClient)
	engine.DistributedLock = redsync.New(pool)

	implement.ServerEngine(config)

}
