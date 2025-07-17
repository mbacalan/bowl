import { queryOptions } from "@tanstack/react-query";
import type { Recipe, RecipeList } from "./types";

export const recipeQueryOptions = (id: string) => {
	return queryOptions({
		queryKey: ["recipes", id],
		queryFn: async () => {
			const resp = await fetch(`http://localhost:3000/recipes/${id}`, {
				method: "GET",
				credentials: "include",
			});

			const data: Recipe = await resp.json();
			return data;
		},
	});
};

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
