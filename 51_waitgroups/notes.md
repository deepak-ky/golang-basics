wg.Add(1)
can also be written as
wg.Add(2)
wg.Add(-1)

BUT THIS IS A VERY BAD PRACTICE

1. Inconsistent Counter:

   - wg.Add(2) increments the wait counter to 2, indicating that you expect two goroutines to signal completion.
   - However, wg.Add(-1) immediately decrements it back to 1, essentially "canceling" one of the goroutines before it even starts.
   - This creates a mismatch between the expected and actual number of goroutines, which can lead to unexpected behavior or race conditions.

2. Potential Errors:

   - If you forget to call wg.Add(-1), the wait counter will remain at 2, causing your main program to wait indefinitely for a nonexistent second goroutine. This can be difficult to debug and might lead to program hangs.
   - If you call wg.Add(-1) more than once, the counter will become negative, triggering the panic: sync: negative WaitGroup counter error.

3. Unnecessary Complexity:

   - Using wg.Add(2) followed by wg.Add(-1) is less readable and maintainable than the simpler approach of using wg.Add(1) per goroutine.
   - This can make your code harder to understand and debug, especially for others who might not be familiar with this specific pattern.

4. Lost Tracking:
   - If you have multiple goroutines executing concurrently, using wg.Add(2) followed by wg.Add(-1) makes it harder to track which goroutine triggered which decrement.
   - This can be problematic for debugging or understanding the flow of your program.

BEST PRACTICE 
- To ensure clarity, accuracy and avoid potential errors, its always recommended to use
wg.Add(x) exactly once for each go routine you spawn. 
- This approach aligns with the intended purpose of Waitgroup and makes your code easier to understand
