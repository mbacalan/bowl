/**
  * @param {PointerEvent} event
*/
function removeInputGroup(event) {
  const inputGroup = event.target.closest(".input-group")
  if (inputGroup.parentElement.children.length <= 1) {
    return
  }

  inputGroup.remove()
}

window.onload = () => {
  const ingredients = document.getElementById("ingredients")
  const addIngredientButton = document.getElementById("add-ingredient")

  addIngredientButton.addEventListener("click", () => {
    ingredients.insertAdjacentHTML("beforeend", `
      <div class="input-group">
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
        ></select>

        <span class="remove remove-ingredient">X</span>
      </div>
    `)

    document.querySelectorAll(".remove-ingredient").forEach(button => {
      button.addEventListener("click", removeInputGroup)
    })

    htmx.process(ingredients)
  })

  const steps = document.getElementById("steps")
  const addStepButton = document.getElementById("add-step")

  addStepButton.addEventListener("click", () => {
    steps.insertAdjacentHTML("beforeend", `
      <div class="input-group">
        <input
          type="text"
          name="step"
          placeholder="Step"
          required
        />

        <span class="remove remove-step">X</span>
      </div>
    `)

    document.querySelectorAll(".remove-step").forEach(button => {
      button.addEventListener("click", removeInputGroup)
    })
  })
}
