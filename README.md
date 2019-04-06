# Programming samples based on SICP reading and exercises

This is a repo where I put sample programs based on the SICP text, and / or solutions to the exercises. The source structure is:

    src/chN/meetup/<lang>
    src/chN/exercises/<lang>

Thus, to find the code for exercise solutions to chapter 8 written in rust, you have to go to `src/ch08/exercises/rust`. Use the `meetup` subdirectory instead for code samples written based on the text read in the meetup.

While SICP teaches a Lisp dialect, I don't limit code examples to Lisp. In fact, I especially use non-Lisp languages and see how far we can extract a reasonable "functional" experience from them. There is value in applying functional thought everywhere, and there is definite value in using a functional solution in whatever programming language you're using. It can make reading code incredibly easier in many cases (parser construction comes to mind).
