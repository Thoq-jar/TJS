let randomNumbers = '';
for (let i = 0; i < 1000; i++) {
  randomNumbers += Math.floor(Math.random() * 100);
  if (i < 999) {
    randomNumbers += '.';
  }
}
console.log("Generated Random Numbers:" + randomNumbers);

function parseRandomNumbers(randomNumbersString) {
  const numbersArray = [];
  let currentNumber = '';

  for (let char of randomNumbersString) {
    if (char === '.') {
      numbersArray.push(Number(currentNumber));
      currentNumber = '';
    } else {
      currentNumber += char;
    }
  }
  if (currentNumber) {
    numbersArray.push(Number(currentNumber));
  }

  return numbersArray;
}

const parsedNumbers = parseRandomNumbers(randomNumbers);
console.log("Parsed Numbers Array:" + parsedNumbers);