import { AppShell } from "@mantine/core";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
	component: App,
});

function App() {
	return (
		<AppShell header={{ height: 60 }} padding="md">
			<AppShell.Header>
				<div>Bowl</div>
			</AppShell.Header>

			<AppShell.Main>Main</AppShell.Main>

			<AppShell.Footer>Bilingual Recipes - 2025</AppShell.Footer>
		</AppShell>
	);
}
