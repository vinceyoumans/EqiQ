# Interface to a SQL DB


In this example, I use an interface to abstract to a traditional SQL DB.
The Idea is to create something like a remote function that can be used as a service for other services.
Internally, there is a "contract" that has a struct that must be passed to the service.

The service responds with result and error.

This solution could act as a gate keeper, by introducing testing logic, for users that wanted to post to the DB.  testing logic would include, does the user have permission to submit to the DB.

I am thinking an immutable ledger where calues cannnot be changed.



