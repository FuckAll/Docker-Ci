/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/11 17:03
 */

package main

import "github.com/FuckAll/Docker-Ci/ci"

func main() {
	ci.AppBuild()
	ci.DockerImageBuild()

}
