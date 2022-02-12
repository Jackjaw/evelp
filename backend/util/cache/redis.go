package cache

import (
	"context"
	"encoding/json"
	"evelp/config/global"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var ctx = context.Background()

func Set(key string, value interface{}, expirationTime time.Duration) error {
	val, err := json.Marshal(value)
	if err != nil {
		return errors.Wrapf(err, "redis marshal %v failed", value)
	}

	if err := global.REDIS.Set(ctx, key, val, expirationTime).Err(); err != nil {
		return errors.Wrapf(err, "redis set %v failed", key)
	}

	return nil
}

func Get(key string, dest interface{}) error {
	val, err := global.REDIS.Get(ctx, key).Result()
	if err != nil {
		return errors.Wrapf(err, "get %v failed", key)
	}
	if err := json.Unmarshal([]byte(val), dest); err != nil {
		return errors.Wrap(err, "redis unmarshal failed")
	}
	return nil
}

func Exist(key ...string) (bool, error) {
	val, err := global.REDIS.Exists(ctx, key...).Result()
	if err != nil {
		return false, errors.Wrapf(err, "check %v exist failed", key)
	}

	if val == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func Key(args ...string) string {
	var build strings.Builder
	for index, arg := range args {
		build.WriteString(arg)
		if !(index == len(args)-1) {
			build.WriteString(":")
		}
	}
	return build.String()
}
