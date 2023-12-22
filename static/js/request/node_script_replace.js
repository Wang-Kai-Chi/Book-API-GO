export default function NodeScriptReplace (node) {
  if (nodeScriptIs(node) === true) {
    node.parentNode.replaceChild(nodeScriptClone(node), node)
  } else {
    let i = -1; const children = node.childNodes
    while (++i < children.length) {
      NodeScriptReplace(children[i])
    }
  }

  return node
}
function nodeScriptClone (node) {
  const script = document.createElement('script')
  script.text = node.innerHTML

  let i = -1
  const attrs = node.attributes

  while (++i < attrs.length) {
    const attr = attrs[i]
    script.setAttribute(attr.name, attr.value)
  }
  return script
}

function nodeScriptIs (node) {
  return node.tagName === 'SCRIPT'
}
