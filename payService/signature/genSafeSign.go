package signature

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
)

func GetGenSafeSign(maps map[string]string, appKey string) string {
	// 1st , key排序
	var keys []string
	for k := range maps {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 2nd , 拼接
	var params = ""
	for _, k := range keys {
		if k == "sign" || k == "sig" {
			// skip
			continue
		}
		if params != "" {
			params = params + "&"
		}
		params = params + k + "=" + maps[k]
	}
	return GetMD5Hash(appKey + params)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
