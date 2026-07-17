package campaign

type Repository interface {
	Save(c *Campaign) error
	Get() ([]Campaign, error)
	GetBy(id string) (*Campaign, error)
}
