import CardRenderer from './util/card_renderer.js'

export default function RecentProduct () {
  const RECENT_PRODUCT_COUNT = 8
  fetch(`/api/v1/product/query/new/${RECENT_PRODUCT_COUNT}`)
    .then(data => data.json())
    .then(value => CardRenderer('#recentProduct').render(value))
    .catch(err => console.log(err))
}
