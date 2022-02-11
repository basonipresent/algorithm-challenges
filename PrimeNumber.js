const isPrime = (number) => {
  if(number < 2){
    return false;
  }

  for(let i = 2; i < number; i++){
    if(number % i === 0){
      return false;
    }
  }
  return true;
}

const printPrime = (N) => {
  let result = [];
  for(let i = 1; i <= N; i++){
    if(isPrime(i)){
      result.push(i);
    }
  }
  return result;
}

console.log(isPrime(97));
console.log(printPrime(100));