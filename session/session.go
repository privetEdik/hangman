package session

import (
	"hangman/game"
	"hangman/io"
)

type GameSession struct {
	Game   *game.Hangman
	Input  *io.Input
	Output *io.OutPut
}

func NewGameSession() *GameSession {
	word := game.FindWord()
	return &GameSession{
		Game:   game.NewGame(word, io.GetHangmanStatesCount()),
		Input:  io.NewInput(),
		Output: io.NewOutput()}
}
func (s *GameSession) Run() {
	for !s.Game.IsOver() {
		s.Output.RenderGame(s.Game)
		letter := s.Input.GetLetter()
		result := s.Game.GuessLetter(letter)
		s.Output.RenderFeedback(result)
	}
	s.Output.RenderResult(s.Game)
}
