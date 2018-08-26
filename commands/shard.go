package commands

import (
	"fmt"
	"strconv"
)

var shard = &Command{
	Aliases: []string{"shard", "debug"},
	Method:  onShard,
}

const shardUsage = "Usage: {{prefix}}shard <guild id> <shard count>"

func onShard(ctx *Context) Response {
	if len(ctx.Args) != 2 {
		r := fmt.Sprintf("This guild is on shard number `%d` of `%d`.",
			ctx.Session.ShardID, ctx.Session.ShardCount)
		return textResponse(r)
	}
	id, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return textResponse(shardUsage)
	}
	count, err := strconv.Atoi(ctx.Args[1])
	if err != nil {
		return textResponse(shardUsage)
	}

	shardID := (id >> 22) % count
	r := fmt.Sprintf("Guild with ID `%d` is on shard `%d`.", id, shardID)
	return textResponse(r)
}
