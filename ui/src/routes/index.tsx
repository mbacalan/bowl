import { Container, Divider, Stack, Title } from "@mantine/core";
import { useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { Link } from "@/components/link";
import { recipesQueryOptions } from "@/queries";
import { RecipesList } from "./-components/recipe-list";

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
						<Link to="/recipes">📃 Recipes</Link>
					</Title>

					<Title order={3}>
						<Link to="/categories">📚 Categories</Link>
					</Title>
					<Divider />
				</Stack>

				{data.length ? (
					<Stack>
						<Title order={4}>Recently created:</Title>

						<RecipesList recipes={data} />
					</Stack>
				) : null}
			</Stack>
		</Container>
	);
}
