## Channels (Unbuffered Channels)


1. Channels in Go are communication primitives that **facilitate the transfer of data** *between goroutines*, 

    -> enabling synchronization and coordination in concurrent programs. 

2. They **allow goroutines to send and receive data safely**, preventing race conditions through blocking operations and providing a clear and efficient means of **inter-goroutine communication.**


## Buffered Channels

1. Simple channels or unbuffered channels do not allow the sending of a value to a channel if there isn't a goroutine available to receive at the same time.

2. **Buffered channels allow this behaviour, you can send values, upto a certain buffer limit, even if there is no goroutine availabe to receive it**

- Buffered channels generally employ FIFO within the buffer


## Question Asked
Question Asked : 
https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922344#questions/21390794