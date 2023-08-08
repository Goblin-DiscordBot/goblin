package race

import "time"

var (
	Races []*Race
)

type Race struct {
	ID          string
	Active      int // FancyDict
	Bets        int // FancyDict
	Config      *Config
	GamesPlayed int64
	Players     int // FancyDictList
	Started     int // FancyDict
	Winners     int // FancyDictList
}

type Config struct {
	BetAllowed bool
	BetMax     int64
	BetMin     int64
	Mode       string
	PayoutMax  int64
	PayoutMin  int64
	Pooling    bool
	Prize      int64
	Wait       time.Duration
}

type Racer struct {
	ID           string
	Character    *Character
	Current      string // track + animal, to keep track of movement
	LastMove     int
	LastPosition int
	Name         string
	Placed       bool
	Position     int
	Results      RacerResults
	Turn         int
	UserName     string
}

type RacerResults struct {
	First  int
	Second int
	Third  int
	Losses int
}

type Characters []*Character

type Character struct {
	Graphic  string
	Movement string
}

// NewRacer returns a racer for a given race
func NewRacer() *Racer {
	racer := &Racer{}
	return racer
}

/*
	Racer:
		animal
		mode (basically the theme; defaults to "clash")
		user
		turn
		position
		lastPosition
		lastMove
		placed (default to False)
		current: Racer.track + self.Animal

	stats
		returns racer stats
			% of races participated in
			1st, 2nd, 3rd place winss
			Losses

	field
		gets the field for the race? not posive

	getPosition
		self.current.find(self.animal)

	updateTrack
		distance = self.move()
		if polition !- 0
			last_move = distance
		self.current = line message with player's animal in the right spot

	updatePosition
		turn++
		updateTrack
		if position !=0
			last_position = position
		position = getPosition

	move
		base on the 'slow', 'fast', etc.


	Race
		bot (will be one of racers)
		system
		config (data/race/race.json)
		version (1.2.0)
		raceTime = -120

	prize
		3 of credits to randomly pick, multiplied by the # of entries in the race
			min, max

	time
		sets time the players have to enter race

	mode
		sets the race mode
			standard: everyone is a turtle
			zoo: random animal
			clash: random clash character

	reset
		reset race parameters

	start
		starts a race and enters yourself as a participant
			wait time between races: 2 minutes

		once the race starts, then allow players to place bets
			- get player's name or nickname
			- use reactions to get which player someone is betting on

		once the race is over, get the winners (1st, 2nd, 3rd) and award prizes
			60% to 1st, 30% to 2nd, 10% to 3rd (prize pooling)
			100% to 1st (no prize pooling)

		once race is over, list out the results of the bets

		Only allow 10 players (plus bot) in the race

	wipe
		deletes all data for the race. Probably want a confirmation about this.
*/
