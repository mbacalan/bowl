class RecipeIngredientStore {
  constructor() {
    this.ingredients = []
  }

  addRecipeIngredient(ri) {
    this.ingredients.push(ri)
  }

  removeRecipeIngredient(ingredient) {
    this.ingredients = this.ingredients.filter(i => i.ingredient == ingredient)
  }
}

const rs = new RecipeIngredientStore()

function renderIngredientTable(ingredients) {
  const ingredientTable = document.getElementById('ingredient-list')

  const template = ingredients.map(i => {
    return `
      <tr>
        <td>
          <input class="table-input" type="text" name="ingredient" readonly value="${i.ingredient}" />
        </td>
        <td>
          <input class="table-input" type="text" name="quantity" readonly value="${i.quantity}" />
        </td>
        <td>
          <input class="table-input" type="text" name="quantity-unit" readonly value="${i.unit}" />
        </td>
      </tr>
    `
  }).join('')

  ingredientTable.innerHTML = template
}

function clearIngredientFields() {
  const ingredientEl = document.getElementById('ingredient')
  const quantityEl = document.getElementById('quantity')
  const quantityUnitEl = document.getElementById('quantity-unit')

  ingredientEl.value = ''
  quantityEl.value = ''
  quantityUnitEl.value = quantityUnitEl.options[0].value
}

function addIngredient() {
  const ingredientEl = document.getElementById('ingredient')
  const quantityEl = document.getElementById('quantity')
  const quantityUnitEl = document.getElementById('quantity-unit')

  rs.addRecipeIngredient({
    ingredient: ingredientEl.value,
    quantity: quantityEl.value,
    unit: quantityUnitEl.value
  })

  clearIngredientFields()
  renderIngredientTable(rs.ingredients)
}
