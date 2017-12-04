// Counting the number of characters
object Counting extends App {
  def readLine: String = io.StdIn.readLine.stripLineEnd
  def readInput(message: String): String = {
    print(message + " ")
    readLine match {
      case "" => readInput("You must enter something:")
      case input => input
    }
  }
  val input = readInput("What is the input string?")
  printf("%s has %d characters.\n", input, input.length)
}
