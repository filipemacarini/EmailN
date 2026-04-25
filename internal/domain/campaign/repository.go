package campaign

type Repository interface {
	Save(c *Campaign) error
}
