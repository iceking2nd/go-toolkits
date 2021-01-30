package info

import (
	"io/ioutil"
	"strings"
)

func getHosts() (map[string][]string, error) {
	bs, err := ioutil.ReadFile(HostsPath)
	if err != nil {
		return nil, err
	}
	hostsMap := map[string][]string{}
	for _, line := range strings.Split(strings.Trim(string(bs), " \t\r\n"), "\n") {
		line = strings.Replace(strings.Trim(line, " \t"), "\t", " ", -1)
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}
		pieces := strings.SplitN(line, " ", 2)
		if len(pieces) > 1 && len(pieces[0]) > 0 {
			if names := strings.Fields(pieces[1]); len(names) > 0 {
				if _, ok := hostsMap[pieces[0]]; ok {
					hostsMap[pieces[0]] = append(hostsMap[pieces[0]], names...)
				} else {
					hostsMap[pieces[0]] = names
				}
			}
		}
	}
	return hostsMap, nil
}
