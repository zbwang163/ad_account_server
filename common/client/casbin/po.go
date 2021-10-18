package casbin

type ADCasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:10;uniqueIndex:unique_index"`
	V0    string `gorm:"size:128;uniqueIndex:unique_index"`
	V1    string `gorm:"size:128;uniqueIndex:unique_index"`
	V2    string `gorm:"size:128;uniqueIndex:unique_index"`
	V3    string `gorm:"size:128;uniqueIndex:unique_index"`
	V4    string `gorm:"size:50;uniqueIndex:unique_index"`
	V5    string `gorm:"size:50;uniqueIndex:unique_index"`
}

func (ADCasbinRule) TableName() string {
	return "ad_role_policy"
}
