class RecipeStore {
  constructor() {
    this.ingredients = []
  }

  addIngredient(ingredient, quantity, unit) {
    this.ingredients.push({ ingredient, quantity, unit })
    console.info(`Added ${quantity} of ${ingredient}: ${this.ingredients}`)
  }

  removeIngredient(ingredient) {
    this.ingredients = this.ingredients.filter(i => i.ingredient == ingredient)
    console.info(`Removed ${ingredient}: ${this.ingredients}`)
  }

  clean() {
    this.ingredients = []
    console.info(`Cleaned ingredients: ${this.ingredients}`)
  }
}

const rs = new RecipeStore()

function renderIngredientTable(ingredients) {
  const ingredientTable = document.getElementById('ingredient-list')

  const template = ingredients.map(i => {
    return `
      <tr>
        <td>
          ${i.ingredient}
        </td>
        <td>
          ${i.quantity} ${i.unit}
        </td>
      </tr>
    `
  }).join('')

  ingredientTable.innerHTML = template
}

function removeSearchResults() {
  const searchResultsEl = document.getElementById('search-results')

  while (searchResultsEl.firstChild) {
    searchResultsEl.removeChild(searchResultsEl.firstChild);
  }
}

function handleSelectIngredient(e) {
  const ingredientEl = document.querySelector('input[id="search"]')

  ingredientEl.value = e.dataset.ingredient
  removeSearchResults()
}

function handleAddIngredient() {
  const ingredientEl = document.querySelector('input[id="search"]')
  const quantityEl = document.getElementById('quantity')
  const quantityUnitEl = document.getElementById('quantity-unit')

  rs.addIngredient(ingredientEl.value, quantityEl.value, quantityUnitEl.value)
  renderIngredientTable(rs.ingredients)
}

window.onload = () => {
  document.body.addEventListener('htmx:configRequest', function (evt) {
    if (evt.target.id == 'recipe-form') {
      console.log('anan')
    }
  });
}
