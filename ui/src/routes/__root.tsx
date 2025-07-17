import {
	AppShell,
	Box,
	Button,
	Group,
	MantineProvider,
	Title,
} from "@mantine/core";
import { useMutation } from "@tanstack/react-query";
import {
	createRootRouteWithContext,
	Outlet,
	useNavigate,
} from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { Link } from "@/components/link";
import type { RouterContext } from "@/types";

export const Route = createRootRouteWithContext<RouterContext>()({
	component: App,
});

// fetch user data in a loader

function App() {
	const navigate = useNavigate();

	const logoutMutation = useMutation({
		mutationFn: async () => {
			fetch("http://localhost:3000/auth/logout", {
				method: "POST",
				credentials: "include",
			});
		},
		onSuccess: async () => {
			await navigate({ to: "/" });
		},
	});

	return (
		<MantineProvider>
			<AppShell header={{ height: 60 }} padding="md">
				<AppShell.Header>
					<Group h="100%" px={"md"} align="center" justify="space-between">
						<Title order={1}>
							<Link to="/">🥗 Bowl</Link>
						</Title>

						{/* if s.IsAdmin {<a href="/admin">Admin</a>}| */}
						<Group>
							<Link to="/recipes/create">+ Recipe</Link>
							{/* TODO: Add conditional if not logged in */}
							<Button onClick={() => logoutMutation.mutateAsync()}>
								Logout
							</Button>
						</Group>
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
