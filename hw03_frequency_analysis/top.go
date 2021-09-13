package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type top struct {
	Str   string
	Count int
}

var topList []top

const cutList = 10

func Top10(input string) []string {
	list := make([]string, 0)
	unique := make(map[string]int)

	regex := regexp.MustCompile(`\S+`).FindAllString(input, -1)

	for _, v := range regex {
		unique[v]++
	}

	for k, v := range unique {
		if !strings.Contains(k, "*") {
			topList = append(topList, top{k, v})
		}
	}

	sort.Slice(topList, func(i, j int) bool {
		if topList[i].Count != topList[j].Count {
			return topList[i].Count > topList[j].Count
		}

		return topList[i].Str < topList[j].Str
	})

	if len(topList) >= cutList {
		topList = topList[:cutList]
	}

	for _, v := range topList {
		list = append(list, v.Str)
	}

	return list
}
