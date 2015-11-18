package core

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
	"net/url"
	"sort"
	"strings"
)

func GetSha1String(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GenGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetSha1String(base64.URLEncoding.EncodeToString(b))
}

func PercentEncode(value string) string {
	return strings.Replace((strings.Replace((strings.Replace(url.QueryEscape(value), "+", "%20", -1)), "*", "%2A", -1)), "%7E", "~", -1)
}

func EncodeParams(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := PercentEncode(k) + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(PercentEncode(v))
		}
	}
	return buf.String()
}
