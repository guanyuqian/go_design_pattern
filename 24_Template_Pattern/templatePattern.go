package _4_Template_Pattern

//步骤 1
//创建一个抽象类
type GameImplement interface {
	initialize() string
	startPlay() string
	endPlay() string
	play() string
}

type GameTemplate struct {
	game GameImplement
}

func (receiver *GameTemplate) play() (res string) {
	res += receiver.game.initialize()
	res += receiver.game.startPlay()
	res += receiver.game.endPlay()
	return
}

func newGameTemplate(implement GameImplement) *GameTemplate {
	return &GameTemplate{implement}
}

//步骤 2
//创建扩展了上述类的实体类。

type CricketGame struct {
	gameTemplate *GameTemplate
}

func (receiver *CricketGame) initialize() string {
	return "Cricket Game Initialized! Start playing."
}

func (receiver *CricketGame) endPlay() string {
	return "Cricket Game Finished!"
}

func (receiver *CricketGame) startPlay() string {
	return "Cricket Game Started. Enjoy the game!"
}
func (receiver *CricketGame) play() string {
	return receiver.gameTemplate.play()
}

func newCricketGame() *CricketGame {
	ret := &CricketGame{}
	gameTemplate := newGameTemplate(ret)
	ret.gameTemplate = gameTemplate
	return ret
}

type FootballGame struct {
	gameTemplate *GameTemplate
}

func (receiver *FootballGame) initialize() string {
	return "Football Game Initialized! Start playing."
}

func (receiver *FootballGame) endPlay() string {
	return "Football Game Finished!"
}

func (receiver *FootballGame) startPlay() string {
	return "Football Game Started. Enjoy the game!"
}

func (receiver *FootballGame) play() string {
	return receiver.gameTemplate.play()
}

func newFootballGame() *FootballGame {
	ret := &FootballGame{}
	gameTemplate := newGameTemplate(ret)
	ret.gameTemplate = gameTemplate
	return ret
}
