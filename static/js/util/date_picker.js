export default function DatePicker () {
  const setDatePicker = (selector) => {
    const publicationDate = document.querySelector(`#${selector}`)

    publicationDate.type = 'date'
    publicationDate.min = '1850-01-01'

    const currentDate = new Date().toJSON().slice(0, 10)
    publicationDate.max = `${currentDate}`
  }

  return {
    set: (htmlId = '') => setDatePicker(htmlId)
  }
}
