import { MantineProvider } from "@mantine/core";
import { createRootRoute, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

export const Route = createRootRoute({
	component: () => (
		<>
			<MantineProvider>
				<Outlet />
				<TanStackRouterDevtools />
			</MantineProvider>
		</>
	),
});
