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
  `created_at` timestamp NOT NULL,
  `updated_by` varchar(45) NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  KEY `appointment_FK` (`created_by`),
  KEY `appointment_FK_1` (`updated_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": "doTEABUpRhFSSgRRMKjabiTyZ",    "topic": "eKuXZWdZOvQQlmAndlREUojnp",    "state": "AUxUQUfBelceVefhDljVDraZm",    "description": "rFyuLOSWnLdWekBAjlSSLRWGZ",    "status": "mAMvBOCLZyhSaLqxFmiXqXykE",    "createdBy": "APVeAEwutamndSuSZbWbmFVNh",    "createdAt": "2143-08-27T05:02:07.35775473+07:00",    "updatedBy": "JapyUIFmtaLlgNlUpnkVyoTgA",    "updatedAt": "2286-09-30T10:10:37.272335239+07:00"}



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
	Description null.String `gorm:"column:description;type:text;size:65535;" json:"description"`
	//[ 4] status                                         varchar(1)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1       default: []
	Status string `gorm:"column:status;type:varchar;size:1;" json:"status"`
	//[ 5] created_by                                     varchar(45)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 45      default: []
	CreatedBy string `gorm:"column:created_by;type:varchar;size:45;" json:"createdBy"`
	//[ 6] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;" json:"createdAt"`
	//[ 7] updated_by                                     varchar(45)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 45      default: []
	UpdatedBy string `gorm:"column:updated_by;type:varchar;size:45;" json:"updatedBy"`
	//[ 8] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;" json:"updatedAt"`
}

// TableName sets the insert table name for this struct type
func (a *Appointment) TableName() string {
	return "appointment"
}
