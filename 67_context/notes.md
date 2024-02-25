1. When you have a process that launches a bunch of goroutines, 
    so when you decide to stop/cancel the process, all the corresponding goroutines,
        should also cancel

2. This is done so that there are no leaking goroutines,
    leaking goroutines use up resources

--- 

### Go Concurrency Patterns: Context

1. In Go servers, each incoming request is handled in its own goroutine. 

2. Request handlers often start additional goroutines to access backends such as databases and RPC services (goroutines of a process 👆👆)

3. The set of goroutines working on a request typically needs access to request-specific values such as identity of the end user, authorization tokens and the requests deadline

4. When a request is cancelled or is timed out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using

5. At Google, we developed a context package that makes it easy to pass request scoped values, cancellation signals and deadlines across API boundaries to all the goroutines involved in handling a request. This package is publicly available as context.

6. request scoped data : 
    a. Information that can only exist once the request has begun.
    b. Good examples include : user IDs extracted from headers, authentication tokens tied to cookies or session IDs, distributed tracing IDs and so on


