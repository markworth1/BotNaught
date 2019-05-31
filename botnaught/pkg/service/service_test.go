package service

import (
	"fmt"
	"context"
	"github.com/gSchool/golang-curriculum-c-6/server/pkg/game"
	"github.com/gSchool/golang-curriculum-c-6/server/pkg/player"
	poker "github.com/chehsunliu/poker"
	"reflect"
	"testing"
)

func Test_basicBotnaughtService_Action(t *testing.T) {
	p := []player.Player{
		player.Player{Name: "p1", Address: ":9501"},
		player.Player{Name: "p2", Address: ":9502"},
		player.Player{Name: "p3", Address: ":9503"},
		player.Player{Name: "p4", Address: ":9504"},
	}
	gm, err := game.StartNewGame(p, 1, 2, 100)
	if err != nil {
		t.Error(err)
	}
	type args struct {
		ctx  context.Context
		game game.Game
	}
	tests := []struct {
		name       string
		b          *basicBotnaughtService
		args       args
		wantAction game.Action
		wantErr    bool
	}{
		{
			name: "PRE-FLOP: Check",
			b: &basicBotnaughtService{},
			args: args{
				game: game.Game{
					GameID:          "CheckPreFlop",
					PokerPlayers:    []game.PokerPlayer{
						game.PokerPlayer{
							Name:                     "Vinnie",
							Chips:                    100,
							HoleCards:            []poker.Card{
										poker.NewCard("Ts"),
										poker.NewCard("2d"),
										},
							HandRankInt:              0,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Jimmy",
							Chips:                    99,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Guido",
							Chips:                    98,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
					},
					HandLog:          gm.HandLog,
					AvailableActions: gm.AvailableActions,
					PotSize:		  gm.PotSize,
					CommunityCards:	  []poker.Card{},
					CurrentBet:       0,
					SmallBlind:       gm.SmallBlind,
					BigBlind:         gm.BigBlind,
					StartingStack:    gm.StartingStack,
				},
			},
			wantAction: game.Action{SelectedAction: "call"},
			wantErr: false,
		},
		{
			name: "PRE-FLOP: Raise (Ace)",
			b: &basicBotnaughtService{},
			args: args{
				game: game.Game{
					GameID:          "RaiseAce",
					PokerPlayers:    []game.PokerPlayer{
						game.PokerPlayer{
							Name:                     "Vinnie",
							Chips:                    100,
							HoleCards:            []poker.Card{
										poker.NewCard("As"),
										poker.NewCard("2d"),
										},
							HandRankInt:              0,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Jimmy",
							Chips:                    99,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Guido",
							Chips:                    98,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
					},
					HandLog:          gm.HandLog,
					AvailableActions: gm.AvailableActions,
					PotSize:		  gm.PotSize,
					CommunityCards:	  []poker.Card{},
					CurrentBet:       0,
					SmallBlind:       gm.SmallBlind,
					BigBlind:         gm.BigBlind,
					StartingStack:    gm.StartingStack,
				},
			},
			wantAction: game.Action{SelectedAction: "raise", Value: 20},
			wantErr: false,
		},
		{
			name: "PRE-FLOP: Raise (Suited Q/K)",
			b: &basicBotnaughtService{},
			args: args{
				game: game.Game{
					GameID:          "Suited",
					PokerPlayers:    []game.PokerPlayer{
						game.PokerPlayer{
							Name:                     "Vinnie",
							Chips:                    100,
							HoleCards:            []poker.Card{
										poker.NewCard("Qd"),
										poker.NewCard("2d"),
										},
							HandRankInt:              0,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Jimmy",
							Chips:                    99,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Guido",
							Chips:                    98,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
					},
					HandLog:          gm.HandLog,
					AvailableActions: gm.AvailableActions,
					PotSize:		  gm.PotSize,
					CommunityCards:	  []poker.Card{},
					CurrentBet:       0,
					SmallBlind:       gm.SmallBlind,
					BigBlind:         gm.BigBlind,
					StartingStack:    gm.StartingStack,
				},
			},
			wantAction: game.Action{SelectedAction: "raise", Value: 20},
			wantErr: false,
		},
		{
			name: "FLOP: Call (Not great hand)",
			b: &basicBotnaughtService{},
			args: args{
				game: game.Game{
					GameID:          "FlopCall",
					PokerPlayers:    []game.PokerPlayer{
						game.PokerPlayer{
							Name:                     "Vinnie",
							Chips:                    100,
							HoleCards:            []poker.Card{
										poker.NewCard("Qs"),
										poker.NewCard("2d"),
										},
							HandRankInt:              0,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Jimmy",
							Chips:                    99,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Guido",
							Chips:                    98,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
					},
					HandLog:          gm.HandLog,
					AvailableActions: gm.AvailableActions,
					PotSize:		  gm.PotSize,
					CommunityCards:	  []poker.Card{
						poker.NewCard("Ts"),
						poker.NewCard("3h"),
						poker.NewCard("7c"),},
					CurrentBet:       0,
					SmallBlind:       gm.SmallBlind,
					BigBlind:         gm.BigBlind,
					StartingStack:    gm.StartingStack,
				},
			},
			wantAction: game.Action{SelectedAction: "call"},
			wantErr: false,
		},
		{
			name: "FLOP: Call (match bet)",
			b: &basicBotnaughtService{},
			args: args{
				game: game.Game{
					GameID:          "FlopMatchCall",
					PokerPlayers:    []game.PokerPlayer{
						game.PokerPlayer{
							Name:                     "Vinnie",
							Chips:                    100,
							HoleCards:            []poker.Card{
										poker.NewCard("Qs"),
										poker.NewCard("2d"),
										},
							HandRankInt:              int(poker.Evaluate([]poker.Card{
								poker.NewCard("Qs"),
								poker.NewCard("2d"),
								poker.NewCard("Ts"),
								poker.NewCard("2h"),
								poker.NewCard("7c"),
								})),
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Jimmy",
							Chips:                    99,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
						game.PokerPlayer{
							Name:                     "Guido",
							Chips:                    98,
							HandRankInt:              -1,
							HandRankString:           "",
							ChipsCommittedThisAction: 0,
							IsPlayingHand:            true,
						},
					},
					HandLog:          gm.HandLog,
					AvailableActions: gm.AvailableActions,
					PotSize:		  gm.PotSize,
					CommunityCards:	  []poker.Card{
						poker.NewCard("Ts"),
						poker.NewCard("2h"),
						poker.NewCard("7c"),},
					CurrentBet:       5,
					SmallBlind:       gm.SmallBlind,
					BigBlind:         gm.BigBlind,
					StartingStack:    gm.StartingStack,
				},
			},
			wantAction: game.Action{SelectedAction: "call"},
			wantErr: false,
		},
		// {
		// 	name: "fold if bet is higher than what we are willing to bet",
		// 	b: &basicBotnaughtService{},
		// 	args: args{
		// 		game: game.Game{
		// 			GameID:          "FoldGame",
		// 			PokerPlayers:    []game.PokerPlayer{
		// 				game.PokerPlayer{
		// 					Name:                     "Vinnie",
		// 					Chips:                    100,
		// 					HoleCards:                poker.NewDeck().Draw(2),
		// 					HandRankInt:              6186,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Jimmy",
		// 					Chips:                    99,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Guido",
		// 					Chips:                    98,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 			},
		// 			HandLog:          gm.HandLog,
		// 			AvailableActions: gm.AvailableActions,
		// 			PotSize:		  gm.PotSize,
		// 			CommunityCards:	  poker.NewDeck().Draw(3),
		// 			CurrentBet:       gm.CurrentBet,
		// 			SmallBlind:       gm.SmallBlind,
		// 			BigBlind:         gm.BigBlind,
		// 			StartingStack:    gm.StartingStack,
		// 		},
		// 	},
		// 	wantAction: game.Action{SelectedAction: "fold"},
		// 	wantErr: false,
		// },
		// {
		// 	name: "call if no community cards are drawn && no pair",
		// 	b: &basicBotnaughtService{},
		// 	args: args{
		// 		game: game.Game{
		// 			GameID:          "CallGame",
		// 			PokerPlayers:    []game.PokerPlayer{
		// 				game.PokerPlayer{
		// 					Name:                     "Vinnie",
		// 					Chips:                    100,
		// 					HoleCards:                []poker.Card{
		// 									poker.NewCard("Ts"),
		// 									poker.NewCard("2s"),
		// 					},
		// 					HandRankInt:              0,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Jimmy",
		// 					Chips:                    99,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Guido",
		// 					Chips:                    98,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 			},
		// 			HandLog:          gm.HandLog,
		// 			AvailableActions: gm.AvailableActions,
		// 			PotSize:		  gm.PotSize,
		// 			CommunityCards:	  gm.CommunityCards,
		// 			CurrentBet:       gm.CurrentBet,
		// 			SmallBlind:       gm.SmallBlind,
		// 			BigBlind:         gm.BigBlind,
		// 			StartingStack:    gm.StartingStack,
		// 		},
		// 	},
		// 	wantAction: game.Action{SelectedAction: "call"},
		// 	wantErr: false,
		// },
		// {
		// 	name: "raise by 20 if starting with a pair",
		// 	b: &basicBotnaughtService{},
		// 	args: args{
		// 		game: game.Game{
		// 			GameID:          "Raise20Game",
		// 			PokerPlayers:    []game.PokerPlayer{
		// 				game.PokerPlayer{
		// 					Name:                     "Vinnie",
		// 					Chips:                    100,
		// 					HoleCards:                []poker.Card{
		// 						poker.NewCard("Ts"),
		// 						poker.NewCard("Td"),
		// 					},
		// 					HandRankInt:              0,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Jimmy",
		// 					Chips:                    99,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Guido",
		// 					Chips:                    98,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 			},
		// 			HandLog:          gm.HandLog,
		// 			AvailableActions: gm.AvailableActions,
		// 			PotSize:		  gm.PotSize,
		// 			CommunityCards:	  gm.CommunityCards,
		// 			CurrentBet:       gm.CurrentBet,
		// 			SmallBlind:       gm.SmallBlind,
		// 			BigBlind:         gm.BigBlind,
		// 			StartingStack:    gm.StartingStack,
		// 		},
		// 	},
		// 	wantAction: game.Action{SelectedAction: "raise", Value: 20},
		// 	wantErr: false,
		// },
		// {
		// 	name: "all in if rank >= 1600",
		// 	b: &basicBotnaughtService{},
		// 	args: args{
		// 		game: game.Game{
		// 			GameID:          "AllInGame",
		// 			PokerPlayers:    []game.PokerPlayer{
		// 				game.PokerPlayer{
		// 					Name:                     "Vinnie",
		// 					Chips:                    100,
		// 					HoleCards:                []poker.Card{
		// 						poker.NewCard("Ts"),
		// 						poker.NewCard("Qd"),
		// 					},
		// 					HandRankInt:              1600,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Jimmy",
		// 					Chips:                    99,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Guido",
		// 					Chips:                    98,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 			},
		// 			HandLog:          gm.HandLog,
		// 			AvailableActions: gm.AvailableActions,
		// 			PotSize:		  gm.PotSize,
		// 			CommunityCards:	  poker.NewDeck().Draw(3),
		// 			CurrentBet:       2,
		// 			SmallBlind:       gm.SmallBlind,
		// 			BigBlind:         gm.BigBlind,
		// 			StartingStack:    gm.StartingStack,
		// 		},
		// 	},
		// 	wantAction: game.Action{SelectedAction: "raise", Value: 98},
		// 	wantErr: false,
		// },
		// {
		// 	name: "fold if currentBet > 35% of our chips and rank is between 1601 - 6185",
		// 	b: &basicBotnaughtService{},
		// 	args: args{
		// 		game: game.Game{
		// 			GameID:          "Fold35%BidGame",
		// 			PokerPlayers:    []game.PokerPlayer{
		// 				game.PokerPlayer{
		// 					Name:                     "Vinnie",
		// 					Chips:                    100,
		// 					HoleCards:                []poker.Card{
		// 						poker.NewCard("Ts"),
		// 						poker.NewCard("Qd"),
		// 					},
		// 					HandRankInt:              3500,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Jimmy",
		// 					Chips:                    99,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Guido",
		// 					Chips:                    98,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 			},
		// 			HandLog:          gm.HandLog,
		// 			AvailableActions: gm.AvailableActions,
		// 			PotSize:		  gm.PotSize,
		// 			CommunityCards:	  poker.NewDeck().Draw(3),
		// 			CurrentBet:       36,
		// 			SmallBlind:       gm.SmallBlind,
		// 			BigBlind:         gm.BigBlind,
		// 			StartingStack:    gm.StartingStack,
		// 		},
		// 	},
		// 	wantAction: game.Action{SelectedAction: "fold"},
		// 	wantErr: false,
		// },
		// {
		// 	name: "call if currentBet < 35% of our chips and rank is between 4000 - 6185",
		// 	b: &basicBotnaughtService{},
		// 	args: args{
		// 		game: game.Game{
		// 			GameID:          "CallMidGame",
		// 			PokerPlayers:    []game.PokerPlayer{
		// 				game.PokerPlayer{
		// 					Name:                     "Vinnie",
		// 					Chips:                    100,
		// 					HoleCards:                []poker.Card{
		// 						poker.NewCard("Ts"),
		// 						poker.NewCard("Qd"),
		// 					},
		// 					HandRankInt:              4500,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Jimmy",
		// 					Chips:                    99,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Guido",
		// 					Chips:                    98,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 			},
		// 			HandLog:          gm.HandLog,
		// 			AvailableActions: gm.AvailableActions,
		// 			PotSize:		  gm.PotSize,
		// 			CommunityCards:	  poker.NewDeck().Draw(3),
		// 			CurrentBet:       30,
		// 			SmallBlind:       gm.SmallBlind,
		// 			BigBlind:         gm.BigBlind,
		// 			StartingStack:    gm.StartingStack,
		// 		},
		// 	},
		// 	wantAction: game.Action{SelectedAction: "call"},
		// 	wantErr: false,
		// },
		// {
		// 	name: "raise if currentBet < 35% of our chips and rank is between 1601 - 3999",
		// 	b: &basicBotnaughtService{},
		// 	args: args{
		// 		game: game.Game{
		// 			GameID:          "RaiseMidGame",
		// 			PokerPlayers:    []game.PokerPlayer{
		// 				game.PokerPlayer{
		// 					Name:                     "Vinnie",
		// 					Chips:                    100,
		// 					HoleCards:                []poker.Card{
		// 						poker.NewCard("Ts"),
		// 						poker.NewCard("Qd"),
		// 					},
		// 					HandRankInt:              3000,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Jimmy",
		// 					Chips:                    99,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 				game.PokerPlayer{
		// 					Name:                     "Guido",
		// 					Chips:                    98,
		// 					HandRankInt:              -1,
		// 					HandRankString:           "",
		// 					ChipsCommittedThisAction: 0,
		// 					IsPlayingHand:            true,
		// 				},
		// 			},
		// 			HandLog:          gm.HandLog,
		// 			AvailableActions: gm.AvailableActions,
		// 			PotSize:		  gm.PotSize,
		// 			CommunityCards:	  poker.NewDeck().Draw(3),
		// 			CurrentBet:       20,
		// 			SmallBlind:       gm.SmallBlind,
		// 			BigBlind:         gm.BigBlind,
		// 			StartingStack:    gm.StartingStack,
		// 		},
		// 	},
		// 	wantAction: game.Action{SelectedAction: "raise", Value: 20},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &basicBotnaughtService{}
			gotAction, err := b.Action(tt.args.ctx, tt.args.game)
			gm, err = game.StartNewGame(p, 1, 2, 100)
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("action returned: %v", gotAction)
			if (err != nil) != tt.wantErr {
				t.Errorf("basicBotnaughtService.Action() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAction, tt.wantAction) {
				t.Errorf("basicBotnaughtService.Action() = %v, want %v", gotAction, tt.wantAction)
			}
		})
	}
}
