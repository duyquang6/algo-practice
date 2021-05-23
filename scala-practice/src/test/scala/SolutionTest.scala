import org.scalatest.FunSuite

class SolutionTest extends FunSuite {

    test("testReplaceDigits") {
        assert(Solution.replaceDigits("a1c1e1") === "abcdef")
        assert(Solution.replaceDigits("a1b2c3d4e") === "abbdcfdhe")
    }

    test("testTruncateSentence") {
        assert(Solution.truncateSentence("Hello how are you Contestant", 4) === "Hello how are you")
        assert(Solution.truncateSentence("What is the solution to this problem", 4) === "What is the solution")
    }

    test("testArraySign") {
      assert(Solution.arraySign(Array(-1,-2,-3,-4,3,2,1)) === 1)
      assert(Solution.arraySign(Array(1,5,0,2,-3)) === 0)
      assert(Solution.arraySign(Array(-1,1,-1,1,-1)) === -1)
    }

    test("testCountPoints") {
      assert(
        Solution.countPoints(Array(Array(1, 3), Array(3, 3), Array(5, 3), Array(2, 2)), Array(Array(2,3,1), Array(4,3,1), Array(1,1,2))).sameElements(Array(3,2,2)))
    }
}
