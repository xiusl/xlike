package utils

import "strconv"

func IntsToStrs(elements []int64) []string {
    res := make([]string, len(elements))
    for i, e := range elements {
        if e == 0 { continue }
        res[i] = strconv.FormatInt(e, 10)
    }
    return res
}
