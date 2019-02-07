package types

// ResultList is the collection of solutions results
type ResultList []Solution

// Solution is a class that was generated on https://mholt.github.io/json-to-go/
type Solution struct {
	ID            string   `json:"_id"`
	Index         int      `json:"index"`
	GUID          string   `json:"guid"`
	IsActive      bool     `json:"isActive"`
	Balance       string   `json:"balance"`
	Picture       string   `json:"picture"`
	Age           int      `json:"age"`
	EyeColor      string   `json:"eyeColor"`
	Name          string   `json:"name"`
	Gender        string   `json:"gender"`
	Company       string   `json:"company"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	Address       string   `json:"address"`
	About         string   `json:"about"`
	Registered    string   `json:"registered"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	Tags          []string `json:"tags"`
	Friends       []Friend `json:"friends"`
	Greeting      string   `json:"greeting"`
	FavoriteFruit string   `json:"favoriteFruit"`
}

// Friend is the you know
type Friend struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Parents []struct {
		Phone int    `json:"phone"`
		Name  string `json:"name"`
	} `json:"parents"`
}
