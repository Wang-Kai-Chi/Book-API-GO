import CardRenderer from './card_renderer.js'

const RECENT_PRODUCT_COUNT = 8
fetch(`/api/v1/product/query/new/${RECENT_PRODUCT_COUNT}`)
  .then(data => data.json())
  .then(value => CardRenderer('#recentProduct').render(value))
  .catch(err => console.log(err))