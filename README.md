# Mocking

# Background

Mocking is an essential part of unit testing, and it is widely used there. A *code unit* under test may have dependencies on other units. In order to isolate and refine the behavior, we want to replace the dependencies with mocks that will simulate possible acts.

In other words, mocking is replacing dependencies with easily controlled code units that simulate the real behavior.

Mocking can be useful when:

1. The dependency is not cooperative (e.g., third-party).
2. It is hard to trigger all side-effect and return values with normal inputs (e.g., triggering a specific error).
3. The dependency is slow (e.g., cryptographic functions)
4. It is hard to set up the dependency
5. There is no implementation of the dependent unit

In addition, mocking allows you to focus solely on testing the unit behavior under specific flows.

# Mocking in Go

## Go Restrictions

1. In non-compiled languages (e.g., Python), mocking is done by replacing the unit in run-time, unfortunately, in a compiled language, once the code is compiled, it cannot be changed easily. 
2. In OOP languages, one can easily mock an object by inheriting it and overriding its functions, unfortunately, Go does not support inheritance.

This forces us to write *testable* code that allows mocking.

## Mocking Techniques in Go

There are several approaches of mocking in Go, however, there isn’t a one-fits-all technique. Some of there a matter of taste and some of them are useless in some cases. Therefore, you should be familiar with all of these techniques and choose the most appropriate one for the situation.

We will walk through all the following techniques using this [examples repository](https://github.com/DanielSolomon/go-mock-techniques).

- Using interfaces (Polymorphism)
    
    > *“Accept interfaces, return structs.”*
    > 
    
    Why accept an interface?
    When providing a function, the supplied inputs are out of the control of the writer, and the caller might find it useful to wrap your structs with additional data and logic. As a function writer, you want to provide an abstract behavior of how this input will be used.
    
    Why return a struct?
    Covering the real struct behind an interface hides information from the caller, but since the caller gets the return value, in order to proceed it will have to cast the struct to a specific type, to enable its full functionality.
    
    - Stubbing
        
        Create a struct that implements the full API of the input’s interface. Pass the stubbed implementation to the function instead of the real production struct.
        
        Advantages:
        
        1. It can maintain a state.
        
        Disadvantages:
        
        1. All functions must be stubbed.
        2. If the interface changes, your stub doesn’t implement it anymore and needs to be fixed.
        3. The stubbed code might grow a lot, remember that this is not the business logic of the system!
        4. To simulate special edge cases, you need to use *magic* inputs or add functionality to the stub that signals that the next call should fail.
    - Embedding
        
        Create a struct that embeds the interface, implement stubs only for the functions you will use.
        
        It revokes the first two disadvantages of the **stubbing** technique.
        
    - Mocking
        - Mock package
            
            The `testify/mock` package provides a mocking struct that can be embedded and used to configure the mock output on each call. Its main benefits are:
            
            1. Trigger any result easily.
            2. Reconfigure each test.
            
            But maintaining a state (like in the **stubbing** technique) is harder.
            
        - Generate
            
            There are some tools that can auto-generate mocks given an interface.
            
            - Using [mockgen](https://github.com/golang/mock)
                
                *mockgen* generates a full mocked struct given an interface, then it exposes a nice API to assert calls and control return values according to the inputs. It also allows building a full mocked flow (asserting X calls happened in order and each call will return a specific response).
                
                *mockgen* comes in handy when we want to fully mock an external interface.
                
            - Using [mockcompose](https://github.com/kelveny/mockcompose)
                
                *mockcompose* allows you to generate a mocked struct that preserves some of the original implementations and mocks the other. This ability can give us fine-grained control in unit testing.
                
- Dependency injection
    
    Instead of calling dependencies directly, receive them as function arguments and use them instead. This allows the caller to control which function will be invoked. This technique is widely used in other programming languages (e.g., Python), although, it is more cumbersome in Go since there are no default arguments, so it means the caller must know which function should be used in production. Another drawback is that it lengthens the function’s signature. 
    
- Monkey patching
    
    Monkey patching is a technique to add, modify, or suppress the default behavior of a piece of code at runtime without changing its original source code**.** It is done by replacing a unit under the package nose. To allow it, define all dependencies as global variables, invoke them indirectly, through the global variables. This allows the caller to *monkey patch* the variables in run-time, to point to the mocked units. 
    
    This technique is very similar to *dependency injection*, in its favor, the caller doesn’t need to know the real dependent value and the function’s signature remains the same. On the other hand, it introduces two side effects:
    
    1. Parallel testing is impossible (without synchronizing the global variables).
    2. If testes outside the package, the variables must be publicly available.
    
    There is also this very cool [monkey](https://github.com/bouk/monkey) package, unfortunately, it was archived and it doesn’t support non-intel chips.
    
- HTTP
    
    `net/http` provides a testing package `net/http/httptest`, it can be used to tests both `http servers` and `http clients`.