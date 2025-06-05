package main

//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//// 统计字符串中出现的某个字符的次数
//// 1 生产者生产随机字符，然后将字符存入管道中，消费者并发处理这些数据
//// 消费者去消费这些数据
//type Producer struct {
//	Id   int
//	Data string
//}
//type Consumer struct {
//	resultCount int    // 字符出现的次数
//	str         string // 字符
//}
//
//func main() {
//	producerChannel := make(chan *Producer, 1024)
//	consumerChannel := make(chan *Consumer, 1024)
//
//	createConsumerPool(2, producerChannel, consumerChannel)
//
//	consumerWorker(consumerChannel)
//	producerWorker(producerChannel)
//}
//
//// 协程池
//func createConsumerPool(count int, producerChan chan *Producer, consumerChan chan *Consumer) {
//	for i := 0; i < count; i++ {
//		go func(producerChan chan *Producer, consumerChan chan *Consumer) {
//			for producer := range producerChan {
//				randstring := producer.Data
//				byteRandString := []byte(randstring)
//				var count int
//				for ii := 0; ii < len(byteRandString); ii++ {
//					if byteRandString[ii] == 's' {
//						count++
//					}
//				}
//				ret := &Consumer{
//					resultCount: count,
//					str:         randstring,
//				}
//				consumerChan <- ret
//			}
//		}(producerChan, consumerChan)
//	}
//}
//
//func consumerWorker(consumerChannel chan *Consumer) {
//	go func(consumerChannel chan *Consumer) {
//		for consumer := range consumerChannel {
//			fmt.Println(consumer)
//		}
//	}(consumerChannel)
//}
//
//// 生产者
//func producerWorker(producerChannel chan *Producer) {
//	id := 0
//	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//	for {
//		rand := randString(100, letterBytes)
//		producer := &Producer{
//			Id:   id,
//			Data: rand,
//		}
//		producerChannel <- producer
//		fmt.Println("producer存入管道: ", producer)
//	}
//}
//
//func randString(count int, letterBytes string) string {
//	rand.Seed(time.Now().UnixNano()) // 初始化随机种子
//	b := make([]byte, count)
//	for i := range b {
//		b[i] = letterBytes[rand.Intn(len(letterBytes))] // 随机选取字符
//	}
//	return string(b)
//}
