-- Write your MySQL query statement below
/*
type Seat struct {
  ID int `json:"id"`
  Student string `json:"student"`
}
*/
SELECT (
	CASE
        WHEN MOD(id,2) != 0 and counts != id THEN id + 1 
	WHEN MOD(id,2) != 0 and counts = id THEN id
	ELSE id - 1 END
	) as id,
	student
	FROM seat, (
		select count(*) as counts from seat) as seat_counts order by id asc;
