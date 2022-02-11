const evaluate = (expr) => {
  let values = [];
  let ops = [];

  for (let i = 0; i < expr.length; i++) {
    if (expr[i] === ' ') {
      continue;
    }

    if (expr[i] >= '0' && expr[i] <= '9') {
      let buffer = '';
      while (i < expr.length && expr[i] >= '0' && expr[i] <= '9') {
        buffer += expr[i++];
      }
      values.push(parseInt(buffer));
      i--;
    } else if (expr[i] === '+' || expr[i] === '-' || expr[i] === '*' || expr[i] === '/') {
      while (ops.length) {
        values.push(applyOp(ops.pop(), values.pop(), values.pop()));
      }
      ops.push(expr[i]);
    }
  }

  while(ops.length){
    values.push(applyOp(ops.pop(), values.pop(), values.pop()));
  }

  return values.pop();
}

const applyOp = (op, b, a) => {
  switch (op) {
    case '+':
      return a + b;
    case '-':
      return a - b;
    case '*':
      return a * b;
    case '/':
      if (b === 0) {
        throw 'Cannot divide by zero';
      }
      return Math.floor(a / b);
  }
  return 0;
};

console.log(evaluate('2+3-5-10'));