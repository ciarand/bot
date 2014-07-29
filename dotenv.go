package bot

import (
	"io/ioutil"
	"os"
	"strings"
)

func parseDotEnv(s string) {
	lines := strings.Split(s, "\n")

	for i := 0; i < len(lines); i += 1 {
		l := lines[i]
		if len(l) < 1 || strings.HasPrefix(l, "#") {
			continue
		}

		// split it on the first =
		slices := strings.SplitN(l, "=", 2)
		if len(slices) < 2 {
			continue
		}

		key := slices[0]
		val := slices[1]

		// if it starts and ends with a quote mark, strip them
		if (strings.HasPrefix(val, `"`) && strings.HasSuffix(val, `"`)) ||
			(strings.HasPrefix(val, `'`) && strings.HasSuffix(val, `'`)) {
			val = val[1 : len(val)-1]
		}

		os.Setenv(key, val)
	}
}

func ParseDotEnvFile(dir string) {
	data, err := ioutil.ReadFile(dir + "/.env")
	if err != nil {
		return
	}

	parseDotEnv(string(data))
}
