## Select Statement

1. In Go language, the select statement is just like switch statement, but in the select statement, **case statement refers to communication, i.e. sent or receive operation on the channel.**

2. The select statement in Go is used for communication between goroutines.

3. It **allows you to wait on multiple channel operations simultaneously.**

4. Each case in a select statement specifies a channel operation (send or receive) along with the corresponding code block to execute if that channel operation can proceed without blocking.

5. **If multiple cases are ready to proceed, one of them is chosen at random.**

6. No cases are ready
    - **If none of the cases are ready, and there's a default case, the default case will be executed.**
    - **If there's no default case, the select statement blocks until at least one of the cases can proceed.**