import type { RecipeList } from "./types";

export const recipesQueryOptions = {
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
