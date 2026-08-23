package main

import "fmt"

type playerStats struct {
	id      int
	name    string
	age     int
	role    string
	runs    int
	wickets int
}

var players []playerStats

func addPlayer() {
	var player playerStats
	fmt.Println("\n--- Add player ---")

	fmt.Print("id:")
	fmt.Scan(&player.id)

	fmt.Print("name:")
	fmt.Scan(&player.name)

	fmt.Print("age:")
	fmt.Scan(&player.age)

	fmt.Print("role:")
	fmt.Scan(&player.role)

	fmt.Print("runs:")
	fmt.Scan(&player.runs)

	fmt.Print("wickets:")
	fmt.Scan(&player.wickets)

	players = append(players, player)
	fmt.Println("Player added successfully!")

}

func playersList() {
	fmt.Println("\n--- players list")

	if len(players) == 0 {
		fmt.Println("no players found")
	}

	for _, player := range players {
		fmt.Println("-------------------------")
		fmt.Println("ID:", player.id)
		fmt.Println("Name:", player.name)
		fmt.Println("Age:", player.age)
		fmt.Println("Role:", player.role)
		fmt.Println("Runs:", player.runs)
		fmt.Println("Wickets:", player.wickets)
	}
}

func showMenu() {
	fmt.Println("\n==============================")

	fmt.Println("   CRICKET MANAGEMENT SYSTEM")
	fmt.Println("==============================")
	fmt.Println("1. Add Player")
	fmt.Println("2. List Players")
	fmt.Println("3. Find Player")
	fmt.Println("4. Update Player Stats")
	fmt.Println("5. Remove Player")
	fmt.Println("6. Exit")
	fmt.Println("==============================")
}

func removePlayer() {

	var id int
	fmt.Print("\nEnter player ID: ")
	fmt.Scan(&id)

	for i, player := range players {
		if player.id == id {
			players = append(players[:i], players[i+1:]...)
			fmt.Println("Player removed successfully!")
			return
		}
	}
	fmt.Println("Player not found.")
}

func findPlayer() {
	var id int

	fmt.Print("\n enter player id")
	fmt.Scan(&id)

	for _, player := range players {
		if player.id == id {
			fmt.Println("\nPlayer found!")

			fmt.Println("Name:", player.name)
			fmt.Println("Age:", player.age)
			fmt.Println("Role:", player.role)
			fmt.Println("Runs:", player.runs)
			fmt.Println("Wickets", player.wickets)
			return
		}

	}
	fmt.Println(" Player not found")

}

func updatePlayersStats() {
	var id int
	fmt.Print("\n Enter Player id: ")
	fmt.Scan(&id)

	for i := range players {
		if players[i].id == id {

			fmt.Print("\n Enter runs to add: ")
			var runs int
			fmt.Scan(&runs)

			fmt.Print("\n Enter wickets to add: ")
			var wickets int
			fmt.Scan(&wickets)

			players[i].runs += runs
			players[i].wickets += wickets

			fmt.Println("Statistics updated!")

			return
		}
	}
	fmt.Println("player not found")
}

func main() {

	for {
		showMenu()

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addPlayer()

		case 2:
			playersList()

		case 3:
			findPlayer()

		case 4:
			updatePlayersStats()

		case 5:
			removePlayer()

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
