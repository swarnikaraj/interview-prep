
🟨 JavaScript
✅ Compilation Phase (Before Execution)
Parsing: JS engine parses the code and converts it into an Abstract Syntax Tree (AST).

Compilation (Just-In-Time):

Modern engines like V8 compile to bytecode and optimize it to machine code.

Uses JIT (Just-In-Time) compilation — combines interpretation + compilation.

Hoisting:

Declarations (like var, function) are hoisted to the top.

JS creates a Global Execution Context (memory setup before execution).

Variable & Function Memory Allocation:

All var, let, const, and functions are stored in memory.

✅ Execution Phase
Runs Line-by-Line (single thread).

Creates Execution Contexts for functions (with scope, this, etc.).

Event Loop kicks in for async tasks:

setTimeout, fetch, Promise, etc. are sent to the Web APIs.

After completion, callback is sent to the Callback Queue, then executed in turn.

🔁 Event Loop Diagram
css
Copy
Edit
 Call Stack        Web APIs         Callback Queue
    ↓                 ↓                  ↓
[Main Code] → [fetch(), timers] → [console.log()]


