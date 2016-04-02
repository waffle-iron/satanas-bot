package travesti

import (
	"fmt"
	"github.com/xdigu/go-bot"
	"math/rand"
	"regexp"
)

const (
	pattern = "(?i)\\b(travesti)\\b"
)

var (
	re          = regexp.MustCompile(pattern)
	frasesTravesti = []string{
		"pegador",
	}
)

func travesti(command *bot.PassiveCmd) (string, error) {
	if re.MatchString(command.Raw) {
		return fmt.Sprintf(" %s", frasesTravesti[rand.Intn(len(frasesTravesti))]), nil
	}
	return "", nil
}

func init() {
	bot.RegisterPassiveCommand(
		"travesti",
		travesti)
}