
const obj = {
  a: '1',
  b: '2',
  c: 'c'
}

const obj2 = {
  d: 'd'
}

const eq = {
  ...obj,
  c: 'dennis',
  ...obj2,
}

console.log("eq", eq)
console.log("obj", obj)
