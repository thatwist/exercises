import scala.util.{Success, Try}

object TaxCalculator extends App {
  def readAmount(message: String): Double = {
    print(message)
    Try(io.StdIn.readDouble) match {
      case Success(n) if n > 0 => n
      case _ => readAmount(message)
    }
  }

  def readState(message: String): String = {
    print(message)
    io.StdIn.readLine.stripLineEnd match {
      case "" => readState(message)
      case state => state
    }
  }

  def stateTax(state: String): Double = state match {
    case "WI" => 0.055
    case _ => 0.0
  }

  val amount = readAmount("What is the order amount? ")
  val taxedAmount = amount * stateTax(readState("What is the state? "))
  var total = amount
  if (taxedAmount > 0) {
    total += taxedAmount
    printf("The subtotal is $%.2f.\n", amount)
    printf("The tax is $%.2f.\n", taxedAmount)
  }
  printf("The total is $%.2f.\n", total)

}