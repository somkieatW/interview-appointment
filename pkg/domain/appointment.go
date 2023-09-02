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


CREATE TABLE `appointment` (
  `id` varchar(45) NOT NULL,
  `topic` text NOT NULL,
  `state` varchar(20) NOT NULL,
  `description` text,
  `status` varchar(1) NOT NULL,
  `created_by` varchar(45) NOT NULL,
  `created_date` timestamp NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": "epCOAEImeIaXyRiakhfbZDhUQ",    "topic": "DMieurJErhZLaiFgosOpaHtLS",    "state": "iEJHPLBbmZvBeUDFLSeepAhhF",    "description": "TThimpYxrwWFAumBXpEZRimWv",    "status": "vfyKkERhywaaCwIcwfyQOHpBS",    "createdBy": "tjVdhygLuMmaMZULjPYGpQnwe",    "createdDate": "2177-06-25T23:37:40.977525944+07:00"}


Comments
-------------------------------------
[ 0] Warning table: appointment does not have a primary key defined, setting col position 1 id as primary key




*/

// Appointment struct is a row record of the appointment table in the recruitment database
type Appointment struct {
	//[ 0] id                                             varchar(45)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 45      default: []
	ID string `gorm:"primary_key;column:id;type:varchar;size:45;" json:"id"`
	//[ 1] topic                                          text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Topic string `gorm:"column:topic;type:text;size:65535;" json:"topic"`
	//[ 2] state                                          varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	State string `gorm:"column:state;type:varchar;size:20;" json:"state"`
	//[ 3] description                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Description null.String `gorm:"column:description;type:text;size:65535;" json:"description" swaggertype:"string"`
	//[ 4] status                                         varchar(1)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1       default: []
	Status string `gorm:"column:status;type:varchar;size:1;" json:"status"`
	//[ 5] created_by                                     varchar(45)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 45      default: []
	CreatedBy string `gorm:"column:created_by;type:varchar;size:45;" json:"createdBy"`
	//[ 6] created_date                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedDate time.Time `gorm:"column:created_date;type:timestamp;" json:"createdDate"`
	UpdatedBy   string    `gorm:"column:updated_by;type:varchar;size:45;" json:"updatedBy"`
	UpdatedDate time.Time `gorm:"column:updated_date;type:timestamp;" json:"updatedDate"`
}

// TableName sets the insert table name for this struct type
func (a *Appointment) TableName() string {
	return "appointment"
}
