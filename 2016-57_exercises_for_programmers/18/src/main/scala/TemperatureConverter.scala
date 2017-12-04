import scala.util.{Try, Success}

object TemperatureConverter extends App {
  object Units extends Enumeration {
    val Celsius = Value("Celsius")
    val Fahrenheit = Value("Fahrenheit")
  }

  def convertUnit(targetUnit: Units.Value)(temp: Double): Double =
    targetUnit match {
      case Units.Celsius => (temp - 32) * 5 / 9
      case Units.Fahrenheit => (temp * 9/5.0) + 32
    }

  def readTemp(message: String): Double = {
    print(message)
    Try(io.StdIn.readDouble) match {
      case Success(x) if x > 0 => x
      case _ => readTemp(message)
    }
  }

  def readConversionUnits(): (Units.Value, Units.Value) = {
    println("Press C to convert from Fahrenheit to Celsius.")
    println("Press F to convert from Celsius to Fahrenheit.")
    print("Your choice: ")

    Try(io.StdIn.readChar()) match {
      case Success('C') => (Units.Fahrenheit, Units.Celsius)
      case Success('F') => (Units.Celsius, Units.Fahrenheit)
      case _ => readConversionUnits()
    }
  }

  val (sourceUnit, targetUnit) = readConversionUnits()
  val temp = readTemp(s"Enter the temperature in $sourceUnit: ")
  val convert = convertUnit(targetUnit)

  printf(s"\nThe temperature in $targetUnit is %.2f.\n", convert(temp))
}