package patmat

import org.scalatest.FunSuite

import org.junit.runner.RunWith
import org.scalatest.junit.JUnitRunner

import patmat.Huffman._

@RunWith(classOf[JUnitRunner])
class HuffmanSuite extends FunSuite {
  trait TestTrees {
    val t1 = Fork(Leaf('a',2), Leaf('b',3), List('a','b'), 5)
    val t2 = Fork(Fork(Leaf('a',2), Leaf('b',3), List('a','b'), 5), Leaf('d',4), List('a','b','d'), 9)
  }

  test("weight of a larger tree") {
    new TestTrees {
      assert(weight(t1) === 5)
    }
  }

  test("chars of a larger tree") {
    new TestTrees {
      assert(chars(t2) === List('a','b','d'))
    }
  }

  test("string2chars(\"hello, world\")") {
    assert(string2Chars("hello, world") === List('h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'))
  }

  test("times(List('a', 'b', 'a'))") {
    assert(times(List('a', 'b', 'a')).sortBy(x => x._1) == List(('a', 2), ('b', 1)))
  }

  test("makeOrderedLeafList for some frequency table") {
    assert(makeOrderedLeafList(List(('t', 2), ('e', 1), ('x', 3))) === List(Leaf('e',1), Leaf('t',2), Leaf('x',3)))
  }

  test("combine of some leaf list") {
    val leaflist = List(Leaf('e', 1), Leaf('t', 2), Leaf('x', 4))
    assert(combine(leaflist) === List(Fork(Leaf('e',1),Leaf('t',2),List('e', 't'),3), Leaf('x',4)))
  }

  test("combine of some leaf list 2") {
    val leaflist = List(Leaf('e', 2), Leaf('t', 3), Leaf('x', 4))
    assert(combine(leaflist) === List(Leaf('x',4), Fork(Leaf('e',2),Leaf('t',3),List('e', 't'),5)))
  }

  test("combine of a singleton") {
    val leaflist = List(Leaf('e', 2))
    assert(combine(leaflist) === List(Leaf('e',2)))
  }

  test("combine of nil") {
    val leaflist = Nil
    assert(combine(leaflist) === Nil)
  }  

  test("until") {
    val t = List(Leaf('c', 1), Leaf('d', 1), Leaf('e', 2), Leaf('f', 2))
    val r = List(Fork(Leaf('f', 2), Fork(Fork(Leaf('c', 1), Leaf('d', 1), List('c', 'd'), 2), Leaf('e', 2), List('c', 'd', 'e'), 4), List('f', 'c', 'd', 'e'), 6))
    assert(until(singleton, combine)(t) === r)
  }

  test("createCodeTree") {
    val chars = List('c', 'd', 'f', 'e', 'f', 'e', 'e')
    val r = Fork(Leaf('e',3),Fork(Fork(Leaf('d',1),Leaf('c',1),List('d', 'c'),2),Leaf('f',2),List('d', 'c', 'f'),4),List('e', 'd', 'c', 'f'),7)
    assert(createCodeTree(chars) === r)
  }

  test("decodedSecret") {
    new TestTrees {
      val d = List('h', 'u', 'f', 'f', 'm', 'a', 'n', 'e', 's', 't', 'c', 'o', 'o', 'l')
      assert(decodedSecret === d)
    }
  }

  test("decode and encode a very short text should be identity") {
    new TestTrees {
      assert(decode(t1, encode(t1)("ab".toList)) === "ab".toList)
    }
  }

  test("decode/encode scala") {
    new TestTrees {
      assert(decode(frenchCode, encode(frenchCode)("scala".toList)) === "scala".toList)
    }
  }

  test("convert") {
    new TestTrees {
      println(convert(frenchCode))
    }
  }

  test("quickEncode") {
    new TestTrees {
      val d = List('h', 'u', 'f', 'f', 'm', 'a', 'n', 'e', 's', 't', 'c', 'o', 'o', 'l')
      assert(quickEncode(frenchCode)(d) === secret)
    }
  }

}
