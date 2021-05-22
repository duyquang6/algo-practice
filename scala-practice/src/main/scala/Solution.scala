import scala.+:

object Solution extends App {
    def subsetXORSum(nums: Array[Int]): Int = {
        var (i, j) = (0, nums.length - 1)
        var res = 0

        def helper(i: Int, cur: Int): Unit = {
            if (i == nums.length) {
                res += cur
                return
            }
            helper(i + 1, cur ^ nums(i))
            helper(i + 1, cur)
        }
        helper(0, 0)
        res
    }

    def sortSentence(s: String): String = {
        s.split(' ')
            .map(_.reverse)
            .sorted
            .map(_.reverse)
            .map(x => x.dropRight(x.length))
            .mkString(" ")
    }

    def decode(encoded: Array[Int], first: Int): Array[Int] = {
        encoded.foldLeft(Array(first)){ (acc, value) =>
            acc :+ (acc.last ^ value)
        }
    }

    def replaceDigits(s: String): String = {
        s.grouped(2).flatMap( x => x.length match {
            case 2 => Seq(x(0), (x(0).toInt + x(1).toInt - 48).toChar)
            case 1 => Seq(x(0))
        }).mkString
    }

    def truncateSentence(s: String, k: Int): String = {
        var count = 0
        s.takeWhile(x => {
            if (x ==  ' ') count += 1
            count < k
        })

        s.split(' ').take(k).mkString(" ")
    }

    def maximumPopulation(logs: Array[Array[Int]]): Int = {
        logs.fold(new Array[Int](2050 - 1950)){ (acc, value) =>
            (value(0) until value(1)).foreach(x => acc(x - 1950) += 1)
            acc
        }.zipWithIndex.maxBy(_._1)._2
    }
}

