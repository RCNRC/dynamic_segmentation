package dynamicsegmentation

type User struct {
	Id int `json:"id"`
}

type UsersSegment struct {
	Service string `json:"slag"`
	TTL     string `json:"ttl"`
}

type UserUpdate struct {
	UserId         int            `json:"id"`
	SegmentsAdd    []UsersSegment `json:"segments_add"`
	SegmentsDelete []UsersSegment `json:"segments_delete"`
}
