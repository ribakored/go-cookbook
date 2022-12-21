package main

import "strings"

type Config struct {
	Chances string
	Scale   string
}

func ParseConfig(propStr string) Config {
	propMap := strings.Split(propStr, "\n")
	appConfig := Config{}
	for _, ele := range propMap {
		propStr := strings.Split(ele, "=")
		// Don't set anything if properties are not in name=value format
		if len(propStr) == 2 {
			if propStr[0] == "chances" {
				appConfig.Chances = propStr[1]
			}
			if propStr[0] == "scale" {
				appConfig.Scale = propStr[1]
			}
		}
	}
	return appConfig
}
