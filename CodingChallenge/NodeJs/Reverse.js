const reverseNumber = (input) => {
  let result = 0;
  while(input > 0){
    result = result * 10 + input % 10;
    input = Math.floor(input / 10);
  }
  return result;
}

const reverseString = (input) => {
  let result = '';
  for(let i = input.length; i > 0; i--) {
    result += input[i - 1];
  }
  return result;
}

console.log(reverseNumber(1234567));
console.log(reverseString('qwertyuiop'));