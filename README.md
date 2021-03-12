# Exam Question

I did not think this would take so long, and I ran out of time.
so I will have to come back to this latter today.




## 1. What are some advantages of Go over other languages?
Single Executable that can run in a Docker scratch container. 
This means there is no RUN Time or JVM, or OS / DLL style dependencies that could get a Breaking “UPGRADE”.   For example, running PHP requires a LAMP stack.  Same curse for Python... Ruby…
A single Executable, running in a Scratch container, running a golang mux for example, also means that you are controlling more vectors of attack.  Exploits are dramatically reduced.  Couple that with running in a VPC/Security Group ( behind Load Balancers / Firewalls ) to connect with other services, such as messaging or Databases, Exploits are limited to inside jobs.

The code can read like a book.  The most complicated style of code follow along would be interfaces, but assuming sound naming conventions, even this can be Pro.  It might be that I have just become accustomed to Procedural style coding, but I find Event style code, such as node, to be confusing.

While golang is a compiled language in the same category as C++, Thankfully, it still lacks C++ “Features” that I do not want anyway. Polymorphism, is an example of a Feature that might have been appropriate back in 32 Bit system times, but I have personally never seen a burning need.  
 
MultiPlexers. Being able to “cook in” a Multiplexer listenAndServe into your code is probably the only feature that something like ruby on rails actually has brought to the game. But GoLang has a variety of MUX, including the new FIBER mux which includes an MVC and an elegant middleware section that can handle Authentication and issues like CORS..   I just started a personal project using goLang/FIber…   https://github.com/vinceyoumans/PTT-server

I was Speaking with Bill Kennedy of Arden Labs about an initiative he is working on, looking for objections.  It is to incorporate some sort of runtime feature to anonymous features that is common to OOP. Generics…  His argument that some OOP academics will never take GoLang seriously with out this feature.  I responded that I fall into the camp that I like keeping everything structured and organized at Compile time and using the Reflection package if I must.  

I used to think that GoRoutines was over used, but I have since started to appreciate it more.  Node tolerates waiting for remote calls with .then.   There are also Promises.  But go addresses bocking or promises by using a goRoutine and sending status messages and responses back as channels.  Without going into detail, this makes for a very elegant combination.
STRUCTS!...   STRUCTS!   STRUCTS!....   As a mobile developer that looks at everything as a noSQL/JSON problem, Structs has been exactly what I need.  
It runs awesome in serverless like AWS Lambdas as well as Docker containers.  For me, as a small team, this gives me the opportunity to build projects far more sophisticated that I could have on my own because now I can use cloud services.  FireBase for example had been my goto DB…  for Mobile, node and GoLang.  

I could ramble on, but right now, I look at any existing project as a working prototype to be rebuilt as a GoLang project.





## 2. What are some disadvantages?
I have been frustrated with the MUX TEMPLATE engine.  I would really like to incorporate even the HTML/JS parts of a project as a template in code, but I am struggling.  I suspect that this really just my lack of Front end skills.  But GoLang is perfect for PWA and reactive web sites, which is better for a team anyway because you can compartmentalize development roles.  The golang devs can focus on mechanical part of system and the API, while completely ignore the Minutia of dealing with the cosmetics and user experience. In some respects, one of the most important skills that a WEb developer can have is patients, dealing with Customers who cannot make up their mind about what color to make the screen.  So perhaps this is actually a blessing.

Compare and contrast.



## 3. Explain a Go pointer and provide some examples of copy by value vs. copy by a pointer.
A go Pointer is an address to a memory location in the program context.  I say Program context, because I am not sure if you could point to a memory location outside of the program. 


TO make a COPY by Value, a new memory Var is being allocated.  Changes to either var will not be reflected in the other.  Example.
A := 10
B := A  // B now equals 10
fmt.Println ( A, “---”, B ) //  “10 -- 10”
There are now two vars that both = 10.
A = 20
fmt.Println ( A, “---”, B ) //  “20 -- 10”


To Copy by Pointer ( reference ), 
A := 10
p := &A  // points to A

fmt.Println( p )  // will show a mem address that looks something like this..  0xc000018050
fmt.Println( *p ) // reads p which is pointer to A  ( 10 )

*p = 200
fmt.Println(A)  //  A was 10,  Now it is 200.




## 4. Is copying by pointer more performant than by value?  
There are a couple of benefits.
Save on memory stack. Big deal if you have large structs
If you pass by value, you double Garbage collection.  Big deal if speed and predictable execution is required.
If you pass by value, then you might have to synch the values together.  Added logic.
Passing by value does required the CPU to MAKE new Vars, which is time/CPU consuming.  To just pass the reference is probably less cpu intensive.

but...
Passing by value also introduces possible Race conditions for example GoRoutines.






## 5. Does Go have an ability to pass by reference? Why or why not?
Yes,   to pass by reference, such as a parameter in a func, you are passing the memory location, aka the pointer.  The pointer effectively becomes the var.  




## 6. Have you worked, or are you familiar, with Go 2.0?
What is your favorite new feature of Go 2.0?
No…  I am at go 1.16. But I did some reading about it for this paper.  My thoughts...
I don’t care about Generics.  Just does not look that compelling a feature.  Might even introduce confusion to me.  But I may not understand it.  Also seems to introduce a contract concept that I would approve.

I take go modules for granted. One of the first things I do is a Go mod Init.
And error handling seems fine as is, but 2.0 seems to improve on reporting. Adding check expression.  Introduces a standard interface, which could be handy.  But again, I would have to study it in more detail.








## 7. Is Go an object-oriented language? Why or why not?
No.  it is not OO.  Which I am fine with.  No Polymorphism, which I do not like anyway.  No inheritance.  But go Has structs and Interfaces, so far more organized at compile, which is how I prefer to think. 
Interfaces do give you Polymorphic behavior.  The classic code example is AREA for a Square vs AREA for a Circle.  AREA would be the interface.





## 8. How can you distribute tasks in Go to different machines?
There are a couple of ways.
- Sockets
- ProtocolBuffers/ gRPC.  Which is Port 80 and takes advantage of long polling http2.
- Messaging such as AWS SQS, SNS, Kinesis.   Or Google PubSub
- AWS Lambdas or Google Functions can make calls to other services.
- Channels, especially when paired with Messaging ques like NATS, Redis...





## x09 - Demonstrate, with code, your favorite aspect of Go. Share an idiom, package, or algorithm that you like.

I went with Channels.  Its a very interesting way to run logic in background without blocking.



## x10 - Abstracting a DB layer with interfaces.

Not sure this is what you asked for 


