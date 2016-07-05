package infrastructure

import (
	"fmt"
	"testing"
	//	"time"
)

//func TestRedisStart(t *testing.T) {
//err := StartRedis()
//if err != nil {
//fmt.Println(err)
//}
//}

//func TestRedisStop(t *testing.T) {
//err := StopRedisContainer()
//if err != nil {
//fmt.Println(err)
//}
//err = RemoveRedisContainer()
//if err != nil {
//fmt.Println(err)
//}
//}
//func TestConsulStart(t *testing.T) {
//err := CreateConsulContainer()
//if err != nil {
//fmt.Println(err)
//}
//err = StartConsulContainer()
//if err != nil {
//fmt.Println(err)
//}

//}

//func TestConsulStop(t *testing.T) {
//err := StopConsulContainer()
//if err != nil {
//fmt.Println(err)
//}
//err = RemoveConsulContainer()
//if err != nil {
//fmt.Println(err)
//}

//}

//func TestPostgresStart(t *testing.T) {
//err := StartPostgres()
//if err != nil {
//fmt.Println(err)
//	}
//err := CreatePostgresContainer()
//if err != nil {
//fmt.Println(err)
//}
//err = StartPostgresContainer()
//if err != nil {
//fmt.Println(err)
//}
//}
func TestEtcdStart(t *testing.T) {
	//err := EtcdInit()
	err := StartEtcd()
	if err != nil {
		fmt.Println(err)
	}
}

//func TestPostgresStop(t *testing.T) {
//err := StopPostgresContainer()
//if err != nil {
//fmt.Println(err)
//}
//err = RemovePostgresContainer()
//if err != nil {
//fmt.Println(err)
//}

//}
