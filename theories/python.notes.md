



==========compilation==================

So, how does Python work under the hood?
✅ Step-by-step process in CPython (the default Python implementation):
You write source code
File: example.py
Contains your human-readable Python code.

Python compiles it into bytecode

Internally, the interpreter converts .py to .pyc (bytecode).

Bytecode = low-level instructions understood by Python Virtual Machine.

This is not machine code (like in C/C++), and not directly run by your OS or CPU.

Bytecode is run by Python Virtual Machine (PVM)

The PVM is the runtime engine.

It reads and executes the bytecode, line by line.

This is what makes Python feel "interpreted".

🧠 Key Terms Explained
1. Compiler
A compiler translates source code into machine code or another intermediate form.

In Python, the .py is compiled to bytecode (.pyc) — not to machine code.

So, Python uses a compiler internally.

2. Interpreter
An interpreter executes code line-by-line, without converting it to machine code first.

CPython interprets the bytecode — so Python is interpreted at runtime.

3. Virtual Machine (PVM)
Think of it as a software CPU that runs bytecode.

It makes Python platform-independent because .pyc bytecode runs on the PVM, not on the actual CPU.

In CPython, the PVM is part of the interpreter.


whwne we download python , we usually download Cpython which interpret bytecode. 
when we wanna use pypy we need to download separalty.  and use pypy script.py  it compiles code into binary on runtime
Instead of interpreting bytecode line by line like CPython, it translates frequently-used code paths into machine code at runtime.Compiles bytecode to machine code (JIT)



=========execution=============

🟩 Python
✅ Compilation Phase
Source Code (.py) is parsed to generate an AST.

Compiled into Bytecode:

Stored in .pyc files (cached in __pycache__/ folder).

Bytecode is intermediate and platform-independent.

✅ Execution Phase
Python Virtual Machine (PVM) reads .pyc and executes instructions.

Execution is line-by-line (blocking by default).

No Event Loop by default:

Async execution via asyncio or multithreading/multiprocessing.

Async operations need explicit setup.



Python gives recusiton error . a recursion depth of 1000 is limited.
✅ How to Solve Python's Recursion Issues?
1️⃣ Increase Recursion Limit (⚠ Not Always Recommended)
You can increase the recursion limit using sys.setrecursionlimit(), but this may cause memory overflow.


User iterative way

3️⃣ Use Memoization (functools.lru_cache) to Reduce Calls
If recursion is unavoidable, memoization (caching results) can improve performance.

🔹 Example: Fibonacci with lru_cache
from functools import lru_cache

@lru_cache(maxsize=None)  # Stores previous results
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

print(fibonacci(50))  # Computes efficiently!




