const main = () => {
  const expressions = [
    ["+", 1, 2],
    ["-", 3, 4],
    ["*", 5, 6],
    ["/", 7, 8],
    [7, 8],
    ["a", 7, 8],
  ];

  expressions.forEach((exp) => {
    const [op, a, b] = exp;
    if (!map.has(op)) {
      console.log(`Unknown operator: ${op}`);
      return;
    }
    if (typeof a !== "number" || typeof b !== "number") {
      console.log(`Invalid expression`);
      return;
    }
    console.log(`${a} ${op} ${b} = ${map.get(op)(a, b)}`);
  });
};

// create map
const map = new Map();
map.set("+", (a, b) => a + b);
map.set("-", (a, b) => a - b);
map.set("*", (a, b) => a * b);
map.set("/", (a, b) => a / b);


main();;