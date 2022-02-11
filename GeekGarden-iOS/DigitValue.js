function printDigit(params) {
  let numbers = params.replace(/[^0-9]/g, '');
  let isFirst = true;
  while (numbers.length > 0) {
    if (isFirst) {
      console.log(numbers);
      isFirst = false;
    } else {
      let square = numbers.length;
      let result = Math.floor(numbers.charAt(0)) * Math.pow(10, square - 1);
      console.log(result);
      numbers = numbers.substring(1);
    }
  }
}

printDigit("9.86-A5.321");