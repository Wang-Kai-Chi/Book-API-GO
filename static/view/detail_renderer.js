/**
 * Bootstrap input group
 *
 * @param {string} [name=""] name of input group
 * @param {string} [id=""] input id
 * @return {string} html element string 
 */
function DetailHTML(name = "", id = "") {
    return /*html*/`
        <div class="form-floating mb-3" id="form${id}">
            <input type="text" class="form-control" id="${id}" placeholder="none" disabled>
            <label for="${id}">${name}</label>
        </div>
    `
}

/**
 * Render list items of detail
 *
 * @param {string} [selector=""] selector of detail list
 * @return {object} 
 */
function DetailRenderer(selector = "") {
    const render = (details) => {
        const items = () => {
            let temp = ""
            const keys = Object.keys(details)
            for (const i in keys) {
                const k = keys[i]
                temp += DetailHTML(details[k], k)
            }

            return temp
        }
        const detailDisplay = document.querySelector(selector)
        detailDisplay.innerHTML = items()
    }
    return {
        render: (details = {}) => render(details)
    }
}