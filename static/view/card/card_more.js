/**
 *Removing product card and delete data from DB
 *
 * @param {*} cardId
 */
function handleDeleteProduct(cardId) {
    setCurrentCardValue(cardId)

    let body = `[${localStorage.getItem("currentProduct")}]`
    if (confirm('Confirm delete?')) {
        document.querySelector(`#card${cardId.id}`).hidden = true
        fetch(`/api/v1/product/delete`, {
            method: "DELETE",
            body: body,
            headers: new Headers({
                "Content-Type": "application/json",
            }),
        }).then(res => res.json())
            .catch(err => console.log(err))
            .then(response => console.log("Success", response))
    }
}
/**
 *Saving current Card Value to LocalStorage
 *
 * @param {string} [cardId=""] id of card
 */
function setCurrentCardValue(cardId = "") {
    const key = "currentProduct"
    localStorage.setItem(key, cardId.innerHTML)
}
