function solution(A) {
  if (A.length === 1) return 0;
  var len = A.length,
    result = 0;

  for (var i = 0; i < len - 1; i++) {
    if (A[i] === A[i + 1]) {
      result++;
    }
  }

  var revers = 0;
  for (var i = 0; i < len; i++) {
    var count = 0;
    if (i > 0) {
      count = A[i - 1] !== A[i] ? count + 1 : count - 1;
    }
    if (i < len - 1) {
      count = A[i] !== A[i + 1] ? count + 1 : count - 1;
    }
    revers = Math.max(revers, count);
  }
  return result + revers;
}
console.log(solution([1, 1, 0, 1, 0, 0, 1, 1])); // 5
console.log(solution([1, 1, 1, 1, 1, 0, 1, 1])); // 7
console.log(solution([1, 0, 1])); // 2
console.log(solution([0, 1, 0])); // 2
console.log(solution([1, 1, 0, 0, 0])); // 3
