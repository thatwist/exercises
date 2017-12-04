import scala.util.{Try, Success, Failure}

object PizzaParty extends App {
  def readNaturalNumber(message: String): Int = {
    print(message + " ")
    Try(io.StdIn.readInt) match {
      case Success(number) if number > 0 => number
      case _ => readNaturalNumber(message)
    }
  }

  val people: Int = readNaturalNumber("How many people?")
  val slices: Int = readNaturalNumber("How many slices of pizza do you have?")
  println(s"$people people with $slices slices of pizza")

  if (slices >= people) {
    printf("Each person gets %d slices of pizza.\n", slices / people)
    printf("There are %d leftover pieces.\n", slices % people)
  } else {
    printf("%d will eat. %d people will not eat.\n", slices, people - slices)
  }
}
