import scala.util.{Success, Try}

object BMICalculator extends App {
  def readNaturalNumber(message: String): Double = {
    print(message)
    Try(io.StdIn.readDouble()) match {
      case Success(n) if n > 0 => n
      case _ => readNaturalNumber(message)
    }
  }

  val height = readNaturalNumber("Enter height (cm): ") * 0.01
  val weight = readNaturalNumber("Enter weight (kg): ")
  val bmi = (weight / (height * height))

  printf("Your BMI is %.2f.\n", bmi)
  bmi match {
    case _ if bmi < 18.5 => println("You are underweight. You should see your doctor.")
    case _ if bmi > 25.0 => println("You are overweight. You should see your doctor.")
    case _ => println("You are within the ideal weight range.")
  }
}