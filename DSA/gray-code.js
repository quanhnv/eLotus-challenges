/**
 * @param {number} n
 * @return {number[]}
 */
var grayCode = function (n) {
  if (n === 0) {
    return [0];
  }

  const result = [];
  const visited = new Set();
  visited.add(0);
  result.push(0);

  backtrack(result, visited, n);

  return result;
};

function backtrack(result, visited, n) {
  if (result.length === Math.pow(2, n)) {
    return true;
  }

  const lastNum = result[result.length - 1];

  for (let i = 0; i < n; i++) {
    const mask = 1 << i;
    const newNum = lastNum ^ mask;

    if (!visited.has(newNum)) {
      visited.add(newNum);
      result.push(newNum);

      if (backtrack(result, visited, n)) {
        return true;
      }

      visited.delete(newNum);
      result.pop();
    }
  }

  return false;
}

const n = 2;
const sequence = grayCode(n);
console.log(sequence);
