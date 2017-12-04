import scala.util.{Success, Try}

object SimpleIntereset extends App {
  def readNaturalNumber(message: String): Int = {
    print(message)
    Try(io.StdIn.readInt) match {
      case Success(n) if n > 0 => n
      case _ => readNaturalNumber(message)
    }
  }

  def readPercentage(message: String): Double = {
    print(message)
    Try(io.StdIn.readDouble) match {
      case Success(p) if p > 0 => p
      case _ => readPercentage(message)
    }
  }

  val principal = readNaturalNumber("Enter the principal: ")
  val rate = readPercentage("Enter the rate of interest (%): ")
  val years = readNaturalNumber("Enter the number of years: ")
  val investment = principal * (1 + years * rate / 100.0)

  printf(s"After $years years at $rate%%, the investment will be worth $$%.2f\n", investment)
}