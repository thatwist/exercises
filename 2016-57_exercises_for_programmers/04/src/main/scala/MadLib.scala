object MadLib extends App {
  def readInput(message: String): String = {
    print(s"$message: ")
    io.StdIn.readLine().stripLineEnd match {
      case "" => readInput(message)
      case entry => entry
    }
  }

  case class Story(noun: String, verb: String, adjective: String, adverb: String) {
    val text= s"Do you $verb your $adjective $noun $adverb?"
    val punchline = "That's hilarious!"
    override def toString() = s"$text $punchline"
  }

  val story = Story(
    noun = readInput("Enter a noun"),
    verb = readInput("Enter a verb"),
    adjective = readInput("Enter an adjective"),
    adverb = readInput("Enter an adverb"))

  println(story)
}