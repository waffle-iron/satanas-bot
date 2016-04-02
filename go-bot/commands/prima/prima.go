package prima

import (
	"fmt"
	"github.com/xdigu/go-bot"
	"math/rand"
	"regexp"
)

const (
	pattern = "(?i)\\b(prima)\\b"
)

var (
	re          = regexp.MustCompile(pattern)
	frasesPrima = []string{
		"Vai pagar 15 boletos sim! Se reclamar eu mando mais 20 :retardat:.",
		"Perai que eu preciso ligar pra minha mãe e avisar que o bb passou no OAB.",
		"Existem alguns programas que só funcionam em processador Windows.",
		"Preciso economizar 8 mil reais para colocar peitos.",
		"Tia! O Alan passou na OAB.",
		"Tia! O que eu faço de janta para meu pai?",
		"Deixa que eu dirijo.",
		"Pensei que ela tinha morrido.",
		"A mina do Zé?",
		"PELOS PODERES DE GRAYSKULL!!!",
		"Essa mina ta ligando pros parente de novo?",
	}
)

func prima(command *bot.PassiveCmd) (string, error) {
	if re.MatchString(command.Raw) {
		return fmt.Sprintf("%s", frasesPrima[rand.Intn(len(frasesPrima))]), nil
	}
	return "", nil
}

func init() {
	bot.RegisterPassiveCommand(
		"prima",
		prima)
}