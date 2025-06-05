package main

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"sync"
	"time"
)

type IosClickData struct {
	Ip          string `json:"ip"`
	Idfa        string `json:"idfa"`
	Ua          string `json:"ua"`
	Model       string `json:"model"`
	Caid        string `json:"caid"`
	CaidVersion string `json:"caid_version"`
}
type AndroidClickData struct {
	Ip    string `json:"ip"`
	Oaid  string `json:"oaid"`
	Ua    string `json:"ua"`
	Model string `json:"model"`
	Imei  string `json:"imei"`
}

// 配置常量
const (
	ServerPort     = ":8888"
	DBMaxConns     = 50
	DBMaxIdleConns = 10
	BatchSize      = 500 // 批量插入大小
	FlushInterval  = 5 * time.Second
)

// 全局变量
var (
	redisClient      *redis.Client
	db               *sql.DB
	AndroidChannel   chan *AndroidClickData
	IOSChannel       chan *IosClickData
	androidClickData *AndroidClickData // android
	iosClickData     *IosClickData     // ios
	wg               sync.WaitGroup
)

// 初始化数据库连接
func initDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/basedata?charset=utf8mb4&parseTime=True"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 优化后的连接池配置
	db.SetMaxOpenConns(100)                // 最大打开连接数
	db.SetMaxIdleConns(30)                 // 最大空闲连接数
	db.SetConnMaxLifetime(2 * time.Hour)   // 连接最大存活时间
	db.SetConnMaxIdleTime(5 * time.Minute) // 空闲连接最大存活时间

	// 验证连接
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("数据库连接验证失败: %v", err)
	}

	log.Println("数据库连接池初始化成功")
}
func initCache() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 密码
		DB:       0,                // 默认 DB

		// 优化后的连接池配置
		PoolSize:     500, // 更适合本地开发的连接数
		MinIdleConns: 50,  // 最小空闲连接
		MaxRetries:   30,  // 添加重试机制
		DialTimeout:  5 * time.Second,
		PoolTimeout:  30 * time.Second,
		// 修正注释
		ConnMaxLifetime: 30 * time.Minute, // 连接最大存活时间 100 分钟

		// 推荐添加以下配置
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// 测试连接（带重试逻辑）
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 添加重试逻辑
	for i := 0; i < 3; i++ {
		_, err = redisClient.Ping(ctx).Result()
		if err == nil {
			break
		}
		time.Sleep(time.Second * time.Duration(i+1))
	}

	if err != nil {
		log.Fatalf("Redis connection failed after retries: %v", err)
	}
	log.Println("Redis connect success")
}
func initChannel() {
	AndroidChannel = make(chan *AndroidClickData, 1024)
	IOSChannel = make(chan *IosClickData, 1024)
}

func main() {
	initDB()
	initCache()
	initChannel()

	for i := 0; i < 4; i++ {
		go AndroidChannelWorker(AndroidChannel)
		go IOSChannelWorker(IOSChannel)
	}
	http.HandleFunc("/adinfo/cpa/ios/4439", IosClickHandler)
	http.HandleFunc("/adinfo/cpa/android/4439", AndroidClickHandler)

	server := &http.Server{
		Addr:              ServerPort,
		ReadTimeout:       2 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1MB
	}
	server.ListenAndServe()
	close(AndroidChannel)
	close(IOSChannel)
}

func AndroidClickHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	data := &AndroidClickData{
		Ip:    query.Get("ip"),
		Oaid:  query.Get("oaid"),
		Ua:    query.Get("ua"),
		Model: query.Get("model"),
		Imei:  query.Get("imei"),
	}
	if data.Imei == "" && data.Oaid == "" {
		log.Println("imei && oaid 都是空的")
		return
	}
	select {
	case AndroidChannel <- data:
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]interface{}{
			"code": 200,
			"msg":  "ok",
		})
	default:
		log.Println("由于AndroidChannel 已满，请求丢弃")
	}
}

func IosClickHandler(writer http.ResponseWriter, request *http.Request) {

}

func IOSChannelWorker(channel chan *IosClickData) {

}

func AndroidChannelWorker(channel chan *AndroidClickData) {
	setKeyImei := "imei"
	setKeyOaid := "oaid"

	for channels := range channel {
		wg.Add(1)
		go func(android *AndroidClickData) {
			count := 0
			defer wg.Done()
			log.Println("结果:", android)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			if android.Imei != "" {
				result, err := redisClient.SAdd(ctx, setKeyImei, android.Imei).Result()
				if err != nil {
					log.Println("imei add filed: %v", err)
					cancel()
				}
				if result == 1 {
					log.Println("imei add success")
					count++
				}
			}
			if android.Oaid != "" {
				iosResult, err := redisClient.SAdd(ctx, setKeyOaid, android.Oaid).Result()
				if err != nil {
					log.Println("oaid remove filed: %v", err)
					cancel()
				}
				if iosResult == 1 {
					log.Println("oaid add success")
					count++
				}
			}
			if count > 0 {
				result, err := db.Exec("INSERT INTO android_click_data (ip, oaid, imei, ua, model, created_at) VALUES (?, ?, ?, ?, ?, ?)",
					android.Ip, android.Oaid, android.Imei, android.Ua, android.Model, time.Now())
				if err != nil {
					log.Printf("insert failed: %v", err) // Fixed typo: "filed" → "failed"
					return
				}
				id, err := result.LastInsertId()
				if err != nil {
					log.Printf("get last insert ID failed: %v", err)
					return
				}
				log.Printf("insert success: %d", id)
			}

		}(channels)

	}
	wg.Wait()
}
