package day2

import (
	"log"

	"github.com/wincus/adventofcode/internal/common"
)

func Solve() {

	data, err := common.GetData(2)

	if err != nil {
		log.Panic(err)
	}

	common.ShowData(data)

}
