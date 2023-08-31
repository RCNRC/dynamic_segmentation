package dynamicsegmentation

type User struct {
	Id int `json:"id"`
}

type UsersSegment struct {
	User    User
	Service string `json:"service"`
	TTL     string `json:"ttl"`
}

type UserSegments struct {
	User     User           `json:"user"`
	Segments []UsersSegment `json:"segments"`
}

type UserUpdate struct {
	User           User           `json:"user"`
	SegmentsAdd    []UsersSegment `json:"segments_add"`
	SegmentsDelete []UsersSegment `json:"segments_delete"`
}
