const printWithoutLoop = (N) => {
  if (N > 0) {
    printWithoutLoop(N - 1);
    console.log(`${N}`);
  }
}

console.log(printWithoutLoop(10));