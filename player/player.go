package main

import (
	"bufio"
	"fmt"
	"os"
)

type player struct {
	playerName []string
	pseudo     []string
	score      int
	nbPlayers  int
}

var myplayer player

func (myplayer *player) addPlayer(name string, pseudo string, score int) {
	myplayer.playerName = append(myplayer.playerName, name)
	myplayer.pseudo = append(myplayer.pseudo, pseudo)
	myplayer.score = score
	myplayer.nbPlayers++
}

func partie() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter player name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter player pseudo: ")
	pseudo, _ := reader.ReadString('\n')
	//fmt.Println(name, pseudo)

	myplayer.addPlayer(name, pseudo, 0)

	for i := 0; i < myplayer.nbPlayers; i++ {
		fmt.Printf("Le joueur %v s'appelle %s , il porte le pseudo %s et possède un score de %d \n", i, myplayer.playerName, myplayer.pseudo, myplayer.score)
	}

}

func main() {
	//myplayer.addPlayer("Nooooon", "Ouuuuui", 500)

	//fmt.Printf("Joueur %v %v %v \n  ", myplayer.playerName, myplayer.pseudo, myplayer.score)

	partie()
}
