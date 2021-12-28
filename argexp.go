package argexp

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Marshall(args []string) (res string) {
	var withEqualCharacter []string
	var multipleBoolean []string
	for _, arg := range args {
		if m, _ := regexp.MatchString(`^(\-{2}(\-*\w)+)=.*$`, arg); m {
			withEqualCharacter = strings.SplitN(arg, "=", 2)
			res += strconv.Quote(withEqualCharacter[0])
			res += strconv.Quote(withEqualCharacter[1])
		} else if m, _ := regexp.MatchString(`^\-(\w+)`, arg); m {
			arg = strings.TrimLeft(arg, "-")
			multipleBoolean = strings.Split(arg, "")
			res += `"-` + strings.Join(multipleBoolean, `""-`) + `"`
		} else {
			res += strconv.Quote(arg)
		}
	}
	return res
}

func GetString(args *string, flag string) string {
	re := regexp.MustCompile(strconv.Quote(flag) + `(".*?[^\\\"]")`)
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
