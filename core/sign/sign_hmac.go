package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"sort"
	"strings"
)

type PlanXSign struct {
	SecretKey      string
	SignStr        string
	ParamsOriginal map[string]string
	Message        string
}

func (h *PlanXSign) Sign() {
	params := make(map[string]string)
	// 剔除sign和空值
	for k, v := range h.ParamsOriginal {
		if k != "sign" && v != "" {
			params[k] = v
		}
	}
	// 排序
	var pKeys []string
	for k := range params {
		pKeys = append(pKeys, k)
	}
	sort.Strings(pKeys)
	// 拼接
	for _, k := range pKeys {
		h.Message = h.Message + k + "=" + params[k] + "&"
	}
	// trim last &
	h.Message = strings.TrimRight(h.Message, "&")
	// sign
	key := []byte(h.SecretKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(h.Message))
	// base64
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	h.SignStr = sign
}

func NewSign(secretKey string) PlanXSign {
	return PlanXSign{
		SecretKey:      secretKey,
		ParamsOriginal: map[string]string{},
	}
}
