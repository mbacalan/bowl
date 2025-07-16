import { MantineProvider } from '@mantine/core'
import { Outlet, createRootRoute } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools'

export const Route = createRootRoute({
  component: () => (
    <>
      <MantineProvider>
        <Outlet />
        <TanStackRouterDevtools />
      </MantineProvider>
    </>
  ),
})
