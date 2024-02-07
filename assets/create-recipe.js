window.onload = () => {
  const ingredients = document.getElementById("ingredients")
  const addIngredientButton = document.getElementById("add-ingredient")

  addIngredientButton.addEventListener("click", () => {
    ingredients.insertAdjacentHTML("beforeend", `
        <input
          type="text"
          name="ingredient"
          placeholder="Ingredient"
          required
        />

        <input
          type="text"
          name="quantity"
          placeholder="Quantity"
          required
        />

        <select
          id="quantity-unit"
          name="quantity-unit"
          hx-get="/quantity-units"
          hx-trigger="load"
        />
    `)

    htmx.process(ingredients)
  })

  const steps = document.getElementById("steps")
  const addStepButton = document.getElementById("add-step")

  addStepButton.addEventListener("click", () => {
    steps.insertAdjacentHTML("beforeend", `
      <input
        type="text"
        name="step"
        placeholder="Step"
        required
      />
    `)

    htmx.process(steps)
  })
}
