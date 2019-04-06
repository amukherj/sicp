puts -> {
  compose = ->(f, g) {
    ->(x) {
      f.(g.(x))
    }
  }

  make_adder = ->(n) {
    ->(x) {
      x + n
    }
  }

  make_multiplier = ->(n) {
    ->(x) {
      x * n
    }
  }

  mul3add1 = compose.(make_multiplier.(3), make_adder.(1))
  mul3add1.(10)
}.()

