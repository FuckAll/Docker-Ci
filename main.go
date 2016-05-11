/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/11 17:03
 */

package main

import (
    "fmt"
    "sync"
    "github.com/FuckAll/Docker-Ci/ci"
    "github.com/FuckAll/Docker-Ci/conf"
)

// var waitGroup = &sync.WaitGroup{}
// var goroutine = func(stageFunc func()){
    // defer waitGroup.Done()
    // stageFunc()
// }

func main(){
    ci.AppBuild()
    // waitGroup.Add(1) // Add 3 task to goroutine sequence
    // fmt.Println("wooooool")
    // go goroutine(ci.Redis)
    // waitGroup.Wait()
    
    //redis
    //pgsql
    //consul
    
    
    //build
    //images
    //AppTest
    //Clean()
    
}