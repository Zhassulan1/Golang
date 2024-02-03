package api

type Film struct {
	ID    int    `json:"ID"`
	Name  string `json:"name"`
	Genre string `json:"string"`
	IMDb  string `json:"IMDb"` // imdb link
}

var Films = []Film{
	{
		ID:    1,
		Name:  "The Shawshank Redemption",
		Genre: "Drama",
		IMDb:  "https://www.imdb.com/title/tt0111161/",
	},
	{
		ID:    2,
		Name:  "The Godfather",
		Genre: "Crime",
		IMDb:  "https://www.imdb.com/title/tt0068646/",
	},
	{
		ID:    3,
		Name:  "The Godfather: Part II",
		Genre: "Crime",
		IMDb:  "https://www.imdb.com/title/tt0071562/",
	},
	{
		ID:    4,
		Name:  "The Dark Knight",
		Genre: "Action",
		IMDb:  "https://www.imdb.com/title/tt0468569/",
	},
	{
		ID:    5,
		Name:  "12 Angry Men",
		Genre: "Drama",
		IMDb:  "https://www.imdb.com/title/tt0050083/",
	},
	{
		ID:    6,
		Name:  "Schindler's List",
		Genre: "Biography",
		IMDb:  "https://www.imdb.com/title/tt0108052/",
	},
	{
		ID:    7,
		Name:  "The Lord of the Rings: The Return of the King",
		Genre: "Adventure",
		IMDb:  "https://www.imdb.com/title/tt0167260/",
	},
	{
		ID:    8,
		Name:  "Pulp Fiction",
		Genre: "Crime",
		IMDb:  "https://www.imdb.com/title/tt0110912/",
	},
	{
		ID:    9,
		Name:  "The Good, the Bad and the Ugly",
		Genre: "Western",
		IMDb:  "https://www.imdb.com/title/tt0060196/",
	},
	{
		ID:    10,
		Name:  "Fight Club",
		Genre: "Drama",
		IMDb:  "https://www.imdb.com/title/tt0137523/",
	},
	{
		ID:    11,
		Name:  "The Matrix",
		Genre: "Action",
		IMDb:  "https://www.imdb.com/title/tt0133093/",
	},
	{
		ID:    12,
		Name:  "Forrest Gump",
		Genre: "Drama",
		IMDb:  "https://www.imdb.com/title/tt0109830/",
	},
	{
		ID:    13,
		Name:  "Inception",
		Genre: "Action",
		IMDb:  "https://www.imdb.com/title/tt1375666/",
	},
	{
		ID:    14,
		Name:  "Goodfellas",
		Genre: "Crime",
		IMDb:  "https://www.imdb.com/title/tt0099685/",
	},
	{
		ID:    15,
		Name:  "One Flew Over the Cuckoo's Nest",
		Genre: "Drama",
		IMDb:  "https://www.imdb.com/title/tt0073486/",
	},
	{
		ID:    16,
		Name:  "Seven Samurai",
		Genre: "Adventure",
		IMDb:  "https://www.imdb.com/title/tt0047478/",
	},
	{
		ID:    17,
		Name:  "Star Wars: Episode V - The Empire Strikes Back",
		Genre: "Action",
		IMDb:  "https://www.imdb.com/title/tt0080684/",
	},
	{
		ID:    18,
		Name:  "The Lord of the Rings: The Fellowship of the Ring",
		Genre: "Adventure",
		IMDb:  "https://www.imdb.com/title/tt0120737/",
	},
	{
		ID:    19,
		Name:  "Rear Window",
		Genre: "Mystery",
		IMDb:  "https://www.imdb.com/title/tt0047396/",
	},
	{
		ID:    20,
		Name:  "The Pianist",
		Genre: "Biography",
		IMDb:  "https://www.imdb.com/title/tt0253474/",
	},
	{
		ID:    21,
		Name:  "The Departed",
		Genre: "Crime",
		IMDb:  "https://www.imdb.com/title/tt0407887/",
	},
	{
		ID:    22,
		Name:  "Terminator 2: Judgment Day",
		Genre: "Action",
		IMDb:  "https://www.imdb.com/title/tt0103064/",
	},
	{
		ID:    23,
		Name:  "Back to the Future",
		Genre: "Adventure",
		IMDb:  "https://www.imdb.com/title/tt0088763/",
	},
	{
		ID:    24,
		Name:  "Whiplash",
		Genre: "Drama",
		IMDb:  "https://www.imdb.com/title/tt2582802/",
	},
	{
		ID:    25,
		Name:  "Gladiator",
		Genre: "Action",
		IMDb:  "https://www.imdb.com/title/tt0172495/",
	},
}
