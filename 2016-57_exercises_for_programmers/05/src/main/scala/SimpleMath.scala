import scala.util.{Failure, Success, Try}

object SimpleMath extends App {
  def readInput(message: String): Int = {
    print(message + " ")
    Try(io.StdIn.readInt()) match {
      case Success(x) => x
      case Failure(_) => readInput(message)
    }
  }

  val d1: Int = readInput("What is the first number?")
  val d2: Int = readInput("What is the second number?")
  printf("%d + %d = %d\n", d1, d2, d1 + d2)
  printf("%d - %d = %d\n", d1, d2, d1 - d2)
  printf("%d * %d = %d\n", d1, d2, d1 * d2)
  printf("%d / %d = %s\n", d1, d2,
    Try(d1 / d2) match {
      case Success(x) => x.toString
      case Failure(_) => "undefined"
    }
  )

}