export default function DatePicker () {
  const setDatePicker = (selector) => {
    const currentDate = new Date().toJSON().slice(0, 10)
    const publicationDate = document.querySelector(selector)

    publicationDate.type = 'date'
    publicationDate.min = '1900-01-01'
    publicationDate.max = `${currentDate}`
  }

  return {
    set: (selector = '') => setDatePicker(selector)
  }
}
