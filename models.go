package main

type User struct {
	Id        int        `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	BirthDate int        `json:"birth_date"`
}

func updateUser(body []byte, rec *User, required bool) bool {
	if !parseId(body, &rec.Id, required) {
		return false
	}
	if !parseString(body, &rec.Email, "email", required) || len(rec.Email) > 100 {
		return false
	}
	if !parseString(body, &rec.FirstName, "first_name", required) || len(rec.FirstName) > 50 {
		return false
	}
	if !parseString(body, &rec.LastName, "last_name", required) || len(rec.LastName) > 50 {
		return false
	}
	if !parseString(body, &rec.Gender, "gender", required) || (rec.Gender != "f" && rec.Gender != "m") {
		return false
	}
	if !parseInt(body, &rec.BirthDate, "birth_date", required) || (rec.BirthDate < -1262325600 || rec.BirthDate > 915123600) {
		return false
	}
	return true
}

type DataUser struct {
	Users []User    `json:"users"`
}

type Location struct {
	Id       int        `json:"id"`
	Place    string    `json:"place"`
	Country  string    `json:"country"`
	City     string    `json:"city"`
	Distance int        `json:"distance"`
}

func updateLocation(body []byte, rec *Location, required bool) bool {
	if !parseId(body, &rec.Id, required) {
		return false
	}
	if !parseString(body, &rec.Place, "place", required) {
		return false
	}
	if !parseString(body, &rec.Country, "country", required) || len(rec.Country) > 50 {
		return false
	}
	if !parseString(body, &rec.City, "city", required) || len(rec.City) > 50 {
		return false
	}
	if !parseInt(body, &rec.Distance, "distance", required) {
		return false
	}
	return true
}

type DataLocation struct {
	Locations []Location    `json:"locations"`
}

type Visit struct {
	Id        int    `json:"id"`
	Location  int    `json:"location"`
	User      int    `json:"user"`
	VisitedAt int    `json:"visited_at"`
	Mark      int    `json:"mark"`
}

func updateVisit(body []byte, rec *Visit, required bool) bool {
	if !parseId(body, &rec.Id, required) {
		return false
	}
	if !parseInt(body, &rec.Location, "location", required) {
		return false
	}
	if _, ok := locations[rec.Location]; !ok {
		return false
	}
	if !parseInt(body, &rec.User, "user", required) {
		return false
	}
	if _, ok := users[rec.User]; !ok {
		return false
	}
	if !parseInt(body, &rec.VisitedAt, "visited_at", required) || (rec.VisitedAt < 946659600 || rec.VisitedAt > 1420045200) {
		return false
	}
	if !parseInt(body, &rec.Mark, "mark", required) || (rec.Mark < 0 || rec.Mark > 5) {
		return false
	}
	return true
}

func visitSetEvent(rec Visit) {
	orig := visits[rec.Id]

	// visits_by_user
	{
		vs, ok := visits_by_user[rec.User]
		if !ok {
			vs = make(map[int]Visit)
		}
		vs[rec.Id] = rec
		visits_by_user[rec.User] = vs
		visits[rec.Id] = rec

		if orig.User != rec.User {
			delete(visits_by_user[orig.User], orig.Id)
		}
	}

	// visits_by_location
	{
		vs, ok := visits_by_location[rec.Location]
		if !ok {
			vs = make(map[int]Visit)
		}
		vs[rec.Id] = rec
		visits_by_location[rec.Location] = vs
		visits[rec.Id] = rec

		if orig.Location != rec.Location {
			delete(visits_by_location[orig.Location], orig.Id)
		}
	}
}

type DataVisit struct {
	Visits []Visit    `json:"visits"`
}

type ShortVisit struct {
	Mark      int        `json:"mark"`
	Place     string    `json:"place"`
	VisitedAt int        `json:"visited_at"`
}

type ShortVisits []ShortVisit

func (s ShortVisits) Len() int {
	return len(s)
}
func (s ShortVisits) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ShortVisits) Less(i, j int) bool {
	return s[i].VisitedAt < s[j].VisitedAt
}

type DataShortVisit struct {
	Visits ShortVisits    `json:"visits"`
}

type DataAvg struct {
	Avg float32    `json:"avg"`
}
