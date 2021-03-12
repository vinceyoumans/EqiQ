# 9. Demonstrate, with code, your favorite aspect of Go. Share an idiom, package, or algorithm that you like.


So pretend there is a long running process with multiple steps and you do not want to block process...   

here we are passing a 
- context
- channels including
    - names, which is receiving a slice of bytes
    - RemoteAA01 which is a channel that is receiving slice of bytes
    - err which is receiving error

The function is waiting for context to return a DONE.  I could have added timeout.
errors and names would print out as they are sent to the channel.

```
func PTT_RemoteProcessAA01(ctx context.Context, 
                    names <-chan []byte,
                    RemoteAA01 chan<- []byte, 
                    errs <-chan error) {

   for {
       select {
       case <-ctx.Done():
           log.Println("finished processing AA01")
           return
       case err := <-errs:
           log.Println("AA01 has a problem:", err)
       case name := <-names:
           greeting := "Test Hello " + string(name)
           greetings <- []byte(greeting)
   }
 }}
```


