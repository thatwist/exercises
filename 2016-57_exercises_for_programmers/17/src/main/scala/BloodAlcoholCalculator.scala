import scala.util.{Success, Try}

object BloodAlcoholCalculator extends App {

  object Genders extends Enumeration {
    val male = Value("M")
    val female = Value("F")
  }

  def readNaturalNumber(message: String): Int = {
    print(message)
    Try(io.StdIn.readInt()) match {
      case Success(n) if n > 0 => n
      case _ => readNaturalNumber(message)
    }
  }

  def readGender(message: String): Genders.Value = {
    print(message)
    Try(Genders.withName(io.StdIn.readLine.stripLineEnd)) match {
      case Success(gender) => gender
      case _ => readGender(message)
    }
  }

  def readDouble(message: String): Double = {
    print(message)
    Try(io.StdIn.readDouble()) match {
      case Success(n) if n > 0 => n
      case _ => readDouble(message)
    }
  }

  val w = readNaturalNumber("Weight (pounds): ")
  val h = readNaturalNumber("Hours since last drink: ")
  val a = readDouble("Amount of alcohol consumed (oz): ")
  val r = readGender(s"Gender (${Genders.male}, ${Genders.female}): ") match {
    case Genders.male => 0.73
    case Genders.female => 0.66
  }

  val bac = (a * 5.14 / w * r)

  println(s"Your BAC is $bac")
  if (bac >= 0.08) println("It is not legal for you to drive.")
}