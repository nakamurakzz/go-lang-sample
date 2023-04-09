class Stack<T> {
  vals: T[];

  constructor() {
    this.vals = [];
  }

  push(val: T) {
    this.vals.push(val);
  }

  pop(): T {
    if (this.vals.length === 0) {
      throw new Error('Stack is empty');
    }
    return this.vals.pop() as T;
  }
}

const main = () => {
  const stack = new Stack<string>();
  stack.push('a');
  stack.push('b');
  stack.push('c');
  for (let index = 0; index < stack.vals.length; index++) {
    console.log(stack.vals[index]);
  }

  const stackInt = new Stack<number>();
  stackInt.push(1);
  stackInt.push(2);
  stackInt.push(3);
  for (let index = 0; index < stackInt.vals.length; index++) {
    console.log(stack.vals[index]);
  }
}

main();