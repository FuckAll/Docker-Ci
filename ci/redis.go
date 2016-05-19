/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 16:23
 */

package ci

import (
	. "github.com/FuckAll/Docker-Ci/cmdrun"
	"github.com/FuckAll/Docker-Ci/conf"
)

func Redis() error {
	_, err := CMD(FMT("docker run -ti -d --net=test -v  --name %s-redis %s", conf.Tracer, conf.RedisImage))
	if err != nil {
		return err
	}
	return nil
}
