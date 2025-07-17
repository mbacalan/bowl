import {
	Anchor,
	AppShell,
	Box,
	Group,
	MantineProvider,
	Title,
} from "@mantine/core";
import { createRootRouteWithContext, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import type { RouterContext } from "@/types";

export const Route = createRootRouteWithContext<RouterContext>()({
	component: App,
});

// fetch user data and recipes in a loader

function App() {
	return (
		<MantineProvider>
			<AppShell header={{ height: 60 }} padding="md">
				<AppShell.Header>
					<Group h="100%" px={"md"} align="center" justify="space-between">
						<Title order={1}>
							<Anchor href="/">🥗 Bowl</Anchor>
						</Title>

						<div>
							{/* if s.IsAdmin {<a href="/admin">Admin</a>}| */}
							<Anchor href="/recipes/create">+ Recipe</Anchor>
						</div>
					</Group>
				</AppShell.Header>

				<AppShell.Main>
					<Outlet />
				</AppShell.Main>

				<AppShell.Footer>
					<Box p={"md"} component="p">
						🥗, Bilingual Recipes - 2025
					</Box>
				</AppShell.Footer>
			</AppShell>

			<TanStackRouterDevtools />
		</MantineProvider>
	);
}
