import scala.util.{Success, Try}

object CompoundInterest extends App {
  def readNaturalNumber(message: String): Int = {
    print(message)
    Try(io.StdIn.readInt()) match {
      case Success(n) if n > 0=> n
      case _ => readNaturalNumber(message)
    }
  }

  def readDouble(message: String): Double = {
    print(message)
    Try(io.StdIn.readDouble()) match {
      case Success(x) => x
      case _ => readDouble(message)
    }
  }

  val amount = readDouble("What is the principal amount? ")
  val rate = readDouble("What is the rate (%)? ")
  val years = readNaturalNumber("What is the number of years? ")
  val compoundedPerYear =
    readNaturalNumber("What is the number of times the interest is compounded per year? ")
  val investment = amount * Math.pow(1 + rate / (100 * compoundedPerYear), compoundedPerYear * years)

  printf(s"$$%.2f invested at %.2f%% for $years years compounded $compoundedPerYear times per year is $$%.2f.\n", amount, rate, investment)
}