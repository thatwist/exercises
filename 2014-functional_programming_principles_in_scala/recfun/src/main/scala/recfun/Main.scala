package recfun
import common._

object Main {
  def main(args: Array[String]) {
    println("Pascal's Triangle")
    for (row <- 0 to 10) {
      for (col <- 0 to row)
        print(pascal(col, row) + " ")
      println()
    }
  }

  /**
   * Exercise 1
   */
  def pascal(c: Int, r: Int): Int = c match {
    case 0   => 1
    case `r` => 1
    case _   => pascal(c - 1, r - 1) + pascal(c, r - 1)
  }

  /**
   * Exercise 2
   */
  def balance(chars: List[Char]): Boolean = {
    def balanced(chars: List[Char], b: Int): Boolean = {
      if (b < 0) {
        false
      } else if (chars.isEmpty) {
        b == 0
      } else {
        balanced(chars.tail, chars.head match {
          case '(' => b + 1
          case ')' => b - 1
          case _ => b
        })
      }
    }       

    balanced(chars, 0)
  }

  /**
   * Exercise 3
   */
  def countChange(money: Int, coins: List[Int]): Int = {
    if (money == 0) {
      1 
    } else if (money < 0) {
      0
    } else if (coins.isEmpty) {
      0
    } else {  
      countChange(money - coins.head, coins) + countChange(money, coins.tail)
    }
  }
}
