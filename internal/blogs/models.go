package blogs

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type Blog struct {
	BlogId        uuid.UUID `json:"blogId" gorm:"type:uuid; default:uuid_generate_v4(); not null"`
	BlogName      string    `json:"blogName"`
	BlogDesc      *string   `json:"blogDesc"`
	BlogAuthor    string    `json:"blogAuthor"`
	NoOfCopies    uint      `json:"noOfCopies"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	BlogPublisher uuid.UUID `json:"blogPublisher" gorm:"type:uuid; default:uuid_generate_v4(); not null"`
}
