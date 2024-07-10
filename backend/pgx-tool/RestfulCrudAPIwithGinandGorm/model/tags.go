package model

type Tags struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

type TagsX struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

// func (t *Tags) Create(conn *pgx.Conn) error {
// 	_, err := conn.Exec(context.Background(), "INSERT INTO tags (id, name) VALUES ($1, $2)", t.Id, t.Name)
// 	return err
// }

// func (t *Tags) Update(conn *pgx.Conn) error {
// 	_, err := conn.Exec(context.Background(), "UPDATE tags SET name = $1 WHERE id = $2", t.Name, t.Id)
// 	return err
// }

// func (t *Tags) Delete(conn *pgx.Conn) error {
// 	_, err := conn.Exec(context.Background(), "DELETE FROM tags WHERE id = $1", t.Id)
// 	return err
// }

// func GetTags(conn *pgx.Conn) ([]Tags, error) {
// 	rows, err := conn.Query(context.Background(), "SELECT * FROM tags")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var tags []Tags
// 	for rows.Next() {
// 		var tag Tags
// 		err := rows.Scan(&tag.Id, &tag.Name)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tags = append(tags, tag)
// 	}

// 	return tags, nil
// }
