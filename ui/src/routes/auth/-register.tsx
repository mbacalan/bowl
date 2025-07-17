import { Button, Stack, TextInput } from "@mantine/core";
import { useForm } from "@tanstack/react-form";
import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { z } from "zod";

const registerSchema = z.object({
	username: z.string(),
	password: z.string(),
});

export function RegisterForm() {
	const navigate = useNavigate();

	const mutation = useMutation({
		mutationFn: async (value: z.infer<typeof registerSchema>) => {
			fetch("http://localhost:3000/auth/signup", {
				method: "POST",
				body: new URLSearchParams(value),
			});
		},
		onSuccess: async () => {
			await navigate({ to: "/" });
		},
	});

	const { Field, handleSubmit } = useForm({
		defaultValues: {
			username: "",
			password: "",
		},
		onSubmit: async ({ value }) => {
			await mutation.mutateAsync(value);
		},
		validators: {
			onSubmit: registerSchema,
		},
	});

	return (
		<form
			onSubmit={(e) => {
				e.preventDefault();
				handleSubmit();
			}}
		>
			<Stack>
				<Field
					name="username"
					children={({ state, handleChange, handleBlur }) => (
						<TextInput
							label="Username"
							defaultValue={state.value}
							onChange={(e) => handleChange(e.target.value)}
							onBlur={handleBlur}
							error={state.meta.errors.join(",")}
						/>
					)}
				/>

				<Field
					name="password"
					children={({ state, handleChange, handleBlur }) => (
						<TextInput
							label="Password"
							type="password"
							defaultValue={state.value}
							onChange={(e) => handleChange(e.target.value)}
							onBlur={handleBlur}
							error={state.meta.errors.join(",")}
						/>
					)}
				/>

				<Button variant="filled" type="submit">
					Register
				</Button>
			</Stack>
		</form>
	);
}
