object Main extends App {
  def makeAdder(n: Int): Int => Int = (x: Int) => x + n
  def makeMultiplier(n: Int): Int => Int = (x: Int) => x * n
  def compose(f: Int => Int, g: Int => Int): Int => Int = (x: Int) => f(g(x))
 
  override def main(args: Array[String]): Unit = {
    println(compose(makeMultiplier(3), makeAdder(1))(10))
  }
}
