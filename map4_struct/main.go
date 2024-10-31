package main

import (
	"errors"
	"fmt"
	"log"
)

type Storage struct {
	team2player map[string]map[string]struct{}
	player2Team map[string]string
}

func NewTeamPlayer() *Storage {
	return &Storage{
		team2player: make(map[string]map[string]struct{}),
		player2Team: make(map[string]string),
	}
}

func (s *Storage) Add(team string, player string) {
	if _, exists := s.team2player[team]; !exists {
		s.team2player[team] = make(map[string]struct{})
	}
	s.team2player[team][player] = struct{}{}

	if _, exists := s.player2Team[player]; !exists {
		s.player2Team[player] = team
	}

}

func (s *Storage) GetPlayersByTeam(team string) ([]string, error) {
	var playersByTeam []string
	if _, exists := s.team2player[team]; exists {
		for players := range s.team2player[team] {
			playersByTeam = append(playersByTeam, players)
		}
		return playersByTeam, nil
	} else {
		return nil, fmt.Errorf("there is no such team")
	}
}

func (s *Storage) GetTeamByPlayer(player string) (string, error) {
	if _, exists := s.player2Team[player]; exists {
		return s.player2Team[player], nil
	} else {
		err := errors.New("there is no such player")
		log.Println("Error")
		return "", err
	}

}

func (s *Storage) TransferPlayer(teamFrom, teamTo, player string) error {
	if _, exists := s.team2player[teamFrom][player]; exists {
		if _, exists := s.team2player[teamTo]; !exists {
			s.team2player[teamTo] = make(map[string]struct{})
		}
		s.team2player[teamTo][player] = struct{}{}
		s.player2Team[player] = teamTo
		delete(s.team2player[teamFrom], player)
		fmt.Printf("%s has been transfered from %s to %s \n", player, teamFrom, teamTo)
		return nil
	} else {
		return fmt.Errorf("impossible to transfer player because there is no player in this team")
	}
}

func main() {
	newTeamPlayer := NewTeamPlayer()

	newTeamPlayer.Add("navi", "Simple")
	newTeamPlayer.Add("navi", "elick")
	newTeamPlayer.Add("navi", "flamie")
	newTeamPlayer.Add("navi", "adsa")
	newTeamPlayer.Add("navi", "dadasda")

	newTeamPlayer.Add("g2", "kennys")
	newTeamPlayer.Add("g2", "pennys")
	newTeamPlayer.Add("g2", "krimz")
	newTeamPlayer.Add("g2", "winroid")
	newTeamPlayer.Add("g2", "dffsdfds")

	newTeamPlayer.Add("fnatic", "jw")
	newTeamPlayer.Add("fnatic", "getright")
	newTeamPlayer.Add("fnatic", "huyzna")
	newTeamPlayer.Add("fnatic", "per")
	newTeamPlayer.Add("fnatic", "fer")

	err := newTeamPlayer.TransferPlayer("fnatic", "navi", "Simple")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newTeamPlayer.GetTeamByPlayer("Simple"))
	newTeamPlayer.TransferPlayer("navi", "g2", "Simple")
	fmt.Println(newTeamPlayer.GetPlayersByTeam("g2"))
	fmt.Println(newTeamPlayer.GetPlayersByTeam("navi"))
	fmt.Println(newTeamPlayer.GetTeamByPlayer("Simple"))
}
