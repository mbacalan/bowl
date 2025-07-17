import { List, Text } from "@mantine/core";
import { Link } from "@/components/link";
import type { RecipeList } from "@/queries/types";

type Props = {
	recipes: RecipeList[];
};

export function RecipesList({ recipes }: Props) {
	return (
		<List>
			{recipes.map((recipe) => (
				<List.Item key={recipe.ID}>
					<Link to={`/recipes/${recipe.ID}`}>{recipe.Name}</Link>
					{" - "}
					<Text span>
						Prep: {recipe.PrepDuration}min{" - "}
						Cooking: {recipe.CookDuration}min
					</Text>
				</List.Item>
			))}
		</List>
	);
}
