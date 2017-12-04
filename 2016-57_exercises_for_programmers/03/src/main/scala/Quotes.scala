object Quotes extends App {
  case class Quote(author: String, text: String)

  val quotes = List(
    Quote("Obi-Wan Kenobi", "These aren't the droids you're looking for."),
    Quote("Yoda", "Do or do not. There is no try.")
  )

  for {
    quote <- quotes
  } yield println(quote.author + " says: \"" + quote.text + "\"")
}