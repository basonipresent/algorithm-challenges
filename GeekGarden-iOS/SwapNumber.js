function swap() {
  let a = 3; let b = 7;
  a = a + b; b = a - b; a = a - b;
  console.log(`a=${a}, b=${b}`)
}

swap()