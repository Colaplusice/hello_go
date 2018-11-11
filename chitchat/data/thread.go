package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func (thread *Thread) Numberreplies(count int) {
	rows, err := Db.Query("select count(*) from posts where thread_id=$1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
		for rows.Next() {
			if err = rows.Scan(&count); err != nil {
				return
			}
		}
		rows.Close()
		return
	}

}
