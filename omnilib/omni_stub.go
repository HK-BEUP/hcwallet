//go:build !omni
// +build !omni

package omnilib

var ChanReqOmToHc = make(chan string)
var ChanRspOmToHc = make(chan string)

func JsonCmdReqHcToOm(strReq string) string {
	return `{"error":"omnilib disabled; rebuild with -tags omni"}`
}

func LoadLibAndInit() {}

func OmniStart(strArgs string, strArgs1 string) {}
