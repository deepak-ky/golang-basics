1. Each goroutine within the loop will attempt to lock the mutex (mu) before accessing and modifying the shared variable (counter).

2. When the goroutine at i = 1 attempts to lock the mutex (mu.Lock())
   a. and finds it already locked by the goroutine at i = 0,
   b. it will indeed wait until the mutex is unlocked by the goroutine at i = 0.
   c. This is the purpose of using a mutex - to ensure that only one goroutine can access the critical section (in this case, accessing/modifying counter) at any given time.

3. The mutex ensures exclusive access to the critical section (modifying counter). Only one goroutine can hold the lock at a time.

4. RWMutex and Mutex

    a. When you call RLock() on an RWMutex, it means you're acquiring a **read lock**. 

    b. Multiple readers can hold this **read lock simultaneously, and while they hold the read lock, other readers can also acquire it.**

    c. What's important is that **no writer can acquire the write lock (Lock()) while there are active read locks.** 

    d. This ensures that readers have shared access and can read the data concurrently without interfering with each other.
