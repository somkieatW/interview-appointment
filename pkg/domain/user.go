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


CREATE TABLE `user` (
  `id` varchar(45) NOT NULL,
  `display_name` varchar(255) NOT NULL,
  `profile_image_url` varchar(255) DEFAULT NULL,
  `status` varchar(1) NOT NULL,
  `created_by` varchar(45) NOT NULL,
  `created_date` timestamp NOT NULL,
  `email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": "tcEZRWoduPKajTIvLRbPBexgZ",    "displayName": "bnmlAgbTARgEKNeXMxqWtiPlW",    "profileImageUrl": "TAbIlfjcDPojPMNbnmdnvNJeS",    "status": "gHoCcTOQHwfqnTTnicncMnsfs",    "createdBy": "aYKMpFbXQvDCJtXpDFcvKDMOE",    "createdDate": "2194-01-21T19:00:10.981226411+07:00",    "email": "qcwPNaBcfBwKRdWONZJFlElWe"}



*/

// User struct is a row record of the user table in the recruitment database
type User struct {
	//[ 0] id                                             varchar(45)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 45      default: []
	ID string `gorm:"primary_key;column:id;type:varchar;size:45;" json:"id"`
	//[ 1] display_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	DisplayName string `gorm:"column:display_name;type:varchar;size:255;" json:"displayName"`
	//[ 2] profile_image_url                              varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProfileImageURL null.String `gorm:"column:profile_image_url;type:varchar;size:255;" json:"profileImageUrl"`
	//[ 3] status                                         varchar(1)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1       default: []
	Status string `gorm:"column:status;type:varchar;size:1;" json:"status"`
	//[ 4] created_by                                     varchar(45)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 45      default: []
	CreatedBy string `gorm:"column:created_by;type:varchar;size:45;" json:"createdBy"`
	//[ 5] created_date                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedDate time.Time `gorm:"column:created_date;type:timestamp;" json:"createdDate"`
	//[ 6] email                                          varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	Email string `gorm:"column:email;type:varchar;size:200;" json:"email"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
