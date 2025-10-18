package engine

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

var DistributedLock *redsync.Redsync
