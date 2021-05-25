package cmd

import (
	"bytes"
)

func ignoreLines(b []byte, prefixs ...string) []byte {
	set := make(map[string]struct{})
	for _, k := range prefixs {
		set[k] = struct{}{}
	}

	splitBytes := bytes.Split(b, []byte("\n"))
	res := make([]byte, 0, len(b))

	for _, line := range splitBytes {
		kvp := bytes.Split(line, []byte(":"))
		if len(kvp) < 2 {
			continue
		}

		k := string(kvp[0])

		if _, exist := set[k]; exist {
			continue
		}

		res = append(res, line...)
		res = append(res, []byte("\n")...)
	}
	return res
}
