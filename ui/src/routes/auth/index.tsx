import { Container, Divider, Stack, Text } from "@mantine/core";
import { createFileRoute } from "@tanstack/react-router";
import { LoginForm } from "./-login";
import { RegisterForm } from "./-register";

export const Route = createFileRoute("/auth/")({
	component: Auth,
});

function Auth() {
	return (
		<Container>
			<Stack>
				<RegisterForm />

				<Divider label={<Text size={"xl"}>Already have an account?</Text>} />

				<LoginForm />
			</Stack>
		</Container>
	);
}
