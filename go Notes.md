
# Go Notes

1. When defining function names, use a capital letter if you want it to be exportable when defining the function in a module you want to import externally.

2. Using `' '` and `" "` are different: the former is for character literals (so it cannot take more than one character/letter). When used, `fmt.Println()` prints the Unicode integer assigned to whatever letter is in the `' '`.

3. `:=` defines a variable while letting Go infer the type of the variable that is being defined.