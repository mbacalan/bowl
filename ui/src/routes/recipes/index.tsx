import { Container, Divider, Stack, Text, Title } from "@mantine/core";
import { useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { Link } from "@/components/link";
import { recipesQueryOptions } from "@/queries";
import { RecipesList } from "../-components/recipe-list";

export const Route = createFileRoute("/recipes/")({
	component: Recipes,
	loader: async ({ context: { queryClient } }) => {
		await queryClient.ensureQueryData(recipesQueryOptions);
	},
});

function Recipes() {
	const { data } = useSuspenseQuery(recipesQueryOptions);

	return (
		<Container>
			{data.length === 0 ? (
				<Stack>
					<Title order={2}>No recipes found :(</Title>
					<Text>
						Why don't you <Link to="/recipes/create">create one</Link>?
					</Text>
				</Stack>
			) : (
				<Stack>
					<Title order={2}>All Recipes</Title>
					<Divider />
					<RecipesList recipes={data} />
				</Stack>
			)}
		</Container>
	);
}
