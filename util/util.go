package util

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"time"

	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/common/math"
)

var Ether = math.BigPow(10, 18)
var Shannon = math.BigPow(10, 9)

var pow256 = math.BigPow(2, 256)
var zeroHash = regexp.MustCompile("^0?x?0+$")

func IsValidBase58Address(s string) bool {
	return common.IsBase58Account(s)
}

func IsZeroHash(s string) bool {
	return zeroHash.MatchString(s)
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTargetHex(diff int64) string {
	difficulty := big.NewInt(diff)

	n := big.NewInt(1)
	n.Lsh(n, 255)
	n.Div(n, difficulty)
	n.Lsh(n, 1)
	diff2 := n

	diff1 := new(big.Int).Div(pow256, difficulty)
	fmt.Println(diff2)
	fmt.Println(diff1)
	return string(common.BytesToHash(diff1.Bytes()).Hex())
}

func TargetHexToDiff(targetHex string) *big.Int {
	targetBytes := common.FromHex(targetHex)
	return new(big.Int).Div(pow256, new(big.Int).SetBytes(targetBytes))
}

func ToHex(n int64) string {
	return "0x0" + strconv.FormatInt(n, 16)
}

func FormatReward(reward *big.Int) string {
	return reward.String()
}

func FormatRatReward(reward *big.Rat) string {
	wei := new(big.Rat).SetInt(Ether)
	reward = reward.Quo(reward, wei)
	return reward.FloatString(8)
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MustParseDuration(s string) time.Duration {
	value, err := time.ParseDuration(s)
	if err != nil {
		panic("util: Can't parse duration `" + s + "`: " + err.Error())
	}
	return value
}

func String2Big(num string) *big.Int {
	n := new(big.Int)
	n.SetString(num, 0)
	return n
}
