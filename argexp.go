package argexp

import (
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

func GetString(flags *string, findFlag string) (res string) {
	rgx := regexp.MustCompile(strconv.Quote(findFlag) + `(".*?[^\\]")`)
	arrRes := rgx.FindStringSubmatch(*flags)
	if len(arrRes) > 0 {
		*flags = rgx.ReplaceAllString(*flags, "")
		res, _ = strconv.Unquote(arrRes[1])
	}
	return
}

func GetBool(flags *string, findFlags ...string) (res bool) {
	for i := 0; i < len(findFlags); i++ {
		pattern := strconv.Quote(findFlags[i])
		if strings.Contains(*flags, pattern) {
			*flags = strings.Replace(*flags, pattern, "", 1)
			res = true
		}
	}
	return
}

func UnMarshall(flags *string) (res []string) {
	arrRes := regexp.MustCompile(`".*?[^\\]"`).FindAllString(*flags, -1)
	for i := 0; i < len(arrRes); i++ {
		strres, _ := strconv.Unquote(arrRes[i])
		res = append(res, strres)
	}
	return
}
