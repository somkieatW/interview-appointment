package domain

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `comment` (
  `id` varchar(45) NOT NULL,
  `recruitment_id` varchar(45) NOT NULL,
  `user_id` varchar(45) NOT NULL,
  `content` text NOT NULL,
  `created_date` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": "sFZyBkUCuMbKqTpmNwMtfPpDc",    "recruitmentId": "gtWAAESocxuKYLwcajUfXETfQ",    "userId": "RtsLYMJPqjxaKjABgQVkTnoQk",    "content": "WwMiBBqvYdOvpeUalijMgJkRP",    "createdDate": "2072-11-18T22:02:14.571782396+07:00"}



*/

// Comments struct is a row record of the comment table in the recruitment database
type Comments struct {
	//[ 0] id                                             varchar(45)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 45      default: []
	ID string `gorm:"primary_key;column:id;type:varchar;size:45;" json:"id"`
	//[ 1] recruitment_id                                 varchar(45)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 45      default: []
	AppointmentID string `gorm:"column:appointment_id;type:varchar;size:45;" json:"appointmentId"`
	//[ 2] user_id                                        varchar(45)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 45      default: []
	UserID string `gorm:"column:user_id;type:varchar;size:45;" json:"userId"`
	//[ 3] content                                        text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Content string `gorm:"column:content;type:text;size:65535;" json:"content"`
	//[ 4] created_date                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;" json:"createdAt"`
}

// TableName sets the insert table name for this struct type
func (c *Comments) TableName() string {
	return "comments"
}
