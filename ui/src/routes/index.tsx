import {
	Anchor,
	Container,
	Divider,
	List,
	Stack,
	Text,
	Title,
} from "@mantine/core";
import { useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

type RecipeList = {
	ID: number;
	Name: string;
	PrepDuration: number;
	CookDuration: number;
};

const recipesQueryOptions = {
	queryKey: ["recipes"],
	queryFn: async () => {
		const resp = await fetch("http://localhost:3000/recipes", {
			method: "GET",
			credentials: "include",
		});

		const data: RecipeList[] = await resp.json();
		return data;
	},
};

export const Route = createFileRoute("/")({
	component: App,
	loader: async ({ context: { queryClient } }) => {
		await queryClient.ensureQueryData(recipesQueryOptions);
	},
});

function App() {
	const { data } = useSuspenseQuery(recipesQueryOptions);

	return (
		<Container>
			<Stack>
				<Stack>
					<Title order={2}>A database of your own recipes</Title>
					<Divider />

					<Title order={3}>
						<Anchor href="/recipes">📃 Recipes</Anchor>
					</Title>

					<Title order={3}>
						<Anchor href="/categories">📚 Categories</Anchor>
					</Title>
					<Divider />
				</Stack>

				{data.length ? (
					<Stack>
						<Title order={4}>Recently created:</Title>

						<List>
							{data.map((recipe) => (
								<List.Item key={recipe.ID}>
									<Anchor href={`/recipes/${recipe.ID}`}>{recipe.Name}</Anchor>
									{" - "}
									<Text span>
										Prep: {recipe.PrepDuration}min{" - "}
										Cooking: {recipe.CookDuration}min
									</Text>
								</List.Item>
							))}
						</List>
					</Stack>
				) : null}
			</Stack>
		</Container>
	);
}
