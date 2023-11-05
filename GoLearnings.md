## Pointer

`&` (Ampersand):
* The `&` symbol is used to create a pointer to a variable or a value. It's the **"address-of"** operator, and it returns the memory address of the operand.
* In your code, `&AppConfig` is used to create a pointer to an **AppConfig struct**. It effectively allocates memory for an **AppConfig struct** and returns a pointer to that memory location. This is often used when you want to create a new instance of a struct and obtain a pointer to that instance.

`*` (Asterisk):

* The `*` symbol is used to declare and work with pointers to values. It's the **"dereference"** operator when used in expressions. It allows you to access the value that the pointer is pointing to.
* In your code, `*AppConfig` is used in the function signature **func LoadConfig() `*`AppConfig**. This specifies that the LoadConfig function returns a pointer to an **AppConfig object**. When you call this function, you'll get a pointer to an **AppConfig object**, and you can access the fields of the **AppConfig** by dereferencing the pointer (e.g., `config.Port` if config is a pointer to AppConfig).

## Functions

* To export a function you have to write it in uppercase, like `GetHelloHandler()` otherwise it will not be exported