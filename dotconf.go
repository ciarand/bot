package bot

import (
	"io/ioutil"
	"strings"
)

func confFromDotConf(s string) *config {
	c := &config{}

	lines := strings.Split(s, "\n")

	for _, l := range lines {
		if len(l) < 1 || strings.HasPrefix(l, "#") {
			continue
		}

		// split it on the first =
		slices := strings.SplitN(l, "=", 2)
		// if we don't have 2 segments someone messed up
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

		c.setFromVar(key, val)
	}

	return c
}

func ParseDotConfFile(dir string) (*config, error) {
	data, err := ioutil.ReadFile(dir + "/.conf")
	if err != nil {
		return nil, err
	}

	return confFromDotConf(string(data)), nil
}
