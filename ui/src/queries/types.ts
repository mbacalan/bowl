type Ingredient = {
	ID: number;
	Name: string;
};

type QuantityUnit = {
	ID: number;
	Name: string;
};

type Category = {
	ID: number;
	Name: string;
};

type Step = {
	ID: number;
	Step: string;
};

type RecipeIngredient = {
	ID: number;
	Ingredient: Ingredient;
	QuantityUnit: QuantityUnit;
	Quantity: number;
};

export type Recipe = {
	ID: number;
	Name: string;
	PrepDuration: number;
	CookDuration: number;
	RecipeIngredients: RecipeIngredient[];
	Steps: Step[];
	Categories: Category[];
};

export type RecipeList = {
	ID: number;
	Name: string;
	PrepDuration: number;
	CookDuration: number;
};
