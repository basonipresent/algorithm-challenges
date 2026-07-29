const printPascal = (N) => {
  for(let line = 1; line <= N; line++){
    let result = [];
    let C = 1;
    for(let i = 1; i <= line; i++){
      result.push(C);
      C = C * (line - i) / i;
    }
    console.log(result);
  }
}

console.log(printPascal(10));