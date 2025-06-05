package config

//import (
//	"context"
//	"database/sql"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/redis/go-redis/v9"
//	"log"
//	"strconv"
//	"sync"
//	"time"
//)
//
//var db *sql.DB
//
//func initDB() {
//
//	dsn := "root:root@tcp(127.0.0.1:3306)/basedata?charset=utf8mb4&parseTime=True"
//
//	var err error
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		log.Fatalf("数据库连接失败: %v", err)
//	}
//
//	// 优化后的连接池配置
//	db.SetMaxOpenConns(100)                // 最大打开连接数
//	db.SetMaxIdleConns(30)                 // 最大空闲连接数
//	db.SetConnMaxLifetime(2 * time.Hour)   // 连接最大存活时间
//	db.SetConnMaxIdleTime(5 * time.Minute) // 空闲连接最大存活时间
//
//	// 验证连接
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	if err := db.PingContext(ctx); err != nil {
//		log.Fatalf("数据库连接验证失败: %v", err)
//	}
//
//	log.Println("数据库连接池初始化成功")
//}
//
//var redisClient *redis.Client
//
//func initCache() {
//	redisClient = redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379", // Redis 地址
//		Password: "",               // 密码
//		DB:       0,                // 默认 DB
//
//		// 优化后的连接池配置
//		PoolSize:     500, // 更适合本地开发的连接数
//		MinIdleConns: 50,  // 最小空闲连接
//		MaxRetries:   30,  // 添加重试机制
//		DialTimeout:  5 * time.Second,
//		PoolTimeout:  30 * time.Second,
//		// 修正注释
//		ConnMaxLifetime: 30 * time.Minute, // 连接最大存活时间 100 分钟
//
//		// 推荐添加以下配置
//		ReadTimeout:  3 * time.Second,
//		WriteTimeout: 3 * time.Second,
//	})
//
//	// 测试连接（带重试逻辑）
//	var err error
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	// 添加重试逻辑
//	for i := 0; i < 3; i++ {
//		_, err = redisClient.Ping(ctx).Result()
//		if err == nil {
//			break
//		}
//		time.Sleep(time.Second * time.Duration(i+1))
//	}
//
//	if err != nil {
//		log.Fatalf("Redis connection failed after retries: %v", err)
//	}
//	log.Println("Redis connect success")
//}
//
//type AndroidData struct {
//	Oaid string `json:"oaid"`
//}
//
//func main() {
//	initDB()
//	initCache()
//
//	AndroidProducerChannel := make(chan *AndroidData, 1024)
//	go AndroidProducerWorker2(AndroidProducerChannel)
//	AndroidWorker2(AndroidProducerChannel)
//
//}
//func AndroidWorker2(channel chan *AndroidData) {
//	var wg sync.WaitGroup
//	setKeyOaid := "oaid"
//	for data := range channel {
//		wg.Add(1)
//		go func(android *AndroidData) {
//			defer wg.Done()
//			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//			log.Println("Oaid:", android.Oaid)
//			result, err := redisClient.SAdd(ctx, setKeyOaid, android.Oaid).Result()
//			if err != nil || result == 0 {
//				log.Printf("oaid add filed: %v", err)
//				cancel()
//			} else {
//				log.Println("oaid result", result)
//			}
//		}(data)
//	}
//	wg.Wait()
//}
//
//func AndroidProducerWorker2(AndroidProducerChannel chan *AndroidData) {
//	count := 0
//	index := 0
//	result, err := redisClient.Get(context.Background(), "Androids_id").Result()
//	if err != nil {
//		log.Printf("oaid add filed: %v", err)
//		count = 0
//	} else {
//		count, _ = strconv.Atoi(result)
//	}
//
//	for {
//		rows, err := db.Query("select oaid from android_click_data where id > ? limit 100", count)
//		if err != nil {
//			log.Fatal("error:", err)
//			continue
//		}
//		for rows.Next() {
//			index++
//			var android AndroidData
//			err := rows.Scan(&android.Oaid)
//			if err != nil {
//				log.Printf("err:%v", err)
//			}
//			log.Println("oaid::", android.Oaid)
//			AndroidProducerChannel <- &android
//		}
//		rows.Close()
//		count += 100
//		redisClient.Set(context.Background(), "Androids_id", count, 0).Result()
//		if count >= 60000000 {
//			break
//		}
//	}
//	close(AndroidProducerChannel)
//}
