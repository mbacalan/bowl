import { Container, Group, List, Stack, Text, Title } from "@mantine/core";
import { useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { Link } from "@/components/link";
import { recipeQueryOptions } from "@/queries";

export const Route = createFileRoute("/recipes/$id")({
	component: RecipeComponent,
	loader: async ({ context: { queryClient }, params: { id } }) => {
		await queryClient.ensureQueryData(recipeQueryOptions(id));
	},
});

function RecipeComponent() {
	const id = Route.useParams().id;
	const { data: recipe } = useSuspenseQuery(recipeQueryOptions(id));

	return (
		<Container>
			<Stack component="article">
				<Group justify="space-between">
					<Title order={2}>{recipe.Name}</Title>

					<Link to={`/recipes/${recipe.ID}/edit`}>Edit</Link>
				</Group>

				<Text>
					Prep: {recipe.PrepDuration}min - Cooking: {recipe.CookDuration}min
				</Text>

				<Stack>
					<Stack>
						<Title order={3}>Ingredients</Title>
						<List>
							{recipe.RecipeIngredients.map((ingredient) => (
								<List.Item key={ingredient.ID}>
									{ingredient.Ingredient.Name} - {ingredient.Quantity},{" "}
									{ingredient.QuantityUnit.Name}
								</List.Item>
							))}
						</List>
					</Stack>

					<Stack>
						<Title order={3}>Steps</Title>
						<List type="ordered">
							{recipe.Steps.map((step) => (
								<List.Item key={step.ID}>{step.Step}</List.Item>
							))}
						</List>
					</Stack>
				</Stack>

				{recipe.Categories.length ? (
					<Stack>
						<Title order={3}>Categories</Title>
						<p>
							{recipe.Categories.map((category) => (
								<Link key={category.ID} to={`/categories/${category.ID}`}>
									{category.Name}
								</Link>
							))}
						</p>
					</Stack>
				) : null}
			</Stack>
		</Container>
	);
}
