https://www.thatjsdude.com/interview/js2.html#insertSpace

console.log("Start");

setTimeout(() => console.log("Timeout"), 0);
setImmediate(() => console.log("Set Immediate"));

process.nextTick(() => console.log("Next Tick"));
Promise.resolve().then(() => console.log("Promise resolved"));

console.log("End");


 "Start" → Synchronous
2️⃣ "End" → Synchronous
3️⃣ "Next Tick" → Microtask (process.nextTick)  ----- supoer high pririty afte sync alls. it can block the main thread
4️⃣ "Promise resolved" → Microtask (Promise)
5️⃣ "Timeout" → Macrotask (setTimeout)
6️⃣ "Set Immediate" → Macrotask (setImmediate)