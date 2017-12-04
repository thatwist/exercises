// Saying Hello
object SayingHello extends App {
  def readName = io.StdIn.readLine().stripLineEnd
  def greet(greeting: String)(target: String): String =
    s"$greeting $target"
  def hello(name: String) = greet("Hello")(name)
  def greeting(name: String): String = name match {
    case "Dejan" => greet("Welcome back")(name)
    case "" => hello("mr. Nobody")
    case _ => hello(name)
  }

  println("What is your name?")
  println(greeting(readName))
}
