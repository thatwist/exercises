import scala.util.{Failure, Success, Try}

object SelfCheckout extends App {

  def readPrice(message: String): Double = {
    print(message)
    Try(io.StdIn.readDouble()) match {
      case Success(n) if n > 0 => n
      case _ => readPrice(message)
    }
  }

  def readQuantity(message: String): Int = {
    print(message)
    Try(io.StdIn.readInt) match {
      case Success(n) if n >= 0 => n
      case _ => readQuantity(message)
    }
  }

  case class Item(quantity: Int, price: Double) {
    def total = quantity * price
  }

  def readItems(accumulator: List[Item] = List()): List[Item] = {
    val index = accumulator.length + 1
    readQuantity(s"Enter the quantity of item $index (0 for checkout):") match {
      case 0 => accumulator
      case q => {
        val item = Item(q, readPrice(s"Enter the price of item $index: "))
        readItems(item :: accumulator)
      }
    }
  }

  val TaxRate = 0.055
  val items: List[Item] = readItems()
  val subtotal: Double = items.map(_.total).sum
  val tax: Double = subtotal * TaxRate
  val total: Double = subtotal + tax

  println("==================")
  printf("Subtotal: $%.2f\n", subtotal)
  printf("Tax: $%.2f\n", tax)
  printf("Total: $%.2f\n", total)
}