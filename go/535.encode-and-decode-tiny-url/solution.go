package encodeanddecodetinyurl

import (
	"encoding/base64"
)

/*
 * @lc app=leetcode id=535 lang=golang
 *
 * [535] Encode and Decode TinyURL
 */

// @lc code=start
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	return base64.StdEncoding.EncodeToString([]byte(longUrl))
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	bs, error := base64.StdEncoding.DecodeString(shortUrl)
	if error != nil {
		return ""
	}
	return string(bs)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

// @lc code=end
