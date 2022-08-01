package pkg

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(machineID int64) (err error) {
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
