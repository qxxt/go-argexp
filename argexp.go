package argexp

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Marshall(flags []string) (res string) {
	rgx := regexp.MustCompile(`^--(\w[-]?)+=.*$`)
	rgx2 := regexp.MustCompile(`^-(\w+)`)
	for i := 0; i < len(flags); i++ {
		if rgx.MatchString(flags[i]) {
			withEqualCharacter := strings.SplitN(flags[i], "=", 2)
			res += strconv.Quote(withEqualCharacter[0])
			res += strconv.Quote(withEqualCharacter[1])
		} else if rgx2.MatchString(flags[i]) {
			flags[i] = strings.TrimLeft(flags[i], "-")
			res += `"-` + strings.Join(strings.Split(flags[i], ""), `""-`) + `"`
		} else {
			res += strconv.Quote(flags[i])
		}
	}
	return
}

func GetString(args *string, flag string) string {
	re := regexp.MustCompile(strconv.Quote(flag) + `(".*?[^\\]")`)
	arres := re.FindStringSubmatch(*args)
	if len(arres) > 0 {
		*args = re.ReplaceAllString(*args, "")
		res, _ := strconv.Unquote(arres[1])
		return res
	}
	return ""
}

func GetBool(args *string, flag string) bool {
	pattern := strconv.Quote(flag)
	if strings.Contains(*args, pattern) {
		*args = strings.Replace(*args, pattern, "", 1)
		return true
	}
	return false
}
