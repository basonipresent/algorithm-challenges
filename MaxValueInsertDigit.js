function solution(N) {
  // write your code in JavaScript (Node.js 8.9.4)
  let intDigit = 5;

  let negative = parseInt(N / Math.abs(N));
  N = Math.abs(N);
  let num = parseInt(N);

  let maxValue = -80005;
  let counter = 0;
  let positions = 1;

  while (parseInt(num) > 0) {
    counter++;
    num = num / 10;
  }

  for (let i = 0; i <= counter; i++) {
    let newValue =
      ((N / positions) * (positions * 10)) +
      (intDigit * positions) + (N % positions);

    if (newValue * negative > maxValue) {
      maxValue = newValue * negative;
    }
    positions = positions * 10;
  }

  return maxValue;
}

console.log(solution(268));
console.log(solution(670));
console.log(solution(-999));
