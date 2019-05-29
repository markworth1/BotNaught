package service

import (
	//"bufio"
	"context"
	"math"
	"fmt"
	//"os"
	"strconv"
	//"strings"

	game "golang-curriculum-c-6/server/pkg/game"
)

// BotnaughtService describes the service.
type BotnaughtService interface {
	Health(ctx context.Context) (err error)
	Action(ctx context.Context, game game.Game) (action game.Action, err error)
}

type basicBotnaughtService struct{}

func (b *basicBotnaughtService) Health(ctx context.Context) (err error) {
	return err
}
func (b *basicBotnaughtService) Action(ctx context.Context, curGame game.Game) (action game.Action, err error) {
	callOrCheck := ""

	println("Game ID: " + curGame.GameID)
	for _, action := range curGame.AvailableActions {
		if action == "call" || action == "check" {
			callOrCheck = action
		}
	}


	myPlayer := game.PokerPlayer{}
	for _, player := range curGame.PokerPlayers {
		if len(player.HoleCards) > 0 {
			myPlayer = player
			break
		}
	}
	println("My Rank:" + strconv.Itoa(myPlayer.HandRankInt))
	println("My Hand:")
	for _, card := range myPlayer.HoleCards {
		fmt.Println(card.String())
	}
	if len(curGame.CommunityCards) == 0 {
		if myPlayer.HoleCards[0].String()[0] == myPlayer.HoleCards[1].String()[0] {
			action.SelectedAction = "raise"
			action.Value = 20
			fmt.Println("Returning action ", action)
			return action, err
		}
		action.SelectedAction = callOrCheck
		fmt.Println("Returning action ", action)
		return action, err
	}

	if myPlayer.HandRankInt > 0 {
		if myPlayer.HandRankInt >= 6186 {
			action.SelectedAction = "fold"
			fmt.Println("Returning action ", action)
			return action, err
		}
		if myPlayer.HandRankInt <= 1600 {
			action.SelectedAction = "raise"
			action.Value = myPlayer.Chips - curGame.CurrentBet
			fmt.Println("ALL IN: Returning action ", action)
			return action, err
		}
		if float32(curGame.CurrentBet) >= float32(myPlayer.Chips) * .35 {
			action.SelectedAction = "fold"
			fmt.Println("Returning action ", action)
			return action, err
		}
		if float32(curGame.CurrentBet) < float32(myPlayer.Chips) * .35 {
			if myPlayer.HandRankInt <= 6185 {
				if myPlayer.HandRankInt >= 4000 {
					action.SelectedAction = callOrCheck
					fmt.Println("Returning action ", action)
					return action, err
				}
				myBet := float64(myPlayer.Chips) * (1 - (float64(myPlayer.HandRankInt)/10000 * 2))
				if int(math.Round(myBet)) > curGame.CurrentBet {
					action.SelectedAction = "raise"
					action.Value = int(math.Round(myBet)) - curGame.CurrentBet
					fmt.Println("RAISE: Returning action ", action)
					return action, err
				}
			}
		}
	}

	action.SelectedAction = callOrCheck
	fmt.Println("Returning action ", action)
	return action, err
}

// NewBasicBotnaughtService returns a naive, stateless implementation of BotnaughtService.
func NewBasicBotnaughtService() BotnaughtService {
	return &basicBotnaughtService{}
}

// New returns a BotnaughtService with all of the expected middleware wired in.
func New(middleware []Middleware) BotnaughtService {
	var svc BotnaughtService = NewBasicBotnaughtService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
