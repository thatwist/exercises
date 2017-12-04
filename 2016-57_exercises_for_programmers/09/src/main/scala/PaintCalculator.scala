import scala.util.{Success, Try}

object PaintCalculator extends App {
  def readNaturalNumber(message: String): Int = {
    print(message)
    Try(io.StdIn.readInt) match {
      case Success(number) if number > 0 => number
      case _ => readNaturalNumber(message)
    }
  }

  val length: Int = readNaturalNumber("Enter room length: ")
  val width: Int = readNaturalNumber("Enter room width: ")

  val ceiling: Int = length * width

  val GallonsPerSquareFeet = 350
  val volume: Int = (ceiling / GallonsPerSquareFeet, ceiling % GallonsPerSquareFeet) match {
    case (v, 0) => v
    case (v, _) => v + 1
  }

  println(s"1 gallon covers $GallonsPerSquareFeet square feet. You will need to purchase $volume gallons of paint to cover the ceiling.")
}