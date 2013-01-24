package main

import (
)

type Utils struct {
}

func (this *Utils) StaticUrl(path string) string {
	return "/assets/" + path
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
