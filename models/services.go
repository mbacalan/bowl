package models

type Services struct {
	AuthService         AuthService
	AdminService        AdminService
	RecipeService       RecipeService
	QuantityUnitService QuantityUnitService
	CategoryService     CategoryService
}

type AuthService interface {
	Signup(name string, password string) (User, error)
	Login(name string, password string) (User, error)
}

type AdminService interface {
	IsAdmin(user uint) bool
	GetIngredients() ([]Ingredient, error)
	CreateIngredient(name string) (Ingredient, error)
	DeleteIngredient(id uint) (bool, error)
	GetQuantityUnits() ([]QuantityUnit, error)
	CreateQuantityUnit(name string) (QuantityUnit, error)
	DeleteQuantityUnit(id uint) (bool, error)
}

type CategoryService interface {
	Get(user uint, id int) (Category, error)
	GetAll(user uint) ([]Category, error)
}

type HomeService interface {
	GetRecent(user uint, limit int) ([]Recipe, error)
}

type QuantityUnitService interface {
	GetAll() ([]QuantityUnit, error)
}

type RecipeService interface {
	Get(user uint, id int) (Recipe, error)
	GetAll(user uint) ([]Recipe, error)
	GetRecent(user uint, limit int) ([]Recipe, error)
	Create(data RecipeData) (Recipe, error)
	Update(id int, data RecipeData) (Recipe, error)
}

type RecipeData struct {
	Name          string
	Steps         []string
	Ingredients   []string
	Quantities    []string
	QuantityUnits []string
	Categories    []string
	PrepDuration  uint
	CookDuration  uint
	UserID        uint
}
