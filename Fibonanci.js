const fibonanciGenerator = (number) => {
    let result = [];
    result[0] = 0;
    result[1] = 1;

    for (i = 2; i <= number; i++) {
        result[i] = result[i - 2] + result[i - 1];
    }

    return result;
}

const fibonanciSum = (array) => {
    let result = 0;
    for(let i = 0; i < array.length; i++){
        result += array[i];
    }
    return result;
}

console.log(fibonanciGenerator(10));
console.log(fibonanciSum(fibonanciGenerator(10)));