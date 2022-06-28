package tool_validator

import (
	"regexp"
	"time"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var domainRegex = regexp.MustCompile(`^((ht|f)tps?):\/\/[\w\-]+(\.[\w\-]+)+([\w\-\.\{\},@?^=%&:\/~\+#]*[\w\-\@?^=%&\/~\+#])?$`)
var ipRegex = regexp.MustCompile(`^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$`)
var deviceIdRegex = regexp.MustCompile(`^[0-9a-zA-Z_-]{1,}$`)
var imeiRegex = regexp.MustCompile(`^[0-9a-zA-z]{15}$`)
var idfaRegex = regexp.MustCompile(`^[0-9a-zA-z]{8}[-][0-9a-zA-z]{4}[-][0-9a-zA-z]{4}[-][0-9a-zA-z]{4}[-][0-9a-zA-z]{12}$`)
var md516Regex = regexp.MustCompile(`^[0-9a-zA-Z]{16}$`)
var md532Regex = regexp.MustCompile(`^[0-9a-zA-Z]{32}$`)

func checkDomain(str string) bool {
	return domainRegex.MatchString(str)
}

//CheckDomainString .
func CheckDomainString(str string) bool {
	if !checkDomain(str) {
		return false
	}
	return true
}

//CheckDomainArray .
func CheckDomainArray(arr []string) bool {
	for _, domain := range arr {
		if !checkDomain(domain) {
			return false
		}
	}
	return true
}

//CheckStringLength .
func CheckStringLength(str string, minLen, maxLen int) bool {
	strLen := utf8.RuneCountInString(str)
	if strLen > maxLen || strLen < minLen {
		return false
	}
	return true
}

// CheckStrLengthInGBK .
func CheckStrLengthInGBK(str string, minLen, maxLen int) bool {
	gbkStr, _ := simplifiedchinese.GBK.NewEncoder().String(str)
	strLen := len(gbkStr)
	if strLen > maxLen || strLen < minLen {
		return false
	}
	return true
}

//CheckIp .
func CheckIp(ip string) bool {
	return ipRegex.MatchString(ip)
}

//CheckDevices .
func CheckDevices(device string) bool {
	return deviceIdRegex.MatchString(device)
}

//CheckIMei .
func CheckIMei(iMei string) bool {
	return imeiRegex.MatchString(iMei)
}

//CheckIdfa .
func CheckIdfa(idfa string) bool {
	return idfaRegex.MatchString(idfa)
}

// CheckMd516Or32 .
func CheckMd516Or32(md string) bool {
	return CheckMd516(md) || CheckMd532(md)
}

// CheckMd516 .
func CheckMd516(md string) bool {
	return md516Regex.MatchString(md)
}

// CheckMd532 .
func CheckMd532(md string) bool {
	return md532Regex.MatchString(md)
}

// CheckDatetime .
func CheckDatetime(str string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return false
	}
	return true
}

// CheckDate .
func CheckDate(str string) bool {
	_, err := time.Parse("2006-01-02", str)
	if err != nil {
		return false
	}
	return true
}
