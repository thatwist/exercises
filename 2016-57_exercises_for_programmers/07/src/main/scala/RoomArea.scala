import scala.util.{Failure, Success, Try}

object RoomArea extends App {

  object UnitOfLength extends Enumeration {
    private val squareConversionFactors = Map(
      Value("feet") -> 1.0,
      Value("meters") -> 0.09290304)
    def convertAreaToFeet(value: Double, unit: Value): Double = value / squareConversionFactors(unit)
    def convertArea(valueInFeet: Double, unit: Value): Double = valueInFeet * squareConversionFactors(unit)
  }

  def readUnit(message: String): UnitOfLength.Value = {
    print(message + " ")
    val input = io.StdIn.readLine.stripLineEnd
    Try(UnitOfLength.withName(input)) match {
      case Success(unit) => unit
      case Failure(_) => readUnit(message)
    }
  }

  def readLength(message: String): Double = {
    print(message + " ")
    Try(io.StdIn.readDouble) match {
      case Success(feet) => feet
      case Failure(_) => readLength(message)
    }
  }

  val unitOptions = UnitOfLength.values.mkString(", ")
  val (unit, length, width) = (
    readUnit(s"What is the unit of length [$unitOptions]?"),
    readLength("What is the length of the room?"),
    readLength("What is the width of the room?"))

  printf(s"You entered dimensions of %.3f $unit by %.3f $unit.\n", length, width)
  println("The area is")
  val areaInFeet = UnitOfLength.convertAreaToFeet(length * width, unit)
  for (unit <- UnitOfLength.values) printf(s"%.3f square $unit\n", UnitOfLength.convertArea(areaInFeet, unit))
}