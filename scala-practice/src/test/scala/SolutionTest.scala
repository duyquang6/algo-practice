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
}
