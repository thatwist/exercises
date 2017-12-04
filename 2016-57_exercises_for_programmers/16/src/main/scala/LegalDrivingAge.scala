import scala.util.{Success, Try}

object LegalDrivingAge extends App {
  val DrivingAge = 16

  def readNaturalNumber(message: String): Int = {
    print(message)
    Try(io.StdIn.readInt) match {
      case Success(n) => n
      case _ => readNaturalNumber(message)
    }
  }
  def eligibleToDrive(age: Int): Boolean = age >= DrivingAge

  val age = readNaturalNumber("What is your age? ")
  if (eligibleToDrive(age)) {
    println("You are old enough to legally drive.")
  } else {
    println("You are not old enough to legally drive.")
  }
}