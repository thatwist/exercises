import org.mindrot.jbcrypt.BCrypt

object PasswordValidation extends App {
  val Passwords = Map(
    "dejan" -> "$2a$12$8Z/ZqAjvQVdrF311MrAVkOVoc5F//2hL.wY7WqzyAJ1sFuYe.UZaa",
    "guest" -> "$2a$12$eJu8mebC.fycd/LVewXkvuDDHIPoPyNvPFRk4667HRfK/H5HXy1AW"
  )

  def readString(message: String): String = {
    print(message)
    io.StdIn.readLine.stripLineEnd match {
      case "" => readString(message)
      case username => username
    }
  }

  def readPassword(message: String): String = {
    new String(System.console.readPassword(message))
  }

  def authenticate(login: String, password: String): Boolean =
    Passwords.get(login) match {
      case Some(x) => BCrypt.checkpw(password, x)
      case _ => false
    }

  val login = readString("login: ")
  val password = readPassword("password: ")
  if (authenticate(login, password)) {
    println("Welcome!")
  } else {
    println("Wrong username/password")
  }

}