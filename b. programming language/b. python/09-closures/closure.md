# Closure Overview

```
def outer_function(outer_variable):
    def inner_function():
        print(outer_variable)

    return inner_function

# Create a closure
closure = outer_function('Hello, World!')

# Call the inner function
closure()  # Output: Hello, World!

```

- Why do we return the function instead of calling it?

  - if you call the inner_function directly within the outer_function instead of returning it, it will not retain access to the outer_variable outside of the outer_function. Let's illustrate this with an example:

    ### Direct Call Example:

    ```
    def outer_function(outer_variable):
      def inner_function():
        print(outer_variable)

      inner_function()
    # Call the inner function directly
    ```

    ### Explanation:

    In this example, inner_function() is called immediately inside outer_function. Here’s what happens:

    - Calling inner_function:

      - When outer_function is executed, inner_function is called and prints outer_variable.
      - After inner_function completes execution, control returns to outer_function.

    - Scope of outer_variable:

      - The outer_variable is accessible to inner_function because it is still within the scope of outer_function when inner_function is called.

    - No Closure:

      - After outer_function finishes executing, inner_function is no longer accessible from outside outer_function.
      - The outer_variable does not exist outside outer_function, so there’s no way to access inner_function or its environment once outer_function completes.

    ### Closure Example:

    To create a closure, you need to return the inner_function:

    ```
    def outer_function(outer_variable):
      def inner_function():
        print(outer_variable)

      return inner_function
      # Return the inner function as a closure

    closure = outer_function('Hello, World!')
    closure() # Output: Hello, World!
    ```

    ### Explanation:

    - Returning inner_function:

      - When outer_function is executed, it returns inner_function, which holds a reference to the outer_variable.
      - This allows inner_function to access outer_variable even after outer_function has completed.

    - Calling closure():

      - closure() is effectively calling inner_function, which has retained access to outer_variable due to the closure.

    ### Summary:

    - Direct Call: If you call inner_function directly within outer_function, it will only have access to outer_variable while outer_function is running. Once outer_function completes, the environment is discarded, and you lose access to inner_function.

    - Returning the Function: By returning inner_function, you create a closure that retains access to the outer_variable even after outer_function has completed. This allows you to call the returned function later with access to the captured environment.

    - Closures are useful for creating functions that need to remember or encapsulate their environment, which is not possible if the inner function is only called and not returned.
