package main

import (
	
	"os"
    "github.com/xdigu/go-bot"
    _ "github.com/xdigu/go-bot/commands/catfacts"
    _ "github.com/xdigu/go-bot/commands/catgif"
    _ "github.com/xdigu/go-bot/commands/chucknorris"
    _ "github.com/xdigu/go-bot/commands/cotacao"
    _ "github.com/xdigu/go-bot/commands/dilma"
    _ "github.com/xdigu/go-bot/commands/gif"
    _ "github.com/xdigu/go-bot/commands/megasena"
    _ "github.com/xdigu/go-bot/commands/treta"
    _ "github.com/xdigu/go-bot/commands/prima"
    _ "github.com/xdigu/go-bot/commands/travesti"
    // Import all the commands you wish to use
)

func main() {
    bot.RunSlack(os.Getenv("SLACK_TOKEN"))
}