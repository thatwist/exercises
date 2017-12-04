import java.time.Year

import scala.util.{Failure, Success, Try}

object RetirementCalculator extends App {

  def readAge(message: String): Int = {
    print(message + " ")
    Try(io.StdIn.readInt) match {
      case Success(age) => age
      case Failure(_) => readAge(message)
    }
  }

  def printRetirementInfo(yearsToRetire: Int): Unit = {
    printf ("You have %d years lef until you can retire.\n", yearsToRetire)
    val currentYear = Year.now.getValue
    val retirementYear = currentYear + yearsToRetire
    printf("It's %d, so you can retire in %d.\n", currentYear, retirementYear)
  }

  val currentAge = readAge("What is your current age?")
  val retirementAge = readAge("At what age would you like to retire?")

  (retirementAge - currentAge) match {
    case yearsToRetire if yearsToRetire > 0 => printRetirementInfo(yearsToRetire)
    case _ => println("You can already retire old man!")
  }

}