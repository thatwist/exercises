import scala.util.{Success, Try}

object CurrencyConversion extends App {

  object Currencies extends Enumeration {
    val rates = Map(
      Value("EUR") -> 1.00,
      Value("USD") -> 0.89,
      Value("RSD") -> 0.0081,
      Value("JPY") -> 0.0079,
      Value("CNY") -> 0.14,
      Value("CHF") -> 0.91,
      Value("RUB") -> 0.013
    )
  }

  def readAmount(message: String): Double = {
    print(message)
    Try(io.StdIn.readDouble()) match {
      case Success(n) if n > 0 => n
      case _ => readAmount(message)
    }
  }

  def readCurrency(message: String): Currencies.Value = {
    print(message)
    Try(Currencies.withName(io.StdIn.readLine.stripLineEnd)) match {
      case Success(currency) => currency
      case _ => readCurrency(message)
    }
  }

  val currency = readCurrency("What currency do you want to exchange? ")
  val amount = readAmount(s"How many $currency are you exchanging? ")
  val exchangeRate = Currencies.rates(currency)
  val exchangedAmount = exchangeRate * amount

  printf("%.2f %s at an exchange rate of %.4f is %.2f EUR\n", amount, currency, exchangeRate, exchangedAmount)
}