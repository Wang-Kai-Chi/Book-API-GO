export default function Filters (filters = []) {
  const operate = FilterOperation('#searchInput')
  const filterItems = () => {
    let temp = ''
    for (const f of filters) {
      temp += FilterItem(f.name, f.value)
    }
    return temp
  }

  const searchFilter = document.querySelector('#searchFilter')

  searchFilter.innerHTML = filterItems()

  for (const f of filters) {
    document.querySelector(`#${f.value}checkbox`).onclick = (e) => operate.activeFilter(e)
    document.querySelector(`#${f.value}`).onchange = (e) => operate.setFilterValue(e)
  }
}
/**
     * Handling Events for FilterItem, such as onclick, onchange, and onfocus...
     *
     * @param {string} [searchSelector=""] css selector of your search input
     * @param {string} [name=""] string name of declared FilterOperation variable
     * @returns {object} FilterOperation object
     */
function FilterOperation (searchSelector = '') {
  const suffix = '='
  const regex = '&'
  const searchInput = document.querySelector(searchSelector)

  const removeFilter = (filter = '') => {
    const values = searchInput.value.split(regex)
    for (const s of values) {
      if (s.includes(filter)) {
        searchInput.value = searchInput.value.replace(s + regex, '')
        break
      }
    }
  }

  const addFilter = (paramId = '', filter = '') => {
    const param = document.querySelector(`#${paramId}`)
    searchInput.value += `${filter}${regex}`
    param.focus()
  }

  const activeFilter = (event) => {
    const checkbox = event.target
    const filter = `${checkbox.value}${suffix}`

    if (checkbox.checked && !searchInput.value.includes(filter)) {
      addFilter(checkbox.value, filter)
    } else {
      removeFilter(filter)
    }
  }

  const setFilterValue = (event) => {
    const param = event.target
    const filter = `${param.id}${suffix}`

    removeFilter(filter)

    addFilter(param.id, filter)

    const input = searchInput.value

    const insertString = (position = 0, str = '') => {
      return input.slice(0, position) + str + input.slice(position)
    }
    if (input.includes(param.id) && !input.includes(param.value)) {
      const insertPosition = input.indexOf(param.id) + param.id.length + suffix.length
      searchInput.value = insertString(insertPosition, param.value)
    }
  }

  return {
    activeFilter: (event) => activeFilter(event),
    setFilterValue: (event) => setFilterValue(event)
  }
}

/**
 *  Returning a html with list item, checkbox, label, and text input.
 *  This is use to generate parameter in another search input.
 *
 * @param {string} [name="name"] name that display on FilterItem
 * @param {string} [value="value"] value that generate on search input
 * @param {string} [FilterOpStr=""] variable name of FilterOperation
 * @return {string}
 */
function FilterItem (name = 'name', value = 'value') {
  return /* html */`
            <li class="list-group-item">
                <input class="form-check-input me-1" id="${value}checkbox" type="checkbox" value="${value}">
                <label class="form-check-label" for="firstCheckbox">${name}</label>
                <input class="" type="text" id="${value}">
            </li>
        `
}
