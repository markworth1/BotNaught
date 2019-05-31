package service

import (
	//"bufio"
	"context"
	"math"
	//"fmt"
	"log"
	"os"
	"strconv"
	//"strings"

	game "github.com/gSchool/golang-curriculum-c-6/server/pkg/game"
	poker "github.com/chehsunliu/poker"
)

// BotnaughtService describes the service.
type BotnaughtService interface {
	Health(ctx context.Context) (err error)
	Action(ctx context.Context, game game.Game) (action game.Action, err error)
}

type basicBotnaughtService struct{
	curGameID string
}

func (b *basicBotnaughtService) Health(ctx context.Context) (err error) {
	return err
}
func (b *basicBotnaughtService) Action(ctx context.Context, curGame game.Game) (action game.Action, err error) {

	f, err := os.OpenFile("botnaught-log.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()


	logger := log.New(f, "", 0)

	callOrCheck := "call"
	if b.curGameID != curGame.GameID{
		logger.Println("Game ID: " + curGame.GameID)
		logger.Println("--------------------------------------")
		logger.Println("--------------------------------------")
		b.curGameID = curGame.GameID
	}
	logger.Println("Previous actions:")
	for _, handlog := range curGame.HandLog {
		logger.Println(handlog)
	}
	logger.Println("")
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
	logger.Println("Player:" + myPlayer.Name)
	logger.Println("My Rank:" + strconv.Itoa(myPlayer.HandRankInt))
	logger.Println("My Hand:")
	for _, card := range myPlayer.HoleCards {
		logger.Print(card.String() + ", ")
	}
	for _, card := range curGame.CommunityCards {
		logger.Print(card.String() + ", ")
	}

	switch myBet := 
			Bet(myPlayer.HoleCards,myPlayer.HandRankInt,myPlayer.Chips,myPlayer.ChipsCommittedThisAction,curGame.CurrentBet,curGame.CommunityCards,logger); {		
		case myBet < 0:
			// FOLD!
			action.SelectedAction = "fold"
		case myBet > 0:
			// RAISE
			action.SelectedAction = "raise"
			action.Value = myBet
		default:
			// CALL or CHECK
			action.SelectedAction = callOrCheck
	}

	logger.Println("Returning action ", action)
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

// Bet - betting function based on input variables
func Bet(myCards []poker.Card, myRank int, myChips int, myCommitted int, currentBet int, communityCards []poker.Card, logger *log.Logger) (int) {
	myBet := -1
	myTotal := myChips + myCommitted
	availChips := myTotal - currentBet
	// ex: 40 chips + 30 committed - 50 current bet = 20 avail
	rankPct := 0.0

	if (myRank > 0) {
		rankPct = 1.0 - (float64(myRank)/7462.0)
		logger.Println("My rank / rank percentage: " + strconv.Itoa(myRank) + ", " + strconv.FormatFloat(rankPct, 'f', 3, 64))
	}
	logger.Println("Current bet: " + strconv.Itoa(currentBet))
	logger.Println("Current chips: " + strconv.Itoa(myChips) + " - Committed chips: " + strconv.Itoa(myCommitted))

	if availChips >= 0 { // We have enough chips to bet...
		flop := len(communityCards) == 3
		turn := len(communityCards) == 4
		river := len(communityCards) == 5


		myHandLead := int32(0)
		// No Community Cards have been dealt (PRE-FLOP)
		if len(communityCards) == 0 {
			// Raise if we have a Pair, Ace or suited K/Q
			if checkRaise(myCards) {
				myBet = int(math.Round(float64(myTotal) * .20))
			} else {
				if float32(currentBet) < float32(myTotal) * .60 {
					// Call
					myBet = currentBet
				}
			}
		} else {
			if river {
			    riverScore := poker.Evaluate(communityCards)
				logger.Println("Community score = " + strconv.Itoa(int(riverScore)) + " Best score = " + strconv.Itoa(myRank))
				myHandLead = riverScore - int32(myRank)
				// if myRank + 500 < int(riverScore) {
				// 	ourHandRocks = true
				// }
			}
			if turn {
			    turnScore1 := poker.Evaluate(append(communityCards, myCards[0]))
			    turnScore2 := poker.Evaluate(append(communityCards, myCards[1]))
				logger.Println("Turn score1 = " + strconv.Itoa(int(turnScore1)) + " Turn score2 = " + strconv.Itoa(int(turnScore2)))
				if turnScore1 > turnScore2 {
					myHandLead = turnScore2 - int32(myRank)
				} else {
					myHandLead = turnScore1 - int32(myRank)
				}
			}
			
			logger.Println("myHandLead: " + strconv.Itoa(int(myHandLead)))
			switch rankPct := rankPct; {
				case rankPct > .7 && myHandLead > 10:
					// ALL IN
					logger.Println("all in")
					myBet = myTotal
				case rankPct > .4 && flop:
					//Bid aggressively FLOP
					logger.Println("aggressive flop")
					myBet = int(math.Round(float64(myTotal) * rankPct))
				case rankPct > .45 && turn && myHandLead > 10:
					//Bid aggressively TURN
					logger.Println("aggressive turn")
					myBet = int(math.Round(float64(myTotal) * rankPct))
				case rankPct > .5 && river && myHandLead > 10:
					//Bid aggressively RIVER
					logger.Println("aggressive river")
					myBet = int(math.Round(float64(myTotal) * rankPct))
				default:
					willing := int(math.Round(float64(myTotal) * rankPct))
					logger.Println("default bet. Max: " + strconv.Itoa(willing))
					if willing >= currentBet {
						myBet = currentBet
					}
			}
		}
		logger.Println("Willing to bet: " + strconv.Itoa(myBet))
		if myHandLead > 500 {
			myBet = int(math.Round(float64(myBet) * 1.5))
			logger.Println("Multiplied by 1.5: " + strconv.Itoa(myBet))
		}
		// if we try to bet more chips than we have
		if myBet > myChips {
			myBet = myChips
		}
		// if current bet is greater than what we're willing to bet
		if (myBet < currentBet && myRank < 5000) {
			myBet = currentBet
		}
		// if current bet is greater than what we're willing to bet
		if (myBet < currentBet && !(myRank < 5000)) {
			myBet = -1
		}
		// if we are only willing to match current bet
		if myBet == currentBet {
			myBet = 0
		}
	}

	return myBet
}

func checkRaise(holeCards []poker.Card) (raiseIt bool) {
	raiseIt = false

	card1 := holeCards[0].String()
	card2 := holeCards[1].String()
	
	if (card1[0] == card2[0]) {
		raiseIt = true
	}
	if card1[0] == 'A' || card2[0] == 'A' {
		raiseIt = true
	}

	if card1[0] == 'K' || card2[0] == 'K' || card1[0] == 'Q' || card2[0] == 'Q'{
		if card1[1] == card2[1] {
			raiseIt = true
		}
	}

	return raiseIt
}